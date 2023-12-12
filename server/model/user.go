package model

import (
	"boardsvr/db"
	"boardsvr/server/dto"
	"context"
	"errors"
	"fmt"
)

func SelectUserByID(id string) (*dto.User, error) {
	conn, err := db.GetInstance()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	sqlStr := `
		SELECT id, pwd, email, created
		FROM user
		WHERE id = ?
	`
	fmt.Println(sqlStr)
	row := conn.QueryRowContext(ctx, db.Parse(sqlStr), id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	user := new(dto.User)
	err = row.Scan(&user.ID, &user.Pwd, &user.Email, &user.Created)
	if err != nil {
		return nil, err
	}

	return user, err
}

func InsertUser(user *dto.User) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	ctx := context.Background()
	sqlStr := `insert into User(id, pwd, email) values(?,?,?)`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, db.Parse(sqlStr), user.ID, user.Pwd, user.Email)
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

func UpdateUser(user *dto.User) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	ctx := context.Background()
	sqlStr := `update user set pwd = ?, email = ? where id = ?`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, db.Parse(sqlStr), "", user.Pwd, user.Email)
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

func DeleteUser(id int) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	ctx := context.Background()
	sqlStr := `delete from user where id = ?`
	fmt.Println(sqlStr)
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
