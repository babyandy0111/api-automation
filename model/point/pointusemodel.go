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
	pointUseFieldNames          = builderx.RawFieldNames(&PointUse{})
	pointUseRows                = strings.Join(pointUseFieldNames, ",")
	pointUseRowsExpectAutoSet   = strings.Join(stringx.Remove(pointUseFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	pointUseRowsWithPlaceHolder = strings.Join(stringx.Remove(pointUseFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cachePointUseIdPrefix = "cache#pointUse#id#"
)

type (
	PointUseModel interface {
		Insert(data PointUse) (sql.Result, error)
		FindOne(id int64) (*PointUse, error)
		Update(data PointUse) error
		Delete(id int64) error
	}

	defaultPointUseModel struct {
		sqlc.CachedConn
		table string
	}

	PointUse struct {
		Id          int64     `db:"id"`
		UserId      int64     `db:"user_id"`     // ref: user.id
		Point       int64     `db:"point"`       // 點數
		Description string    `db:"description"` // 點數相關描述
		UseType     int64     `db:"use_type"`    // 使用類型 1: gift shop, 2: donation
		RelateId    int64     `db:"relate_id"`   // 關聯 ID
		OperatorId  int64     `db:"operator_id"` // ref: sys_account.id
		CreatedAt   time.Time `db:"created_at"`
	}
)

func NewPointUseModel(conn sqlx.SqlConn, c cache.CacheConf) PointUseModel {
	return &defaultPointUseModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`point_use`",
	}
}

func (m *defaultPointUseModel) Insert(data PointUse) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, pointUseRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserId, data.Point, data.Description, data.UseType, data.RelateId, data.OperatorId, data.CreatedAt)

	return ret, err
}

func (m *defaultPointUseModel) FindOne(id int64) (*PointUse, error) {
	pointUseIdKey := fmt.Sprintf("%s%v", cachePointUseIdPrefix, id)
	var resp PointUse
	err := m.QueryRow(&resp, pointUseIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pointUseRows, m.table)
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

func (m *defaultPointUseModel) Update(data PointUse) error {
	pointUseIdKey := fmt.Sprintf("%s%v", cachePointUseIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pointUseRowsWithPlaceHolder)
		return conn.Exec(query, data.UserId, data.Point, data.Description, data.UseType, data.RelateId, data.OperatorId, data.CreatedAt, data.Id)
	}, pointUseIdKey)
	return err
}

func (m *defaultPointUseModel) Delete(id int64) error {

	pointUseIdKey := fmt.Sprintf("%s%v", cachePointUseIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, pointUseIdKey)
	return err
}

func (m *defaultPointUseModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePointUseIdPrefix, primary)
}

func (m *defaultPointUseModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pointUseRows, m.table)
	return conn.QueryRow(v, query, primary)
}
