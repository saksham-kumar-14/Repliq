package main

import (
	"log"

	"github.com/saksham-kumar-14/Repliq/backend/internal/env"
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

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
