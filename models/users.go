package models

import (
	"errors"

	"example.com/rest_api/db"
	"example.com/rest_api/utils"
)

type User struct {
	Id       int64
	Email    string  `binding:"required"`
	Password string   `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email , password) 
	VALUES(? , ?)`

	stmt, err := db.DB.Prepare(query)
	 if err != nil {
		return err
	 }

	 defer stmt.Close()
	 hashedPassword , err := utils.HashPassword(u.Password)
	 if err != nil {
		return err
	 }
     _ , err = stmt.Exec(u.Email , hashedPassword)
	 	if err != nil {
			return err
		}
    return nil	
}

func (u *User) ValidateCredentials() error{
  query := `SELECT Id, Password FROM users WHERE Email = ?`
  row := db.DB.QueryRow(query , u.Email)

  var hashedPassword string
  err := row.Scan(&u.Id, &hashedPassword)
      if err != nil {
		return errors.New("Invlaid credentials")
	  }
  
	  isValid := utils.CheckHashedPassword(u.Password , hashedPassword)

	  if !isValid {
		return errors.New("Invalid credentials")
	  }
  	  return nil
}