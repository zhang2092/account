package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zhang2092/account/pkg/random"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Username:       random.RandomName(),
		HashedPassword: random.RandomPassword(),
		Email:          random.RandomEmail(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.HashedPassword, account.HashedPassword)
	require.Equal(t, arg.Email, account.Email)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	source := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), source.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, source.ID, account.ID)
	require.Equal(t, source.Username, account.Username)
	require.Equal(t, source.HashedPassword, account.HashedPassword)
	require.Equal(t, source.Email, account.Email)
	require.WithinDuration(t, source.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	source := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:             source.ID,
		HashedPassword: random.RandomPassword(),
		Email:          random.RandomEmail(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, source.ID, account.ID)
	require.Equal(t, source.Username, account.Username)
	require.Equal(t, arg.HashedPassword, account.HashedPassword)
	require.Equal(t, arg.Email, account.Email)
	require.WithinDuration(t, source.CreatedAt, account.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	source := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), source.ID)
	require.NoError(t, err)

	account, err := testQueries.GetAccount(context.Background(), source.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Equal(t, len(accounts), 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
