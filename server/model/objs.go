package model

import (
	"database/sql"
	"time"
)

type BoardDTO struct {
	Id      int       `json:"board_id" uri:"board_id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Author  string    `json:"author" form:"author"`
	Updated time.Time `json:"updated"`
}

func (dto *BoardDTO) ToEntity() *BoardEntity {
	var id sql.NullInt32
	var title sql.NullString
	var content sql.NullString
	var author sql.NullString
	var updated sql.NullTime

	err := id.Scan(dto.Id)
	if err != nil {
		return nil
	}

	err = title.Scan(dto.Title)
	if err != nil {
		return nil
	}

	err = content.Scan(dto.Content)
	if err != nil {
		return nil
	}

	err = author.Scan(dto.Author)
	if err != nil {
		return nil
	}

	err = updated.Scan(dto.Updated)
	if err != nil {
		return nil
	}

	return &BoardEntity{
		Id:      id,
		Title:   title,
		Content: content,
		Author:  author,
		Updated: updated,
	}
}

type BoardEntity struct {
	Id      sql.NullInt32
	Title   sql.NullString
	Content sql.NullString
	Author  sql.NullString
	Updated sql.NullTime
}

func (entity *BoardEntity) ToDTO() *BoardDTO {
	if !entity.Id.Valid {
		return nil
	}

	if !entity.Title.Valid {
		return nil
	}

	if !entity.Content.Valid {
		return nil
	}

	if !entity.Author.Valid {
		return nil
	}

	if !entity.Updated.Valid {
		return nil
	}

	return &BoardDTO{
		Id:      int(entity.Id.Int32),
		Title:   entity.Title.String,
		Content: entity.Content.String,
		Author:  entity.Content.String,
		Updated: entity.Updated.Time,
	}
}

type UserDTO struct {
	ID      string    `json:"user_id"`
	Pwd     string    `json:"user_pwd"`
	Email   string    `json:"user_email"`
	Role    string    `json:"user_role"`
	Created time.Time `json:"user_created"`
}

func (dto *UserDTO) ToEntity() *UserEntity {
	entity := new(UserEntity)

	if err := entity.ID.Scan(dto.ID); err != nil {
		return nil
	}

	if err := entity.Pwd.Scan(dto.Pwd); err != nil {
		return nil
	}

	if err := entity.Email.Scan(dto.Email); err != nil {
		return nil
	}

	if err := entity.Created.Scan(dto.Created); err != nil {
		return nil
	}

	return entity
}

type UserEntity struct {
	ID      sql.NullString
	Pwd     sql.NullString
	Email   sql.NullString
	Created sql.NullTime
}

func (entity *UserEntity) ToDTO() *UserDTO {
	if !entity.ID.Valid {
		return nil
	}

	if !entity.Pwd.Valid {
		return nil
	}

	if !entity.Email.Valid {
		return nil
	}

	if !entity.Created.Valid {
		return nil
	}

	return &UserDTO{
		ID:      entity.ID.String,
		Pwd:     entity.Pwd.String,
		Email:   entity.Email.String,
		Created: entity.Created.Time,
	}
}
