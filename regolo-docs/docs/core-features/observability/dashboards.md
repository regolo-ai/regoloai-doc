# Dashboards

View real-time metrics and insights with interactive dashboards.

## Default Dashboards

### Overview Dashboard

High-level summary of your API activity:

- **Requests**: Total requests and request rate (req/min)
- **Tokens**: Input and output tokens consumed
- **Cost**: Daily and monthly spend
- **Latency**: Average, P95, and P99 response times
- **Errors**: Error rate and recent errors

### Models Dashboard

Breakdown by model:

- Requests per model
- Token consumption per model
- Cost per model
- Error rate per model
- Latency distribution by model

### Projects Dashboard

If using multiple projects:

- Requests per project
- Costs per project
- Team member activity
- API key usage

## Creating Custom Dashboards

### 1. Add Dashboard

From the dashboard homepage:

1. Click **Create Dashboard**
2. Name your dashboard
3. Select widgets to add

### 2. Add Widgets

Available widgets:

- **Metric cards** - Single metrics (requests, tokens, cost)
- **Line charts** - Trends over time
- **Bar charts** - Comparisons across dimensions
- **Pie charts** - Distribution breakdown
- **Heatmaps** - Usage patterns by hour/day
- **Tables** - Raw data with filtering

### 3. Configure Filters

Limit data by:

- Time range
- Model
- Project
- API key
- User
- Status code

## Example Dashboards

### Development Monitoring

```
Layout:
- Requests (line chart, 7 days)
- Errors (metric card)
- P99 Latency (metric card)
- Tokens by Model (bar chart)
```

### Production Monitoring

```
Layout:
- Request Rate (metric card)
- Error Rate (metric card)
- P95/P99 Latency (metric cards)
- Cost Trend (line chart, 30 days)
- Tokens by Model (pie chart)
- Recent Errors (table)
```

### Cost Optimization

```
Layout:
- Daily Cost (line chart)
- Cost by Model (pie chart)
- Tokens per Dollar (metric card)
- Cost by User (bar chart)
- Top Models (table)
```

## Metric Definitions

| Metric | Description |
|--------|-------------|
| **Requests** | Total API requests (includes retries) |
| **Tokens** | Input + output tokens processed |
| **Cost** | Estimated cost based on token usage |
| **Latency (P50)** | Median response time |
| **Latency (P95)** | 95th percentile response time |
| **Latency (P99)** | 99th percentile response time |
| **Error Rate** | Percentage of requests that failed |
| **Throughput** | Requests per minute |

## Dashboard Sharing

### Share with Team

1. Click **Share** on dashboard
2. Select team members
3. Set permissions (view-only or edit)
4. Copy shareable link

### Public Dashboard

Create public dashboards for status pages:

```
https://dashboard.regolo.ai/public/dashboards/abc123
```

## API Access

Access dashboard data programmatically:

```python
import regolo

# Get dashboard metrics
metrics = regolo.get_dashboard_metrics(
    dashboard_id="dashboard_123",
    time_range="7d"
)

print(f"Requests: {metrics['requests']}")
print(f"Cost: ${metrics['cost']:.2f}")
print(f"Error rate: {metrics['error_rate']:.2%}")
```