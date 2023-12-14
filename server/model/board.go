package model

import (
	"boardsvr/db"
	"boardsvr/server/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func SelectBoardAll() ([]*BoardEntity, error) {
	entitys := make([]*BoardEntity, 0)
	conn, err := db.GetInstance()
	if err != nil {
		return entitys, err
	}

	sqlStr := `
	select id, title, content, author, updated 
	from board 
	order by updated desc
	`
	fmt.Println(sqlStr)
	ctx := context.Background()
	rows, err := conn.QueryContext(ctx, helper.ParseSql(sqlStr))
	if err != nil {
		return entitys, err
	}

	for rows.Next() {
		entity := new(BoardEntity)
		err = rows.Scan(&entity.Id, &entity.Title, &entity.Content, &entity.Author, &entity.Updated)
		if err != nil {
			return entitys, err
		}

		entitys = append(entitys, entity)
	}

	return entitys, nil
}

func SelectBoardByID(id int) (*BoardEntity, error) {
	conn, err := db.GetInstance()
	if err != nil {
		return nil, err
	}

	sqlStr := `
		select id, title, content, author, updated 
		from board 
		where id = ?
	`
	fmt.Println(sqlStr)
	ctx := context.Background()
	row := conn.QueryRowContext(ctx, helper.ParseSql(sqlStr), id)
	if row.Err() != nil {
		return nil, err
	}

	entity := new(BoardEntity)
	err = row.Scan(&entity.Title, &entity.Content, &entity.Author, &entity.Updated)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func SelectBoardByAuthor(author string) ([]*BoardEntity, error) {
	entitys := make([]*BoardEntity, 0)
	conn, err := db.GetInstance()
	if err != nil {
		return entitys, err
	}

	sqlStr := `
		select id, title, content, author, updated 
		from board 
		where author = ?
		order by updated desc
	`
	fmt.Println(sqlStr)
	ctx := context.Background()
	rows, err := conn.QueryContext(ctx, helper.ParseSql(sqlStr), author)
	if err != nil {
		return entitys, err
	}

	for rows.Next() {
		entity := new(BoardEntity)
		err = rows.Scan(&entity.Id, &entity.Title, &entity.Content, &entity.Author, &entity.Updated)
		if err != nil {
			return entitys, err
		}

		entitys = append(entitys, entity)
	}

	return entitys, nil
}

func InsertBoard(board *BoardEntity) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	ctx := context.Background()
	sqlStr := `insert into board(title, content, author) values(?,?,?)`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, helper.ParseSql(sqlStr), board.Title, board.Content, board.Author)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n < 1 {
		return errors.New("no affected rows")
	}

	return nil
}

func UpdateBoard(board *BoardEntity) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	sqlStr := `update board set title = ?, content = ? where id = ?`
	fmt.Println(sqlStr)
	ctx := context.Background()
	res, err := conn.ExecContext(ctx, helper.ParseSql(sqlStr), board.Title, board.Content, board.Id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n < 1 {
		return errors.New("no affected rows")
	}

	return nil
}

func DeleteBoard(id int) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	sqlStr := `delete from board where id = ?`
	fmt.Println(sqlStr)
	ctx := context.Background()
	res, err := conn.ExecContext(ctx, helper.ParseSql(sqlStr), id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n < 1 {
		return errors.New("no affected rows")
	}

	return nil
}
