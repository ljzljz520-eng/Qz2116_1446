package config

import "os"

type Config struct {
	DBPath  string
	Host    string
	Port    string
	Limit   int
	Modulus int
}

func Default() Config {
	return Config{DBPath: "streetlight.db", Host: "0.0.0.0", Port: "8080", Limit: 1000, Modulus: 40}
}
func Load() Config {
	c := Default()
	if v := os.Getenv("STREETLIGHT_DB"); v != "" {
		c.DBPath = v
	}
	if v := os.Getenv("STREETLIGHT_HOST"); v != "" {
		c.Host = v
	}
	if v := os.Getenv("STREETLIGHT_PORT"); v != "" {
		c.Port = v
	}
	return c
}
func (c Config) Address() string { return c.Host + ":" + c.Port }
func (c Config) Valid() bool {
	return c.DBPath != "" && c.Host != "" && c.Port != "" && c.Limit > 0 && c.Modulus > 0
}
func (c *Config) SetLimit(n int) bool {
	if n < 1 {
		return false
	}
	c.Limit = n
	return true
}
func (c *Config) SetModulus(n int) bool {
	if n < 1 {
		return false
	}
	c.Modulus = n
	return true
}
