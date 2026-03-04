# Latency & Performance

Monitor latency and performance metrics across your API calls.

## Latency Metrics

### Request Latency

Time from request submission to response:

```
Total Latency = Network Delay + Queue Time + Processing Time + Network Return
```

### Processing Time

Time spent in the model:

```
Processing Time = Token Generation Time + Model Overhead
```

## Performance Tiers

### P50 (Median)

50% of requests complete faster:

- **Good**: < 200ms for completions
- **Excellent**: < 100ms for embeddings

### P95

95% of requests complete faster:

- **Good**: < 500ms for completions
- **Excellent**: < 300ms for embeddings

### P99

99% of requests complete faster:

- **Good**: < 1000ms for completions
- **Excellent**: < 1000ms for embeddings

## Performance Factors

### Model Size

Larger models = slower responses:

```
Llama-2-7B:       P50=80ms
Llama-3.3-70B:    P50=150ms
GPT-4:            P50=300ms
```

### Request Complexity

More tokens = longer processing:

```
100 input tokens:  P50=80ms
1000 input tokens: P50=150ms
```

### Output Length

More output tokens = longer generation:

```
10 output tokens:   P50=50ms
100 output tokens:  P50=150ms
500 output tokens:  P50=500ms
```

## Performance Monitoring

### Dashboard Widgets

```
Latency Trend (7 days)
├─ Line chart showing P50, P95, P99
├─ Color coding: Green (good), Yellow (acceptable), Red (poor)
└─ Hover for exact values

Latency by Model
├─ Bar chart
├─ Sortable by metric
└─ Drill-down to see detailed stats

Latency Distribution
├─ Histogram
├─ Shows bucket distribution
└─ Identify outliers
```

### API Access

```python
import regolo

# Get performance metrics
metrics = regolo.get_performance_metrics(
    model="Llama-3.3-70B-Instruct",
    time_range="24h"
)

print(f"P50: {metrics.p50}ms")
print(f"P95: {metrics.p95}ms")
print(f"P99: {metrics.p99}ms")
print(f"Error rate: {metrics.error_rate:.2%}")
```

## Optimization Tips

### 1. Reduce Input Size

```python
# Before: 2000 token prompt
response = regolo.static_chat_completions(
    messages=[{
        "role": "user",
        "content": very_long_prompt
    }]
)
# P50 latency: 300ms

# After: 500 token optimized prompt
response = regolo.static_chat_completions(
    messages=[{
        "role": "user",
        "content": optimized_prompt
    }]
)
# P50 latency: 80ms
```

### 2. Use Smaller Models

Trade quality for speed:

```python
# Accurate but slow
model = "Llama-3.3-70B-Instruct"  # P50=150ms

# Faster alternative
model = "Llama-2-7B"  # P50=80ms
```

### 3. Batch Processing

For non-latency-critical workloads:

```python
# Before: Sequential requests
for prompt in prompts:
    response = regolo.static_chat_completions(messages=prompt)
    
# After: Batch processing
responses = regolo.batch_completions(requests=prompts)
```

### 4. Enable Caching

Cache frequent requests:

```python
response = regolo.static_chat_completions(
    messages=messages,
    cache_ttl=3600
)
# Cache hit: <10ms latency
```

## Performance SLA

Our performance guarantees:

| Tier | P50 | P95 | P99 |
|------|-----|-----|-----|
| Standard | <300ms | <800ms | <2000ms |
| Premium | <200ms | <500ms | <1000ms |
| Enterprise | <100ms | <300ms | <500ms |

## Troubleshooting Slow Responses

### Check Model Performance

```bash
Is the model slower than baseline?
→ Check model-specific metrics
→ Compare with historical performance
```

### Check Network

```bash
Is network delay significant?
→ Check from different regions
→ Try private endpoint (if available)
```

### Check Queue

```bash
Is request queuing?
→ Check system load
→ Consider upgrade or reserved capacity
```

### Check Request

```bash
Is the prompt unusually large?
→ Optimize prompt size
→ Reduce input tokens
```