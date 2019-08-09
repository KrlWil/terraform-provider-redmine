package redmine

import (
	"github.com/hashicorp/terraform/helper/schema"
	redmine "github.com/mattn/go-redmine"
	"github.com/pkg/errors"
	"strconv"
)

// resourceProject used to map a Redmine project to a terraform schema
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
			"parent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"created_on": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_on": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// resourceProjectCreate creates a new redmine project using the redmine REST api
func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	name := d.Get("name").(string)
	identifier := d.Get("identifier").(string)
	parent := d.Get("parent").(int)
	description := d.Get("description").(string)
	createdOn := d.Get("created_on").(string)
	updatedOn := d.Get("updated_on").(string)

	var parentStruct *redmine.Project
	var err error
	if parent != 0 {
		parentStruct, err = config.redmineClient.Project(parent)
		if err != nil {
			return errors.Wrap(err, "parent id not found")
		}
	}

	i := redmine.Project{
		Name:        name,
		Identifier:  identifier,
		Parent:      parentStruct,
		Description: description,
		CreatedOn:   createdOn,
		UpdatedOn:   updatedOn,
	}

	if parent != 0 {
		i.ParentID = parent
	}

	project, err := config.redmineClient.CreateProject(i)
	if err != nil {
		return errors.Wrap(err, "creating redmine project failed")
	}

	// convert project.ID to string, setId() takes string
	s1 := strconv.Itoa(project.ID)

	d.SetId(s1)

	return resourceProjectRead(d, m)
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)

	// convert string to int
	i1, err := strconv.Atoi(d.Id())

	project, err := config.redmineClient.Project(i1)
	if err != nil {
		return errors.Wrap(err, "fetching redmine project failed")
	}

	d.Set("name", project.Name)
	d.Set("identifier", project.Identifier)

	if project.Parent != nil {
		d.Set("parent", project.Parent.ID)
	}
	if project.Description != "" {
		d.Set("description", project.Description)
	}
	if project.CreatedOn != "" {
		d.Set("created_on", project.CreatedOn)
	}
	if project.UpdatedOn != "" {
		d.Set("updated_on", project.UpdatedOn)
	}

	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	name := d.Get("name").(string)
	identifier := d.Get("identifier").(string)
	parent := d.Get("parent").(int)
	description := d.Get("description").(string)
	createdOn := d.Get("created_on").(string)
	updatedOn := d.Get("updated_on").(string)

	// convert string to int
	i1, err1 := strconv.Atoi(d.Id())

	if err1 != nil {
		return errors.Wrap(err1, "converting string failed")
	}

	var parentStruct *redmine.Project
	var err error
	if parent != 0 {
		parentStruct, err = config.redmineClient.Project(parent)
		if err != nil {
			return errors.Wrap(err, "parent id not found")
		}
	}

	i := redmine.Project{
		ID:          i1,
		Name:        name,
		Identifier:  identifier,
		Parent:      parentStruct,
		Description: description,
		CreatedOn:   createdOn,
		UpdatedOn:   updatedOn,
	}

	if parent != 0 {
		i.ParentID = parent
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
