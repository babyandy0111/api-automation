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
	voucherFieldNames          = builderx.RawFieldNames(&Voucher{})
	voucherRows                = strings.Join(voucherFieldNames, ",")
	voucherRowsExpectAutoSet   = strings.Join(stringx.Remove(voucherFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	voucherRowsWithPlaceHolder = strings.Join(stringx.Remove(voucherFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheVoucherIdPrefix = "cache#voucher#id#"
)

type (
	VoucherModel interface {
		Insert(data Voucher) (sql.Result, error)
		FindOne(id int64) (*Voucher, error)
		Update(data Voucher) error
		Delete(id int64) error
	}

	defaultVoucherModel struct {
		sqlc.CachedConn
		table string
	}

	Voucher struct {
		Name          string    `db:"name"`      // 商品名稱
		ImageUrl      string    `db:"image_url"` // 商品圖片連結
		Sort          int64     `db:"sort"`      // 排序
		Remark        string    `db:"remark"`    // 兌換說明
		CreatedAt     time.Time `db:"created_at"`
		UpdatedAt     time.Time `db:"updated_at"`
		Note          string    `db:"note"`       // 說明欄位 (使用步驟須知)
		Status        int64     `db:"status"`     // 1: 上架, 2: 下架
		ExpiredDt     time.Time `db:"expired_dt"` // 使用期限
		Country       string    `db:"country"`    // 國家 (ISO 3166)
		Id            int64     `db:"id"`
		Point         int64     `db:"point"`          // 需要多少點數兌換
		ExchangeType  int64     `db:"exchange_type"`  // 兌換種類 1: 寄送, 2: 親領, 3: Redeem code
		TermCondition string    `db:"term_condition"` // 注意事項
		IsExpired     int64     `db:"is_expired"`     // 是否已過期
	}
)

func NewVoucherModel(conn sqlx.SqlConn, c cache.CacheConf) VoucherModel {
	return &defaultVoucherModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`voucher`",
	}
}

func (m *defaultVoucherModel) Insert(data Voucher) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, voucherRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Name, data.ImageUrl, data.Sort, data.Remark, data.CreatedAt, data.UpdatedAt, data.Note, data.Status, data.ExpiredDt, data.Country, data.Point, data.ExchangeType, data.TermCondition, data.IsExpired)

	return ret, err
}

func (m *defaultVoucherModel) FindOne(id int64) (*Voucher, error) {
	voucherIdKey := fmt.Sprintf("%s%v", cacheVoucherIdPrefix, id)
	var resp Voucher
	err := m.QueryRow(&resp, voucherIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", voucherRows, m.table)
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

func (m *defaultVoucherModel) Update(data Voucher) error {
	voucherIdKey := fmt.Sprintf("%s%v", cacheVoucherIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, voucherRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.ImageUrl, data.Sort, data.Remark, data.CreatedAt, data.UpdatedAt, data.Note, data.Status, data.ExpiredDt, data.Country, data.Point, data.ExchangeType, data.TermCondition, data.IsExpired, data.Id)
	}, voucherIdKey)
	return err
}

func (m *defaultVoucherModel) Delete(id int64) error {

	voucherIdKey := fmt.Sprintf("%s%v", cacheVoucherIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, voucherIdKey)
	return err
}

func (m *defaultVoucherModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheVoucherIdPrefix, primary)
}

func (m *defaultVoucherModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", voucherRows, m.table)
	return conn.QueryRow(v, query, primary)
}
