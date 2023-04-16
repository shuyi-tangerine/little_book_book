package thrift

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/little_book_book/filesystem"
	"github.com/shuyi-tangerine/little_book_book/gen-go/base"
	"github.com/shuyi-tangerine/little_book_book/gen-go/tangerine/little_book_book"
	"github.com/shuyi-tangerine/little_book_book/top"
)

type LittleBookBooker struct {
	contentService top.ContentService
}

func NewLittleBookBooker() little_book_book.LittleBookBooker {
	contentService := filesystem.NewContentServiceWithDefaultBookDir()
	return &LittleBookBooker{
		contentService: contentService,
	}
}

func (m *LittleBookBooker) SaveContent(ctx context.Context, req *little_book_book.SaveContentRequest) (resp *little_book_book.SaveContentResponse, err error) {
	resp = little_book_book.NewSaveContentResponse()
	resp.Base = base.NewRPCResponse()
	err = m.contentService.Save(ctx, &top.Content{Text: req.Text})
	if err != nil {
		resp.Base.Code = -1
		resp.Base.Message = fmt.Sprintf("%v", err)
		return resp, nil
	}
	resp.Content = &little_book_book.ContentData{Text: req.Text}
	return resp, nil
}

func (m *LittleBookBooker) GetContent(ctx context.Context, req *little_book_book.GetContentRequest) (resp *little_book_book.GetContentResponse, err error) {
	resp = little_book_book.NewGetContentResponse()
	resp.Base = base.NewRPCResponse()
	content, err := m.contentService.Get(ctx)
	if err != nil {
		resp.Base.Code = -1
		resp.Base.Message = fmt.Sprintf("%v", err)
		return resp, nil
	}

	resp.Content = content.ToContentData()
	return resp, nil
}
