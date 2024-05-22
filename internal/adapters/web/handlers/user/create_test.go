package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tapiaw38/tweet-app/internal/adapters/web/handlers/user"
	"github.com/tapiaw38/tweet-app/internal/domain"
	usecase "github.com/tapiaw38/tweet-app/internal/usecases/user"
	"github.com/tapiaw38/tweet-app/internal/usecases/user/mocks"
	"go.uber.org/mock/gomock"
)

func TestNewCreateHandler(t *testing.T) {
	type fields struct {
		usecase *mocks.MockCreateUsecase
		w       *httptest.ResponseRecorder
		r       *http.Request
	}

	validUserInput := domain.User{
		Email:     "johndoe@example.com",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	validRequestBody, _ := json.Marshal(validUserInput)

	tests := map[string]struct {
		requestBody  []byte
		prepare      func(f *fields)
		expectStatus int
	}{
		"when everything works fine": {
			requestBody: validRequestBody,
			prepare: func(f *fields) {
				f.usecase.EXPECT().Execute(gomock.Any(), validUserInput).Return(usecase.CreateOutput{
					Data: usecase.UserOutputData{
						Email:     "johndoe@example.com",
						UpdatedAt: time.Time{},
						CreatedAt: time.Time{},
					},
				}, nil)
			},
			expectStatus: http.StatusCreated,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				usecase: mocks.NewMockCreateUsecase(ctrl),
				w:       httptest.NewRecorder(),
				r:       httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(tc.requestBody)),
			}

			if tc.prepare != nil {
				tc.prepare(&f)
			}

			router := gin.Default()
			router.POST("/api/users", user.NewCreateHandler(f.usecase))

			router.ServeHTTP(f.w, f.r)

			assert.Equal(t, tc.expectStatus, f.w.Code)

			if f.w.Code == http.StatusCreated {
				var responseBody domain.User
				err := json.Unmarshal(f.w.Body.Bytes(), &responseBody)
				assert.NoError(t, err)
				assert.Equal(t, validUserInput.FirstName, responseBody.FirstName)
			}
		})
	}
}
