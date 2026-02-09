package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/sawkyawwalarhtwe/ecom-api/internal/env"
)

func main() {

	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}
	//logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	api := application{
		config: cfg,
		db:     conn,
	}

	logger.Info("Connected to the database", "dsn", cfg.db.dsn)

	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
}
