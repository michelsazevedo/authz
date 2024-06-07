package domain

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/michelsazevedo/authz/config"
	"golang.org/x/crypto/bcrypt"
)

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
	user, err := s.authenticate(ctx, signInParams)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	token, err := createJWTClaims(ctx, user)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &SignInResponse{Token: token, User: *user}, err
}

func (s *userService) SignUp(ctx context.Context, user *User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("error to encrypt password")
	}

	user.Password = string(password)
	return s.userRepository.Create(ctx, user)
}

func (s *userService) Refresh(ctx context.Context, jwtToken JwtToken) (*SignInResponse, error) {
	user, err := s.userRepository.FindOne(ctx, jwtToken.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	token, err := createJWTClaims(ctx, user)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &SignInResponse{Token: token, User: *user}, err
}

func (s *userService) authenticate(ctx context.Context, signInParams *SignInParams) (*User, error) {
	user, err := s.userRepository.FindOne(ctx, signInParams.Email)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInParams.Password))

	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func createJWTClaims(ctx context.Context, user *User) (string, error) {
	expiresAt := time.Now().Add(15 * time.Minute)

	token := JwtToken{
		UserID:    user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	secret := config.FromContext(ctx).Secret

	return jwtToken.SignedString([]byte(secret))
}
