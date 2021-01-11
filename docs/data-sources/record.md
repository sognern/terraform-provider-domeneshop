---
page_title: "domeneshop_record Data Source - terraform-provider-domeneshop"
subcategory: ""
description: |-
  Use this data source to retrieve information about a DNS record.
---

# Data Source `domeneshop_record`

Use this data source to retrieve information about a DNS record.

## Example Usage

```terraform
data "domeneshop_record" "example" {
  domain_id = 1234
  record_id = 56789
}
```

## Schema

### Required

- **domain_id** (Number) ID of the domain.
- **record_id** (Number) ID of DNS the record.

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **alg** (Number) DS record algorithm.
- **data** (String) The value of the record.
- **digest** (Number) DS record digest type.
- **dtype** (Number) TLSA record matching type.
- **flags** (Number) CAA record flags.
- **host** (String) The host/subdomain the DNS record applies to.
- **port** (Number) SRV record port. The port where the service is found.
- **priority** (Number) MX/SRV record priority, also known as preference. Lower values are usually preferred first, but this is not guaranteed.
- **selector** (Number) TLSA record selector.
- **tag** (Number) CAA/DS record tag.
- **ttl** (Number) TTL of DNS record in seconds.
- **type** (String) The type of the record.
- **usage** (Number) TLSA record certificate usage.
- **weight** (Number) SRV record weight. Relevant if multiple records have same preference.


