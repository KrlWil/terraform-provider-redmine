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
            "identifier": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default:  "",
            },
            "created_on": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default:  "",
            },
            "updated_on": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default:  "",
            },
        },
    }
}

// resourceProjectCreate creates a new redmine project using the redmine REST api
func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
        config := m.(*Config)
        name := d.Get("name").(string)
        identifier := d.Get("identifier").(string)
        description := d.Get("description").(string)
        createdOn := d.Get("created_on").(string)
        updatedOn := d.Get("updated_on").(string)

        i := redmine.Project{
                        Name: name,
                        Identifier: identifier,
                        Description: description,
                        CreatedOn: createdOn,
                        UpdatedOn: updatedOn,
        }
        project, err := config.redmineClient.CreateProject(i)
        if err != nil {
                return errors.Wrap(err, "creating redmine project failed")
        }

        //convert project.Id to string, setId() takes string
        s1 := strconv.Itoa(project.Id)

        d.SetId(s1)

        return resourceProjectRead(d, m)
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
