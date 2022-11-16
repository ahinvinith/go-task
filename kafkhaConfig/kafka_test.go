package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainSuccess(t *testing.T) {
	c, err := kafkaconfig()
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, c.Address, "IDFE")
}
func TestMainFailure(t *testing.T) {
	c, err := kafkaconfig()
	assert.Nil(t, err)
	assert.NotNil(t, c)
}
