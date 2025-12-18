# Cost Analytics

Analyze and optimize your API costs.

## Cost Breakdown

Regolo pricing is based on tokens consumed:

```
Cost = (Input Tokens × Input Rate) + (Output Tokens × Output Rate)
```

### Example

```
Input: 100 tokens × $0.00001 = $0.001
Output: 50 tokens × $0.00002 = $0.001
Total: $0.002 per request
```

## Cost Analytics Dashboard

Access cost analytics at [dashboard.regolo.ai/costs](https://dashboard.regolo.ai/costs)

### Cost by Model

See which models drive your costs:

```
Llama-3.3-70B-Instruct: $1,250 (65%)
GPT-4 Vision: $450 (23%)
Embedding Model: $300 (15%)
```

### Cost Over Time

Track spending trends:

```
Week 1: $100
Week 2: $115
Week 3: $142
Week 4: $198
Month Total: $555
```

### Cost by Project

Allocate costs to projects:

```
Production API: $400 (72%)
Development: $120 (22%)
Testing: $35 (6%)
```

### Cost by API Key

Track usage by API key:

```
api_key_prod: $350
api_key_staging: $150
api_key_dev: $55
```

## Optimization Recommendations

### 1. Model Selection

Switching to cost-efficient models:

```
Current: Llama-3.3-70B @ $0.00002/token
Optimal: Llama-2-7B @ $0.00001/token
Savings: 50%
```

### 2. Batch Processing

Group requests for better efficiency:

```
Before: 1000 individual requests
After: 10 batch requests
Latency: Acceptable (non-critical)
Savings: 20%
```

### 3. Caching

Cache frequent requests:

```
Common query: "Summarize [document]"
Hit rate: 60%
Cost reduction: 40% on this query
```

### 4. Prompt Optimization

Reduce unnecessary tokens:

```
Before: 500 token prompt with examples
After: 150 token prompt, tuned
Tokens saved: 70%
```

## Cost Alerts

Set up alerts for unexpected spending:

=== "Daily Limit"

    Alert when daily cost exceeds $100

=== "Weekly Spike"

    Alert if weekly cost increases by 50%

=== "Monthly Budget"

    Alert when monthly spend reaches 80% of $5,000 budget

## Cost Export

### CSV Export

Export cost data for external analysis:

```bash
curl -X GET https://api.regolo.ai/v1/analytics/costs \
    -H "Authorization: Bearer YOUR_API_KEY" \
    -H "Accept: text/csv" \
    -G \
    -d "start_date=2024-12-01" \
    -d "end_date=2024-12-18" \
    > costs.csv
```

### JSON Export

```python
import regolo

cost_data = regolo.export_costs(
    start_date="2024-12-01",
    end_date="2024-12-18",
    format="json"
)

import json
with open('costs.json', 'w') as f:
    json.dump(cost_data, f)
```

## Billing

### Invoice Frequency

- Monthly invoices on the 1st
- Usage calculated from previous month
- Billed to associated payment method

### Payment Methods

Accepted:
- Credit card (Visa, Mastercard, Amex)
- Bank transfer (ACH)
- Purchase order (enterprise)

### Volume Discounts

| Monthly Tokens | Discount |
|---|---|
| 1M - 10M | 0% |
| 10M - 100M | 5% |
| 100M - 500M | 10% |
| 500M+ | 15% |

Contact sales for custom volume discounts.