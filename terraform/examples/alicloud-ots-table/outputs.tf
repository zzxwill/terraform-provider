output "table_name" {
  value = "${alicloud_ots_table.table.table_name}"
}

output "time_to_live" {
  value = "${alicloud_ots_table.table.time_to_live}"
}

output "max_version" {
  value = "${alicloud_ots_table.table.max_version}"
}