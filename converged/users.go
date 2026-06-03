package converged

import (
	"context"

	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// Users defines the interface for IAM Users API operations.
type Users[User any] interface {
	// Getter is the interface for Get operations.
	Getter[User]

	// Lister is the interface for List operations.
	Lister[User]

	// Creator is the interface for Create operations.
	Creator[User]
	
	// UpdateUserState is the interface for updating user state operations.
	UpdateUserState(uuid string, status *authn.UserStateUpdate) (*authn.UserStateUpdateResponse, error)

	// ListUserKeys is the interface for listing user keys operations.
	ListUserKeys(ctx context.Context, userExtId string, opts ...ODataOption) ([]authn.Key, error)

	// CreateUserKey is the interface for creating user keys operations.
	CreateUserKey(ctx context.Context, userExtId string, key *authn.Key) (*authn.Key, error)
	
	// GetUserKeyById is the interface for getting user keys operations.
	GetUserKeyById(ctx context.Context, userExtId string, keyID string) (*authn.Key, error)
	
	// DeleteUserKeyById is the interface for deleting user keys operations.
	DeleteUserKeyById(ctx context.Context, userExtId string, keyID string) error

	// RevokeUserKey is the interface for revoking user keys operations.
	RevokeUserKey(ctx context.Context, userExtId string, keyID string) error
}
