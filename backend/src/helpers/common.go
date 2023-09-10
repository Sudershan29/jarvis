
package helpers

import (
	"os"
)

// Referencing https://stackoverflow.com/questions/40326540/how-to-assign-default-value-if-env-var-is-empty
func GetEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}