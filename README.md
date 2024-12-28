Terraform Provider Dodo Pizza
=============================

This repository contains a sample Terraform provider. Can be used for learning purposes.
It is a fully working solution. Uses [dodo-go](https://github.com/Nmishin/dodo-go) client library to interact with the [Dodo Global API v2](https://globalapi.dodopizza.com/api/index.html?urls.primaryName=Dodo%20Global%20API%20v2).


__For more information about how to write a custom terraform provider, please check out my article:
<br>[My Path for Terraform Provider Creation](https://hackernoon.com/my-path-for-terraform-provider-creation)__

This terraform provider is available for both repositories: terraform and opentofu.

Using the Provider
------------------
```tcl
cat provider.tf

terraform {
  required_providers {
    dodo = {
      source = "Nmishin/dodo"
      version = "0.0.5"
    }
  }
}

provider "dodo" {}
```

```tcl
cat main.tf

data "dodo_brand" "test" {}
```

```tcl
cat outputs.tf

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

