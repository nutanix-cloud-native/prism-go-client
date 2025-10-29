package converged

type Categories[Category, CategoryListParams any] interface {
	// Getter is the interface for Get operations.
	Getter[Category]

	// Lister is the interface for List operations.
	Lister[Category, CategoryListParams]

	// Creator is the interface for Create operations.
	Creator[Category]

	// Updater is the interface for Update operations.
	Updater[Category]

	// Deleter is the interface for Delete operations.
	Deleter[Category]
}
