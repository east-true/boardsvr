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

func SelectUserByID(id string) (*UserEntity, error) {
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

	entity := new(UserEntity)
	err = row.Scan(&entity.ID, &entity.Pwd, &entity.Email, &entity.Created)
	if err != nil {
		return nil, err
	}

	return entity, err
}

func InsertUser(user *UserDTO) error {
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

func UpdateUser(user *UserDTO) error {
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
