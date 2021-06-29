package main

import "time"

type dateVar struct {
	time.Time
}

const dateFormat = "2006-01-02"

// String is the method to format the flag's value, part of the flag.Value interface.
func (d *dateVar) String() string {
	return d.Format(dateFormat)
}

// Set is the method to set the flag value, part of the flag.Value interface.
func (d *dateVar) Set(v string) error {
	t, err := time.Parse(dateFormat, v)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}
