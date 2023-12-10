package model

import (
	"boardsvr/db"
	"boardsvr/server/dto"
	"context"
	"errors"
)

func SelectBoardByID(id string) (*dto.Board, error) {
	conn, err := db.GetInstance()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	sql := `select title, content, author, ts from board where id = ?`
	board := new(dto.Board)
	row := conn.QueryRowContext(ctx, db.Parse(sql), id)
	if row.Err() != nil {
		return board, err
	}

	err = row.Scan(&board.Title, &board.Content, &board.Author, &board.Created)
	if err != nil {
		return board, err
	}

	return board, nil
}

// TODO : fix err unsupported Scan, storing driver.Value type []uint8 into type *time.Time
func SelectBoardAll() ([]*dto.Board, error) {
	conn, err := db.GetInstance()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	sql := `select title, content, author, ts from board order by ts desc`
	rows, err := conn.QueryContext(ctx, db.Parse(sql))
	if err != nil {
		return nil, err
	}

	boards := make([]*dto.Board, 0)
	for rows.Next() {
		board := new(dto.Board)
		err = rows.Scan(&board.Title, &board.Content, &board.Author, &board.Created)
		if err != nil {
			return nil, err
		}

		boards = append(boards, board)
	}

	return boards, nil
}

func InsertBoard(board *dto.Board) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	ctx := context.Background()
	sql := `insert into board(title, content, author) values(?,?,?)`
	res, err := conn.ExecContext(ctx, db.Parse(sql), board.Title, board.Content, board.Author)
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
