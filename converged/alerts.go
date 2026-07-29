package converged

import "context"

// AlertAction is the action to perform on an alert through the manage-alert
// operation. Its values map one-to-one to the actionType parameter of the
// .../serviceability/alerts/{extId}/$actions/manage-alert API.
type AlertAction string

const (
	// AlertActionResolve resolves an open alert.
	AlertActionResolve AlertAction = "RESOLVE"

	// AlertActionAcknowledge acknowledges an open alert.
	AlertActionAcknowledge AlertAction = "ACKNOWLEDGE"

	// AlertActionUnknown represents an unknown value.
	AlertActionUnknown AlertAction = "$UNKNOWN"

	// AlertActionRedacted represents a redacted value.
	AlertActionRedacted AlertAction = "$REDACTED"
)

// Alerts defines the interface for managing alerts.
type Alerts[Alert any] interface {
	// Getter is the interface for Get operations.
	Getter[Alert]

	// Lister is the interface for List operations.
	Lister[Alert]

	// ManageAlert performs the given action (e.g. acknowledge or resolve) on
	// the entity identified by the given external identifier, waits for the
	// task to complete, and returns the updated entity.
	ManageAlert(ctx context.Context, uuid string, action AlertAction) (*Alert, error)

	// ManageAlertAsync performs the given action on the entity identified by
	// the given external identifier and returns an Operation to track the task.
	ManageAlertAsync(ctx context.Context, uuid string, action AlertAction) (Operation[Alert], error)
}
