package mysql

import (
	"database/sql"
	"fmt"
	"shopstoretest/entity"
)

func (mysql MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := mysql.DB.QueryRow("select * from users where phone_number = ?", phoneNumber)

	user := entity.User{}
	var createdAt []uint8

	uErr := row.Scan(&user.ID, &user.Name, &user.Password, &user.Credit, &user.PhoneNumber, &createdAt, &user.Role)

	if uErr != nil {
		if uErr == sql.ErrNoRows {

			return true, nil
		}

		return false, uErr
	}

	return false, nil
}

func (mysql MySQLDB) Register(user entity.User) (entity.User, error) {
	res, exErr := mysql.DB.Exec("insert into users(name, hashed_password, phone_number) values(?, ?, ?)",
		user.Name, user.Password, user.PhoneNumber)
	if exErr != nil {

		return entity.User{}, exErr
	}

	id, iErr := res.LastInsertId()
	if iErr != nil {

		return entity.User{}, fmt.Errorf("cant get last id after save new user in table %w", iErr)
	}

	user.ID = uint(id)

	return user, nil
}

func (mysql MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	row := mysql.DB.QueryRow("select * from users where phone_number = ?", phoneNumber)
	user := entity.User{}
	var createdAt []uint8
	var roleString string

	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Credit, &user.PhoneNumber, &createdAt, &roleString)
	if err != nil {

		return entity.User{}, err
	}

	return entity.User{
		ID:          user.ID,
		Role:        entity.MapToRoleEntity(roleString),
		Name:        user.Name,
		Password:    user.Password,
		Credit:      user.Credit,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (mysql MySQLDB) GetUserByID(ID uint) (entity.User, error) {
	row := mysql.DB.QueryRow("select * from users where id = ?", ID)
	user := entity.User{}
	var createdAt []uint8
	var roleString string

	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Credit, &user.PhoneNumber, &createdAt, &roleString)
	if err != nil {

		return entity.User{}, err
	}

	return entity.User{
		ID:          user.ID,
		Role:        entity.MapToRoleEntity(roleString),
		Name:        user.Name,
		Password:    user.Password,
		Credit:      user.Credit,
		PhoneNumber: user.PhoneNumber,
	}, nil
}
