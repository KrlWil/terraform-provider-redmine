package redmine

import (
    "log"
    "os"
    redmine "github.com/mattn/go-redmine"
)

type Config struct {
    redmineClient *redmine.Client
}

func (c *Config) createAndAuthenticateClient() {
    log.Printf("[INFO] creating redmine client")
    redmineClient := redmine.NewClient(os.Getenv("REDMINE_URL"), os.Getenv("REDMINE_APIKEY"))

    c.redmineClient = redmineClient

}
