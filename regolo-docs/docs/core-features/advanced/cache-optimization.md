# Cache & Optimization

Optimize performance and reduce costs with intelligent caching strategies.

## Cache Types

### Request Caching

Automatically cache identical requests and return cached results:

```python
import regolo

response = regolo.static_chat_completions(
    messages=[{"role": "user", "content": "What is 2+2?"}],
    cache_control="max-age=3600"  # Cache for 1 hour
)
```

**Benefits:**
- Reduce API costs for repeated queries
- Instant response times for cached results
- Reduced server load

### Prompt Caching

Cache expensive prompt prefixes to optimize costs:

```python
response = regolo.static_chat_completions(
    messages=[
        {"role": "system", "content": "You are a helpful assistant..."},  # Cached
        {"role": "user", "content": "User question"}
    ],
    cache_prompt=True
)
```

**Use cases:**
- Long system prompts (documentation, templates)
- Multi-turn conversations
- Batch processing with common prefixes

## Cache Configuration

### TTL (Time to Live)

```python
# Cache for 1 hour
response = regolo.static_chat_completions(
    messages=messages,
    cache_ttl=3600
)

# Cache until manual expiration
response = regolo.static_chat_completions(
    messages=messages,
    cache_ttl="forever"
)
```

### Cache Keys

Customize cache key strategy:

```python
response = regolo.static_chat_completions(
    messages=messages,
    cache_key="user_123:conversation_456"
)
```

## Cache Statistics

Monitor cache performance:

```python
response = regolo.static_chat_completions(messages=messages)

print(f"Cache hit: {response.cache_hit}")
print(f"Cache age: {response.cache_age}")
print(f"Tokens saved: {response.cache_tokens_saved}")
```

## Cost Optimization

### Before Caching

```
1000 requests × 50 tokens = 50,000 tokens
Cost: $0.50
```

### After Caching (90% hit rate)

```
100 new requests × 50 tokens = 5,000 tokens
900 cached requests = 0 tokens
Cost: $0.05 (90% reduction)
```

## Best Practices

1. **Cache high-frequency queries**: Focus on queries executed many times
2. **Use appropriate TTL**: Balance freshness vs. cost savings
3. **Monitor hit rates**: Track cache effectiveness
4. **Segment by user**: Separate caches per user when appropriate
5. **Test thoroughly**: Ensure cached responses meet quality standards

## Advanced Optimization

### Batch Processing

```python
requests = [
    {"role": "user", "content": "Query 1"},
    {"role": "user", "content": "Query 2"},
    # ...
]

responses = regolo.batch_completions(
    requests=requests,
    enable_cache=True
)
```

### Distributed Caching

For multi-region deployments:

```python
response = regolo.static_chat_completions(
    messages=messages,
    cache_strategy="distributed",  # Shared across regions
    cache_ttl=7200
)
```

## Cache Invalidation

Manually clear cache when needed:

```python
# Clear all caches
regolo.clear_cache()

# Clear specific cache
regolo.clear_cache(key="user_123:*")

# Clear by age
regolo.clear_cache(older_than=3600)  # Older than 1 hour
```