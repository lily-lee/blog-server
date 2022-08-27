package types

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Birthday string

var dateFormat = "2006-01-02"

func (b Birthday) Value() (driver.Value, error) {
	if b == "" {
		return nil, nil
	}

	_, err := time.Parse(dateFormat, string(b))
	if err != nil {
		return nil, errors.New("invalid birthday")
	}

	return string(b), nil
}

func (b *Birthday) Scan(v interface{}) error {
	if v == nil {
		return nil
	}

	val, ok := v.(*time.Time)
	if !ok {
		return nil
	}

	*b = Birthday(val.Format(dateFormat))

	return nil
}
