package usecase

import (
	aport "clirzy/auth/internal/application/port"
	"clirzy/auth/internal/domain"
	"clirzy/auth/internal/domain/valueobject"
	testfakes "clirzy/auth/internal/testing/fakes"
	"context"
	"errors"
	"testing"
	"time"
)

func validUser() *aport.UserDTO {
	return &aport.UserDTO{
		ID:           "676fec45-6645-47eb-91dd-73f809a3cde0",
		Email:        "test@mail.com",
		PasswordHash: "hash",
	}
}

//
// ===== TESTS =====
//

func TestLoginUseCase_Execute_Success(t *testing.T) {
	now := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	refreshRepo := &testfakes.FakeRefreshTokenRepo{}

	uc := NewLoginUseCase(
		&testfakes.FakeUserGateway{
			User: validUser(),
		},
		&testfakes.FakeRoleAssignmentRepo{
			Roles: []valueobject.RoleCode{
				valueobject.UserRole,
				valueobject.AdminRole,
			},
		},
		&testfakes.FakePasswordVerifier{
			Ok: true,
		},
		refreshRepo,
		&testfakes.FakeTokenSigner{},
		&testfakes.FakeClock{
			FakeNow: now,
		},
		time.Hour*24,
	)

	res, err := uc.Execute(context.Background(), LoginCommand{
		Email:    "test@mail.com",
		Password: "password",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.AccessToken != "access.jwt" {
		t.Fatalf("unexpected access token: %s", res.AccessToken)
	}

	if res.RefreshToken != "refresh.jwt" {
		t.Fatalf("unexpected refresh token: %s", res.RefreshToken)
	}

	if refreshRepo.Token == nil {
		t.Fatal("refresh token was not saved")
	}

	if refreshRepo.Token.ExpiresAt.Before(now) {
		t.Fatal("refresh token expiration is invalid")
	}
}

func TestLoginUseCase_InvalidCredentials(t *testing.T) {
	tests := []struct {
		name     string
		user     *aport.UserDTO
		password bool
	}{
		{"user not found", nil, true},
		{"password mismatch", validUser(), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewLoginUseCase(
				&testfakes.FakeUserGateway{
					User: tt.user,
				},
				&testfakes.FakeRoleAssignmentRepo{},
				&testfakes.FakePasswordVerifier{
					Ok: tt.password,
				},
				&testfakes.FakeRefreshTokenRepo{},
				&testfakes.FakeTokenSigner{},
				&testfakes.FakeClock{
					FakeNow: time.Now(),
				},
				time.Hour,
			)

			_, err := uc.Execute(context.Background(), LoginCommand{
				Email:    "x@y.z",
				Password: "123",
			})

			if err != domain.ErrInvalidCredentials {
				t.Fatalf("expected ErrInvalidCredentials, got %v", err)
			}
		})
	}
}

func TestLoginUseCase_PropagatesErrors(t *testing.T) {
	expectedErr := errors.New("boom")

	tests := []struct {
		name string
		uc   *LoginUseCase
	}{
		{
			name: "user repo error",
			uc: NewLoginUseCase(
				&testfakes.FakeUserGateway{
					Err: expectedErr,
				},
				&testfakes.FakeRoleAssignmentRepo{},
				&testfakes.FakePasswordVerifier{Ok: true},
				&testfakes.FakeRefreshTokenRepo{},
				&testfakes.FakeTokenSigner{},
				&testfakes.FakeClock{FakeNow: time.Now()},
				time.Hour,
			),
		},
		{
			name: "roles repo error",
			uc: NewLoginUseCase(
				&testfakes.FakeUserGateway{
					User: validUser(),
				},
				&testfakes.FakeRoleAssignmentRepo{
					Err: expectedErr,
				},
				&testfakes.FakePasswordVerifier{Ok: true},
				&testfakes.FakeRefreshTokenRepo{},
				&testfakes.FakeTokenSigner{},
				&testfakes.FakeClock{FakeNow: time.Now()},
				time.Hour,
			),
		},
		{
			name: "sign access error",
			uc: NewLoginUseCase(
				&testfakes.FakeUserGateway{
					User: validUser(),
				},
				&testfakes.FakeRoleAssignmentRepo{},
				&testfakes.FakePasswordVerifier{Ok: true},
				&testfakes.FakeRefreshTokenRepo{},
				&testfakes.FakeTokenSigner{
					AccessErr: expectedErr,
				},
				&testfakes.FakeClock{FakeNow: time.Now()},
				time.Hour,
			),
		},
		{
			name: "save refresh error",
			uc: NewLoginUseCase(
				&testfakes.FakeUserGateway{
					User: validUser(),
				},
				&testfakes.FakeRoleAssignmentRepo{},
				&testfakes.FakePasswordVerifier{Ok: true},
				&testfakes.FakeRefreshTokenRepo{
					Err: expectedErr,
				},
				&testfakes.FakeTokenSigner{},
				&testfakes.FakeClock{FakeNow: time.Now()},
				time.Hour,
			),
		},
		{
			name: "sign refresh error",
			uc: NewLoginUseCase(
				&testfakes.FakeUserGateway{
					User: validUser(),
				},
				&testfakes.FakeRoleAssignmentRepo{},
				&testfakes.FakePasswordVerifier{Ok: true},
				&testfakes.FakeRefreshTokenRepo{},
				&testfakes.FakeTokenSigner{
					RefreshErr: expectedErr,
				},
				&testfakes.FakeClock{FakeNow: time.Now()},
				time.Hour,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.uc.Execute(context.Background(), LoginCommand{
				Email:    "a@b.c",
				Password: "123",
			})

			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}
