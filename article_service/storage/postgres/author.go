package postgres

import (
	"errors"
	"fmt"
	"time"
	"yorqinbek/microservices/Blogpost/article_service/protogen/blogpost"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetAuthorByID ...
func (stg Postgres) GetAuthorByID(id string) (*blogpost.Author, error) {
	result := &blogpost.Author{}
	var createdAt, updatedAt *time.Time
	err := stg.db.QueryRow(`SELECT 
		id,
		fullname,
		created_at,
		updated_at
    FROM author WHERE id = $1`, id).Scan(
		&result.Id,
		&result.Fullname,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return result, err
	}

	if createdAt != nil {
		result.CreatedAt = timestamppb.New(*createdAt)
	}

	if updatedAt != nil {
		result.UpdatedAt = timestamppb.New(*updatedAt)
	}

	fmt.Printf("%+v", result)

	return result, nil
}

// AddAuthor ...
func (stg Postgres) AddAuthor(id string, entity *blogpost.CreateAuthorRequest) error {
	_, err := stg.db.Exec(`INSERT INTO author 
	(
		id,
		firstname,
		lastname,
		middlename
	) VALUES (
		$1,
		$2,
		$3
	)`,
		id,
		entity.Fullname,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetAuthorList ...
func (stg Postgres) GetAuthorList(offset, limit int, search string) (*blogpost.GetAuthorListResponse, error) {
	// resp = im.Db.InMemoryAuthorData
	resp := &blogpost.GetAuthorListResponse{
		Authors: make([]*blogpost.Author, 0),
	}
	rows, err := stg.db.Queryx(`SELECT
	id,
	fullname,
	created_at,
	updated_at,
	deleted_at
	FROM author WHERE deleted_at IS NULL AND ((firstname ILIKE '%' || $1 || '%') OR (lastname ILIKE '%' || $1 || '%'))
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)
	if err != nil {
		return resp, err
	}
	//var tempMiddlename *string
	for rows.Next() {
		a := &blogpost.Article{
			Content: &blogpost.Content{},
		}

		var updatedAt *string

		err := rows.Scan(
			&a.ID,
			&a.Fullname,
			//&a.Lastname,
			//&tempMiddlename,
			&a.CreatedAt,
			&updatedAt,
			&a.DeletedAt,
		)
		if err != nil {
			return resp, err
		}

		if updatedAt != nil {
			a.UpdatedAt = *updatedAt
		}

		resp.Authors = append(resp.Authors, a)
	}
	return resp, err
}

func (stg Postgres) UpdateAuthor(entity *blogpost.UpdateAuthorRequest) error {
	res, err := stg.db.NamedExec("UPDATE author  SET fullname=:f, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": entity.ID,
		"f":  entity.Fullname,
	})
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("Authr not found")
}

// DeleteAuthor ...
func (stg Postgres) DeleteAuthor(id string) error {
	res, err := stg.db.Exec("UPDATE author SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("Author not found")
}
