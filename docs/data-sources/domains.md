---
page_title: "domeneshop_domains Data Source - terraform-provider-domeneshop"
subcategory: ""
description: |-
  Use this data source to retrieve a list of domains.
---

# Data Source `domeneshop_domains`

Use this data source to retrieve a list of domains.

## Example Usage

```terraform
data "domeneshop_domains" "example" {
  domain = ".no"
}
```

## Schema

### Optional

- **domain** (String) Only return domains whose `domain` field includes this string.
- **id** (String) The ID of this resource.

### Read-only

- **domains** (List of Object) List of domains. (see [below for nested schema](#nestedatt--domains))

<a id="nestedatt--domains"></a>
### Nested Schema for `domains`

Read-only:

- **domain** (String)
- **expiry_date** (String)
- **id** (Number)
- **nameservers** (List of String)
- **registered_date** (String)
- **registrant** (String)
- **renew** (Boolean)
- **services** (List of Object) (see [below for nested schema](#nestedobjatt--domains--services))
- **status** (String)

<a id="nestedobjatt--domains--services"></a>
### Nested Schema for `domains.services`

Read-only:

- **dns** (Boolean)
- **email** (Boolean)
- **registrar** (Boolean)
- **webhotel** (String)


