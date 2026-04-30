# FAQ

Common questions and answers about using the Regolo API.

## Overview

This page addresses frequently asked questions about API parameters, model limits, and troubleshooting common issues.

---

## Questions

### How do I disable model fallbacks?

By default, if the requested model is unavailable, the API automatically falls back to a compatible alternative. To disable this behavior and receive an error instead, set `disable_fallbacks` to `true`.

!!! warning "Important"
    When `disable_fallbacks` is `true`, requests will fail with an error if the selected model is unavailable. Use this when you need guaranteed model consistency.

```bash
curl -L -X POST 'https://api.regolo.ai/v1/chat/completions' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer YOUR_API_KEY' \
  -d '{
    "messages": [
      {
        "role": "user",
        "content": "List 5 important events in the 19th century"
      }
    ],
    "model": "qwen3.5-122b",
    "disable_fallbacks": true
  }'
```

---

### What is the maximum value for `max_tokens`?

The maximum value for `max_tokens` varies by model. Setting a value higher than the model's limit will automatically bcrash the request.

```bash
curl -L -X POST 'https://api.regolo.ai/v1/chat/completions' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer YOUR_API_KEY' \
  -d '{
    "messages": [
      {
        "role": "user",
        "content": "Write a long story"
      }
    ],
    "model": "qwen3.5-122b",
    "max_tokens": 120000
  }'
```

!!! note "Cost Consideration"
    Higher `max_tokens` values increase latency and token usage. Set only what you need for your use case.

---

### Can I generate multiple completions at once?

Yes, use the `n` parameter to generate multiple response variations in a single request.

```bash
curl -L -X POST 'https://api.regolo.ai/v1/chat/completions' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer YOUR_API_KEY' \
  -d '{
    "messages": [
      {
        "role": "user",
        "content": "Suggest a name for a new API product"
      }
    ],
    "model": "qwen3.5-122b",
    "n": 3
  }'
```

The response will contain 3 different choices in the `choices` array.

---

### How do I choose between creativity and accuracy?

Use the `temperature` parameter to control the balance:

| Temperature | Use Case |
|-------------|---------|
| 0.0 - 0.3 | Factual, accurate responses |
| 0.4 - 0.7 | Balanced (default) |
| 0.8 - 1.0 | Creative, varied outputs |

```bash
# For factual answers
curl -L -X POST 'https://api.regolo.ai/v1/chat/completions' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer YOUR_API_KEY' \
  -d '{
    "messages": [
      {
        "role": "user",
        "content": "What year did WW2 end?"
      }
    ],
    "model": "qwen3.5-122b",
    "temperature": 0.2
  }'
```

---

### How do I reduce repetition in model outputs?

Use `frequency_penalty` to reduce token repetition and `presence_penalty` to encourage new topics.

```bash
# Reduce repetition
curl -L -X POST 'https://api.regolo.ai/v1/chat/completions' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer YOUR_API_KEY' \
  -d '{
    "messages": [
      {
        "role": "user",
        "content": "List 10 different programming languages"
      }
    ],
    "model": "qwen3.5-122b",
    "frequency_penalty": 0.5,
    "presence_penalty": 0.4
  }'
```

| Parameter | Range | Effect |
|-----------|-------|--------|
| frequency_penalty | -2.0 to 2.0 | Reduces repeated tokens |
| presence_penalty | -2.0 to 2.0 | Encourages new topics |
