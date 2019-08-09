module github.com/inuits/terraform-provider-redmine

go 1.12

require (
	github.com/hashicorp/terraform v0.12.6
	github.com/mattn/go-redmine v0.0.0-20181021123913-2267b9239bac
	github.com/pkg/errors v0.8.1
)

replace github.com/mattn/go-redmine => github.com/inuits/go-redmine v0.0.0-20181021123913-2267b9239bac
