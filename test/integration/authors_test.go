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

	ctx := context.Background()
	author, err := testClient.CreateAuthor(ctx, &model.CreateAuthorRequest{
		FirstName: "Mark",
		LastName:  "Danielewski",
	})
	require.NoError(t, err)
	require.Equal(t, "Mark", author.FirstName)
	require.Equal(t, "Danielewski", author.LastName)
	require.NotNil(t, author.UUID)
}
