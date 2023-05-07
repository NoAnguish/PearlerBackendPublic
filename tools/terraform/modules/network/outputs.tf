output "net_id" {
  value = resource.yandex_vpc_network.pearler_network.id
}

output "subnet_ids" {
  value = {
    zone_a = resource.yandex_vpc_subnet.pearler_subnet_zone_a.id,
    zone_b = resource.yandex_vpc_subnet.pearler_subnet_zone_b.id,
    zone_c = resource.yandex_vpc_subnet.pearler_subnet_zone_c.id,
  }
}