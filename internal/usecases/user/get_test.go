package user_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tapiaw38/tweet-app/internal/adapters/datasources/repositories/user/mocks"
	"github.com/tapiaw38/tweet-app/internal/domain"
	usecase "github.com/tapiaw38/tweet-app/internal/usecases/user"
	"go.uber.org/mock/gomock"
)

func TestGet(t *testing.T) {
	type fields struct {
		repository *mocks.MockRepository
	}

	var (
		validCreatedAt, _ = time.Parse("01-02-2006", "01-07-2023")
		validUpdatedAt, _ = time.Parse("01-02-2006", "01-07-2023")
	)

	tests := map[string]struct {
		id        uint64
		prepare   func(f *fields)
		expect    *usecase.GetOutput
		expectErr error
	}{
		"when everything works fine": {
			id: 1,
			prepare: func(f *fields) {
				f.repository.EXPECT().Get(gomock.Any(), uint64(1)).Return(&domain.User{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Email:     "johndoe@example.com",
					CreatedAt: validCreatedAt,
					UpdatedAt: validUpdatedAt,
				}, nil)
			},
			expect: &usecase.GetOutput{
				Data: usecase.UserOutputData{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Email:     "johndoe@example.com",
					CreatedAt: validCreatedAt,
					UpdatedAt: validUpdatedAt,
				},
			},
			expectErr: nil,
		},
		"when an error occurs": {
			id: 1,
			prepare: func(f *fields) {
				f.repository.EXPECT().Get(gomock.Any(), uint64(1)).Return(nil, fmt.Errorf("error"))
			},
			expect:    nil,
			expectErr: fmt.Errorf("error"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				repository: mocks.NewMockRepository(ctrl),
			}

			if tc.prepare != nil {
				tc.prepare(&f)
			}

			usecase := usecase.NewGetUsecase(f.repository)

			actual, actualErr := usecase.Execute(context.Background(), tc.id)
			assert.Equal(t, tc.expect, actual)
			assert.Equal(t, tc.expectErr, actualErr)
		})
	}
}
