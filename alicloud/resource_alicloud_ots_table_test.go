package alicloud

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func TestAccAlicloudOtsTable_Basic(t *testing.T) {
	var table tablestore.DescribeTableResponse
	fmt.Printf("Starting...")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {

			testAccPreCheck(t)
		},

		// module name
		IDRefreshName: "alicloud_ots_table.basic",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckOtsTableDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccOtsTable,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOtsTableExist(
						"alicloud_ots_table.basic", &table),
				),
			},
		},
	})

}

func testAccCheckOtsTableExist(n string, table *tablestore.DescribeTableResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Table ID is set")
		}

		client := testAccProvider.Meta().(*AliyunClient)
		conn := client.otsconn

		response, _ := conn.DescribeTable(&tablestore.DescribeTableRequest{
			TableName: rs.Primary.ID,
		})

		log.Printf("[WARN] Ots table name is: %#v", rs.Primary.ID)

		if response != nil && response.TableMeta != nil {
			table = response
			return nil
		}
		return fmt.Errorf("Error finding OTS table %#v", rs.Primary.ID)
	}
}

func testAccCheckOtsTableDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "alicloud_ots_table" {
			continue
		}

		client := testAccProvider.Meta().(*AliyunClient)
		conn := client.otsconn

		response, _ := conn.DescribeTable(&tablestore.DescribeTableRequest{
			TableName: rs.Primary.ID,
		})

		if response != nil && response.TableMeta != nil{
			return fmt.Errorf("Error! Ots table still exists")
		}
	}

	return nil
}

const testAccOtsTable = `
resource "alicloud_ots_table" "basic" {
  table_name = "ots_table_c"
  primary_key {
    name = "pk1"
    type = "INTEGER"
  }
  table_option {
    time_to_live = -1
    max_version = 3
  }
}
`
