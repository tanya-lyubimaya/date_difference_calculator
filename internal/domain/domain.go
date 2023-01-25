package domain

import "context"

type UseCase interface {
	Close()
	CalculateDateDifference(ctx context.Context, year int) (string, error)
}
