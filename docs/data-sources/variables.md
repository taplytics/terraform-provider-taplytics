---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "taplytics_variables Data Source - terraform-provider-taplytics"
subcategory: ""
description: |-
  
---

# taplytics_variables (Data Source)

-> Visit the [Taplytics UAPI Docs](https://universal-docs.taplytics.com/) for more information on how to use Variables.

The feature flag data source allow you to check the value of all variables for a given user

## Example Usage

```terraform
data "taplytics_variables" "all" {
  userid = "1235abcada"
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- **userid** (String)

### Read-Only

- **variables** (List of Object) (see [below for nested schema](#nestedatt--variables))

<a id="nestedatt--variables"></a>

### Nested Schema for `variables`

Read-Only:

- **name** (String)
- **value** (String)


