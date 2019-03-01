package redmine

import (
	redmine "github.com/mattn/go-redmine"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

// resourceProject: used to map a Redmine project to a terraform schema
func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// resourceProjectCreate creates a new redmine project using the redmine REST api
func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
        config := m.(*Config)
        name := d.Get("name").(string)
        description := d.Get("description").(string)

        i := redmine.Project{
                        Name: projectId,
                        Description: description,
        }
        project, err := config.redmineClient.CreateProject(i)
        if err != nil {
                return errors.Wrap(err, "creating redmine project failed")
        }
        if project != nil {
                return nil
        }

        return nil
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
       return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
       return nil
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
       return nil
}
