package entity

import (
	"context"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

type (
	Users struct {
		ID          uint
		Name        string
		Age         uint
		Gender      string
		Location    string
		Interest    string
		Preferences string
	}

	IUserRepository interface {
		GetUserByID(context context.Context, id uint) (*Users, error)
		GetUserByCriteria(c context.Context, request RequestFilterUsers) (users []Users, count int64, cursor paginator.Cursor, err error)
	}

	RequestFilterUsers struct {
		Name      string            `form:"name"`
		Gender    string            `form:"gender"`
		Age       []int             `form:"age"`
		Size      int64             `json:"size" form:"size,default=10" example:"10"`   // Optional, will fill with default value 10
		Cursor    string            `json:"cursor" form:"cursor,default=0" example:"0"` // Optional, will fill with default value 0
		CursorDir CursorDirection   `json:"cursorDir" form:"cursorDir,default=next"`    // Optional, will fill with default value NEXT
		SortBy    CustomerURLSortBy `json:"sortBy" form:"sortBy,default=id"`            // "id" is the same as "created at"
		SortDir   CustomerSortDir   `json:"sortDir" form:"sortDir,default=desc"`        // Default value is asc
	}
)

func (s *RequestFilterUsers) ToCursorInfo(cursor paginator.Cursor, count int64) CursorInfo {
	cursorInfo := CursorInfo{
		Size:      s.Size,
		Cursor:    s.Cursor,
		CursorDir: s.CursorDir,
		Count:     count,
	}

	if cursor.Before != nil {
		cursorInfo.PrevCursor = *cursor.Before
		cursorInfo.HasPrev = true
	}

	if cursor.After != nil {
		cursorInfo.NextCursor = *cursor.After
		cursorInfo.HasNext = true
	}

	return cursorInfo
}

func (s *RequestFilterUsers) SetDefaultValue() {
	if _, ok := CustomerSortByValues[s.SortBy]; !ok {
		// set ID as a default order by
		s.SortBy = CustomerSortByID
	}

	if _, ok := shortenURLSortDirValues[s.SortDir]; !ok {
		// set ascending as a default sort direction
		s.SortDir = CustomerSortDirDescending
	}

	if s.CursorDir == "" {
		s.CursorDir = CursorDirectionNext
	}

	if s.Cursor == "0" {
		s.Cursor = ""
	}

	if s.Size <= 0 {
		s.Size = 10
	}

	if s.Size > 20 {
		s.Size = 20
	}
}
