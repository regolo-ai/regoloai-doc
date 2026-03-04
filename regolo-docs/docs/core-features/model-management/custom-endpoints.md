# Custom Endpoints

Deploy custom models and create endpoints with full control over configuration and resources.

## Overview

Custom endpoints allow you to:

- Deploy any Hugging Face model from the library
- Configure resource allocation and GPU selection
- Create private endpoints for your models
- Control request/response handling
- Set custom SLAs
- Isolate traffic for different use cases

## Adding Hugging Face Models

### Upload Model from Hugging Face

Follow these 5 steps to add any Hugging Face model to your custom library:

#### Step 1: Start Upload
Click **"Add model"** button in the top-right corner of your Custom Model Library.

#### Step 2: Enter Model URL
Enter the Hugging Face model URL (e.g., `https://huggingface.co/meta-llama/Llama-2-7b-hf`).

#### Step 3: Deploy Model
Click **"Deploy model"** in your library to initialize the deployment.

#### Step 4: Choose Instance
Select the instance that fits your needs based on:
- Model size (7B, 13B, 70B, etc.)
- Required GPU resources (NVIDIA A100, H100, etc.)
- Expected throughput
- Budget constraints

#### Step 5: Get API Access
Use your regolo **'All models'** active API key to access the deployed model.

### Using Your Custom Models

```python
import regolo

# Initialize client with your API key
client = regolo.RegoloClient(api_key="YOUR_API_KEY")

# Make requests to your custom model
response = client.static_chat_completions(
    model="your-custom-model-name",
    messages=[
        {"role": "user", "content": "Hello, how are you?"}
    ]
)

print(response.choices[0].message.content)
```

### Custom Model Endpoint URL

Your custom models are accessible at:

```
https://api.regolo.ai/custom-model/v1/chat/completions/
```

Example with cURL:

```bash
curl https://api.regolo.ai/custom-model/v1/chat/completions/ \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "your-custom-model-name",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

## Pricing for Custom Models

### Hourly Usage Model

Custom models use hourly prepaid billing:

- **Charged from first hour** - You pay hourly once your model is deployed
- **Monthly invoice** - Receive dedicated invoice at month-end showing total hours used and charges
- **No per-request charges** - All requests included in hourly rate

### Cost Estimation

```
Hourly rate depends on:
- Model size (larger models cost more)
- GPU instance type selected (A100, H100, etc.)
- Replicas/concurrent instances

Example:
- Llama-2-7B on single A100: ~$0.50/hour = $360/month
- Llama-2-70B on multiple A100s: ~$2.00/hour = $1,440/month
```

## Creating Endpoints

### Basic Endpoint

```python
import regolo

endpoint = regolo.create_endpoint(
    name="sentiment-api",
    model="sentiment-classifier",
    version="1.0",
    replicas=2  # Number of instances
)

print(f"Endpoint ID: {endpoint.id}")
print(f"URL: {endpoint.url}")
```

### Advanced Configuration

```python
import regolo

endpoint = regolo.create_endpoint(
    name="sentiment-api-prod",
    model="sentiment-classifier",
    version="1.0",
    
    # Scaling
    replicas=5,
    min_replicas=2,
    max_replicas=20,
    target_utilization=0.7,
    
    # Resources
    cpu_request="2",
    memory_request="4Gi",
    cpu_limit="4",
    memory_limit="8Gi",
    
    # Timeout
    request_timeout_seconds=30,
    
    # Network
    enable_cache=True,
    enable_compression=True
)
```

## Deploying Models

### Deploy Model Version

```python
import regolo

endpoint = regolo.create_endpoint(
    name="classifier-v1-1",
    model="sentiment-classifier",
    version="1.1",
    replicas=3
)

print(f"Endpoint created: {endpoint.status}")
```

### Canary Deployment

Gradually roll out new versions:

```python
import regolo

# Deploy new version to canary
canary = regolo.create_canary(
    endpoint_id="sentiment-api",
    new_version="1.1",
    traffic_percentage=5  # Start with 5% traffic
)

print(f"Canary deployment started")
print(f"New version: {canary.new_version}")
print(f"Traffic: {canary.traffic_percentage}%")
```

## Managing Endpoints

### List Endpoints

```python
import regolo

endpoints = regolo.list_endpoints()

for endpoint in endpoints:
    print(f"{endpoint.name} ({endpoint.id})")
    print(f"  Model: {endpoint.model}:{endpoint.version}")
    print(f"  Status: {endpoint.status}")
    print(f"  Replicas: {endpoint.replicas}")
    print(f"  URL: {endpoint.url}")
```

### Update Endpoint

```python
import regolo

regolo.update_endpoint(
    endpoint_id="sentiment-api",
    replicas=10,  # Scale up
    request_timeout_seconds=60
)
```

### Delete Endpoint

```python
import regolo

regolo.delete_endpoint(endpoint_id="sentiment-api")
print("Endpoint deleted")
```

## Traffic Management

### Split Traffic

```python
import regolo

# A/B testing: 50% traffic to each version
regolo.update_endpoint_traffic(
    endpoint_id="sentiment-api",
    version_traffic={
        "1.0": 0.5,
        "1.1": 0.5
    }
)
```

### Blue-Green Deployment

```python
import regolo

# Route all traffic to green (new version)
regolo.update_endpoint_traffic(
    endpoint_id="sentiment-api",
    version_traffic={
        "1.0": 0,     # Blue (old) - no traffic
        "1.1": 1.0    # Green (new) - all traffic
    }
)
```

## Using Endpoints

### Making Requests

```python
import requests

endpoint_url = "https://sentiment-api.regolo.ai/v1"

response = requests.post(
    f"{endpoint_url}/chat/completions",
    headers={
        "Authorization": "Bearer YOUR_API_KEY",
        "Content-Type": "application/json"
    },
    json={
        "model": "sentiment-classifier",
        "messages": [{"role": "user", "content": "This is great!"}]
    }
)

result = response.json()
print(result['choices'][0]['message']['content'])
```

### Using Python Client

```python
import regolo

client = regolo.RegoloClient(
    api_key="YOUR_API_KEY",
    endpoint_url="https://sentiment-api.regolo.ai"
)

response = client.static_chat_completions(
    messages=[{"role": "user", "content": "This is great!"}]
)
```

## Monitoring Endpoints

### Endpoint Metrics

```python
import regolo

metrics = regolo.get_endpoint_metrics(
    endpoint_id="sentiment-api",
    time_range="24h"
)

print(f"Requests: {metrics.request_count}")
print(f"Average latency: {metrics.avg_latency_ms}ms")
print(f"Error rate: {metrics.error_rate:.2%}")
print(f"Availability: {metrics.availability:.2%}")
```

### Health Checks

```python
import regolo

health = regolo.check_endpoint_health(endpoint_id="sentiment-api")

print(f"Status: {health.status}")
print(f"Replicas ready: {health.ready_replicas} / {health.total_replicas}")
print(f"Last check: {health.last_check}")
```

## Auto-Scaling

### Configure Autoscaling

```python
import regolo

regolo.configure_autoscaling(
    endpoint_id="sentiment-api",
    metrics=[
        {
            "metric": "cpu_usage",
            "target": 0.7,
            "scale_up_threshold": 0.8,
            "scale_down_threshold": 0.3
        },
        {
            "metric": "request_rate",
            "target": 100,  # requests/min
            "scale_up_threshold": 150,
            "scale_down_threshold": 50
        }
    ],
    min_replicas=2,
    max_replicas=50
)
```

## Cost Optimization

### Right-Size Resources

```python
import regolo

# Monitor and optimize
metrics = regolo.get_endpoint_metrics(endpoint_id="sentiment-api")

if metrics.cpu_usage < 0.3:
    print("CPU is underutilized, consider reducing")
    regolo.update_endpoint(
        endpoint_id="sentiment-api",
        cpu_request="1"
    )
```

### Batch Requests

```python
import regolo

# Use batch API for non-time-critical work
regolo.submit_batch(
    requests="batch.jsonl",
    endpoint_id="sentiment-api"
)
```

## Best Practices

1. **Resource Allocation** - Monitor and adjust resources based on usage
2. **Autoscaling** - Enable for variable workloads
3. **Health Checks** - Monitor endpoint health regularly
4. **Error Handling** - Implement proper error handling in clients
5. **Load Testing** - Test endpoints under expected load
6. **Cost Monitoring** - Track and optimize endpoint costs
7. **Versioning** - Use canary deployments for new versions