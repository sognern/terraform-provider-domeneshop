---
page_title: "domeneshop_records Data Source - terraform-provider-domeneshop"
subcategory: ""
description: |-
  Use this data source to retrieve a list of DNS records.
---

# Data Source `domeneshop_records`

Use this data source to retrieve a list of DNS records.

## Example Usage

```terraform
data "domeneshop_records" "example" {
  domain_id = 1234
}
```

## Schema

### Optional

- **domain_id** (Number) Only return domains whose `domain` field includes this string.
- **host** (String) Only return records whose `host` field matches this string.
- **id** (String) The ID of this resource.
- **type** (String) Only return records whose `type` field matches this string.

### Read-only

- **records** (List of Object) List of records. (see [below for nested schema](#nestedatt--records))

<a id="nestedatt--records"></a>
### Nested Schema for `records`

Read-only:

- **alg** (Number)
- **data** (String)
- **digest** (Number)
- **dtype** (Number)
- **flags** (Number)
- **host** (String)
- **id** (Number)
- **port** (Number)
- **priority** (Number)
- **selector** (Number)
- **tag** (Number)
- **ttl** (Number)
- **type** (String)
- **usage** (Number)
- **weight** (Number)


