package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URL struct {
	ID        int64
	Url       string
	CreatedAt time.Time
}

type URLModel struct {
	DB *pgxpool.Pool
}

func (m *URLModel) Insert(ctx context.Context, url string) (int64, error) {
	var id int64
	r := m.DB.QueryRow(ctx,
		`INSERT INTO urls (url) VALUES ($1) RETURNING id`,
		url)
	if err := r.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (m *URLModel) GetIdByUrl(ctx context.Context, url string) (int64, error) {
	var id int64
	err := m.DB.QueryRow(ctx,
		`SELECT id FROM urls WHERE url = $1`, url).Scan(&id)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return 0, ErrNotFound
		default:
			return 0, err
		}
	}
	return id, nil
}

func (m *URLModel) GetById(ctx context.Context, id int64) (string, error) {
	var url string
	err := m.DB.QueryRow(ctx,
		`SELECT url FROM urls WHERE id = $1`, id).Scan(&url)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return "", ErrNotFound
		default:
			return "", err
		}
	}
	return url, nil
}
