# Version Control

Manage model versions and track changes throughout your model lifecycle.

## Overview

Version control allows you to:

- Track model iterations over time
- Compare different versions
- Rollback to previous versions
- Document changes and experiments
- Collaborate with team members

## Creating Versions

### Automatic Versioning

Each time you update a model, a new version is created:

```python
import regolo

# Version 1.0 is created automatically
model = regolo.register_model(
    name="sentiment-classifier",
    model_data=model_artifact,
    version="1.0"
)
```

### Manual Version Creation

Create versions for experiments:

```python
import regolo

# Version 1.1 - Experiment with new data
model_v1_1 = regolo.create_model_version(
    model_id="sentiment-classifier",
    version="1.1",
    model_data=updated_artifact,
    metadata={
        "experiment": "expanded-training-data",
        "accuracy": 0.92,
        "f1_score": 0.89
    }
)
```

## Version Metadata

Each version stores important information:

```python
version = regolo.get_model_version(
    model_id="sentiment-classifier",
    version="1.0"
)

print(f"Created: {version.created_at}")
print(f"Author: {version.created_by}")
print(f"Size: {version.size_mb} MB")
print(f"Status: {version.status}")
print(f"Metrics: {version.metadata['accuracy']}")
```

## Comparing Versions

### Side-by-Side Comparison

```python
import regolo

comparison = regolo.compare_model_versions(
    model_id="sentiment-classifier",
    version_a="1.0",
    version_b="1.1"
)

print(f"Size difference: {comparison.size_diff_mb} MB")
print(f"Accuracy improvement: {comparison.metrics['accuracy']['diff']:.2%}")
print(f"F1 score improvement: {comparison.metrics['f1_score']['diff']:.2%}")
```

### Performance Comparison

```python
metrics_comparison = {
    "1.0": {"accuracy": 0.89, "latency_ms": 150},
    "1.1": {"accuracy": 0.92, "latency_ms": 165},
}

for version, metrics in metrics_comparison.items():
    print(f"Version {version}:")
    print(f"  Accuracy: {metrics['accuracy']:.2%}")
    print(f"  Latency: {metrics['latency_ms']}ms")
```

## Versioning Strategy

### Semantic Versioning

Use MAJOR.MINOR.PATCH format:

```
1.0.0  - Initial release
1.1.0  - New features
1.1.1  - Bug fix
2.0.0  - Breaking changes
```

### By Environment

```
dev-1.0   - Development
staging-1.0 - Staging
prod-1.0  - Production
```

## Rollback

### Rollback to Previous Version

```python
import regolo

# Issue found with v1.1, rollback to v1.0
regolo.rollback_endpoint(
    endpoint_id="sentiment-endpoint",
    to_version="1.0"
)

print("Rolled back to version 1.0")
```

### Gradual Rollback

```python
# Canary deployment: 10% traffic to new version
regolo.update_endpoint_traffic(
    endpoint_id="sentiment-endpoint",
    version_traffic={
        "1.0": 0.9,   # 90% traffic
        "1.1": 0.1    # 10% traffic
    }
)

# Monitor metrics
metrics = regolo.get_endpoint_metrics(endpoint_id="sentiment-endpoint")
print(f"Error rate (v1.1): {metrics.version['1.1']['error_rate']:.2%}")

# If errors low, increase traffic to v1.1
if metrics.version['1.1']['error_rate'] < 0.01:
    regolo.update_endpoint_traffic(
        endpoint_id="sentiment-endpoint",
        version_traffic={"1.0": 0.5, "1.1": 0.5}  # 50/50 split
    )
```

## Best Practices

1. **Version Naming** - Use consistent naming scheme
2. **Document Changes** - Record what changed in each version
3. **Test Thoroughly** - Validate versions before production
4. **Monitor Metrics** - Track performance of each version
5. **Clean Up** - Remove old versions you don't need

## Listing Versions

```python
import regolo

# Get all versions
versions = regolo.list_model_versions(
    model_id="sentiment-classifier"
)

for version in versions:
    print(f"{version.version} - {version.created_at} - {version.status}")
```

## Version Promotion

### Promote to Production

```python
import regolo

# After testing in staging, promote to production
regolo.promote_model_version(
    model_id="sentiment-classifier",
    version="1.1",
    environment="production",
    replicas=10  # Number of instances
)
```