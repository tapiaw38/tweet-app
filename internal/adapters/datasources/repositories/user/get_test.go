package user_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user"
	"github.com/tapiaw38/tweet-app/internal/domain"
)

func TestGet(t *testing.T) {
	columns := []string{
		"id",
		"first_name",
		"last_name",
		"email",
		"created_at",
		"updated_at",
	}

	type fields struct {
		db   *sql.DB
		mock sqlmock.Sqlmock
	}

	var (
		validCreatedAt, _ = time.Parse("01-02-2006", "01-07-2023")
		validUpdatedAt, _ = time.Parse("01-02-2006", "01-07-2023")
	)

	tests := map[string]struct {
		id        uint64
		prepare   func(f *fields)
		expect    *domain.User
		expectErr error
	}{
		"when everything works fine": {
			id: 1,
			prepare: func(f *fields) {
				f.mock.ExpectQuery(
					`^SELECT 
						id, first_name, last_name, email, 
						created_at, updated_at 
					FROM users 
					WHERE id = \$1$`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows(columns).
						AddRow(
							1,
							"John",
							"Doe",
							"johndoe@example.com",
							validCreatedAt,
							validUpdatedAt,
						))
			},
			expect: &domain.User{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "johndoe@example.com",
				CreatedAt: validCreatedAt,
				UpdatedAt: validUpdatedAt,
			},
			expectErr: nil,
		},
		"when an error occurs while fetching the user": {
			id: 2,
			prepare: func(f *fields) {
				f.mock.ExpectQuery(
					`^SELECT 
						id, first_name, last_name, email, 
						created_at, updated_at 
					FROM users 
					WHERE id = \$1$`).
					WithArgs(2).
					WillReturnError(sql.ErrConnDone)
			},
			expect:    nil,
			expectErr: sql.ErrConnDone,
		},
		"when the user does not exist": {
			id: 3,
			prepare: func(f *fields) {
				f.mock.ExpectQuery(
					`^SELECT 
						id, first_name, last_name, email, 
						created_at, updated_at 
					FROM users 
					WHERE id = \$1$`).
					WithArgs(3).
					WillReturnRows(sqlmock.NewRows(columns))
			},
			expect:    nil,
			expectErr: sql.ErrNoRows,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			f := fields{
				db,
				mock,
			}

			if tc.prepare != nil {
				tc.prepare(&f)
			}

			repository := user.NewRepository(f.db)

			actual, actualErr := repository.Get(context.Background(), tc.id)
			assert.Equal(t, tc.expect, actual)
			assert.Equal(t, tc.expectErr, actualErr)
		})
	}
}
