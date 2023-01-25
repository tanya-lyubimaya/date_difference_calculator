package app

import (
	"github.com/sirupsen/logrus"
	"github.com/tanya-lyubimaya/date_difference_calculator/internal/delivery/http"
	"github.com/tanya-lyubimaya/date_difference_calculator/internal/usecase"
)

type Application struct {
	server *http.Server
	logger *logrus.Logger
}

func (a *Application) Serve(port string) error {
	return a.server.Serve(port)
}

func (a *Application) GracefulStop() {
	a.server.GracefulShutdown()
}

func New(logger *logrus.Logger) (*Application, error) {
	uc, err := usecase.New(logger)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	server, err := http.New(uc, logger)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return &Application{server: server, logger: logger}, nil
}
