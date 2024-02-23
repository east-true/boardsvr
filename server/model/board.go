package model

import (
	"boardsvr/server/helper"
	"context"
	"errors"
	"fmt"
)

func (m *Model) SelectBoardAll() ([]*BoardEntity, error) {
	entitys := make([]*BoardEntity, 0)
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return entitys, err
	}

	sqlStr := `
	select id, title, content, author, updated 
	from board 
	order by updated desc
	`
	fmt.Println(sqlStr)
	rows, err := conn.QueryContext(ctx, helper.FormatSql(sqlStr))
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

func (m *Model) SelectBoardByID(id int) (*BoardEntity, error) {
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return nil, err
	}

	sqlStr := `
		select id, title, content, author, updated 
		from board 
		where id = ?
	`
	fmt.Println(sqlStr)
	row := conn.QueryRowContext(ctx, helper.FormatSql(sqlStr), id)
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

func (m *Model) SelectBoardByAuthor(author string) ([]*BoardEntity, error) {
	entitys := make([]*BoardEntity, 0)
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return nil, err
	}

	sqlStr := `
	select id, title, content, author, updated 
	from board 
		where author = ?
		order by updated desc
	`
	fmt.Println(sqlStr)
	rows, err := conn.QueryContext(ctx, helper.FormatSql(sqlStr), author)
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

func (m *Model) InsertBoard(board *BoardDTO) error {
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return err
	}

	sqlStr := `insert into board(title, content, author) values(?,?,?)`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, helper.FormatSql(sqlStr), board.Title, board.Content, board.Author)
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

func (m *Model) UpdateBoard(board *BoardDTO) error {
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return err
	}

	sqlStr := `update board set title = ?, content = ? where id = ?`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, helper.FormatSql(sqlStr), board.Title, board.Content, board.Id)
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

func (m *Model) DeleteBoard(id int) error {
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return err
	}

	sqlStr := `delete from board where id = ?`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, helper.FormatSql(sqlStr), id)
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
