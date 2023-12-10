package model

import (
	"boardsvr/db"
	"boardsvr/server/dto"
	"context"
)

func SelectBoardByID(id string) (*dto.Board, error) {
	conn, err := db.GetInstance()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	sql := `select title, content, author from board where id = ?`
	row := conn.QueryRowContext(ctx, db.Parse(sql), id)
	if row.Err() != nil {
		return nil, err
	}

	board := new(dto.Board)
	err = row.Scan(board.Title, board.Content, board.Author.Email)
	if err != nil {
		return nil, err
	}

	return board, nil
}

func SelectBoardAll() ([]*dto.Board, error) {
	conn, err := db.GetInstance()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	sql := `select title, content, author from board order by creted desc`
	rows, err := conn.QueryContext(ctx, db.Parse(sql))
	if err != nil {
		return nil, err
	}

	var boards []*dto.Board
	for rows.Next() {
		board := new(dto.Board)
		err = rows.Scan(board.Title, board.Content, board.Author.Email)
		if err != nil {
			return nil, err
		}

	}

	return boards, nil
}
