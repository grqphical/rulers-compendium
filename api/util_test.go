package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckLimit(t *testing.T) {
	valid_limit := "1"
	_, err := CheckLimit(valid_limit)

	assert.Nil(t, err, "Should be a valid limit")

	invalid_integer := "-1"
	_, err = CheckLimit(invalid_integer)

	assert.NotNil(t, err, "Should be an invalid limit")

	zero := "0"
	_, err = CheckLimit(zero)

	assert.NotNil(t, err, "Should be an invalid limit")

	not_a_number := "asdf"
	_, err = CheckLimit(not_a_number)

	assert.NotNil(t, err, "Should be an invalid limit")

	not_an_integer := "3.14159"
	_, err = CheckLimit(not_an_integer)

	assert.NotNil(t, err, "Should be an invalid limit")
}
