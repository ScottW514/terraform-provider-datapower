---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_amqpsourceprotocolhandler Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  AMQP handler
  CLI Alias: source-amqp
---

# datapower_amqpsourceprotocolhandler (Resource)

AMQP handler
  - CLI Alias: `source-amqp`

## Example Usage

```terraform
resource "datapower_amqpsourceprotocolhandler" "test" {
  id         = "AMQPSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  broker     = "TestAccAMQPBroker"
  from       = "amqpfrom"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `broker` (String) AMQP broker
  - CLI Alias: `broker`
  - Reference to: `datapower_amqpbroker:id`
- `from` (String) Source terminus
  - CLI Alias: `from`
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `credit` (Number) Credit
  - CLI Alias: `credit`
  - Range: `1`-`3600`
  - Default value: `100`
- `ignore_reply_to` (Boolean) Ignore reply-to
  - CLI Alias: `ignore-reply-to`
  - Default value: `true`
- `to` (String) Target terminus
  - CLI Alias: `to`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
