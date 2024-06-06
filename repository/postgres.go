package repository

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"
	"sync"

	"github.com/go-playground/validator"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

var validate *validator.Validate
var dbInstance *PostgresRepository
var once sync.Once

type PostgresRepository struct {
	db *pgx.Conn
}

func GetDbInstance(URL string) *PostgresRepository {
	once.Do(func() {
		conn, _ := pgConn(context.Background(), URL)
		dbInstance = &PostgresRepository{db: conn}
	})

	return dbInstance
}

func pgConn(ctx context.Context, URL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, URL)
	if err != nil {
		log.Fatal().Msg("Cannot connect to the Database")
		panic(err.Error())
	}

	if err = createSchema(conn); err != nil {
		log.Fatal().Msg(err.Error())
		return nil, err
	}

	if _, err = conn.Exec(ctx, "SELECT 1"); err != nil {
		panic(err.Error())
	}

	return conn, nil
}

func createSchema(db *pgx.Conn) error {
	databaseURL := db.Config().ConnString()

	m, err := migrate.New("file://./database/migrations", databaseURL)
	if err != nil {
		return err
	}

	if err = m.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
