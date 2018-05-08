output "db_name" {
  value = "${alicloud_drds_db.db.db_name}"
}

output "drds_instance" {
  value = "${alicloud_drds_db.db.drds_instance_id}"
}

output "rds_instances" {
  value = "${alicloud_drds_db.db.rds_instances}"
}