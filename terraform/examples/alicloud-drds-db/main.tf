provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_drds_db" "db" {
  provider = "alicloud"
  drds_instance_id = "${var.drds_instance_id}"
  db_name = "${var.db_name}"
  encode = "${var.encode}"
  password = "${var.password}"
  rds_instances = "${var.rds_instances}"
}