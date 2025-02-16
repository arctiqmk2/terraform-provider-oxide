---
page_title: "oxide_image Resource - terraform-provider-oxide"
---

# oxide_image (Resource)

This resource manages images.

## Example Usage

To create an image it's necessary to define its source by setting one of `source_url` or `source_snapshot_id`.

```hcl
resource "oxide_image" "example2" {
  project_id         = "c1dee930-a8e4-11ed-afa1-0242ac120002"
  description        = "a test image"
  name               = "myimage2"
  source_snapshot_id = "eb65d5cb-d8c5-4eae-bcf3-a0e89a633042"
  os                 = "ubuntu"
  version            = "20.04"
  timeouts = {
    read   = "1m"
    create = "3m"
  }
}
```

## Schema

### Required

- `description` (String) Description for the image.
- `os` (String) OS image distribution. Example: "alpine".
- `project_id` (String) ID of the project that will contain the image.
- `version` (String) OS image version. Example: "3.16".
- `name` (String) Name of the image.

### Optional

- `source_snapshot_id` (String) Snapshot ID of the image source.
- `timeouts` (Attribute, Optional) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `block_size` (Number) Size of blocks in bytes.
- `digest` (Object) Hash of the image contents, if applicable (see [below for nested schema](#nestedobject--digest)).
- `id` (String) Unique, immutable, system-controlled identifier of the image.
- `size` (Number) Total size in bytes.
- `time_created` (String) Timestamp of when this image was created.
- `time_modified` (String) Timestamp of when this image was last modified.

<a id="nestedatt--timeouts"></a>

### Nested Schema for `timeouts`

Optional:

- `create` (String, Default `10m`)
- `read` (String, Default `10m`)

### Nested Schema for `digest`

Read-Only:

- `type` (String) Digest type.
- `value` (String) Digest type value.
