package goflyway

import "errors"

var (
	ErrDatabaseConnectionNull    = errors.New("database connection is null")
	ErrUnsupportedDatabaseDriver = errors.New("unsupported database driver")
	ErrRunnerNotInitialized      = errors.New("runner not initialized")
	ErrLocationCannotBeEmpty     = errors.New("migration location cannot be empty")
)

var (
	warnNoMigrationFound = "no migrations found, are your location set up correctly?"
)

type ErrMigration struct {
	message string
}

func (e *ErrMigration) Error() string {
	return e.message
}

func throwErrMigration(err error) error {
	e := ErrMigration{
		message: err.Error(),
	}

	return &e
}
