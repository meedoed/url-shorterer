package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/meedoed/url-shorterer/internal/urlcollection"
	"github.com/meedoed/url-shorterer/pkg/client/postgresql"
	"github.com/meedoed/url-shorterer/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, url urlcollection.URL) error {
	q := `INSERT INTO urlcollection (source_url, short_url) VALUES ($1, $2)`
	err := r.CheckURL(ctx, url)
	if errors.Is(pgx.ErrNoRows, err) {
		r.client.QueryRow(ctx, q, url.SourceURL, url.ShortURL)
		return nil
	}
	return err
}

func (r *repository) Find(ctx context.Context, shortURL string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, shortURL string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) urlcollection.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) CheckURL(ctx context.Context, url urlcollection.URL) error {
	q := `
			SELECT short_url 
			FROM urlcollection 
			WHERE source_url = $1`
	err := r.client.QueryRow(ctx, q, url.SourceURL).Scan(url.ShortURL)
	if err != nil {
		return err
	}
	return nil
}
