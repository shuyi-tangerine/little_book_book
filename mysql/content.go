package mysql

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shuyi-tangerine/little_book_book/top"
	"strings"
	"time"
)

const (
	TableNameContent = "content"
)

type ContentDao struct {
	DB *sqlx.DB
}

func NewContentDao() (dao *ContentDao, err error) {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(mysql.shuyi.com:3306)/shuyi?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	dao = &ContentDao{
		DB: db,
	}
	// 校验继承关系
	var _ top.ContentDao = dao
	return dao, nil
}

func (m *ContentDao) Select(ctx context.Context, req *top.ContentPO) (pos []*top.ContentPO, err error) {
	var params []interface{}
	var where []string
	if req.ID != 0 {
		where = append(where, "`id`=?")
		params = append(params, req.ID)
	}
	if req.ContentID != 0 {
		where = append(where, "`content_id`=?")
		params = append(params, req.ContentID)
	}
	if req.CreatedAtTimeRange != nil {
		if req.CreatedAtTimeRange.S > 0 {
			where = append(where, "`created_at`>=?")
			params = append(params, time.Unix(req.CreatedAtTimeRange.S, 0))
		}
		if req.CreatedAtTimeRange.E > 0 {
			where = append(where, "`created_at`<?")
			params = append(params, time.Unix(req.CreatedAtTimeRange.E, 0))
		}
	}
	limit := ""
	if req.Limit != 0 {
		limit = "limit ?, ?"
		params = append(params, req.Offset, req.Limit)
	}

	err = m.DB.SelectContext(ctx, &pos, fmt.Sprintf("select * from %s where %s %s", TableNameContent, strings.Join(where, " and "), limit), params...)
	if err != nil {
		return
	}
	return
}

func (m *ContentDao) SelectOne(ctx context.Context, req *top.ContentPO) (po *top.ContentPO, err error) {
	pos, err := m.Select(ctx, req)
	if err != nil {
		return
	}

	if len(pos) > 1 {
		return nil, fmt.Errorf("selelct ont not one[%d]", len(pos))
	}

	if len(pos) == 0 {
		return
	}

	return pos[0], nil
}

func (m *ContentDao) Insert(ctx context.Context, req *top.ContentPO) (id int64, err error) {
	params := []interface{}{
		req.ContentID, req.Text, req.ContentType, req.Title, req.Backup, req.Extra,
		req.CreatedBy, req.UpdatedBy,
	}
	fields := []string{
		`content_id`, `text`, `content_type`, `title`, `backup`, `extra`,
		`created_by`, `updated_by`,
	}

	var placeholders []string
	for i := 0; i < len(fields); i++ {
		placeholders = append(placeholders, "?")
	}

	sql := fmt.Sprintf("insert into %s(%s) values(%s)", TableNameContent, strings.Join(fields, ","), strings.Join(placeholders, ","))
	res, err := m.DB.ExecContext(ctx, sql, params...)
	if err != nil {
		return
	}
	return res.LastInsertId()
}

func (m *ContentDao) Update(ctx context.Context, req *top.ContentPO) (affectedRows int64, err error) {
	if req.ID == 0 && req.ContentID == 0 {
		return 0, fmt.Errorf("update no unique key(id or content_id)")
	}

	var params []interface{}
	var setFields []string

	if req.Text != "" {
		params = append(params, req.Text)
		setFields = append(setFields, "`text`=?")
	}

	if req.ContentType != 0 {
		params = append(params, req.ContentType)
		setFields = append(setFields, "`content_type`=?")
	}

	if req.Title != "" {
		params = append(params, req.Title)
		setFields = append(setFields, "`title`=?")
	}

	if req.Backup != nil {
		params = append(params, req.Backup)
		setFields = append(setFields, "`backup`=?")
	}

	if req.Extra != nil {
		params = append(params, req.Extra)
		setFields = append(setFields, "`extra`=?")
	}

	var where []string
	if req.ID != 0 {
		where = append(where, "`id`=?")
		params = append(params, req.ID)
	}
	if req.ContentID != 0 {
		where = append(where, "`content_id`=?")
		params = append(params, req.ContentID)
	}

	sql := fmt.Sprintf("update %s set %s where %s", TableNameContent, strings.Join(setFields, ","), strings.Join(where, " and "))
	res, err := m.DB.ExecContext(ctx, sql, params...)
	if err != nil {
		return
	}

	return res.RowsAffected()
}
