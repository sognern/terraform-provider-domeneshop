resource "domeneshop_record" "example" {
  domain_id = 1234
  host      = "example"
  type      = "A"
  data      = "192.0.2.56"
  ttl       = 300
}
