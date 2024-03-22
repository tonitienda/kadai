package authentication

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonitienda/kadai/backend-golang-rest/pkg/common"
)

func TestAuthenticateNewUser(t *testing.T) {
	auth := Auth{
		db: NewInMemoryAuthDB(),
	}

	userId1, userfound1 := auth.GetUserIdBySub("sub1")

	assert.True(t, userfound1)

	userId2, userfound2 := auth.GetUserIdBySub("sub1")

	assert.True(t, userfound2)

	assert.Equal(t, userId1, userId2)

	assert.True(t, common.IsValidUUID(userId1))

}
