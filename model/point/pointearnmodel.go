package point

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	pointEarnFieldNames          = builderx.RawFieldNames(&PointEarn{})
	pointEarnRows                = strings.Join(pointEarnFieldNames, ",")
	pointEarnRowsExpectAutoSet   = strings.Join(stringx.Remove(pointEarnFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	pointEarnRowsWithPlaceHolder = strings.Join(stringx.Remove(pointEarnFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cachePointEarnIdPrefix = "cache#pointEarn#id#"
)

type (
	PointEarnModel interface {
		Insert(data PointEarn) (sql.Result, error)
		FindOne(id int64) (*PointEarn, error)
		Update(data PointEarn) error
		Delete(id int64) error
		QueryEarnSUM(userID int64) (int64, error)
	}

	defaultPointEarnModel struct {
		sqlc.CachedConn
		table string
	}

	PointEarn struct {
		Id          int64     `db:"id"`
		Description string    `db:"description"` // 點數相關描述
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
		UserId      int64     `db:"user_id"`     // ref: user.id
		Point       int64     `db:"point"`       // 點數
		ExpiredAt   time.Time `db:"expired_at"`  // 點數有效期限
		IsExpired   int64     `db:"is_expired"`  // 是否已過期
		OperatorId  int64     `db:"operator_id"` // ref: sys_account.id
	}
)

func NewPointEarnModel(conn sqlx.SqlConn, c cache.CacheConf) PointEarnModel {
	return &defaultPointEarnModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`point_earn`",
	}
}

func (m *defaultPointEarnModel) Insert(data PointEarn) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, pointEarnRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Description, data.CreatedAt, data.UpdatedAt, data.UserId, data.Point, data.ExpiredAt, data.IsExpired, data.OperatorId)

	return ret, err
}

func (m *defaultPointEarnModel) FindOne(id int64) (*PointEarn, error) {
	pointEarnIdKey := fmt.Sprintf("%s%v", cachePointEarnIdPrefix, id)
	var resp PointEarn
	err := m.QueryRow(&resp, pointEarnIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pointEarnRows, m.table)
		return conn.QueryRow(v, query, id)
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

func (m *defaultPointEarnModel) Update(data PointEarn) error {
	pointEarnIdKey := fmt.Sprintf("%s%v", cachePointEarnIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pointEarnRowsWithPlaceHolder)
		return conn.Exec(query, data.Description, data.CreatedAt, data.UpdatedAt, data.UserId, data.Point, data.ExpiredAt, data.IsExpired, data.OperatorId, data.Id)
	}, pointEarnIdKey)
	return err
}

func (m *defaultPointEarnModel) Delete(id int64) error {

	pointEarnIdKey := fmt.Sprintf("%s%v", cachePointEarnIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, pointEarnIdKey)
	return err
}

func (m *defaultPointEarnModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePointEarnIdPrefix, primary)
}

func (m *defaultPointEarnModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pointEarnRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultPointEarnModel) QueryEarnSUM(userId int64) (int64, error) {
	query := fmt.Sprintf(""+
		"SELECT SUM(total_earn - total_use) total "+
		"FROM("+
		"SELECT user_id, SUM(point) total_earn, 0 AS total_use FROM %s WHERE user_id = ? AND expired_at > now() GROUP BY user_id "+
		"UNION ALL "+
		"SELECT user_id, 0 AS total_earn, SUM(point) total_use FROM %s WHERE user_id = ? GROUP BY user_id) tmp", m.table, "point_use")
	var count int64
	err := m.QueryRowNoCache(&count, query, userId, userId)
	if err != nil {
		return 0, err
	}
	return count, nil
}
