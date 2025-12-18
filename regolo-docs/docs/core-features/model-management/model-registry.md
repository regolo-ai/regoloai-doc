# Model Registry

Centralized hub for discovering, managing, and organizing your models.

## Overview

The Model Registry is your central repository for:

- All your trained and deployed models
- Model metadata and documentation
- Performance metrics and benchmarks
- Access control and permissions
- Model lineage and dependencies

## Registering Models

### Register from Local Files

```python
import regolo

model = regolo.register_model(
    name="sentiment-classifier",
    version="1.0",
    model_path="./models/sentiment.pkl",
    description="Sentiment analysis model trained on 50K tweets",
    tags=["nlp", "sentiment", "production"],
    metadata={
        "framework": "scikit-learn",
        "accuracy": 0.92,
        "f1_score": 0.89,
        "training_data": "tweets-50k"
    }
)

print(f"Model registered: {model.id}")
```

### Register from Different Frameworks

```python
import regolo

# PyTorch model
regolo.register_model(
    name="pytorch-classifier",
    framework="pytorch",
    model_path="./models/model.pt"
)

# TensorFlow model
regolo.register_model(
    name="tf-classifier",
    framework="tensorflow",
    model_path="./models/model"
)

# ONNX model
regolo.register_model(
    name="onnx-classifier",
    framework="onnx",
    model_path="./models/model.onnx"
)
```

## Searching Models

### Search by Name

```python
import regolo

models = regolo.search_models(
    query="sentiment",
    limit=10
)

for model in models:
    print(f"{model.name} ({model.id}) - v{model.latest_version}")
```

### Filter by Tags

```python
# Find all NLP models
models = regolo.search_models(
    tags=["nlp"]
)

# Find production-ready models
models = regolo.search_models(
    tags=["production"],
    status="active"
)
```

### Advanced Search

```python
# Search with multiple criteria
models = regolo.search_models(
    query="classifier",
    tags=["nlp", "production"],
    framework="pytorch",
    min_accuracy=0.85,
    sort_by="created_at",
    sort_order="desc"
)
```

## Model Metadata

### View Model Details

```python
import regolo

model = regolo.get_model(model_id="sentiment-classifier")

print(f"Name: {model.name}")
print(f"Description: {model.description}")
print(f"Framework: {model.metadata['framework']}")
print(f"Accuracy: {model.metadata['accuracy']:.2%}")
print(f"Tags: {', '.join(model.tags)}")
print(f"Created by: {model.created_by}")
print(f"Created at: {model.created_at}")
```

### Update Metadata

```python
import regolo

regolo.update_model(
    model_id="sentiment-classifier",
    description="Updated: now supports 10 languages",
    tags=["nlp", "sentiment", "multilingual"],
    metadata={
        "languages": 10,
        "accuracy": 0.94  # Improved
    }
)
```

## Organizing Models

### Collections

Group related models:

```python
import regolo

# Create a collection
collection = regolo.create_collection(
    name="sentiment-analysis",
    description="All sentiment analysis models",
    tags=["nlp", "sentiment"]
)

# Add models to collection
regolo.add_to_collection(
    collection_id=collection.id,
    model_ids=["sentiment-classifier", "aspect-sentiment-model"]
)
```

### Teams

Manage access by team:

```python
import regolo

# Create a team
team = regolo.create_team(
    name="NLP Team",
    description="Natural language processing models"
)

# Add team members
regolo.add_team_member(
    team_id=team.id,
    user_email="engineer@company.com",
    role="editor"
)

# Grant team access to model
regolo.grant_model_access(
    model_id="sentiment-classifier",
    team_id=team.id,
    permission="edit"
)
```

## Model Benchmarks

### View Benchmarks

```python
import regolo

benchmarks = regolo.get_model_benchmarks(
    model_id="sentiment-classifier"
)

for benchmark in benchmarks:
    print(f"Benchmark: {benchmark.name}")
    print(f"  Score: {benchmark.score:.2%}")
    print(f"  Date: {benchmark.date}")
```

### Compare with Baselines

```python
import regolo

# Get baseline models
baselines = regolo.search_models(
    tags=["baseline"],
    category="sentiment"
)

for baseline in baselines:
    comparison = regolo.compare_models(
        model_a="sentiment-classifier",
        model_b=baseline.id,
        metric="accuracy"
    )
    print(f"vs {baseline.name}: {comparison.improvement:.2%} improvement")
```

## Access Control

### Public Models

Share models publicly:

```python
import regolo

regolo.set_model_visibility(
    model_id="sentiment-classifier",
    visibility="public",
    require_approval=True  # Require approval for use
)
```

### Private Models

Restrict access to team:

```python
regolo.set_model_visibility(
    model_id="internal-model",
    visibility="private",
    allowed_users=["team-member@company.com"]
)
```

## Best Practices

1. **Clear Naming** - Use descriptive, consistent names
2. **Document Purpose** - Include description and use cases
3. **Tag Models** - Use tags for easy discovery
4. **Version Consistently** - Follow semantic versioning
5. **Track Metrics** - Record accuracy, latency, etc.
6. **Set Permissions** - Control who can access models
7. **Keep Updated** - Remove deprecated models