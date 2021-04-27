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
	voucherUserFieldNames          = builderx.RawFieldNames(&VoucherUser{})
	voucherUserRows                = strings.Join(voucherUserFieldNames, ",")
	voucherUserRowsExpectAutoSet   = strings.Join(stringx.Remove(voucherUserFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	voucherUserRowsWithPlaceHolder = strings.Join(stringx.Remove(voucherUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheVoucherUserIdPrefix = "cache#voucherUser#id#"
)

type (
	VoucherUserModel interface {
		Insert(data VoucherUser) (sql.Result, error)
		FindOne(id int64) (*VoucherUser, error)
		Update(data VoucherUser) error
		Delete(id int64) error
	}

	defaultVoucherUserModel struct {
		sqlc.CachedConn
		table string
	}

	VoucherUser struct {
		IsExpired int64     `db:"is_expired"` // 是否已過期
		UpdatedAt time.Time `db:"updated_at"`
		Id        int64     `db:"id"`
		VoucherId int64     `db:"voucher_id"` // ref: voucher.id
		SerialNum string    `db:"serial_num"` // 票券序號
		ExpiredDt time.Time `db:"expired_dt"` // 票券有效期限
		UserId    int64     `db:"user_id"`    // ref: user.id
		Status    int64     `db:"status"`     // 0: 未使用, 1: 已使用
		CreatedAt time.Time `db:"created_at"`
	}
)

func NewVoucherUserModel(conn sqlx.SqlConn, c cache.CacheConf) VoucherUserModel {
	return &defaultVoucherUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`voucher_user`",
	}
}

func (m *defaultVoucherUserModel) Insert(data VoucherUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, voucherUserRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.IsExpired, data.UpdatedAt, data.VoucherId, data.SerialNum, data.ExpiredDt, data.UserId, data.Status, data.CreatedAt)

	return ret, err
}

func (m *defaultVoucherUserModel) FindOne(id int64) (*VoucherUser, error) {
	voucherUserIdKey := fmt.Sprintf("%s%v", cacheVoucherUserIdPrefix, id)
	var resp VoucherUser
	err := m.QueryRow(&resp, voucherUserIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", voucherUserRows, m.table)
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

func (m *defaultVoucherUserModel) Update(data VoucherUser) error {
	voucherUserIdKey := fmt.Sprintf("%s%v", cacheVoucherUserIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, voucherUserRowsWithPlaceHolder)
		return conn.Exec(query, data.IsExpired, data.UpdatedAt, data.VoucherId, data.SerialNum, data.ExpiredDt, data.UserId, data.Status, data.CreatedAt, data.Id)
	}, voucherUserIdKey)
	return err
}

func (m *defaultVoucherUserModel) Delete(id int64) error {

	voucherUserIdKey := fmt.Sprintf("%s%v", cacheVoucherUserIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, voucherUserIdKey)
	return err
}

func (m *defaultVoucherUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheVoucherUserIdPrefix, primary)
}

func (m *defaultVoucherUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", voucherUserRows, m.table)
	return conn.QueryRow(v, query, primary)
}
