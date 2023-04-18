package repository

import "errors"

var (
	ErrRepoNoData     = errors.New("no data")
	ErrRepoNoSuchData = errors.New("no such data")
	ErrRepoAlready    = errors.New("already exists")
	ErrNotFound       = errors.New("id not found")
)
