package entity

// DateLayout date layout yyyy-mm-dd
const DateLayout = "2006-01-02"

type CursorDirection string

const (
	CursorDirectionNext CursorDirection = "next"
	CursorDirectionPrev CursorDirection = "prev"
)

type CursorInfo struct {
	Size       int64           `json:"size" example:"10"`
	Count      int64           `json:"count" example:"20"`
	HasNext    bool            `json:"hasNext" example:"true"`
	HasPrev    bool            `json:"hasPrev" example:"true"`
	Cursor     string          `json:"cursor" example:"1696466522533538969"`
	CursorDir  CursorDirection `json:"cursorType" example:"next"`
	PrevCursor string          `json:"prevCursor" example:"1696415865308136181"`
	NextCursor string          `json:"nextCursor" example:"1695785802835854036"`
}
