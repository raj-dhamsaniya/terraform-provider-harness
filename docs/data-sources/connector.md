---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_connector Data Source - terraform-provider-harness"
subcategory: ""
description: |-
  Data source for retrieving a Harness connector This resource is part of the Harness nextgen platform.
---

# harness_connector (Data Source)

Data source for retrieving a Harness connector This resource is part of the Harness nextgen platform.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **branch** (String) The specified branch of the connector.
- **ccm_connector_filter** (Block List, Max: 1) The ccm connector filter. (see [below for nested schema](#nestedblock--ccm_connector_filter))
- **connectivity_statuses** (List of String) The connectivity status of the connector. Available options are SUCCESS, FAILURE, PARTIAL, UNKNOWN
- **first_result** (Boolean) When set to true if the query returns more than one result the first item will be selected. When set to false (default) this will return an error.
- **get_default_from_other_repo** (Boolean) Whether to get default from other repo.
- **get_distinct_from_branches** (Boolean) Whether to get distinct from branches.
- **id** (String) The ID of this resource.
- **identifier** (String) Unique identifier of the connector.
- **include_all_connectors_available_at_scope** (Boolean) Whether to include all connectors available at scope.
- **inheriting_credentials_from_delegate** (Boolean) Whether the connector inherits credentials from the delegate.
- **org_id** (String) Unique identifier of the organization.
- **project_id** (String) Unique identifier of the project.
- **repo_id** (String) Unique identifier of the repository.
- **search_term** (String) The search term used to find the connector.
- **tags** (List of String) The tags of the connector.

### Read-Only

- **description** (String) The description of the connector.
- **name** (String) The name of the connector.
- **type** (String) The type of the selected connector.
- **types** (List of String) The type of the connector. Available values are K8sCluster, Git, Splunk, AppDynamics, Prometheus, Dynatrace, Vault, AzureKeyVault, DockerRegistry, Local, AwsKms, GcpKms, AwsSecretManager, Gcp, Aws, Artifactory, Jira, Nexus, Github, Gitlab, Bitbucket, Codecommit, CEAws, CEAzure, GcpCloudCost, CEK8sCluster, HttpHelmRepo, NewRelic, Datadog, SumoLogic, PagerDuty

<a id="nestedblock--ccm_connector_filter"></a>
### Nested Schema for `ccm_connector_filter`

Optional:

- **aws_account_id** (String) The AWS account identifier.
- **azure_subscription_id** (String) The Azure subscription identifier.
- **azure_tenant_id** (String) The Azure tenant identifier.
- **features_enabled** (List of String) The CCM features that are enabled. Valid options are BILLING, OPTIMIZATION, VISIBILITY.
- **gcp_project_id** (String) The GCP project identifier.
- **k8s_connector_ref** (String) The Kubernetes connector reference.

