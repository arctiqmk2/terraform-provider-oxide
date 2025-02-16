// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type dataSourceImageConfig struct {
	BlockName         string
	SupportBlockName  string
	SupportBlockName2 string
}

var dataSourceImageConfigTpl = `
data "oxide_project" "{{.SupportBlockName}}" {
	name = "tf-acc-test"
}

data "oxide_images" "{{.SupportBlockName2}}" {
  project_id = data.oxide_project.{{.SupportBlockName}}.id
}

data "oxide_image" "{{.BlockName}}" {
  project_name = "tf-acc-test"
  name = element(tolist(data.oxide_images.{{.SupportBlockName2}}.images[*].name), 0)
  timeouts = {
    read = "1m"
  }
}
`

// NB: The project must be populated with at least one image for this test to pass
func TestAccDataSourceImage_full(t *testing.T) {
	blockName := newBlockName("datasource-image")
	config, err := parsedAccConfig(
		dataSourceImageConfig{
			BlockName:         blockName,
			SupportBlockName:  newBlockName("support"),
			SupportBlockName2: newBlockName("support"),
		},
		dataSourceImageConfigTpl,
	)
	if err != nil {
		t.Errorf("error parsing config template data: %e", err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: checkDataSourceImage(
					fmt.Sprintf("data.oxide_image.%s", blockName),
				),
			},
		},
	})
}

func checkDataSourceImage(dataName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc([]resource.TestCheckFunc{
		resource.TestCheckResourceAttrSet(dataName, "id"),
		resource.TestCheckResourceAttr(dataName, "timeouts.read", "1m"),
		resource.TestCheckResourceAttrSet(dataName, "block_size"),
		resource.TestCheckResourceAttrSet(dataName, "description"),
		resource.TestCheckResourceAttrSet(dataName, "os"),
		resource.TestCheckResourceAttrSet(dataName, "project_id"),
		resource.TestCheckResourceAttrSet(dataName, "project_name"),
		resource.TestCheckResourceAttrSet(dataName, "name"),
		resource.TestCheckResourceAttrSet(dataName, "size"),
		resource.TestCheckResourceAttrSet(dataName, "time_created"),
		resource.TestCheckResourceAttrSet(dataName, "time_modified"),
		resource.TestCheckResourceAttrSet(dataName, "version"),
	}...)
}
