// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/JPZ13/service/model"
	"github.com/stretchr/testify/require"
)

func TestCreateAuthor(t *testing.T) {
	t.Parallel()

	createAuthor(t, "Mark", "Danielewski")
}

func createAuthor(t *testing.T, firstname, lastname string) {
	ctx := context.Background()
	author, err := testClient.CreateAuthor(ctx, &model.CreateAuthorRequest{
		FirstName: firstname,
		LastName:  lastname,
	})

	require.NoError(t, err)
	require.Equal(t, firstname, author.FirstName)
	require.Equal(t, lastname, author.LastName)
	require.NotNil(t, author.UUID)
}
