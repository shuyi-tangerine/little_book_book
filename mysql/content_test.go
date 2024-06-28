package mysql

import (
	"context"
	"fmt"
	"github.com/shuyi-tangerine/little_book_book/top"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContentDao_Insert(t *testing.T) {
	dao, err := NewContentDao()
	if !assert.Nil(t, err) {
		return
	}

	extra := "{}"

	ContentPOs, err := dao.Insert(context.Background(), &top.ContentPO{
		ContentID:   1,
		Text:        "你好",
		ContentType: 1,
		Title:       "记仇",
		Backup:      nil,
		Extra:       &extra,
		CreatedBy:   "chenshuyi",
		UpdatedBy:   "chenshuyi",
	})
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println(ContentPOs)
}

func TestContentDao_Select(t *testing.T) {
	dao, err := NewContentDao()
	if !assert.Nil(t, err) {
		return
	}

	ContentPOs, err := dao.Select(context.Background(), &top.ContentPO{
		ID:        2020,
		ContentID: 1,
	})
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println(ContentPOs)
}

func TestContentDao_Update(t *testing.T) {
	dao, err := NewContentDao()
	if !assert.Nil(t, err) {
		return
	}

	backup := "{}"

	affectedRows, err := dao.Update(context.Background(), &top.ContentPO{
		ID:          2020,
		ContentID:   1,
		Text:        "哟哟",
		ContentType: 2,
		Title:       "嘿哈",
		Backup:      &backup,
		UpdatedBy:   "chenshuyi02",
	})
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println(affectedRows)
}
