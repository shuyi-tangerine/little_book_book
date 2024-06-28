package filesystem

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/little_book_book/top"
	"testing"
)

func TestContentService_Save(t *testing.T) {
	service := NewContentService("/tmp")
	err := service.Save(context.Background(), &top.ContentPO{Text: "你好才是真的好\n广州好滴   \n  "})
	fmt.Println(err)
}

func TestContentService_Get(t *testing.T) {
	service := NewContentService("/tmp")
	content, err := service.Get(context.Background())
	fmt.Println(content, err)
}
