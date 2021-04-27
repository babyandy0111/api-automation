package point

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userPointFieldNames          = builderx.RawFieldNames(&UserPoint{})
	userPointRows                = strings.Join(userPointFieldNames, ",")
	userPointRowsExpectAutoSet   = strings.Join(stringx.Remove(userPointFieldNames, "`create_time`", "`update_time`"), ",")
	userPointRowsWithPlaceHolder = strings.Join(stringx.Remove(userPointFieldNames, "`user_id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserPointUserIdPrefix = "cache#userPoint#userId#"
)

type (
	UserPointModel interface {
		Insert(data UserPoint) (sql.Result, error)
		FindOne(userId int64) (*UserPoint, error)
		Update(data UserPoint) error
		Delete(userId int64) error
	}

	defaultUserPointModel struct {
		sqlc.CachedConn
		table string
	}

	UserPoint struct {
		Version          int64 `db:"version"`            // 樂觀鎖
		UserId           int64 `db:"user_id"`            // ref: user.id
		Point            int64 `db:"point"`              // 總點數
		SoonExpiredPoint int64 `db:"soon_expired_point"` // 即將過期的點數
	}
)

func NewUserPointModel(conn sqlx.SqlConn, c cache.CacheConf) UserPointModel {
	return &defaultUserPointModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_point`",
	}
}

func (m *defaultUserPointModel) Insert(data UserPoint) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userPointRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Version, data.UserId, data.Point, data.SoonExpiredPoint)

	return ret, err
}

func (m *defaultUserPointModel) FindOne(userId int64) (*UserPoint, error) {
	userPointUserIdKey := fmt.Sprintf("%s%v", cacheUserPointUserIdPrefix, userId)
	var resp UserPoint
	err := m.QueryRow(&resp, userPointUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userPointRows, m.table)
		return conn.QueryRow(v, query, userId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserPointModel) Update(data UserPoint) error {
	userPointUserIdKey := fmt.Sprintf("%s%v", cacheUserPointUserIdPrefix, data.UserId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userPointRowsWithPlaceHolder)
		return conn.Exec(query, data.Version, data.Point, data.SoonExpiredPoint, data.UserId)
	}, userPointUserIdKey)
	return err
}

func (m *defaultUserPointModel) Delete(userId int64) error {

	userPointUserIdKey := fmt.Sprintf("%s%v", cacheUserPointUserIdPrefix, userId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.Exec(query, userId)
	}, userPointUserIdKey)
	return err
}

func (m *defaultUserPointModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserPointUserIdPrefix, primary)
}

func (m *defaultUserPointModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userPointRows, m.table)
	return conn.QueryRow(v, query, primary)
}
