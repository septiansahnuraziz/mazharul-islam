package entity

import (
	"context"
	"github.com/mazharul-islam/utils"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"time"
)

type (
	CustomerStatus string

	Customer struct {
		ID         uint64    `gorm:"primary_key;auto_increment"`
		Name       string    `json:"name"`
		Identifier string    `json:"identifier"`
		Status     string    `json:"status"`
		UpdatedAt  time.Time `json:"updated_at"`
		CreatedAt  time.Time `json:"created_at"`
	}

	ICustomerService interface {
		CreateCustomer(ctx context.Context, request RequestCreateCustomer) error
		GetCustomers(c context.Context, requestFilter RequestFilterCustomer) ([]Customer, CursorInfo, error)
		GetCustomersWithES(c context.Context, requestFilter RequestFilterCustomer) ([]Customer, CursorInfo, error)
		IndexCustomerESDocumentByCustomerID(context context.Context, customer Customer) error
	}

	ICustomerRepository interface {
		Create(c context.Context, customer Customer) (Customer, error)
		GetCustomerByID(context context.Context, id uint) (*Customer, error)
		GetAll(c context.Context, request RequestFilterCustomer) (customers []Customer, count int64, cursor paginator.Cursor, err error)
		GetAllWithES(context context.Context, request RequestFilterCustomer) (customers []GetCustomerByCriteriaElasticsearchQueryDTO, count uint, err error)

		IndexCustomerES(c context.Context, customer Customer) error
	}

	RequestFilterCustomer struct {
		Name       string            `form:"name"`
		Identifier string            `form:"identifier"`
		Size       int64             `json:"size" form:"size,default=10" example:"10"`   // Optional, will fill with default value 10
		Cursor     string            `json:"cursor" form:"cursor,default=0" example:"0"` // Optional, will fill with default value 0
		CursorDir  CursorDirection   `json:"cursorDir" form:"cursorDir,default=next"`    // Optional, will fill with default value NEXT
		SortBy     CustomerURLSortBy `json:"sortBy" form:"sortBy,default=id"`            // "id" is the same as "created at"
		SortDir    CustomerSortDir   `json:"sortDir" form:"sortDir,default=desc"`        // Default value is asc
	}

	RequestCreateCustomer struct {
		Name       string `json:"name" validate:"required" example:"John"`
		Identifier string `json:"identifier" example:"john@mail.com"`
	}

	GetCustomerByCriteriaElasticsearchQueryDTO struct {
		ID uint `json:"id"`
	}
)

// CustomerStatus constants
const (
	CustomerStatusActive   CustomerStatus = "ACTIVE"
	CustomerStatusInactive CustomerStatus = "INACTIVE"
)

func (request RequestCreateCustomer) Validate() error {
	if err := validate.Struct(request); err != nil {
		return err
	}

	return nil
}

func (request RequestCreateCustomer) ToCustomerEntity() Customer {
	return Customer{
		Name:       request.Name,
		Identifier: request.Identifier,
		Status:     utils.ExpectedString(CustomerStatusInactive), // Customer default status is CustomerStatusInactive
	}
}

func (s *RequestFilterCustomer) ToCursorInfo(cursor paginator.Cursor, count int64) CursorInfo {
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

type CustomerURLSortBy string

type CustomerSortDir string

const (
	CustomerSortByID          CustomerURLSortBy = "id"
	CustomerSortDirAscending  CustomerSortDir   = "asc"
	CustomerSortDirDescending CustomerSortDir   = "desc"
)

var CustomerSortByValues = map[CustomerURLSortBy]bool{
	CustomerSortByID: true,
}

var shortenURLSortDirValues = map[CustomerSortDir]bool{
	CustomerSortDirAscending:  true,
	CustomerSortDirDescending: true,
}

func (s *RequestFilterCustomer) SetDefaultValue() {
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
