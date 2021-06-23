---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "taplytics_variation Data Source - terraform-provider-taplytics"
subcategory: ""
description: |-
  
---

# taplytics_variation (Data Source)

-> Visit the [Taplytics UAPI Docs](https://universal-docs.taplytics.com/) for more information on how to use Variations.

The feature flag data source allow you to check the variation of a given experiment for a specific user.

## Example Usage

```terraform
data "taplytics_variation" "specific" {
  userid = "1235abcada"
  experiment = "experiment name"
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- **experiment** (String)
- **userid** (String)

### Read-Only

- **variation** (String)

