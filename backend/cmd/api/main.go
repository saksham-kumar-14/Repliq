package main

import (
	"github.com/saksham-kumar-14/Repliq/backend/internal/env"
	"go.uber.org/zap"
)

const version = "0.1"

func main() {
	cfg := config{
		addr:        ":8000",
		frontendURL: ":5173",
		db: dbconfig{
			addr:         env.GetString("DB_URL", "postgresql://admin:admin@localhost5432/repliq?sslmode=disable"),
			maxOpenConns: env.GetInt("MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("MAX_IDLE_CONNS", 30),
			maxIdleTime:  "15m",
		},
	}

	// logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// main app
	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := app.mount()
	logger.Info(app.run(mux))
}
