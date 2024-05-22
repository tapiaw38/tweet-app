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

func TestCreate(t *testing.T) {
	type fields struct {
		repository *mocks.MockRepository
	}

	var (
		validCreatedAt, _ = time.Parse("01-02-2006", "01-07-2023")
		validUpdatedAt, _ = time.Parse("01-02-2006", "01-07-2023")
	)

	validUserInput := domain.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Password:  "password",
		CreatedAt: validCreatedAt,
		UpdatedAt: validUpdatedAt,
	}

	tests := map[string]struct {
		userInput domain.User
		prepare   func(f *fields)
		expect    usecase.CreateOutput
		expectErr error
	}{
		"when everything works fine": {
			userInput: validUserInput,
			prepare: func(f *fields) {
				f.repository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(uint64(1), nil)
				f.repository.EXPECT().Get(gomock.Any(), uint64(1)).Return(&domain.User{
					ID:        1,
					FirstName: "John",
					LastName:  "Doe",
					Email:     "johndoe@example.com",
					CreatedAt: validCreatedAt,
					UpdatedAt: validUpdatedAt,
				}, nil)
			},
			expect: usecase.CreateOutput{
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
		"when an error occurs while creating": {
			prepare: func(f *fields) {
				f.repository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(uint64(0), fmt.Errorf("internal error"))
			},
			expect:    usecase.CreateOutput{},
			expectErr: fmt.Errorf("internal error"),
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

			usecase := usecase.NewCreateUsecase(f.repository)

			actual, actualErr := usecase.Execute(context.Background(), tc.userInput)
			assert.Equal(t, tc.expect, actual)
			assert.Equal(t, tc.expectErr, actualErr)
		})
	}
}
