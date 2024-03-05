package helpers

import (
	"os"
	"strconv"
)

const prepend = "jarvisproposal"

// Referencing https://stackoverflow.com/questions/40326540/how-to-assign-default-value-if-env-var-is-empty
func GetEnv(key string) string {
	return GetEnvWithDefault(key, "")
}

func GetEnvWithDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func IdToCalendarId(id int64) string {
	return prepend + strconv.FormatInt(id, 10)
}

func CalendarIdToId(idStr string) int64 {
	id, err := strconv.ParseInt(idStr[len(prepend):], 10, 0)
	if err != nil {
		return -1 // -1 represents default
	}
	return id
}
