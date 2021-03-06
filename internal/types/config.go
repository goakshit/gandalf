package types

// PostgresConfig - Postgres config
type PostgresConfig struct {
	User         string
	Passwd       string
	DatabaseName string
	Host         string
	Port         int
	SSLMode      string
	Timezone     string
}

type KafkaConfig struct {
	Topic  string
	Server string
}

// Config - Complete application configuration
type Config struct {
	Database       PostgresConfig
	MessageService KafkaConfig
}
