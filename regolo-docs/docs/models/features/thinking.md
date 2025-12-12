# Thinking

Thinking is a feature that allows models to reason through problems step by step, showing their internal thought process before providing a final answer.

## Overview

The thinking feature enables models to break down complex problems into smaller steps, making their reasoning process transparent and allowing for better understanding of how they arrive at their conclusions.

## Usage

To enable thinking, you can use the `thinking` parameter in your API requests.

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
        {"role": "user", "content": "What color was Napoleon's white horse?"}
    ],
    "thinking": True
}

response = requests.post(api_url, headers=headers, json=data)
result = response.json()

# Extract main content and reasoning
message = result.get("choices", [{}])[0].get("message", {})
content = message.get("content", "")
reasoning = message.get("reasoning_content", "")

if reasoning:
    print("=== Reasoning ===")
    print(reasoning)
    print("\n=== Final Answer ===")

print(content)
```

## Parameters

### `reasoning_effort`

Controls the depth and detail of the reasoning process. Available values:

- `low`: Minimal reasoning effort, faster responses with brief reasoning
- `medium`: Balanced reasoning effort (default)
- `high`: Maximum reasoning effort, more detailed and thorough reasoning

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Solve this step by step: What is 15% of 240?"}
    ],
    "thinking": True,
    "reasoning_effort": "high"  # low, medium, or high
}
```

!!! note "Standard Parameters"
    Standard API parameters like `temperature`, `max_tokens`, `top_p`, `frequency_penalty`, and `presence_penalty` can also influence the thinking process. See [Response Parameters](response-parameters.md) for detailed documentation.

## Benefits

- **Transparency**: See how the model reasons through problems
- **Debugging**: Understand where the model might make mistakes
- **Education**: Learn problem-solving strategies from the model's reasoning
- **Quality**: Better results for complex, multi-step problems

