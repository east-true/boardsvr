package model

import (
	"boardsvr/server/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type Model struct {
	mysql.Config

	instance *sql.DB
}

func (m *Model) Open() {
	db, err := sql.Open("mysql", m.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	m.instance = db
}

func (m *Model) Close() {
	m.instance.Close()
}

// BOARD
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

// USER

func (m *Model) SelectUserByID(id string) (*UserEntity, error) {
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return nil, err
	}

	sqlStr := `
		SELECT id, pwd, email, created
		FROM user
		WHERE id = ?
	`
	fmt.Println(sqlStr)
	row := conn.QueryRowContext(ctx, helper.FormatSql(sqlStr), id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	entity := new(UserEntity)
	err = row.Scan(&entity.ID, &entity.Pwd, &entity.Email, &entity.Created)
	if err != nil {
		return nil, err
	}

	return entity, err
}

func (m *Model) InsertUser(user *UserDTO) error {
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return err
	}

	sqlStr := `insert into User(id, pwd, email) values(?,?,?)`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, helper.FormatSql(sqlStr), user.ID, user.Pwd, user.Email)
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

func (m *Model) UpdateUser(user *UserDTO) error {
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return err
	}

	sqlStr := `update user set pwd = ?, email = ? where id = ?`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, helper.FormatSql(sqlStr), "", user.Pwd, user.Email)
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

func (m *Model) DeleteUser(id int) error {
	ctx := context.Background()
	conn, err := m.instance.Conn(ctx)
	if err != nil {
		return err
	}

	sqlStr := `delete from user where id = ?`
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
