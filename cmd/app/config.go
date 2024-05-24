package app

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	database database
	DB       *sql.DB
	cors     struct {
		trustedOrigins []string
	}
}

type database struct {
	admin        string
	password     string
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (a *Application) setConfiguration() error {
	db := database{}
	cfg := Config{database: db}

	flag.StringVar(&cfg.database.admin, "admin", os.Getenv("BILLING_ADMIN"), "admin name")
	flag.StringVar(&cfg.database.password, "password", os.Getenv("BILLING_ADMIN_PASSWORD"), "database administrator password")

	flag.IntVar(&cfg.database.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.database.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.database.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	flag.Parse()

	if cfg.database.admin == "" || cfg.database.password == "" {
		return errors.New(`
      empty database admin or database password fields, cannot connect to database. 
      Set environment variables 'BILLING_ADMIN' and 'BILLING_ADMIN_PASSWORD'
      in order to match to your Postgres instance
      `)
	}

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@brawny-gecko-14732.7tt.aws-us-east-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full",
		cfg.database.admin, cfg.database.password)

	cfg.database.dsn = dsn

	a.Config = &cfg

	return nil
}

func (a *Application) setDBPool() error {
	db, err := sql.Open("postgres", a.Config.database.dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(a.Config.database.maxOpenConns)
	db.SetMaxIdleConns(a.Config.database.maxIdleConns)
	duration, err := time.ParseDuration(a.Config.database.maxIdleTime)
	if err != nil {
		return err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	a.Config.DB = db
	a.Logger.Info().Msg("Database connected")

	return nil
}
