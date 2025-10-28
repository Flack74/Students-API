package app

import (
	"net/http"

	"github.com/Flack74/Students-API/internal/config"
	"github.com/Flack74/Students-API/internal/http/handlers/router"
	"github.com/Flack74/Students-API/internal/storage/sqlite"
)

type Application struct {
	Server  *http.Server
	Storage *sqlite.Sqlite
}

func New() (*Application, error) {
	cfg := config.MustLoad()

	db, err := sqlite.New(cfg)
	if err != nil {
		return nil, err
	}

	r := router.New(db)

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: r,
	}

	return &Application{
		Server:  server,
		Storage: db,
	}, nil
}
