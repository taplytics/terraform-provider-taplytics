---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "taplytics_featureflag Data Source - terraform-provider-taplytics"
subcategory: ""
description: |-
  
---

# taplytics_featureflag (Data Source)

-> Visit the [Taplytics UAPI Docs](https://universal-docs.taplytics.com/) for more information on how to use Feature
Flags.

The feature flag data source allow you to check whether a specific feature flag is enabled for a given user

## Example Usage

```terraform
data "taplytics_featureflag" "specific" {
  userid = "1235abcada"
  featureflag_key = "specifickey"
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- **featureflag_key** (String)
- **userid** (String)

### Read-Only

- **enabled** (Boolean)

