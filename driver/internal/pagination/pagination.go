package pagination

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/escoutdoor/linko/common/pkg/errwrap"
)

type Cursor struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func DecodeCursor(token string) (*Cursor, error) {
	if token == "" {
		return nil, nil
	}

	d, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return nil, errwrap.Wrap("decode string", err)
	}

	var cursor Cursor
	if err := json.Unmarshal(d, &cursor); err != nil {
		return nil, errwrap.Wrap("json unmarshal", err)
	}

	return &cursor, nil
}

func EncodeCursor(id string, createdAt time.Time) string {
	cursor := Cursor{
		ID:        id,
		CreatedAt: createdAt,
	}

	d, _ := json.Marshal(cursor)
	return base64.RawURLEncoding.EncodeToString(d)
}
