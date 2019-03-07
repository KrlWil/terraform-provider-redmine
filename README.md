# terraform-provider-redmine

## Progression

| CRUD   | Implements |
| ------ | ---------: |
| Create |        50% |
| Read   |        50% |
| Update |        50% |
| Delete |        50% |

# Install

- go build or Download `terraform-provider-redmine` from [Releases](https://github.com/KrlWil/terraform-provider-redmine/releases)

- copy to plugins directory

  `cp terraform-provider-redmine $HOME/.terraform.d/plugins/terraform-provider-redmine`

# Use

- Set environment variables

  ```hcl
  export REDMINE_URL=http://localhost:8009
  export REDMINE_APIKEY=123456789abcdef
  ```

- write terraform file


```hcl
resource "redmine_project" "example" {
  `name = "test1"`
  `description = "test1 created with terraform"`
  `identifier = "idtest"`
}
```

`terraform init`

`terraform plan`
`terraform apply`

