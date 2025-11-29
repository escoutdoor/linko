package driver

import (
	"time"

	"github.com/escoutdoor/linko/common/pkg/errwrap"
	"github.com/escoutdoor/linko/driver/internal/entity"
)

type Driver struct {
	ID             string    `db:"id"`
	UserID         string    `db:"user_id"`
	TotalRatingSum float64   `db:"total_rating_sum"`
	ReviewCount    int32     `db:"review_count"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (e Driver) ToServiceEntity() entity.Driver {
	// this way we get driver's rating
	rating := e.TotalRatingSum / float64(e.ReviewCount)

	return entity.Driver{
		ID:          e.ID,
		UserID:      e.UserID,
		Rating:      float32(rating),
		ReviewCount: e.ReviewCount,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

type Drivers []Driver

func (e Drivers) ToServiceEntities() []entity.Driver {
	list := make([]entity.Driver, len(e))
	for _, d := range e {
		list = append(list, d.ToServiceEntity())
	}

	return list
}

func buildSQLError(err error) error {
	return errwrap.Wrap("build sql", err)
}

func executeSQLError(err error) error {
	return errwrap.Wrap("execute sql", err)
}

func scanRowError(err error) error {
	return errwrap.Wrap("scan row", err)
}

func scanRowsError(err error) error {
	return errwrap.Wrap("scan rows", err)
}
