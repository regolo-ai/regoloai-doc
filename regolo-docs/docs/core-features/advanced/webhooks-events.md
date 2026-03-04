# Webhooks & Events

Integrate with webhooks and real-time event streams.

## Event Types

Regolo sends events for key system activities:

### Inference Events

- `inference.started` - Inference request initiated
- `inference.completed` - Inference request completed
- `inference.failed` - Inference request failed
- `inference.timeout` - Request timeout

### Batch Events

- `batch.submitted` - Batch job submitted
- `batch.processing` - Batch processing started
- `batch.completed` - Batch processing completed
- `batch.failed` - Batch processing failed

### Account Events

- `account.usage_warning` - Usage quota approaching
- `account.quota_exceeded` - Usage quota exceeded
- `billing.invoice_created` - Invoice generated

## Setting Up Webhooks

### 1. Create Webhook Endpoint

Your endpoint must:
- Accept POST requests
- Return 200-299 status code within 30 seconds
- Process events idempotently

```python
from flask import Flask, request

app = Flask(__name__)

@app.route('/webhooks/regolo', methods=['POST'])
def handle_webhook():
    event = request.json
    
    # Verify signature
    verify_webhook_signature(request)
    
    # Process event
    if event['type'] == 'inference.completed':
        handle_inference_complete(event['data'])
    
    return {'success': True}, 200
```

### 2. Register Webhook

=== "API"

    ```bash
    curl -X POST https://api.regolo.ai/v1/webhooks \
        -H "Authorization: Bearer YOUR_API_KEY" \
        -H "Content-Type: application/json" \
        -d '{
            "url": "https://your-app.com/webhooks/regolo",
            "events": ["inference.completed", "batch.completed"],
            "active": true
        }'
    ```

=== "Python"

    ```python
    import regolo
    
    webhook = regolo.create_webhook(
        url="https://your-app.com/webhooks/regolo",
        events=["inference.completed", "batch.completed"],
        active=True
    )
    ```

## Event Structure

```json
{
  "id": "evt_abc123",
  "type": "inference.completed",
  "timestamp": "2024-12-18T10:30:00Z",
  "data": {
    "request_id": "req_xyz789",
    "model": "Llama-3.3-70B-Instruct",
    "status": "success",
    "tokens_used": 150,
    "latency_ms": 1250
  },
  "request": {
    "id": "req_xyz789",
    "webhook_id": "wh_12345"
  }
}
```

## Webhook Security

### Signature Verification

All webhooks are signed with HMAC-SHA256:

```python
import hmac
import hashlib

def verify_webhook_signature(request, webhook_secret):
    signature = request.headers.get('X-Regolo-Signature')
    timestamp = request.headers.get('X-Regolo-Timestamp')
    body = request.get_data()
    
    # Check timestamp (prevent replay attacks)
    if abs(time.time() - int(timestamp)) > 300:
        return False
    
    # Verify signature
    signed_content = f"{timestamp}.{body.decode()}"
    expected_signature = hmac.new(
        webhook_secret.encode(),
        signed_content.encode(),
        hashlib.sha256
    ).hexdigest()
    
    return hmac.compare_digest(signature, expected_signature)
```

## Retry Policy

Failed webhook deliveries are retried:

| Attempt | Delay | Status |
|---------|-------|--------|
| 1st | Immediate | Sent |
| 2nd | 5 seconds | Retrying |
| 3rd | 30 seconds | Retrying |
| 4th | 2 minutes | Final attempt |

After 4 failed attempts, the webhook is marked as failed.

## Monitoring Webhooks

```python
# Get webhook details
webhook = regolo.get_webhook(webhook_id)
print(f"Status: {webhook.status}")
print(f"Last delivery: {webhook.last_delivery_at}")
print(f"Failed attempts: {webhook.failed_attempts}")

# List webhook events
events = regolo.list_webhook_events(
    webhook_id=webhook_id,
    limit=100
)

for event in events:
    print(f"{event.type} - {event.status} - {event.delivered_at}")
```

## Best Practices

1. **Verify signatures**: Always verify webhook signatures
2. **Process asynchronously**: Queue events for async processing
3. **Idempotent handling**: Handle duplicate events safely
4. **Fast responses**: Return 200 quickly, process in background
5. **Monitor delivery**: Track failed webhook deliveries
6. **Secure endpoints**: Use HTTPS and authentication