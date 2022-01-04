package config

type DatabaseConfiguration struct {
    DBDriver    string
    DBName      string
    DBUsername  string
    DBPassword  string
    DBHostname  string
    DBPort      string
    LoggingMode bool
}
