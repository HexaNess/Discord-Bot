package utils

import (
	"os"
	"strconv"

	"github.com/disgoorg/snowflake/v2"
)

func GetSnowflakeIDFromEnv(envVar string) snowflake.ID {
	valStr := os.Getenv(envVar)
	valInt, _ := strconv.ParseInt(valStr, 10, 64)
	return snowflake.ID(valInt)
}
