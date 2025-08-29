package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TestCase struct {
	TestNum int    `json:"test_num"`
	Input   string `json:"input"`
	Output  string `json:"output"`
}

type TestCases []TestCase

func (tc *TestCases) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("cannot decode test_case column of tests table. expected bytes, got %T", src)
	}
	return json.Unmarshal(bytes, tc)
}

func (tc TestCases) Value() (driver.Value, error) {
	return json.Marshal(tc)
}
