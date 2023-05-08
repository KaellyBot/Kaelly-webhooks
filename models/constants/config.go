package constants

import "github.com/rs/zerolog"

const (
	ConfigFileName = ".env"

	// MySQL URL with the following format: HOST:PORT.
	MySQLURL = "MYSQL_URL"

	// MySQL user.
	MySQLUser = "MYSQL_USER"

	// MySQL password.
	MySQLPassword = "MYSQL_PASSWORD"

	// MySQL database name.
	MySQLDatabase = "MYSQL_DATABASE"

	// RabbitMQ address.
	RabbitMQAddress = "RABBITMQ_ADDRESS"

	// Discord Bot Token.
	Token = "TOKEN"

	// Metric port.
	MetricPort = "METRIC_PORT"

	// Zerolog values from [trace, debug, info, warn, error, fatal, panic].
	LogLevel = "LOG_LEVEL"

	// Boolean; used to register commands at development guild level or globally.
	Production = "PRODUCTION"

	// Default values.
	defaultMySQLURL        = "localhost:3306"
	defaultMySQLUser       = ""
	defaultMySQLPassword   = ""
	defaultMySQLDatabase   = "kaellybot"
	defaultRabbitMQAddress = "amqp://localhost:5672"
	defaultToken           = ""
	defaultMetricPort      = 2112
	defaultLogLevel        = zerolog.InfoLevel
	defaultProduction      = false
)

func GetDefaultConfigValues() map[string]any {
	return map[string]any{
		MySQLURL:        defaultMySQLURL,
		MySQLUser:       defaultMySQLUser,
		MySQLPassword:   defaultMySQLPassword,
		MySQLDatabase:   defaultMySQLDatabase,
		RabbitMQAddress: defaultRabbitMQAddress,
		Token:           defaultToken,
		MetricPort:      defaultMetricPort,
		LogLevel:        defaultLogLevel.String(),
		Production:      defaultProduction,
	}
}
