package postgres

import (
	"News/pkg/storage"
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

func (s *Store) Posts(limit int) ([]storage.Posts, error) {
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
	var posts []storage.Posts //массив структур
	for rows.Next() {
		var post storage.Posts //структура
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
		// добавление переменной в массив результатов
		posts = append(posts, post)
	}
	return posts, rows.Err()
}

func (s *Store) NewPost(Posts []storage.Posts) error {
	posts := Posts
	var post storage.Posts
	for _, v := range posts {
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
