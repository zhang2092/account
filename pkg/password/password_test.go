package password

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zhang2092/account/pkg/random"
	"golang.org/x/crypto/bcrypt"
)

func TestBcryptPassword(t *testing.T) {
	password := random.RandomString(6)

	hashedPassword1, err := BcryptHashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = BcryptComparePassword(hashedPassword1, password)
	require.NoError(t, err)

	wrongPassword := random.RandomString(6)
	err = BcryptComparePassword(hashedPassword1, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := BcryptHashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}

func TestScryptPassword(t *testing.T) {
	password := random.RandomString(6)

	hashedPassword1, err := ScryptHashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	b, err := ScryptComparePassword(hashedPassword1, password)
	require.NoError(t, err)
	require.Equal(t, b, true)

	wrongPassword := random.RandomString(6)
	b, err = ScryptComparePassword(hashedPassword1, wrongPassword)
	require.Error(t, err)
	require.Equal(t, b, false)

	hashedPassword2, err := ScryptHashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
