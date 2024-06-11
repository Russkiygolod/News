package internal

import (
	"News/internal/postgres"
	posts "News/pkg/model"
)

type Inter struct {
	ps postgres.Store
}

func New(ps postgres.Store) Inter {
	var Inters Inter
	Inters.ps = ps
	return Inters
}

func (i *Inter) GetPosts(limit int) ([]posts.Posts, error) {
	return (i.ps.GetPosts(limit))
}
