package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"pencatatan/internal/config"
	"strconv"
	"sync"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
	Close() error
	DB() *sql.DB
}

type service struct {
	db *sql.DB
}

var (
	dbInstance *service
	dbOne      sync.Once
	initErr    error
)

func New(cfg *config.Config) (Service, error) {
	dbOne.Do(func() {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		db, err := sql.Open("pgx", connStr)
		if err != nil {
			initErr = err
			return
		}

		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(time.Hour)
		db.SetConnMaxLifetime(30 * time.Minute)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = db.PingContext(ctx); err != nil {
			initErr = fmt.Errorf("failed to ping database: %w", err)
			return
		}

		dbInstance = &service{
			db: db,
		}

		log.Println("Database connection pool initialized")
	})

	if initErr != nil {
		return nil, initErr
	}

	return dbInstance, nil
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Printf("db down: %v", err)
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "It's healthy"

	dbStats := s.db.Stats()
	stats["open_connection"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	if dbStats.OpenConnections > 40 {
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 100 {
		stats["message"] = "The database has a high."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

func (s *service) DB() *sql.DB {
	return s.db
}

func (s *service) Close() error {
	log.Printf("Disconnected from database")
	return s.db.Close()
}
