package top

import (
	"context"
	"github.com/shuyi-tangerine/little_book_book/gen-go/tangerine/little_book_book"
)

type ContentService interface {
	// Save 保存
	Save(ctx context.Context, req *Content) (err error)
	// Get 获取
	Get(ctx context.Context) (content *Content, err error)
}

type Content struct {
	Text string `json:"text"`
}

func (m *Content) ToContentData() (commentData *little_book_book.ContentData) {
	return &little_book_book.ContentData{
		Text: m.Text,
	}
}
