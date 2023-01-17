// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const updateValues = `-- name: UpdateValues :exec
UPDATE myschema.foo SET a = $1, b = $2
`

type UpdateValuesParams struct {
	A pgtype.Text
	B pgtype.Int4
}

func (q *Queries) UpdateValues(ctx context.Context, arg UpdateValuesParams) error {
	_, err := q.db.Exec(ctx, updateValues, arg.A, arg.B)
	return err
}