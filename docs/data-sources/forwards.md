---
page_title: "domeneshop_forwards Data Source - terraform-provider-domeneshop"
subcategory: ""
description: |-
  Use this data source to retrieve a list of HTTP forwards.
---

# Data Source `domeneshop_forwards`

Use this data source to retrieve a list of HTTP forwards.

## Example Usage

```terraform
data "domeneshop_forwards" "example" {
  domain_id = 1234
}
```

## Schema

### Required

- **domain_id** (Number) ID of the domain.

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **forwards** (List of Object) List of HTTP forwards. (see [below for nested schema](#nestedatt--forwards))

<a id="nestedatt--forwards"></a>
### Nested Schema for `forwards`

Read-only:

- **frame** (Boolean)
- **host** (String)
- **url** (String)


