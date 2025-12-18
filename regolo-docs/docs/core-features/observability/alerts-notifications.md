# Alerts & Notifications

Set up alerts and notifications for important events and thresholds.

## Alert Types

### Quota Alerts

Notify when approaching usage limits:

```
Daily token quota: 1M tokens
Alert at 80%: 800K tokens
Alert at 90%: 900K tokens
```

### Cost Alerts

Notify when spending exceeds budget:

```
Daily budget: $100
Alert at $80 (80%)
Alert at $100 (100%)
```

### Performance Alerts

Notify on performance degradation:

```
P95 latency threshold: 500ms
Alert when exceeded

Error rate threshold: 1%
Alert when exceeded
```

### Availability Alerts

Notify on service issues:

```
Consecutive failed requests: 5
Alert triggered

Sustained error rate > 5%
Alert triggered
```

## Creating Alerts

### Via Dashboard

1. Go to [dashboard.regolo.ai/alerts](https://dashboard.regolo.ai/alerts)
2. Click **Create Alert**
3. Select alert type
4. Set threshold
5. Choose notification method
6. Save

### Via API

```python
import regolo

alert = regolo.create_alert(
    name="Daily quota warning",
    type="quota",
    threshold=800000,  # 80% of 1M
    comparison="greater_than",
    notification_method="email",
    cooldown_minutes=30
)
```

## Notification Channels

### Email

Receive alerts via email:

```
Alert triggered at 10:30 AM
Email sent to your registered address
```

### SMS

Critical alerts via SMS:

```
[Regolo] Error rate exceeded 5% - Check dashboard
```

### Slack

Integrate with Slack:

```
[Regolo Alert] Cost limit reached
Daily spend: $100 / $100
Details: <link to dashboard>
```

### Webhook

Custom webhooks:

```python
alert = regolo.create_alert(
    type="cost",
    threshold=100,
    notification_method="webhook",
    webhook_url="https://your-app.com/alerts"
)
```

## Alert Configuration

### Severity Levels

| Level | Description | Notification |
|-------|-------------|---------------|
| Info | FYI events | Dashboard only |
| Warning | Approaching limits | Email |
| Critical | Urgent action needed | Email + SMS |

### Cooldown Periods

Prevent alert spam:

```
Alert fires at 10:30 AM
Cooldown: 30 minutes
Next alert: Can fire at 11:00 AM earliest
```

### Alert Grouping

Group related alerts:

```
[Regolo] 5 alerts triggered
├─ Error rate: 5.2%
├─ P95 latency: 600ms
├─ Daily quota: 95%
├─ Cost limit: 80%
└─ Failed requests: 10
```

## Example Alert Configurations

### Development Monitoring

```
Alert 1: Error rate > 1% → Email
Alert 2: P99 latency > 5000ms → Email
Alert 3: Failed requests > 5 → Dashboard
```

### Production Monitoring

```
Alert 1: Error rate > 0.5% → Email + SMS
Alert 2: P95 latency > 500ms → Email
Alert 3: Daily cost > $500 → Email
Alert 4: Failed requests > 3 → SMS
Alert 5: Quota > 90% → Email
```

### Cost Control

```
Alert 1: Daily cost > 70% budget → Email
Alert 2: Daily cost > 90% budget → Email + SMS
Alert 3: Monthly cost > 50% budget → Email
Alert 4: Monthly cost > 90% budget → Email + SMS
```

## Slack Integration

### Setup

1. Go to **Integrations** in dashboard
2. Click **Connect Slack**
3. Authorize Regolo app
4. Select channel for alerts
5. Configure which alerts to send

### Custom Messages

```python
alert = regolo.create_alert(
    type="cost",
    threshold=100,
    notification_method="slack",
    slack_config={
        "channel": "#alerts",
        "mention": "@devops",
        "message_template": "Cost threshold exceeded: ${{ actual }} / ${{ threshold }}"
    }
)
```

## Alert Status

### Dismissing Alerts

Temporarily dismiss repeated alerts:

```
Alert is dismissed
Will reappear in 24 hours or if condition changes
```

### Muting Alerts

Disable alerts during maintenance:

```python
regolo.mute_alerts(
    duration_minutes=60,
    alert_type="cost"  # Specific type
)
```

## Alert Analytics

### Alert History

View past alerts:

```
December 18:
  10:30 - Cost limit reached
  11:00 - Error rate spike
  15:30 - Quota alert

December 17:
  09:30 - Daily quota reset
```

### Alert Trends

Analyze alert frequency:

```
Last 7 days: 5 alerts
Last 30 days: 12 alerts
Most common: Cost limit alerts
```