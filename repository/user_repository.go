package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5"
	"github.com/michelsazevedo/authz/config"
	"github.com/michelsazevedo/authz/domain"
	"github.com/rs/zerolog/log"
)

const createUser = `INSERT INTO users(first_name, last_name, email, password, locale)
VALUES ($1, $2, $3, $4, $5) RETURNING id`

const getUser = `SELECT id, first_name, last_name, email, locale, created_at, updated_at
FROM users WHERE email = $1`

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(conf *config.Config) domain.UserRepository {
	postgresDb := GetDbInstance(conf.GetDatabaseURL())
	return &userRepository{db: postgresDb.db}
}

func (u *userRepository) FindOne(ctx context.Context, value string) (*domain.User, error) {
	user := &domain.User{}

	err := u.db.QueryRow(ctx, getUser, value).Scan(user.GetColumns()...)
	if err != nil {
		log.Info().Msg("here: " + err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	if err := validatesUser(user); err != nil {
		return err
	}

	if err := u.db.QueryRow(ctx, createUser, user.GetAttributes()...).Scan(&user.Id); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func validatesUser(user *domain.User) error {
	validate = validator.New()
	err := validate.Struct(user)

	if err != nil {
		buf := bytes.NewBufferString("Validation Error: ")

		for _, err := range err.(validator.ValidationErrors) {
			buf.WriteString(fmt.Sprintf("Field %s is %s. ", err.Field(), err.ActualTag()))
		}

		return errors.New(buf.String())
	}

	return nil
}
