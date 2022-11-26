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

func (r *repository) Create(ctx context.Context, url urlcollection.URL) (string, error) {
	q := `INSERT INTO urlcollection (source_url, short_url) VALUES ($1, $2)`
	var err error
	url.ShortURL, err = r.CheckURL(ctx, url.SourceURL)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			url.GenURL()
			r.client.QueryRow(ctx, q, url.SourceURL, url.ShortURL)
			return url.ShortURL, nil
		}
		return "", err
	}
	return url.ShortURL, err
}

func (r *repository) Find(ctx context.Context, shortURL string) (string, error) {
	q := `SELECT source_url FROM urlcollection WHERE short_url = $1`
	var sourceURL string
	err := r.client.QueryRow(ctx, q, shortURL).Scan(&sourceURL)
	if err != nil {
		return "", err
	}
	return sourceURL, nil
}

/*func (r *repository) Update(ctx context.Context, shortURL string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}*/

func NewRepository(client postgresql.Client, logger *logging.Logger) urlcollection.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) CheckURL(ctx context.Context, sourceURL string) (string, error) {
	q := `
			SELECT short_url 
			FROM urlcollection 
			WHERE source_url = $1`
	var shortURL string
	err := r.client.QueryRow(ctx, q, sourceURL).Scan(shortURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}
