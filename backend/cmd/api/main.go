package main

import (
	"github.com/saksham-kumar-14/Repliq/backend/internal/db"
	"github.com/saksham-kumar-14/Repliq/backend/internal/env"
	"github.com/saksham-kumar-14/Repliq/backend/internal/store"
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

	// database
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Database Connected!")

	if err := db.AutoMigrate(&store.User{}); err != nil {
		logger.Fatalw("failed to migrate database", "error", err)
	}
	if err := db.AutoMigrate(&store.Comment{}); err != nil {
		logger.Fatalw("failed to migrate database", "erorr", err)
	}
	logger.Info("Database migrated successfully!")

	// store
	store := store.NewDbStorage(db)

	// main app
	app := &application{
		config: cfg,
		logger: logger,
		store:  store,
	}

	mux := app.mount()
	logger.Info(app.run(mux))
}
