resource "yandex_vpc_network" "pearler_network" {
    name = "PearlerNetwork"
}

resource "yandex_vpc_subnet" "pearler_subnet_zone_a" {
  name           = "PearlerNetworkZoneA"
  zone           = "ru-central1-a"
  network_id     = yandex_vpc_network.pearler_network.id
  v4_cidr_blocks = ["10.10.0.0/24"]
}
resource "yandex_vpc_subnet" "pearler_subnet_zone_b" {
  name           = "PearlerNetworkZoneB"
  zone           = "ru-central1-b"
  network_id     = yandex_vpc_network.pearler_network.id
  v4_cidr_blocks = ["10.11.0.0/24"]
}
resource "yandex_vpc_subnet" "pearler_subnet_zone_c" {
  name           = "PearlerNetworkZoneC"
  zone           = "ru-central1-c"
  network_id     = yandex_vpc_network.pearler_network.id
  v4_cidr_blocks = ["10.12.0.0/24"]
}