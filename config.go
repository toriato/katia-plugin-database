package main

const (
	PostgreSQL DatabaseType = "postgresql"
)

type DatabaseType string

type Config struct {
	Type   DatabaseType `yaml:"type"`
	Source string       `yaml:"source"`
}
