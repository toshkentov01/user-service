package utils

import (
	"fmt"
	"github.com/toshkentov01/alif-tech-task/user-service/config"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	var conf = config.Get()
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			conf.PostgresHost,
			conf.PostgresPort,
			conf.PostgresUser,
			conf.PostgresPassword,
			conf.PostgresDatabase,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
