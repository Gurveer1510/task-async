package persistance

import (
	"context"
	"fmt"
	"log"

	"github.com/Gurveer1510/task-scheduler/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	DB *pgxpool.Pool
}

func NewDatabase() (*Database, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s&channel_binding=require", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_NAME, config.DB_SSLMODE)

	fmt.Println("DATABASE URL:", dsn)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	cfg.MaxConns = 10

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		log.Println("New connection established")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database successfully")

	return &Database{DB:pool}, nil
}