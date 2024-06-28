package top

import (
	"context"
	"github.com/shuyi-tangerine/little_book_book/gen-go/tangerine/little_book_book"
	"time"
)

type ContentDao interface {
	Select(ctx context.Context, req *ContentPO) (pos []*ContentPO, err error)
	SelectOne(ctx context.Context, req *ContentPO) (po *ContentPO, err error)
	Insert(ctx context.Context, req *ContentPO) (id int64, err error)
	Update(ctx context.Context, req *ContentPO) (affectedRows int64, err error)
}

type ContentService interface {
	// Save 保存
	Save(ctx context.Context, req *ContentPO) (err error)
	// Get 获取
	Get(ctx context.Context) (content *ContentPO, err error)
}

type ContentPO struct {
	ID          int64     `json:"id" db:"id"`
	ContentID   int64     `json:"content_id" db:"content_id"`
	Text        string    `json:"text" db:"text"`
	ContentType int64     `json:"content_type" db:"content_type"`
	Title       string    `json:"title" db:"title"`
	Backup      *string   `json:"backup" db:"backup"`
	Extra       *string   `json:"extra" db:"extra"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by"`

	// 查询的时候当参数用
	CreatedAtTimeRange *little_book_book.TimeRange `json:"-"`
	Limit              int64                       `json:"-"`
	Offset             int64                       `json:"-"`
}

func (m *ContentPO) ToContentData() (commentData *little_book_book.ContentData) {
	return &little_book_book.ContentData{
		Text: m.Text,
	}
}
