---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "taplytics_bucketing Data Source - terraform-provider-taplytics"
subcategory: ""
description: |-
  
---

# taplytics_bucketing (Data Source)

-> Visit the [Taplytics UAPI Docs](https://universal-docs.taplytics.com/) for more information on how to use Bucketing.

The bucketing data source allows you to retrieve all experiements and variations that apply to a given user.

## Example Usage

```terraform
data "taplytics_bucketing" "all" {
  userid = "1235abcada"
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- **userid** (String)

### Read-Only

- **bucketing** (List of Object) (see [below for nested schema](#nestedatt--bucketing))

<a id="nestedatt--bucketing"></a>

### Nested Schema for `bucketing`

Read-Only:

- **experiment** (String)
- **variation** (String)


