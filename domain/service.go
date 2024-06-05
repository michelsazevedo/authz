package domain

import "context"

type UserService interface {
	SignIn(ctx context.Context, params *SignInParams) (*SignInResponse, error)
	SignUp(ctx context.Context, user *User) error
	Refresh(ctx context.Context, token JwtToken) (*SignInResponse, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) SignIn(ctx context.Context, signInParams *SignInParams) (*SignInResponse, error) {
	return &SignInResponse{}, nil
}

func (s *userService) SignUp(ctx context.Context, user *User) error {
	return nil
}

func (s *userService) Refresh(ctx context.Context, token JwtToken) (*SignInResponse, error) {
	return &SignInResponse{}, nil
}
