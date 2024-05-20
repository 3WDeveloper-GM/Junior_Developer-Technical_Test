package jsonbobjects

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JsonObjects is an abstraction on top of a map[string]interface{}
// This abstraction makes it so that the jsonb field in the database
// can be filled with a map of said peculiarities.
type JsonObjects map[string]interface{}

func (j JsonObjects) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (a *JsonObjects) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
