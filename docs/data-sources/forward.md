---
page_title: "domeneshop_forward Data Source - terraform-provider-domeneshop"
subcategory: ""
description: |-
  Use this data source to retrieve information about a HTTP forwarding.
---

# Data Source `domeneshop_forward`

Use this data source to retrieve information about a HTTP forwarding.

## Example Usage

```terraform
data "domeneshop_forward" "example" {
  domain_id = 1234
  host      = "www"
}
```

## Schema

### Required

- **domain_id** (Number) ID of the domain.
- **host** (String) The subdomain this forward applies to, without the domain part.

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **frame** (Boolean) Whether to enable frame forwarding using an iframe embed. NOT recommended for a variety of reasons.
- **url** (String) The URL to forward to. Must include scheme, e.g. `https://` or `ftp://`.


