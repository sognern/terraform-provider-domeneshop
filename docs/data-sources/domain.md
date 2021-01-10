---
page_title: "domeneshop_domain Data Source - terraform-provider-domeneshop"
subcategory: ""
description: |-
  Use this data source to retrieve information about a domain.
---

# Data Source `domeneshop_domain`

Use this data source to retrieve information about a domain.

## Example Usage

```terraform
data "domeneshop_domain" "example" {
  domain_id = 1234
}
```

## Schema

### Required

- **domain_id** (Number) ID of the domain.

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **domain** (String) Name of the domain.
- **expiry_date** (String) Expiry date.
- **nameservers** (List of String) List of nameservers.
- **registered_date** (String) Registered date.
- **registrant** (String) Name of the registrant.
- **renew** (Boolean) Whether the domain should be renewed.
- **services** (List of Object) Domain services. (see [below for nested schema](#nestedatt--services))
- **status** (String) Domain status.

<a id="nestedatt--services"></a>
### Nested Schema for `services`

Read-only:

- **dns** (Boolean)
- **email** (Boolean)
- **registrar** (Boolean)
- **webhotel** (String)


