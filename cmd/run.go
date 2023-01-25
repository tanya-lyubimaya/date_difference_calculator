package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/tanya-lyubimaya/date_difference_calculator/internal/app"
	"os"
	"os/signal"
)

func Execute() error {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.TraceLevel)

	application, err := app.New(logger)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	go func() {
		_ = application.Serve(":8080")
		logger.Traceln("application was closed!")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	application.GracefulStop()
	return nil
}
