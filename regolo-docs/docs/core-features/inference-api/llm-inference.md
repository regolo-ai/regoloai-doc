# LLM Inference

Advanced LLM inference capabilities for production workloads.

## Supported Models

Regolo supports leading open-source and commercial LLMs:

### Large Language Models

| Model | Parameters | Context | Speed | Best For |
|-------|-----------|---------|-------|----------|
| Llama-3.3-70B-Instruct | 70B | 8K | Medium | General purpose, instruction-following |
| Llama-2-70B-Chat | 70B | 4K | Medium | Chat, conversational |
| Llama-2-7B | 7B | 4K | Fast | Edge, mobile |
| Mistral-7B | 7B | 8K | Fast | Efficient, cost-effective |

## Request Format

### Basic Completion

```python
import regolo

response = regolo.static_chat_completions(
    model="Llama-3.3-70B-Instruct",
    messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "Explain quantum computing"}
    ]
)

print(response.choices[0].message.content)
```

### Advanced Parameters

```python
response = regolo.static_chat_completions(
    model="Llama-3.3-70B-Instruct",
    messages=[
        {"role": "user", "content": "Write a poem about AI"}
    ],
    max_tokens=100,           # Limit output length
    temperature=0.7,          # Creativity (0=deterministic, 1=random)
    top_p=0.9,               # Nucleus sampling
    top_k=40,                # Top-k sampling
    frequency_penalty=0,      # Reduce repetition
    presence_penalty=0        # Encourage new topics
)
```

## Message Format

### Message Roles

```python
messages = [
    # System message sets behavior
    {"role": "system", "content": "You are an expert Python developer."},
    
    # Previous conversation for context
    {"role": "user", "content": "What's the capital of France?"},
    {"role": "assistant", "content": "The capital of France is Paris."},
    
    # Current user question
    {"role": "user", "content": "What's its population?"}
]
```

### Content Types

```python
# Text only
{"role": "user", "content": "Hello"}

# Multiple content blocks
{"role": "user", "content": [
    {"type": "text", "text": "What's in this image?"},
    {"type": "image_url", "image_url": {"url": "https://..."}}
]}
```

## Streaming Responses

For real-time response streaming:

```python
stream = regolo.static_chat_completions(
    model="Llama-3.3-70B-Instruct",
    messages=[{"role": "user", "content": "Tell a long story"}],
    stream=True
)

for chunk in stream:
    if chunk.choices[0].delta.content:
        print(chunk.choices[0].delta.content, end="", flush=True)
```

## Response Format

```json
{
  "id": "chatcmpl-abc123",
  "object": "chat.completion",
  "created": 1703001234,
  "model": "Llama-3.3-70B-Instruct",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Quantum computing uses quantum mechanics..."
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 25,
    "completion_tokens": 150,
    "total_tokens": 175
  }
}
```

## Finish Reasons

| Reason | Meaning |
|--------|----------|
| `stop` | Model reached natural end |
| `length` | `max_tokens` limit reached |
| `content_filter` | Content policy violation |
| `function_call` | Function call triggered |

## Error Handling

### Common Errors

```python
import regolo

try:
    response = regolo.static_chat_completions(
        model="Llama-3.3-70B-Instruct",
        messages=[{"role": "user", "content": "Hello"}]
    )
except regolo.APIError as e:
    print(f"API Error: {e.status_code} - {e.message}")
    # Handle API errors
except regolo.RateLimitError as e:
    print(f"Rate limited, retry after {e.retry_after} seconds")
    # Implement exponential backoff
except regolo.TimeoutError as e:
    print("Request timed out, try again")
    # Handle timeout
```

### Retry Logic

```python
import time
import random

def call_with_retry(messages, max_retries=3):
    for attempt in range(max_retries):
        try:
            return regolo.static_chat_completions(
                model="Llama-3.3-70B-Instruct",
                messages=messages
            )
        except regolo.RateLimitError:
            wait_time = 2 ** attempt + random.random()
            print(f"Rate limited, retrying in {wait_time:.1f}s")
            time.sleep(wait_time)
        except Exception as e:
            if attempt == max_retries - 1:
                raise
            time.sleep(2 ** attempt)
```

## Performance Optimization

### 1. Model Selection

Choose the right model for your use case:

```python
# Fast but less capable
response = regolo.static_chat_completions(
    model="Llama-2-7B",  # 80ms, cheaper
    messages=messages
)

# Slower but more capable
response = regolo.static_chat_completions(
    model="Llama-3.3-70B-Instruct",  # 150ms, more accurate
    messages=messages
)
```

### 2. Prompt Optimization

```python
# Before: 500 token prompt
prompt = """
You are an expert AI assistant with years of experience.
Please follow these guidelines carefully:
1. Be accurate
2. Be concise
...(many more guidelines)...
"""

# After: 50 token optimized prompt
prompt = "Expert AI assistant. Accurate, concise answers."

# Result: 10x faster, 10x cheaper
```

### 3. Token Estimation

```python
# Estimate before request
estimated = regolo.estimate_tokens(messages=messages)
print(f"Estimated cost: ${estimated * 0.00001:.4f}")

# Adjust if necessary
if estimated > 1000:
    # Simplify prompt
    messages = simplify_messages(messages)
```

### 4. Caching

```python
# Cache common system prompts
system_prompt = "You are a helpful assistant."

response = regolo.static_chat_completions(
    model="Llama-3.3-70B-Instruct",
    messages=[
        {"role": "system", "content": system_prompt, "cache_ttl": 3600},
        {"role": "user", "content": "Your question"}
    ]
)
```

## Production Best Practices

### 1. Rate Limiting

```python
from ratelimit import limits, sleep_and_retry
import regolo

@sleep_and_retry
@limits(calls=100, period=60)  # 100 requests per minute
def call_llm(messages):
    return regolo.static_chat_completions(
        model="Llama-3.3-70B-Instruct",
        messages=messages
    )
```

### 2. Monitoring

```python
import logging

logger = logging.getLogger(__name__)

response = regolo.static_chat_completions(
    model="Llama-3.3-70B-Instruct",
    messages=messages
)

logger.info(
    f"LLM call complete",
    extra={
        "model": "Llama-3.3-70B-Instruct",
        "tokens": response.usage.total_tokens,
        "latency_ms": response.usage.latency
    }
)
```

### 3. Fallback Strategy

```python
def get_response(messages):
    try:
        # Try primary model
        return regolo.static_chat_completions(
            model="Llama-3.3-70B-Instruct",
            messages=messages
        )
    except Exception as e:
        logger.warning(f"Primary model failed: {e}")
        try:
            # Fallback to faster model
            return regolo.static_chat_completions(
                model="Llama-2-7B",
                messages=messages
            )
        except Exception as e:
            logger.error(f"Both models failed: {e}")
            return None
```

## Advanced Patterns

### Function Calling

```python
tools = [
    {
        "type": "function",
        "function": {
            "name": "get_weather",
            "description": "Get the current weather",
            "parameters": {
                "type": "object",
                "properties": {
                    "location": {"type": "string"}
                },
                "required": ["location"]
            }
        }
    }
]

response = regolo.static_chat_completions(
    model="Llama-3.3-70B-Instruct",
    messages=messages,
    tools=tools
)
```

### Few-Shot Prompting

```python
messages = [
    {"role": "system", "content": "Classify sentiment as positive, negative, or neutral."},
    
    # Examples
    {"role": "user", "content": "I love this!"},
    {"role": "assistant", "content": "positive"},
    
    {"role": "user", "content": "This is terrible."},
    {"role": "assistant", "content": "negative"},
    
    # Actual request
    {"role": "user", "content": "Pretty good product"}
]

response = regolo.static_chat_completions(
    model="Llama-3.3-70B-Instruct",
    messages=messages
)
```