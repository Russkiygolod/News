package postgres

import (
	posts "News/pkg/model"
	"database/sql"
)

// Хранилище данных.
type Store struct {
	db *sql.DB
}

// Конструктор объекта хранилища.
func New(db *sql.DB) *Store {
	var postgres Store
	postgres.db = db
	return &postgres
}

func (s *Store) GetPosts(limit int) ([]posts.Posts, error) {
	rows, err := s.db.Query(`
	SELECT
	id,
	title,
	description,
	date,
	link

	FROM
	posts

	ORDER BY id DESC LIMIT $1;
`,
		limit,
	)
	if err != nil {
		return nil, err
	}
	var Posts []posts.Posts
	for rows.Next() {
		var post posts.Posts
		err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Description,
			&post.PubDate,
			&post.Link,
		)
		if err != nil {
			return nil, err
		}
		Posts = append(Posts, post)
	}
	return Posts, rows.Err()
}

func (s *Store) NewPost(Posts []posts.Posts) error {
	//posts := Posts
	var post posts.Posts
	for _, v := range Posts {
		post = v
		_, err := s.db.Exec(`
			
		INSERT INTO posts (title,description,date,link)
		VALUES ($1, $2, $3, $4);
		`,
			post.Title,
			post.Description,
			post.PubDate,
			post.Link,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
