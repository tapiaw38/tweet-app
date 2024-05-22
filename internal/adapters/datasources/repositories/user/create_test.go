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

func TestCreate(t *testing.T) {
	type fields struct {
		db   *sql.DB
		mock sqlmock.Sqlmock
		user domain.User
	}

	var (
		validCreatedAt, _ = time.Parse("01-02-2006", "01-07-2023")
		validUpdatedAt, _ = time.Parse("01-02-2006", "01-07-2023")
	)

	tests := map[string]struct {
		prepare        func(f *fields)
		expectedUserID uint64
		expectedError  error
	}{
		"when everything works fine while inserting": {
			prepare: func(f *fields) {
				f.mock.ExpectQuery(`^INSERT INTO users \(first_name, last_name, username, email, password, created_at, updated_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\) RETURNING id;$`).WithArgs(
					f.user.FirstName,
					f.user.LastName,
					f.user.Username,
					f.user.Email,
					f.user.Password,
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			expectedUserID: 1,
			expectedError:  nil,
		},
		"when an error occurs while inserting": {
			prepare: func(f *fields) {
				f.mock.ExpectQuery(`^INSERT INTO users \(first_name, last_name, username, email, password, created_at, updated_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\) RETURNING id;$`).WithArgs(
					f.user.FirstName,
					f.user.LastName,
					f.user.Username,
					f.user.Email,
					f.user.Password,
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnError(sql.ErrConnDone)
			},
			expectedError: sql.ErrConnDone,
		},
		"when the query returns an empty result": {
			prepare: func(f *fields) {
				f.mock.ExpectQuery(`^INSERT INTO users \(first_name, last_name, username, email, password, created_at, updated_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\) RETURNING id;$`).WithArgs(
					f.user.FirstName,
					f.user.LastName,
					f.user.Username,
					f.user.Email,
					f.user.Password,
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnRows(sqlmock.NewRows([]string{"id"}))
			},
			expectedError: sql.ErrNoRows,
		},
		"when a error occur while inserting on users table": {
			prepare: func(f *fields) {
				f.mock.ExpectQuery(`^INSERT INTO users \(first_name, last_name, username, email, password, created_at, updated_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\) RETURNING id;$`).WithArgs(
					f.user.FirstName,
					f.user.LastName,
					f.user.Username,
					f.user.Email,
					f.user.Password,
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
				).WillReturnError(sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			validUser := domain.User{
				FirstName: "John",
				LastName:  "Doe",
				Username:  "johndoe",
				Email:     "johndoe@example.com",
				CreatedAt: validCreatedAt,
				UpdatedAt: validUpdatedAt,
			}

			f := fields{
				db:   db,
				mock: mock,
				user: validUser,
			}

			if tc.prepare != nil {
				tc.prepare(&f)
			}

			repository := user.NewRepository(f.db)

			actualUserID, actualError := repository.Create(context.Background(), validUser)
			assert.Equal(t, tc.expectedUserID, actualUserID)
			assert.Equal(t, tc.expectedError, actualError)
		})
	}
}
