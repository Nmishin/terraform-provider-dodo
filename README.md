Terraform Provider Dodo Pizza
=============================

This repository contains a sample Terraform provider. Can be used for learning purpose.
This is fully working solution. Used [dodo-go](https://github.com/Nmishin/dodo-go) client library for iteract with [Dodo Global API v2](https://globalapi.dodopizza.com/api/index.html?urls.primaryName=Dodo%20Global%20API%20v2).

For more info about how to write custom terraform provider, please check out my article:
[My Path for Terraform Provider Creation](https://hackernoon.com/my-path-for-terraform-provider-creation)

This terraform provider available for both repositories: terraform and opentofu.

Using the Provider
------------------
`cat provider.tf`
```tcl
terraform {
  required_providers {
    dodo = {
      source = "Nmishin/dodo"
      version = "0.0.1"
    }
  }
}

provider "dodo" {}
```

`cat main.tf`
```tcl
data "dodo_brand" "test" {}
```

`cat outputs.tf`
```tcl
output "brands" {
  value = data.dodo_brand.test.names
}
```

```tcl
tofu plan

data.dodo_brand.test: Reading...
data.dodo_brand.test: Read complete after 1s [id=bf4b2096fab07279e4a6a9db5fb704b5]

Changes to Outputs:
  + brands = [
      + "dodopizza",
      + "doner42",
      + "drinkit",
    ]

You can apply this plan to save these new output values to the OpenTofu state, without changing any real infrastructure.
```
