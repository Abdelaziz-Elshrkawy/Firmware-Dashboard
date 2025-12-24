package authService

import (
	"errors"
	"firmware_server/database"
	authDtos "firmware_server/dtos/auth"
	"firmware_server/tables"
	"firmware_server/tables/tablesName"
	"firmware_server/utils"
	"fmt"
	"strings"
)

// the tables.Users will have only (id) and (username)

func Login(user authDtos.Creds) (authDtos.JwtUserResponse, error) {

	var userRow tables.Users

	res := database.DB.QueryRow(fmt.Sprintf("select id, username, password from %s where username = ?", tablesName.Users), user.UserName)

	err := res.Scan(&userRow.Id, &userRow.Username, &userRow.Password)

	if err != nil {
		return authDtos.JwtUserResponse{}, err
	}

	err = utils.ComparePassword([]byte(userRow.Password), user.Password)

	if err != nil {
		return authDtos.JwtUserResponse{}, errors.New("invalid credentials")
	}

	return authDtos.JwtUserResponse{
		Id:       userRow.Id,
		Username: userRow.Username,
	}, nil
}

func SignUp(user authDtos.Creds) (any, error) {
	var userRow tables.Users

	res := database.DB.QueryRow(fmt.Sprintf("select id, username, password from %s where username = ?", tablesName.Users), user.UserName)

	err := res.Scan(&userRow.Id, &userRow.Username, &userRow.Password)

	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, err
	}

	hash, err := utils.CreatePassword(user.Password)

	if err != nil {
		return nil, err
	}

	_, err = database.DB.Exec(fmt.Sprintf("insert into %s (username, password) values (?, ?)", tablesName.Users), user.UserName, hash)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, errors.New("username is used")
		}
		return nil, err
	}

	return "Registered", nil
}
