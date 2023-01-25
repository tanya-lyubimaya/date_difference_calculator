package usecase

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tanya-lyubimaya/date_difference_calculator/internal/domain"
	"strconv"
	"time"
)

type useCase struct {
	logger *logrus.Logger
} // Тут бы могли быть репозитории, но задача их не требует

func New(logger *logrus.Logger) (*useCase, error) {
	return &useCase{logger: logger}, nil
}

func (uc *useCase) CalculateDateDifference(ctx context.Context, year int) (string, error) {
	// Будем считать корректными годы после 1 года н.э., так как 0-ой год отсутствует в григорианском и юлианском календарях, а годы до н.э. мы не учитываем
	if year <= 0 {
		return "", domain.ErrInvalidYear
	}
	janFirstOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	currentDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	diff := currentDate.Sub(janFirstOfYear).Hours() / 24
	if diff > 0 {
		return "Days gone: " + strconv.Itoa(int(diff)), nil
	} else {
		return "Days left: " + strconv.Itoa(-int(diff)), nil
	}
}

func (uc *useCase) Close() {
	uc.logger.Traceln("close use-case")
	// Если бы были репозитории (repository), здесь бы вызывались их методы Close()
}
