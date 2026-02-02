package models

import (
	"time"

	"example.com/rest_api/db"
)

	type Event struct {
		ID          int64
		Name        string   `binding:"required"`  // this makes the gin context(ShouldBindJSON) to validate the field as required 
		Description string   `binding:"required"`
		Location    string   `binding:"required"`
		DateTime    time.Time   `binding:"required"`
		UserID      int
	}


func (e *Event ) Save ()  error {
	query := `INSERT INTO events (name , description , location , dateTime , user_id)
	VALUES (? , ? , ? , ? , ?)
	`

	stmt , err :=db.DB.Prepare(query) // we can simply use db.DB.Exec(query , e.Name , ..) but prepare is better for security like sql injection and performance if we are executing multiple times
	if err != nil {
		return err
	}
	defer stmt.Close()
    dateTimeStr := e.DateTime.Format(time.RFC3339)
	result , err := stmt.Exec(e.Name ,e.Description , e.Location , dateTimeStr , e.UserID)
	if err != nil {
		return err
	}

	resultID , err := result.LastInsertId() // this is for when we create a struct it needs a full data so we are include the id , it nothing do about the sql
	e.ID = resultID
	if err != nil {
		return err
	}
	return nil
}

func GetAllEvents () ([]Event , error) {
	query := `SELECT * FROM events`
	rows , err :=db.DB.Query(query) // if the database file is not change we use query , but like insert and update we use exec (Exec) , but before using Query() we can prepare like Prepare()

		if err != nil {
		return nil , err
		}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event 
		var dateTimeStr string
		err := rows.Scan(&event.ID , &event.Name , &event.Description , &event.Location , &dateTimeStr , &event.UserID)
		
		if err != nil {
			return nil , err
		}
		event.DateTime, err = time.Parse(time.RFC3339, dateTimeStr)
		if err != nil {
            return nil, err
		}
		events = append(events , event)
	}
	return  events , nil
}
