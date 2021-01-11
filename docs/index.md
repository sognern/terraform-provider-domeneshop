---
page_title: "Provider: Domeneshop"
subcategory: ""
description: |-
  Terraform provider for Domeneshop (Domainnameshop).
---

# Domeneshop Provider

Terraform provider for Domeneshop (Domainnameshop).

## Example Usage

```terraform
variable "domeneshop_token" {
  type      = string
  sensitive = true
}

variable "domeneshop_secret" {
  type      = string
  sensitive = true
}

provider "domeneshop" {
  token  = var.domeneshop_token
  secret = var.domeneshop_secret
}
```

## Schema

### Optional

- **secret** (String, Sensitive) A Domeneshop API secret. This can also be set with the `DOMENESHOP_SECRET` environment variable.
- **token** (String, Sensitive) A Domeneshop API token. This can also be set with the `DOMENESHOP_TOKEN` environment variable.