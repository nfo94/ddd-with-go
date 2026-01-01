package chapter4

import (
	"errors"
	"log"
)

type Car interface {
	BeepBeep()
}

type BMW struct {
	// Hell no with this aberration
	heatedSeatSubscriptionEnabled bool
}

func (B BMW) BeepBeep() {
	// TODO implement
}

type Tesla struct {
	autoPilotEnabled bool
}

func (t Tesla) BeepBeep() {
	// TODO implement
}

func BuildCar(carType string) (Car, error) {
	switch carType {
	case "bmw":
		return BMW{heatedSeatSubscriptionEnabled: true}, nil
	case "tesla":
		return Tesla{autoPilotEnabled: true}, nil
	default:
		return nil, errors.New("unknown car type")
	}
}

func main() {
	myCar, err := BuildCar("tesla")
	if err != nil {
		log.Fatal(err)
	}
	// TODO
}
