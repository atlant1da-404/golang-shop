package server

import (
	"context"
	"database/sql"
	"golang-shop/config"
	"golang-shop/internal/handler"
	"golang-shop/internal/models"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	cfg            = config.Init()
	hdl            = handler.NewHandler()
	rTime          = 10 * time.Second
	wTime          = 10 * time.Second
	MaxHeaderBytes = 1 << 20
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {

	a.httpServer = &http.Server{
		Addr:           ":" + cfg.HTTP.Port,
		Handler:        hdl.InitRoutes(),
		ReadTimeout:    rTime,
		WriteTimeout:   wTime,
		MaxHeaderBytes: MaxHeaderBytes,
	}
	err := a.ConnectDb()
	if err != nil {
		logrus.Fatalf("error by connect to db %+v", err.Error())
	}

	go func() {
		err := a.httpServer.ListenAndServe()
		if err != nil {
			logrus.Fatalf("error by starting server: %+v", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func (a *App) ConnectDb() error {
	sql, err := sql.Open("postgres", cfg.String())
	if err != nil {
		return err
	}

	err = sql.Ping()
	if err != nil {
		return err
	}
	models.NewSqlPostgres(sql)
	return nil
}
