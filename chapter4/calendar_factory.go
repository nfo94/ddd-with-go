package chapter4

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	id            uuid.UUID
	userID        uuid.UUID
	from          time.Time
	to            time.Time
	hairDresserID uuid.UUID
}

func CreateBooking(from, to time.Time, userID, hairDresserID uuid.UUID) (*Booking, error) {
	// To define your own format, write down what the reference time would look like formatted
	// your way; see the values of constants like ANSIC, StampMicro or Kitchen for examples.
	// The model is to demonstrate what the reference time looks like so that the Format and
	// Parse methods can apply the same transformation to a general time value.
	closingTime, _ := time.Parse(time.Kitchen, "17:00pm")

	if from.After(closingTime) {
		return nil, errors.New("no appointments after closing time")
	}

	return &Booking{
		hairDresserID: hairDresserID,
		id:            uuid.New(),
		userID:        userID,
		from:          from,
		to:            to,
	}, nil
}
