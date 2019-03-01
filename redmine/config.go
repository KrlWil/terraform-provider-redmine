package redmine

import (
    "log"
    redmine "github.com/mattn/go-redmine"
)

type Config struct {
    redmineClient *redmine.Client
}

func (c *Config) createAndAuthenticateClient() {
    log.Printf("[INFO] creating redmine client")
    redmineClient := redmine.NewClient("http://localhost:8009", "4473c836594fe329b27b13a55331770d43da4367")

    c.redmineClient = redmineClient

}
