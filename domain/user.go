package domain

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type User struct {
	Id        pgtype.UUID `json:"id"`
	FirstName string      `json:"first_name" validate:"required"`
	LastName  string      `json:"last_name" validate:"required"`
	Email     string      `json:"email" validate:"required,email"`
	Password  string      `json:"password" validate:"required"`
	Locale    string      `json:"locale" valid:"required"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt time.Time   `json:"-"`
}

func (u *User) GetAttributes() []interface{} {
	return []interface{}{u.FirstName, u.LastName, u.Email, u.Password, u.Locale}
}

func (u *User) GetColumns() []interface{} {
	return []interface{}{
		&u.Id,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Locale,
		&u.CreatedAt,
		&u.UpdatedAt,
	}
}

// SignInParams ...
type SignInParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// SignInResponse ...
type SignInResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// JwtToken ...
type JwtToken struct {
	UserID    pgtype.UUID
	FirstName string
	LastName  string
	Email     string
	*jwt.RegisteredClaims
}
