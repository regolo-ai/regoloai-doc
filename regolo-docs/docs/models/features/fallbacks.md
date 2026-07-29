# Fallbacks

When a requested model is unavailable, Regolo can automatically fall back to a compatible alternative so your requests keep working without interruption.

## Overview

Model availability can change unexpectedly due to maintenance, capacity limits, or provider outages. By default, the Regolo API is designed for resilience: if the model you requested cannot be served, the request is transparently routed to a compatible model that can fulfill it.

This behavior is ideal for production workloads where uptime matters more than pinning an exact model. If you need guaranteed model consistency instead, for example for reproducible evaluations or strict compliance, you can disable fallbacks and get a clear error in return.

## How it works

When a fallback occurs:

- The request is completed using a **compatible alternative** model.
- Your API call still succeeds, returning a standard response.
- This happens transparently, with no change required to your integration.

!!! note "Compatibility"
    Fallback models are chosen to be compatible with the requested one, so request and response shapes remain consistent with the standard Chat Completions format.

## Disabling fallbacks

To turn off automatic fallbacks and receive an error when the requested model is unavailable, set `disable_fallbacks` to `true`.

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

## When to disable fallbacks

- **Reproducibility**: evaluations, benchmarks, and tests where the exact model must be pinned.
- **Compliance**: regulated workflows that require a specific, declared model.
- **Debugging**: isolating issues to a single model without ambiguity.

For most everyday and production traffic, leaving fallbacks enabled is recommended to maximize availability.
