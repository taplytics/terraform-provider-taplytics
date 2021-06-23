---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "taplytics Provider"
subcategory: ""
description: |-
  
---

# taplytics Provider

The Taplytics provider is used to interact with the Taplytics UAPI via Terraform. The following features are supported:

1. Feature Flags
1. Bucketing
1. Variations
1. Variables

-> Visit the [Taplytics UAPI Docs](https://universal-docs.taplytics.com/) for more information on how to use these
features.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "taplytics" {
  api_token = "<SDK API TOKEN HERE>"
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- **api_token** (String)