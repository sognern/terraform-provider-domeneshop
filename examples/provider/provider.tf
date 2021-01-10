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
