package model

import (
	"boardsvr/db"
	"boardsvr/server/dto"
	"context"
	"errors"
)

// TODO : ts column rename to updated
func SelectBoardByID(id int) (*dto.Board, error) {
	conn, err := db.GetInstance()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	sqlStr := `select id, title, content, author, ts 
				from board 
				where id = ?`
	board := new(dto.Board)
	row := conn.QueryRowContext(ctx, db.Parse(sqlStr), id)
	if row.Err() != nil {
		return board, err
	}

	err = row.Scan(&board.Title, &board.Content, &board.Author, &board.Updated)
	if err != nil {
		return board, err
	}

	return board, nil
}

func SelectBoardAll() ([]*dto.Board, error) {
	conn, err := db.GetInstance()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	sqlStr := `select id, title, content, author, ts 
				from board 
				order by ts desc`
	rows, err := conn.QueryContext(ctx, db.Parse(sqlStr))
	if err != nil {
		return nil, err
	}

	boards := make([]*dto.Board, 0)
	for rows.Next() {
		board := new(dto.Board)
		err = rows.Scan(&board.Id, &board.Title, &board.Content, &board.Author, &board.Updated)
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
	sqlStr := `insert into board(title, content, author) values(?,?,?)`
	res, err := conn.ExecContext(ctx, db.Parse(sqlStr), board.Title, board.Content, board.Author)
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

func UpdateBoard(board *dto.Board) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	ctx := context.Background()
	sqlStr := `update board set title = ?, content = ? where id = ?`
	res, err := conn.ExecContext(ctx, db.Parse(sqlStr), board.Title, board.Content, board.Id)
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

	ctx := context.Background()
	sqlStr := `delete from board where id = ?`
	res, err := conn.ExecContext(ctx, db.Parse(sqlStr), id)
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
