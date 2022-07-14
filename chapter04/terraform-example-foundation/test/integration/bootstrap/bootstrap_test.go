// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bootstrap

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/utils"
	"github.com/tidwall/gjson"
	"github.com/stretchr/testify/assert"
)

// getResultFieldStrSlice parses a field of a results list into a string slice
func getResultFieldStrSlice(rs []gjson.Result, field string) []string {
	s := make([]string, 0)
	for _, r := range rs {
		s = append(s, r.Get(field).String())
	}
	return s
}

func TestBootstrap(t *testing.T) {

	bootstrap := tft.NewTFBlueprintTest(t,
		tft.WithTFDir("../../../0-bootstrap"),
	)

	cloudSourceRepos := []string{
		"gcp-org",
		"gcp-environments",
		"gcp-networks",
		"gcp-projects",
		"gcp-policies",
	}

	triggerRepos := []string{
		"gcp-org",
		"gcp-environments",
		"gcp-networks",
		"gcp-projects",
	}

	branchesRegex := `^(development|non\\-production|production)$`

	activateApis := []string{
		"serviceusage.googleapis.com",
		"servicenetworking.googleapis.com",
		"compute.googleapis.com",
		"logging.googleapis.com",
		"bigquery.googleapis.com",
		"cloudresourcemanager.googleapis.com",
		"cloudbilling.googleapis.com",
		"iam.googleapis.com",
		"admin.googleapis.com",
		"appengine.googleapis.com",
		"storage-api.googleapis.com",
		"monitoring.googleapis.com",
		"pubsub.googleapis.com",
		"securitycenter.googleapis.com",
		"accesscontextmanager.googleapis.com",
	}

	saOrgLevelRoles := []string{
		"roles/accesscontextmanager.policyAdmin",
		"roles/billing.user",
		"roles/compute.networkAdmin",
		"roles/compute.xpnAdmin",
		"roles/iam.securityAdmin",
		"roles/iam.serviceAccountAdmin",
		"roles/logging.configWriter",
		"roles/orgpolicy.policyAdmin",
		"roles/resourcemanager.folderAdmin",
		"roles/securitycenter.notificationConfigEditor",
		"roles/resourcemanager.organizationViewer",
	}

	orgID := utils.ValFromEnv(t, "TF_VAR_org_id")

	bootstrap.DefineVerify(
		func(assert *assert.Assertions) {

			// cloud build project
			cbProjectID := bootstrap.GetStringOutput("cloudbuild_project_id")
			bucketName := bootstrap.GetStringOutput("gcs_bucket_cloudbuild_artifacts")

			prj := gcloud.Runf(t, "projects describe %s", cbProjectID)
			assert.True(prj.Exists(), "project %s should exist", cbProjectID)

			gcAlphaOpts := gcloud.WithCommonArgs([]string{"--project", cbProjectID, "--json"})
			bkt := gcloud.Run(t, fmt.Sprintf("alpha storage ls --buckets gs://%s", bucketName), gcAlphaOpts).Array()[0]
			assert.True(bkt.Exists(), "bucket %s should exist", bucketName)

			for _, repo := range cloudSourceRepos {
				sourceRepoFullName := fmt.Sprintf("projects/%s/repos/%s", cbProjectID, repo)
				sourceRepo := gcloud.Runf(t, "source repos describe %s --project %s", repo, cbProjectID)
				assert.Equal(sourceRepoFullName, sourceRepo.Get("name").String(), fmt.Sprintf("repository %s should exist", repo))
			}

			for _, triggerRepo := range triggerRepos {
				for _, filter := range []string{
					fmt.Sprintf("trigger_template.branch_name='%s' AND  trigger_template.repo_name='%s' AND substitutions._TF_ACTION='apply'", branchesRegex, triggerRepo),
					fmt.Sprintf("trigger_template.branch_name='%s' AND  trigger_template.repo_name='%s' AND substitutions._TF_ACTION='plan' AND trigger_template.invert_regex=true", branchesRegex, triggerRepo),
				} {
					cbOpts := gcloud.WithCommonArgs([]string{"--project", cbProjectID, "--filter", filter, "--format", "json"})
					cbTriggers := gcloud.Run(t, "beta builds triggers list", cbOpts).Array()
					assert.Equal(1, len(cbTriggers), fmt.Sprintf("cloud builds trigger with filter %s should exist", filter))
				}
			}

			// seed project
			seedProjectID := bootstrap.GetStringOutput("seed_project_id")
			tfStateBucketName := bootstrap.GetStringOutput("gcs_bucket_tfstate")

			seedPrj := gcloud.Runf(t, "projects describe %s", seedProjectID)
			assert.True(seedPrj.Exists(), "project %s should exist", seedProjectID)

			enabledAPIS := gcloud.Runf(t, "services list --project %s", seedProjectID).Array()
			listApis := getResultFieldStrSlice(enabledAPIS, "config.name")
			assert.Subset(listApis, activateApis, "APIs should have been enabled")

			seedAlphaOpts := gcloud.WithCommonArgs([]string{"--project", seedProjectID, "--json"})
			tfStateBkt := gcloud.Run(t, fmt.Sprintf("alpha storage ls --buckets gs://%s", tfStateBucketName), seedAlphaOpts).Array()[0]
			assert.True(tfStateBkt.Exists(), "bucket %s should exist", tfStateBucketName)

			terraformSAEmail := bootstrap.GetStringOutput("terraform_service_account")
			terraformSAName := fmt.Sprintf("projects/%s/serviceAccounts/%s", seedProjectID, terraformSAEmail)
			terraformSA := gcloud.Runf(t, "iam service-accounts describe %s --project %s", terraformSAEmail, seedProjectID)
			assert.Equal(terraformSAName, terraformSA.Get("name").String(), fmt.Sprintf("service account %s should exist", terraformSAEmail))

			iamFilter := fmt.Sprintf("bindings.members:'serviceAccount:%s'", terraformSAEmail)
			iamOpts := gcloud.WithCommonArgs([]string{"--flatten", "bindings", "--filter", iamFilter, "--format", "json"})
			orgIamPolicyRoles := gcloud.Run(t, fmt.Sprintf("organizations get-iam-policy %s", orgID), iamOpts).Array()
			listRoles := getResultFieldStrSlice(orgIamPolicyRoles, "bindings.role")
			assert.Subset(listRoles, saOrgLevelRoles, fmt.Sprintf("service account %s should have organization level roles", terraformSAEmail))
		})
	bootstrap.Test()
}
