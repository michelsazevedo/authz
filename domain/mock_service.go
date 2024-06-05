package domain

import (
	"context"
	"errors"
)

type MockService struct{}

func NewMockService() UserService {
	return &MockService{}
}

func (*MockService) SignIn(ctx context.Context, signInParams *SignInParams) (*SignInResponse, error) {
	if signInParams.Email == "peter.parker@marvel.com" && signInParams.Password == "password123" {
		return nil, errors.New("user not found")
	}

	return &SignInResponse{
		Token: "Secret",
		User: User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Locale:    "Americas/Sao_Paulo",
		},
	}, nil
}

func (*MockService) SignUp(ctx context.Context, user *User) error {
	if user.FirstName == "" {
		return errors.New("first name is required")
	}

	if user.LastName == "" {
		return errors.New("last name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (*MockService) Refresh(ctx context.Context, token JwtToken) (*SignInResponse, error) {
	return &SignInResponse{
		Token: "Secret",
		User: User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
			Locale:    "Americas/Sao_Paulo",
		},
	}, nil
}
