# terraform-provider-redmine

## progression

|CRUD               |Implements|
|-------------------|---------:|
|Create             |       30%|
|Read               |        0%|
|Update             |        0%|
|Delete             |        0%|


#Use

go build
cp terraform-provider-redmine $HOME/.terraform.d/plugins/terraform-provider-redmine
terraform init

write terraform file:

`resource "redmine_project" "example" {`
  `name = "test1"`
  `description = "test1 created with terraform"`
  `identifier = "idtest"`
`}`

terraform plan
terraform apply
