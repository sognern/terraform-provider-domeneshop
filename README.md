![](https://domene.shop/svg/logo-no.svg)

# Terraform Provider Domeneshop

Available in the [Terraform Registry](https://registry.terraform.io/providers/innovationnorway/domeneshop/latest).

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.15

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

```hcl
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

data "domeneshop_domains" "example" {
  domain = "example.com"
}

resource "domeneshop_record" "example" {
  domain_id = data.domeneshop_domains.test.domains.0.id
  host      = "foo"
  type      = "A"
  data      = "192.0.2.56"
  ttl       = 300
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

The following environment variables must be set to run acceptance tests:
- `DOMENESHOP_TOKEN`
- `DOMENESHOP_SECRET`
- `DOMENESHOP_DOMAIN`
