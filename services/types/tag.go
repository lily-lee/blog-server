package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Tag []string

func (t Tag) Value() (driver.Value, error) {
	if len(t) == 0 {
		return `[]`, nil
	}

	b, e := json.Marshal(t)
	return string(b), e
}

func (t *Tag) Scan(v interface{}) error {
	if v == nil {
		*t = make([]string, 0)
		return nil
	}

	val, ok := v.([]byte)
	if !ok {
		return errors.New("invalid scan source")
	}

	if len(val) == 0 {
		*t = make([]string, 0)
		return nil
	}

	return json.Unmarshal(val, t)
}
