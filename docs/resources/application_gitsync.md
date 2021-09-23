---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_application_gitsync Resource - terraform-provider-harness"
subcategory: ""
description: |-
  Resource for configuring application git sync.
---

# harness_application_gitsync (Resource)

Resource for configuring application git sync.

## Example Usage

```terraform
data "harness_secret_manager" "default" {
  default = true
}

resource "harness_encrypted_text" "github_token" {
  name              = "github_token"
  value             = "<TOKEN>"
  secret_manager_id = data.harness_secret_manager.default.id
}

resource "harness_git_connector" "myrepo" {
  name                 = "myrepo"
  url                  = "https://github.com/someorg/myrepo"
  branch               = "main"
  generate_webhook_url = true
  username             = "someuser"
  password_secret_id   = harness_encrypted_text.github_token.id
  url_type             = "REPO"
}

resource "harness_application" "example" {
  name = "example-app"
}

resource "harness_application_gitsync" "example" {
  app_id       = harness_application.example.id
  connector_id = harness_git_connector.myrepo.id
  branch       = "main"
  enabled      = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **app_id** (String) The id of the application.
- **branch** (String) The branch of the git repository to sync to.
- **connector_id** (String) The id of the git connector to use.

### Optional

- **enabled** (Boolean) Whether or not to enable git sync.
- **id** (String) The ID of this resource.
- **repository_name** (String) The name of the git repository to sync to. This is only used if the git connector is for an account and not an individual repository.

## Import

Import is supported using the following syntax:

```shell
# Import using the Harness application id
terraform import harness_application_gitsync.myapp Xyz123
```