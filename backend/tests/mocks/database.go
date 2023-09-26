
package mocks

import (
	"os"
)

func SetDatabaseCredentials() {
	os.Setenv("ENV", "TESTING")
	os.Setenv("DATABASE_PORT", "5432")
	os.Setenv("DATABASE_USER", "postgres")
	os.Setenv("DATABASE_PASSWORD", "postgres")
	os.Setenv("DATABASE_NAME", "peachtree_test")
	os.Setenv("DATABASE_HOST", "peachtree-database-1")
}