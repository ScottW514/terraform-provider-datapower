# Dependency Actions

Dependency actions allow you to define automated operations on other DataPower resources (targets) when performing create, update, or delete operations on the current resource. This is useful for managing dependencies, such as quiescing services before modifications or restarting domains after changes.

These actions are configured as a list of nested blocks within supported resources. Actions are executed in two phases:
- **Pre-process**: Before the main operation (quiesce actions during create/update/delete).
- **Post-process**: After the main operation (e.g., unquiesce, cache flushing or restart actions).

**Note**: Not all actions are supported for every target type. Refer to the specific resource documentation for supported actions and their effects. Validation ensures only valid actions are used.

## Schema

The `dependency_actions` attribute is an optional list of nested objects with the following structure:

- **target_id** (String, Optional): The ID of the target resource. Required for all target types except `datapower_domain`. Must be 1-128 characters and match the regex `^[a-zA-Z0-9_-]+$`.
- **target_domain** (String, Required): The application domain of the target resource. Must be 1-128 characters and match the regex `^[a-zA-Z0-9_-]+$`.
- **target_type** (String, Required): The type of the target resource (e.g., `datapower_domain`, `xml_management_interface`).
- **on_create** (Boolean, Optional, Default: `false`): Execute the action when creating the current resource.
- **on_update** (Boolean, Optional, Default: `true`): Execute the action when updating the current resource.
- **on_delete** (Boolean, Optional, Default: `false`): Execute the action when deleting the current resource.
- **action** (String, Required): The action to perform on the target (e.g., `quiesce`, `restart`). Supported actions vary by `target_type`; see the resource's documentation page for details.

## Validation

During configuration validation:
- `target_id` is required unless `target_type` is `datapower_domain`.
- The specified `action` must be supported for the `target_type`.
- Unsupported `target_type` values are rejected.

If validation fails, Terraform will report attribute errors under the `dependency_actions` path.

## Execution Flow

1. **Pre-Process Phase**: Triggered before the main create/update/delete operation. For `quiesce` actions, it checks the target's quiesce state and submits the action if needed (e.g., quiesce to prepare for changes). Duplicate actions are ignored.
2. **Main Operation**: The resource's core create/update/delete logic runs.
3. **Post-Process Phase**: Triggered after the main operation. Submits remaining actions (e.g., unquiesce or restart).

Actions like `quiesce` and `restart` include polling to ensure the target reaches the desired state (e.g., quiesced/unquiesced, or domain restarted). Timeouts are set to 90 seconds, with retries every 500ms.

Errors during action submission or state checks are reported as diagnostics and halt the process.

## Examples

### Basic Quiesce on Update

Quiesce a service in the same domain before updating a related resource, and unquiesce afterward.

```hcl
resource "datapower_example_resource" "example" {
  # ... other attributes ...

  dependency_actions = [
    {
      target_id     = "my-service-id"
      target_domain = "my-domain"
      target_type   = "multi_protocol_gateway"  # Example target_type; check resource docs for support
      on_create     = false
      on_update     = true
      on_delete     = false
      action        = "quiesce"
    }
  ]
}
```

In this example:
- On update, the service is quiesced pre-process and unquiesced post-process.
- No action on create or delete.

### Restart Domain on Create and Delete

Restart the entire domain after creating or deleting the resource.

```hcl
resource "datapower_example_resource" "example" {
  # ... other attributes ...

  dependency_actions = [
    {
      target_domain = "my-domain"
      target_type   = "datapower_domain"
      on_create     = true
      on_update     = false
      on_delete     = true
      action        = "restart"
    }
  ]
}
```

In this example:
- `target_id` is omitted since the target is a domain.
- The domain restarts post-process on create and delete.

### Multiple Actions on Different Targets

Combine actions across targets, such as quiescing a service and restarting a domain on update.

```hcl
resource "datapower_example_resource" "example" {
  # ... other attributes ...

  dependency_actions = [
    {
      target_id     = "my-service-id"
      target_domain = "my-domain"
      target_type   = "multi_protocol_gateway"
      on_create     = false
      on_update     = true
      on_delete     = false
      action        = "quiesce"
    },
    {
      target_domain = "my-domain"
      target_type   = "datapower_domain"
      on_create     = false
      on_update     = true
      on_delete     = false
      action        = "restart"
    }
  ]
}
```

In this example:
- On update, the service is quiesced/unquiesced, and the domain is restarted post-process.

## Troubleshooting

- **Action Not Supported**: Ensure the `action` is valid for the `target_type` (e.g., `restart` may only work on domains).
- **Timeouts**: If actions take longer than 90 seconds, adjust your DataPower configuration or check for underlying issues.
- **State Checks Fail**: For quiesce/restart, verify the target's status via DataPower APIs manually.
- **Domain Save Errors**: Authentication issues (e.g., 401) are ignored; other errors are logged.

For resource-specific supported actions and target types, refer to the individual resource documentation.
