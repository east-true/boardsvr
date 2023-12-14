package model

import (
	"boardsvr/db"
	"boardsvr/server/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	ID      string       `json:"user_id"`
	Pwd     string       `json:"user_pwd"`
	Email   string       `json:"user_email"`
	Created sql.NullTime `json:"user_created"`
}

func SelectUserByID(id string) (*User, error) {
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
	row := conn.QueryRowContext(ctx, helper.ParseSql(sqlStr), id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	user := new(User)
	err = row.Scan(&user.ID, &user.Pwd, &user.Email, &user.Created)
	if err != nil {
		return nil, err
	}

	return user, err
}

func InsertUser(user *User) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	ctx := context.Background()
	sqlStr := `insert into User(id, pwd, email) values(?,?,?)`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, helper.ParseSql(sqlStr), user.ID, user.Pwd, user.Email)
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

func UpdateUser(user *User) error {
	conn, err := db.GetInstance()
	if err != nil {
		return err
	}

	ctx := context.Background()
	sqlStr := `update user set pwd = ?, email = ? where id = ?`
	fmt.Println(sqlStr)
	res, err := conn.ExecContext(ctx, helper.ParseSql(sqlStr), "", user.Pwd, user.Email)
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
