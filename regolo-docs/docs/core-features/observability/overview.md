# Observability & Monitoring

Monitor your API usage, performance, and costs with comprehensive observability tools.

## Key Metrics

Regolo provides real-time visibility into:

- **Request Metrics**: Volume, latency, error rates
- **Token Usage**: Prompt and completion tokens by model
- **Cost Analytics**: Per-request, per-model, per-user costs
- **Performance**: P50, P95, P99 latencies
- **Errors**: Error types, failure rates, stack traces

## Access Observability

### Dashboard

Visit [dashboard.regolo.ai](https://dashboard.regolo.ai) to access:

- Real-time metrics dashboards
- Cost analytics and breakdown
- Request and token tracking
- Alert configuration

### API

Access metrics programmatically:

```python
import regolo

# Get usage metrics
metrics = regolo.get_metrics(
    start_time="2024-12-01",
    end_time="2024-12-18",
    group_by="model"
)

for metric in metrics:
    print(f"{metric.model}: {metric.tokens} tokens, ${metric.cost}")
```

## Key Features

- **Real-time dashboards** - Live metrics and performance data
- **Cost analytics** - Break down costs by model, user, project
- **Token tracking** - Monitor token consumption patterns
- **Latency analysis** - Understand performance bottlenecks
- **Alerts** - Get notified of issues and quota limits
- **Export data** - Download reports and raw metrics