package filesystem

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/little_book_book/top"
	"os"
)

type ContentService struct {
	bookDir string // 小本本所在目录
}

func NewContentService(bookDir string) top.ContentService {
	return &ContentService{
		bookDir: bookDir,
	}
}

func NewContentServiceWithDefaultBookDir() top.ContentService {
	bookDir := os.Getenv("BOOK_DIR")
	if bookDir == "" {
		bookDir = "/tmp"
	}
	return NewContentService(bookDir)
}

func (m *ContentService) Save(ctx context.Context, req *top.Content) (err error) {
	// 指定目录创建文件

	// 打开文件
	// O_WRONLY: 只写, O_TRUNC: 清空文件
	f, err := os.OpenFile(m.bookPath(), os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	// 将信息覆盖写入
	_, err = f.WriteString(req.Text)
	if err != nil {
		return err
	}

	return nil
}

func (m *ContentService) bookPath() string {
	return fmt.Sprintf("%s/book_content.txt", m.bookDir)
}

func (m *ContentService) Get(ctx context.Context) (content *top.Content, err error) {
	bts, err := os.ReadFile(m.bookPath())
	if err != nil {
		return
	}

	content = &top.Content{Text: string(bts)}
	return
}
