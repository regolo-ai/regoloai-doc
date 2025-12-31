# Response Parameters

The Regolo API follows the OpenAI Chat Completions API specification and supports standard parameters that control the behavior and quality of model responses.

## Overview

These parameters allow you to fine-tune how the model generates responses, controlling aspects like creativity, length, diversity, and repetition.

## Standard Parameters

### `temperature`

Controls the randomness of the model's output. 

- **Type**: `float`
- **Range**: 0.0 to 2.0
- **Default**: 1.0

Lower values (e.g., 0.2) make the output more deterministic and focused, while higher values (e.g., 0.8) increase creativity and variability.

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Tell me a story"}
    ],
    "temperature": 0.7  # Balanced creativity
}
```

### `max_tokens`

Sets the maximum number of tokens that can be generated in the response.

- **Type**: `integer`
- **Default**: Varies by model

Higher values allow for longer responses but may increase latency and costs. Lower values provide more concise outputs.

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Explain quantum computing"}
    ],
    "max_tokens": 1000  # Limit response length
}
```

### `top_p`

Nucleus sampling parameter that controls the diversity of tokens considered.

- **Type**: `float`
- **Range**: 0.0 to 1.0
- **Default**: 1.0

Lower values make the model more focused on likely tokens, while higher values allow more diverse choices.

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Write a creative poem"}
    ],
    "top_p": 0.9  # Allow more diverse word choices
}
```

### `frequency_penalty`

Reduces the likelihood of the model repeating frequent tokens.

- **Type**: `float`
- **Range**: -2.0 to 2.0
- **Default**: 0

Positive values reduce repetition, while negative values increase it. Useful for avoiding repetitive phrases.

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "List 10 different ideas"}
    ],
    "frequency_penalty": 0.5  # Reduce repetition
}
```

### `presence_penalty`

Reduces the likelihood of the model reusing tokens that are already present in the text.

- **Type**: `float`
- **Range**: -2.0 to 2.0
- **Default**: 0

Positive values encourage the model to use new tokens and explore different approaches.

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Generate diverse solutions"}
    ],
    "presence_penalty": 0.6  # Encourage new approaches
}
```

### `stop`

Array of strings that cause the model to stop generating when encountered.

- **Type**: `array[string]` or `string`
- **Default**: `null`

Useful for controlling where the response ends or preventing certain phrases.

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Count to 10"}
    ],
    "stop": ["11", "twelve"]  # Stop at these sequences
}
```

### `n`

Number of completions to generate for each prompt.

- **Type**: `integer`
- **Default**: 1

Generates multiple response variations.

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Suggest a name"}
    ],
    "n": 3  # Generate 3 different suggestions
}
```

### `seed`

Seed for reproducible outputs.

- **Type**: `integer`
- **Default**: `null`

When set, the model will produce more deterministic outputs, useful for testing and reproducibility.

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Generate a random number"}
    ],
    "seed": 42  # Reproducible output
}
```

## Complete Example

```python
import requests

api_url = "https://api.regolo.ai/v1/chat/completions"
headers = {
    "Content-Type": "application/json",
    "Authorization": "Bearer YOUR_REGOLO_KEY"
}
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Write a creative story"}
    ],
    "temperature": 0.8,
    "max_tokens": 500,
    "top_p": 0.9,
    "frequency_penalty": 0.3,
    "presence_penalty": 0.4
}

response = requests.post(api_url, headers=headers, json=data)
result = response.json()
print(result["choices"][0]["message"]["content"])
```

## Best Practices

- **For factual tasks**: Use lower `temperature` (0.2-0.4) for more deterministic outputs
- **For creative tasks**: Use higher `temperature` (0.7-0.9) for more varied responses
- **For long responses**: Set appropriate `max_tokens` to avoid truncation
- **To reduce repetition**: Use `frequency_penalty` (0.3-0.7) and `presence_penalty` (0.4-0.6)
- **For consistency**: Use `seed` when you need reproducible outputs






