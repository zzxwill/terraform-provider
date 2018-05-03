package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/drds"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAlicloudDRDSDb_Basic(t *testing.T) {
	var db drds.DescribeDrdsDBResponse
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		IDRefreshName: "alicloud_drds_db.basic",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckDRDSDbDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDrdsDb,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDRDSDbExist(
						"alicloud_drds_db.basic", &db),
				),
			},
		},
	})
}

func testAccCheckDRDSDbExist(n string, db *drds.DescribeDrdsDBResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no DRDS Db ID is set")
		}

		client := testAccProvider.Meta().(*AliyunClient).drdsconn
		req := drds.CreateDescribeDrdsDBRequest()
		req.DrdsInstanceId = rs.Primary.ID
		//req.DbName = rs.Get("db_name").(string)

		response, err := client.DescribeDrdsDB(req)

		//if err == nil && response != nil && response.Data.DrdsDbId != "" {
		//	db = response
		//	return nil
		//}

		if err == nil && response != nil {
			db = response
			return nil
		}

		return fmt.Errorf("error finding DRDS db %#v", rs.Primary.ID)
	}
}

func testAccCheckDRDSDbDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "alicloud_drds_db" {
			continue
		}

		client := testAccProvider.Meta().(*AliyunClient)
		conn := client.drdsconn
		req := drds.CreateDescribeDrdsDBRequest()
		req.DrdsInstanceId = rs.Primary.ID



		response, err := conn.DescribeDrdsDB(req)

		//if err == nil && response != nil && response.Data.Status != "5" {
		//	return fmt.Errorf("error! DRDS db still exists")
		//}
		if err == nil && response != nil {
			return fmt.Errorf("error! DRDS db still exists")
		}
	}

	return nil
}

const testAccDrdsDb = `
provider "alicloud" {
	region = "cn-hangzhou"
}
resource "alicloud_drds_db" "basic" {
  provider = "alicloud"
  drds_instance_id = "drdsxzru72io1j0f"
  db_name = "test"
  encode = "utf8"
  password = "Admin123"
  rds_instances = "[\"DBInstanceId\": \"rm-bp1865y7whv5u665k\"]"
}
`
