package urlcollection

import "context"

type Repository interface {
	Create(ctx context.Context, sourceURL string) error
	Find(ctx context.Context, shortURL string) (string, error)
	Update(ctx context.Context, shortURL string) error
	Delete(ctx context.Context, id string) error
}
