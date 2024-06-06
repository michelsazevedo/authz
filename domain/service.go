package domain

import "context"

type UserService interface {
	SignIn(ctx context.Context, params *SignInParams) (*SignInResponse, error)
	SignUp(ctx context.Context, user *User) error
	Refresh(ctx context.Context, token JwtToken) (*SignInResponse, error)
}
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindOne(ctx context.Context, value string) (*User, error)
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{userRepository: userRepository}
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
