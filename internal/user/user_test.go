package user

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/SamStalschus/secrets-api/infra/bcrypt"
	"github.com/SamStalschus/secrets-api/internal"
	"github.com/golang/mock/gomock"

	apierr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/infra/log"
	"github.com/SamStalschus/secrets-api/infra/mongodb/user_repo"
)

func TestService_CreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	logger := log.NewMockProvider(mockCtrl)
	repository := user_repo.NewMockIRepository(mockCtrl)
	apiErr := apierr.NewMockProvider(mockCtrl)
	bcryptMck := bcrypt.NewMockProvider(mockCtrl)

	service := NewService(logger, repository, apiErr, bcryptMck)

	type structure struct {
		name       string
		prepare    func(t structure)
		ctx        context.Context
		user       *internal.User
		wantApiErr *apierr.Message
	}

	testCases := []structure{
		{
			name: "Create user with success",
			prepare: func(tt structure) {
				repository.EXPECT().FindUserByEmail(tt.ctx, tt.user.Email).Return(nil, nil)
				bcryptMck.EXPECT().EncryptPassword(tt.user.Password).Return([]byte("$2a$10$isZtzwTnub0jp1HgZi/4xO9RpGaWsx4GUcpVEA1DycepyoqiV0sH."), nil)
				repository.EXPECT().CreateUser(tt.ctx, gomock.Any()).Return("6355fd6995b4c8d74085a286", nil)
				logger.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any())
			},
			ctx: context.Background(),
			user: &internal.User{
				Email:    "zeze@email.com",
				Password: "123456789",
				Name:     "Zeze",
			},
			wantApiErr: nil,
		},
		{
			name: "Error because user already exists",
			prepare: func(tt structure) {
				repository.EXPECT().FindUserByEmail(tt.ctx, tt.user.Email).Return(tt.user, nil)
				apiErr.EXPECT().BadRequest(tt.wantApiErr.ErrorMessage, gomock.Any()).Return(tt.wantApiErr)
			},
			ctx: context.Background(),
			user: &internal.User{
				Email:    "zeze@email.com",
				Password: "123456789",
				Name:     "Zeze",
			},
			wantApiErr: &apierr.Message{
				ErrorMessage: "User already exists",
			},
		},
		{
			name: "Error because error in encrypt password",
			prepare: func(tt structure) {
				repository.EXPECT().FindUserByEmail(tt.ctx, tt.user.Email).Return(nil, nil)
				bcryptMck.EXPECT().EncryptPassword(tt.user.Password).Return(nil, fmt.Errorf(""))
				apiErr.EXPECT().InternalServerError(fmt.Errorf("")).Return(tt.wantApiErr)
			},
			ctx: context.Background(),
			user: &internal.User{
				Email:    "zeze@email.com",
				Password: "123456789",
				Name:     "Zeze",
			},
			wantApiErr: &apierr.Message{
				ErrorMessage: "Internal Server Error",
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt)
			if gotApiErr := service.CreateUser(tt.ctx, tt.user); !reflect.DeepEqual(gotApiErr, tt.wantApiErr) {
				t.Errorf("CreateUser() = %v, want %v", gotApiErr, tt.wantApiErr)
			}
		})
	}
}

func TestService_GetUserByEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	repository := user_repo.NewMockIRepository(mockCtrl)
	apiErr := apierr.NewMockProvider(mockCtrl)

	service := NewService(nil, repository, apiErr, nil)

	type structure struct {
		name       string
		prepare    func(t structure)
		ctx        context.Context
		userEmail  string
		wantUser   *internal.User
		wantApiErr *apierr.Message
	}

	testCases := []structure{
		{
			name: "User getted with success",
			prepare: func(tt structure) {
				repository.EXPECT().FindUserByEmail(tt.ctx, tt.userEmail).Return(tt.wantUser, nil)
			},
			ctx:       context.Background(),
			userEmail: "zeze@email.com",
			wantUser: &internal.User{
				Email: "zeze@email.com",
				Name:  "Zeze",
			},
			wantApiErr: nil,
		},
		{
			name: "Error because user don't exists",
			prepare: func(tt structure) {
				repository.EXPECT().FindUserByEmail(tt.ctx, tt.userEmail).Return(tt.wantUser, nil)
				apiErr.EXPECT().BadRequest(tt.wantApiErr.ErrorMessage, gomock.Any()).Return(tt.wantApiErr)
			},
			ctx:       context.Background(),
			userEmail: "zeze@email.com",
			wantUser:  nil,
			wantApiErr: &apierr.Message{
				ErrorMessage: "User don't exists",
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt)

			gotUser, gotApiErr := service.GetUserByEmail(tt.ctx, tt.userEmail)
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUserByEmail() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
			if !reflect.DeepEqual(gotApiErr, tt.wantApiErr) {
				t.Errorf("GetUserByEmail() gotApiErr = %v, want %v", gotApiErr, tt.wantApiErr)
			}
		})
	}
}
