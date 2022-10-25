package user

import (
	"context"
	"reflect"
	"testing"

	"secrets-api/domain"
	apierr "secrets-api/infra/errors"
	"secrets-api/infra/log"
	"secrets-api/infra/mongodb/user_repo"
)

func TestNewService(t *testing.T) {
	type args struct {
		logger     log.Provider
		repository user_repo.IRepository
		apiErr     apierr.Provider
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.logger, tt.args.repository, tt.args.apiErr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateUser(t *testing.T) {
	type fields struct {
		logger     log.Provider
		repository user_repo.IRepository
		apiErr     apierr.Provider
	}
	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantApiErr *apierr.Message
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				logger:     tt.fields.logger,
				repository: tt.fields.repository,
				apiErr:     tt.fields.apiErr,
			}
			if gotApiErr := s.CreateUser(tt.args.ctx, tt.args.user); !reflect.DeepEqual(gotApiErr, tt.wantApiErr) {
				t.Errorf("CreateUser() = %v, want %v", gotApiErr, tt.wantApiErr)
			}
		})
	}
}

func TestService_GetUserByEmail(t *testing.T) {
	type fields struct {
		logger     log.Provider
		repository user_repo.IRepository
		apiErr     apierr.Provider
	}
	type args struct {
		ctx       context.Context
		userEmail string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantUser   *domain.User
		wantApiErr *apierr.Message
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				logger:     tt.fields.logger,
				repository: tt.fields.repository,
				apiErr:     tt.fields.apiErr,
			}
			gotUser, gotApiErr := s.GetUserByEmail(tt.args.ctx, tt.args.userEmail)
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUserByEmail() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
			if !reflect.DeepEqual(gotApiErr, tt.wantApiErr) {
				t.Errorf("GetUserByEmail() gotApiErr = %v, want %v", gotApiErr, tt.wantApiErr)
			}
		})
	}
}
