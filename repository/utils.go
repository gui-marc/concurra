package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func toPgTimestamp(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}

func toPgUUID(id string) pgtype.UUID {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return pgtype.UUID{
			Valid: false,
		}
	}
	return pgtype.UUID{
		Bytes: parsedID,
		Valid: true,
	}
}
