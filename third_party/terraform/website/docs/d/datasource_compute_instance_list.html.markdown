---
layout: "google"
page_title: "Google: google_compute_instance_list"
sidebar_current: "docs-google-datasource-compute-instance-list-x"
description: |-
  Get a list of VM instances within a GCE zone.
---

# google\_compute\_instance\_list

Get information about VM instances within a GCE zone. For more information see
[the official documentation](https://cloud.google.com/compute/docs/instances)
and
[API](https://cloud.google.com/compute/docs/reference/latest/instances).


## Example Usage

```hcl
data "google_compute_instance_list" "instances" {
	zone = "us-central1-a"
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) A filter string to use when getting the instance list. See the
  [API](https://cloud.google.com/compute/docs/reference/latest/instances) for more information
  about how to filter for certain values.

* `project` - (Optional) The ID of the project in which the resources belong.

* `zone` - (Required) The zone of the instances.

## Attributes Reference

* `self_link` - The URI of the created resource.

* `instances` - A list of instances created in the given zone inside the project.
  Structure is documented below.

---

The `instances` block supports:

* `name` - The name of the instance.

* `project` - The ID of the project in which the resource belongs.

* `zone` - The zone of the instance.

* `self_link` - The self link of the instance.
