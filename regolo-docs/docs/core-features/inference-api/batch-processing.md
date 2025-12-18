# Batch Processing

Process large batches of requests efficiently with lower latency requirements.

## Overview

Batch processing is ideal for:

- Processing large datasets asynchronously
- Reducing per-request latency requirements
- Cost optimization through bulk processing
- Off-peak processing for non-time-critical workloads

## When to Use Batch Processing

### Ideal Use Cases

✅ Processing logs for analysis  
✅ Bulk content classification  
✅ Dataset augmentation  
✅ Daily report generation  
✅ Offline analysis of large documents  

### Not Recommended

❌ Real-time customer-facing responses  
❌ Interactive applications  
❌ Time-sensitive operations  

## Submitting a Batch Job

### Step 1: Prepare Requests

Format your requests as JSONL (JSON Lines):

```json
{"messages": [{"role": "user", "content": "What is AI?"}], "model": "Llama-3.3-70B-Instruct"}
{"messages": [{"role": "user", "content": "Define machine learning"}], "model": "Llama-3.3-70B-Instruct"}
{"messages": [{"role": "user", "content": "Explain neural networks"}], "model": "Llama-3.3-70B-Instruct"}
```

### Step 2: Create Batch Job

=== "Python"

    ```python
    import regolo
    
    # Read your JSONL file
    with open('requests.jsonl', 'r') as f:
        batch_data = f.read()
    
    # Submit batch job
    batch = regolo.submit_batch(
        requests=batch_data,
        model="Llama-3.3-70B-Instruct"
    )
    
    print(f"Batch ID: {batch.id}")
    print(f"Status: {batch.status}")
    ```

=== "cURL"

    ```bash
    curl -X POST https://api.regolo.ai/v1/batches \
        -H "Authorization: Bearer YOUR_API_KEY" \
        -H "Content-Type: application/json" \
        -d '{
            "input_file_id": "file-abc123",
            "endpoint": "/v1/chat/completions",
            "timeout_minutes": 1440
        }'
    ```

### Step 3: Monitor Progress

```python
import regolo

batch = regolo.get_batch(batch_id="batch_123")
print(f"Status: {batch.status}")
print(f"Processed: {batch.request_counts.completed} / {batch.request_counts.total}")
```

### Step 4: Retrieve Results

```python
# Get results when complete
results = regolo.get_batch_results(batch_id="batch_123")

for result in results:
    print(result)
    # {
    #   "id": "batch-req-123",
    #   "custom_id": "request-1",
    #   "response": {
    #     "body": {
    #       "choices": [{"message": {"content": "AI is..."}}],
    #       "usage": {"total_tokens": 150}
    #     }
    #   },
    #   "error": null
    # }
```

## Batch Job Lifecycle

```
Submitted
    ↓
Validating
    ↓
Queued
    ↓
Processing (in progress)
    ↓
Completed / Failed
```

## Pricing Benefits

Batch processing offers cost savings:

### Regular Processing

```
1000 requests × $0.00002/token = Cost based on tokens
```

### Batch Processing

```
1000 requests in batch = 10% discount
Cost: 90% of regular pricing
```

## Request Format

### Basic Format

```json
{
  "custom_id": "request-1",
  "params": {
    "model": "Llama-3.3-70B-Instruct",
    "messages": [
      {"role": "user", "content": "Hello"}
    ]
  }
}
```

### With Advanced Options

```json
{
  "custom_id": "request-1",
  "params": {
    "model": "Llama-3.3-70B-Instruct",
    "messages": [{"role": "user", "content": "Summarize this document"}],
    "max_tokens": 100,
    "temperature": 0.7,
    "top_p": 0.9
  }
}
```

## Response Format

```json
{
  "id": "batch-req-123",
  "custom_id": "request-1",
  "response": {
    "status_code": 200,
    "request_id": "req-123",
    "body": {
      "id": "chatcmpl-abc",
      "object": "chat.completion",
      "created": 1703001234,
      "model": "Llama-3.3-70B-Instruct",
      "choices": [
        {
          "index": 0,
          "message": {"role": "assistant", "content": "Response content"},
          "finish_reason": "stop"
        }
      ],
      "usage": {
        "prompt_tokens": 10,
        "completion_tokens": 50,
        "total_tokens": 60
      }
    }
  },
  "error": null
}
```

## Complete Example

### Create Batch File

```python
import json

requests = [
    {"custom_id": f"req-{i}", "params": {
        "model": "Llama-3.3-70B-Instruct",
        "messages": [{"role": "user", "content": f"Question {i}"}]
    }} for i in range(1000)
]

# Save as JSONL
with open('batch_requests.jsonl', 'w') as f:
    for req in requests:
        f.write(json.dumps(req) + '\n')
```

### Submit and Monitor

```python
import regolo
import time

# Submit batch
with open('batch_requests.jsonl', 'r') as f:
    batch = regolo.submit_batch(requests=f.read())

print(f"Batch {batch.id} submitted")

# Monitor progress
while True:
    batch = regolo.get_batch(batch.id)
    print(f"Status: {batch.status}")
    print(f"Progress: {batch.request_counts.completed}/{batch.request_counts.total}")
    
    if batch.status == "completed":
        break
    
    time.sleep(10)  # Check every 10 seconds

# Get results
results = regolo.get_batch_results(batch.id)
for result in results:
    if result['error'] is None:
        content = result['response']['body']['choices'][0]['message']['content']
        print(f"{result['custom_id']}: {content}")
    else:
        print(f"{result['custom_id']}: Error - {result['error']}")
```

## Best Practices

1. **Use custom IDs** - Track results with meaningful identifiers
2. **Batch size** - 1000-10000 requests per batch is optimal
3. **Error handling** - Check each result for errors
4. **Cleanup** - Process results and delete when done
5. **Timeout** - Set appropriate timeout for your workload

## Limits

| Parameter | Limit |
|-----------|-------|
| Batch size | 10,000 requests |
| File size | 500 MB |
| Timeout | 24 hours |
| Concurrent batches | 10 per account |

## Troubleshooting

### Batch Stuck in Processing

```python
# Cancel if needed
regolo.cancel_batch(batch_id="batch_123")
```

### High Error Rate

- Check request format
- Verify model availability
- Review error messages in results

### Slow Processing

- Submit during off-peak hours
- Reduce batch size
- Check system load