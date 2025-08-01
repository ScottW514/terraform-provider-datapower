---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_kafkasourceprotocolhandler Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Kafka handler
  CLI Alias: source-kafka
---

# datapower_kafkasourceprotocolhandler (Resource)

Kafka handler
  - CLI Alias: `source-kafka`

## Example Usage

```terraform
resource "datapower_kafkasourceprotocolhandler" "test" {
  id             = "KafkaSourceProtocolHandler_name"
  app_domain     = "acc_test_domain"
  cluster        = "TestAccKafkaCluster"
  request_topic  = "topic"
  consumer_group = "consumer"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `cluster` (String) Kafka cluster
  - CLI Alias: `cluster`
  - Reference to: `datapower_kafkacluster:id`
- `consumer_group` (String) Consumer group
  - CLI Alias: `consumer-group`
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `request_topic` (String) Request topic
  - CLI Alias: `request-topic`

### Optional

- `batch_size` (Number) Batch size
  - CLI Alias: `batch-size`
  - Range: `1`-`65535`
  - Default value: `1`
- `response_topic` (String) Response topic
  - CLI Alias: `response-topic`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
