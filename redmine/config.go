package redmine

import (
	redmine "github.com/mattn/go-redmine"
	"log"
	"os"
)

// Config contains configuration and a reference to the redmine client struct.
type Config struct {
	redmineClient *redmine.Client
}

func (c *Config) createAndAuthenticateClient() {
	log.Printf("[INFO] creating redmine client")
	redmineClient := redmine.NewClient(os.Getenv("REDMINE_URL"), os.Getenv("REDMINE_APIKEY"))

	c.redmineClient = redmineClient

}
