package config

import (
	"github.com/example-golang-projects/clean-architecture/packages/env"
)

type CommonConfig struct {
	ProjectName      string      `json:"project_name"`
	ProjectDir       string      `json:"project_dir"`
	OrganizationName string      `json:"organization_name"`
	ImageURL         string      `json:"image_url"`
	Environment      env.EnvType `json:"environment"`
}

// Config contains User service config
type Config struct {
	Common CommonConfig `json:"common"`

	ServiceName     string   `json:"service_name"`
	ServicePort     string   `json:"service_port"`
	Database        DBConfig `json:"database"`
	UserServiceAddr string   `json:"user_service_addr"`
}

type DBConfig struct {
	UserDB PostgresConfig `json:"user_db"`
}

// PostgresConfig contains database config
type PostgresConfig struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
	Timeout  int    `json:"timeout"`

	MaxOpenConns    int `json:"max_open_conns"`
	MaxIdleConns    int `json:"max_idle_conns"`
	MaxConnLifetime int `json:"max_conn_lifetime"`

	GoogleAuthFile string `json:"google_auth_file"`
}
