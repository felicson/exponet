package expo

import "time"

//Expo struct
type Expo struct {
	Name, Announce, Description, City string
	DateStart, DateEnd                time.Time
}

//Valid check data
func (exh Expo) Valid() bool {
	if exh.Announce != "" &&
		exh.Name != "" &&
		exh.Description != "" &&
		exh.City != "" &&
		!exh.DateStart.IsZero() &&
		!exh.DateEnd.IsZero() {
		return true
	}
	return false
}
