package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	maker, err := New()
	assert.NoError(t, err)

	token, _, err := maker.CreateToken(1, "username-test", 1*time.Hour)
	assert.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "username-test", payload.Username)
}

func TestJwtExpire(t *testing.T) {
	maker, err := New()
	assert.NoError(t, err)

	token, _, err := maker.CreateToken(1, "username-test", 1*time.Second)
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)

	_, err = maker.VerifyToken(token)
	assert.NotErrorIs(t, err, ErrExpiredToken)
}
