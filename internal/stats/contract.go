package stats

import "context"

type Storage interface {
	Add(ctx context.Context, userId, authorId int) error
	GetStats(ctx context.Context, authorId int) (int, error)
}
