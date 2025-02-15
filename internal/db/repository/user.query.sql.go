// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.query.sql

package repository

import (
	"context"
)

const findUsers = `-- name: FindUsers :many
SELECT id, email, role, status, created_at, updated_at FROM "user"
`

func (q *Queries) FindUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, findUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Role,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
