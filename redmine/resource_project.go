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
        config :=m.(*Config)

        //convert string to int
        i1, err := strconv.Atoi(d.Id())

        project, err := config.redmineClient.Project(i1)
        if err != nil {
                return errors.Wrap(err, "creating redmine project failed")
        }

        d.Set("name", project.Name)
        d.Set("identifier", project.Identifier)

        if project.Description !="" {
                 d.Set("description", project.Description)
        }
        if project.CreatedOn !="" {
                 d.Set("created_on", project.CreatedOn)
        }
        if project.UpdatedOn !="" {
                 d.Set("updated_on", project.UpdatedOn)
        }

        return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
        config := m.(*Config)
        name := d.Get("name").(string)
        identifier := d.Get("identifier").(string)
        description := d.Get("description").(string)
        createdOn := d.Get("created_on").(string)
        updatedOn := d.Get("updated_on").(string)

        //convert string to int
        i1, err1 := strconv.Atoi(d.Id())

        if err1 != nil {
                return errors.Wrap(err1, "converting string failed")
        }

        i := redmine.Project{
                        Id: i1,
                        Name: name,
                        Identifier: identifier,
                        Description: description,
                        CreatedOn: createdOn,
                        UpdatedOn: updatedOn,
        }
        err2 := config.redmineClient.UpdateProject(i)

        if err2 != nil {
                return errors.Wrap(err2, "updating redmine project failed")
        }
        return nil
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
        config := m.(*Config)

        id, err1 := strconv.Atoi(d.Id())
        if err1 != nil {
                return errors.Wrap(err1, "converting string failed")
        }

        err2 := config.redmineClient.DeleteProject(id)
        if err2 != nil {
            return errors.Wrap(err2, "deleting redmine issue failed")
        }
        return nil
}
