package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func (d DateRange) Hours() float64 {
	return d.End.Sub(d.Start).Hours()
}

func NewDateRange(start time.Time, end time.Time) (DateRange, error) {
	if start.IsZero() || end.IsZero() {
		return DateRange{}, errors.New("start or end data is empty")
	}

	if end.Before(start) {
		return DateRange{}, errors.New("end date ends after start date")
	}

	dr := DateRange{
		Start: start,
		End:   end,
	}

	return dr, nil
}

func main() {
	/*
		lifetime := DateRange{
			Start: time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
			End:   time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		}

		fmt.Println(lifetime.Hours())

		travelInTime := DateRange{
			Start: time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
			End:   time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		}

		fmt.Println(travelInTime.Hours())
	*/

	start := time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC)
	end := time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC)
	lifetime, err := NewDateRange(start, end)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(lifetime.Hours())
	}

	start = time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC)
	end = time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC)
	travelInTime, err := NewDateRange(start, end)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(travelInTime.Hours())
	}

}
