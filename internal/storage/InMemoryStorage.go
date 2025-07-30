package storage

// 2 ручки
// 1-я Добавляет записть о просмотре материала автора пользователем
// 2-я Возвращает количество уникальных просмотров за вчера (уникальными пользователями)

import (
	"context"
	"sync"
	"time"
)

type data struct {
	UserId    int
	AuthorId  int
	Timestamp time.Time
}

type InMemoryStorage struct {
	store []data
	mu    *sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		store: make([]data, 0, 10000000),
		mu:    &sync.RWMutex{},
	}
}

func (s *InMemoryStorage) Add(ctx context.Context, userId, authorId int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store = append(s.store, data{
		UserId:    userId,
		AuthorId:  authorId,
		Timestamp: time.Now(),
	})

	return nil
}

// SELECT COUNT(*) FROM stats 
// WHERE stats.timestamp >= @minTime AND stats.timestamp <= @maxTime
// GROUP BY stats.authorId

func (s *InMemoryStorage) GetStats(ctx context.Context, authorId int) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	count := 0
	now := time.Now()
	minTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(-24 * time.Hour)
	maxTime := minTime.Add(24*time.Hour - 1)

	users := make(map[int]bool)

	for _, data := range s.store {
		if data.Timestamp.Compare(minTime) < 0 || data.Timestamp.Compare(maxTime) > 0 {
			continue
		}

		if data.AuthorId != authorId {
			continue
		}

		if _, ok := users[data.UserId]; ok {
			continue
		}

		users[data.UserId] = true
		count++
	}
	return count, nil
}
