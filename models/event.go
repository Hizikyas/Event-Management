package models

import "time"

func main() {
	type Event struct {
		ID          int
		Name        string
		Description string
		Location    string
		DateTime    time.Time
		UserID      int
	}
}

func (e Event ) Add () {
// add event to the database
}