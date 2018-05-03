package alicloud

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/drds"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"time"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func resourceAliCloudDRDSDb() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudDRDSDbCreate,
		Read:   resourceAliCloudDRDSDbRead,
		Update: resourceAliCloudDRDSDbFullTableUpdate,
		Delete: resourceAliCloudDRDSDbDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"drds_instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc:validateStringLengthInRange(0,25),
			},
			"encode": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc:validateAllowedStringValue([]string{string(UTF8Encode), string(GBKEncode),
				string(Latin1Encode), string(Utf8mb4Encode)}),
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc:validateStringLengthInRange(7,31),
			},
			"rds_instances": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"full_table_scan": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudDRDSDbCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient).drdsconn

	req := drds.CreateCreateDrdsDBRequest()
	instanceId := d.Get("drds_instance_id").(string)
	req.DrdsInstanceId = instanceId
	req.DbName = d.Get("db_name").(string)
	req.Encode = d.Get("encode").(string)
	req.Password = d.Get("password").(string)
	req.RdsInstances = d.Get("rds_instances").(string)

	_, err := client.CreateDrdsDB(req)

	if err != nil {
		return fmt.Errorf("failed to create DRDS instance with error: %s", err)
	}

	d.SetId(instanceId)

	return resourceAliCloudDRDSDbFullTableUpdate(d, meta)
}

func resourceAliCloudDRDSDbRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient).drdsconn

	req := drds.CreateDescribeDrdsDBRequest()
	req.DrdsInstanceId = d.Id()
	req.DbName = d.Get("db_name").(string)

	res, err := client.DescribeDrdsDB(req)
	data := res.Data

	if err != nil || res == nil  {
		return fmt.Errorf("failed to describe DRDS database with error: %s", err)
	}
	fmt.Print(data)

	// `description` isn't returned somehow, reported a bug https://connect.aliyun.com/suggestion/39734.
	//d.Set("description", data.Description)

	// As describe only return `type` 0 or 1, convert `type`. https://help.aliyun.com/document_detail/51126.html
	//d.Set("type", convertTypeValue(data.Type, d.Get("type").(string)))
	//d.Set("zone_id", data.ZoneId)

	return nil
}

func resourceAliCloudDRDSDbFullTableUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient).drdsconn
	update := false

	req := drds.CreateModifyFullTableScanRequest()
	req.DrdsInstanceId = d.Id()

	if d.HasChange("full_table_scan") && !d.IsNewResource() {
		update = true
		req.FullTableScan = d.Get("full_table_scan").(requests.Boolean)
	}

	if update {
		_, err := client.ModifyFullTableScan(req)

		if err != nil {
			return fmt.Errorf("failed to update Drds instance with error: %s", err)
		}
	}

	return resourceAliCloudDRDSDbRead(d, meta)
}

func resourceAliCloudDRDSDbDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AliyunClient).drdsconn
	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		req := drds.CreateDescribeDrdsDBRequest()
		instanceId := d.Id()
		dbName :=  d.Get("db_name").(string)
		req.DrdsInstanceId = instanceId
		req.DbName = dbName

		res, err := client.DescribeDrdsDB(req)

		if err != nil {
			if NotFoundError(err) {
				return nil
			}
		}

		//if res == nil || res.Data.DrdsInstanceId == "" {
		//	return nil
		//}
		if res == nil {
			return nil
		}

		removeReq := drds.CreateDeleteDrdsDBRequest()
		removeReq.DrdsInstanceId = instanceId
		removeReq.DbName = dbName

		removeRes, removeErr := client.DeleteDrdsDB(removeReq)
		if removeErr != nil || (removeRes != nil && !removeRes.Success) {
			return resource.RetryableError(fmt.Errorf("failed to delete instance timeout "+
				"and got an error: %#v", err))
		}

		return nil
	})
}
