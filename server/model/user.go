package model

import (
	"boardsvr/server/helper"
	"context"
	"errors"
	"fmt"
)

type UserAdaptor interface {
	SelectUserByID(id string) (*UserEntity, error)
	InsertUser(user *UserDTO) error
	UpdateUser(user *UserDTO) error
	DeleteUser(id int) error
}

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
