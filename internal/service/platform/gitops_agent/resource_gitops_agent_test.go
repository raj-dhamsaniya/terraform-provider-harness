package gitops_agent_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/antihax/optional"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGitopsAgent(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	orgId := "gitopstest"
	projectId := "gitopsagenttest"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "harness_platform_gitops_agent.test"
	agentName := id
	namespace := "tf-test"
	updatedNamespace := namespace + "-updated"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceGitopsAgentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGitopsAgent(id, accountId, projectId, orgId, agentName, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", agentName),
				),
			},
			{
				Config: testAccResourceGitopsAgent(id, accountId, projectId, orgId, agentName, updatedNamespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "metadata.0.namespace", updatedNamespace),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_id", "type"},
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccGetAgent(resourceName string, state *terraform.State) (*nextgen.V1Agent, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	agentIdentifier := r.Primary.Attributes["identifier"]

	resp, _, err := c.AgentServiceApi.AgentServiceGet(ctx, agentIdentifier, &nextgen.AgentServiceApiAgentServiceGetOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     optional.NewString(r.Primary.Attributes["org_identifier"]),
		ProjectIdentifier: optional.NewString(r.Primary.Attributes["project_identifier"]),
	})

	if err != nil {
		return nil, err
	}

	if resp.Type_ == nil {
		return nil, nil
	}

	return &resp, nil
}

func testAccResourceGitopsAgentDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		agent, _ := testAccGetAgent(resourceName, state)
		if agent != nil {
			return fmt.Errorf("Found Agent: %s", agent.Identifier)
		}
		return nil
	}

}

func testAccResourceGitopsAgent(agentId string, accountId string, projectId string, orgId string, agentName string, namespace string) string {
	return fmt.Sprintf(`
		resource "harness_platform_gitops_agent" "test" {
			identifier = "%[1]s"
			account_id = "%[2]s"
			project_id = "%[3]s"
			org_id = "%[4]s"
			name = "%[5]s"
			type = "CONNECTED_ARGO_PROVIDER"
			metadata {
        namespace = "%[6]s"
        high_availability = true
    	}
		}
		`, agentId, accountId, projectId, orgId, agentName, namespace)
}
