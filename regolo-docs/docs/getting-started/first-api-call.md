# Your First API Call

Make your first API call to Regolo AI.

## Prerequisites

- A Regolo AI account ([sign up here](https://dashboard.regolo.ai))
- An API key ([learn how to create one](sign-up.md))

## Chat Completions API

The Chat Completions API is the primary way to interact with Regolo AI models. It accepts a list of messages and returns a model-generated response.

### Request Format

=== "Python"

    ```python
    import regolo

    regolo.default_key = "YOUR_API_KEY"
    regolo.default_chat_model = "Llama-3.3-70B-Instruct"

    response = regolo.static_chat_completions(
        messages=[
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": "What is the capital of Italy?"}
        ]
    )
    print(response)
    ```

=== "Python (requests)"

    ```python
    import requests

    api_url = "https://api.regolo.ai/v1/chat/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_API_KEY"
    }
    data = {
        "model": "Llama-3.3-70B-Instruct",
        "messages": [
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": "What is the capital of Italy?"}
        ]
    }

    response = requests.post(api_url, headers=headers, json=data)
    print(response.json())
    ```

=== "cURL"

    ```bash
    curl -X POST https://api.regolo.ai/v1/chat/completions \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer YOUR_API_KEY" \
        -d '{
            "model": "Llama-3.3-70B-Instruct",
            "messages": [
                {"role": "system", "content": "You are a helpful assistant."},
                {"role": "user", "content": "What is the capital of Italy?"}
            ]
        }'
    ```

### Response Format

The API returns a JSON response with the following structure:

```json
{
    "id": "chatcmpl-abc123",
    "object": "chat.completion",
    "created": 1234567890,
    "model": "Llama-3.3-70B-Instruct",
    "choices": [
        {
            "index": 0,
            "message": {
                "role": "assistant",
                "content": "The capital of Italy is Rome."
            },
            "finish_reason": "stop"
        }
    ],
    "usage": {
        "prompt_tokens": 25,
        "completion_tokens": 8,
        "total_tokens": 33
    }
}
```

### Message Roles

| Role | Description |
|------|-------------|
| `system` | Sets the behavior and context for the assistant |
| `user` | The human user's input |
| `assistant` | Previous responses from the model (for context) |

## Error Handling

Always handle potential errors in your API calls:

```python
import requests

try:
    response = requests.post(api_url, headers=headers, json=data)
    response.raise_for_status()
    result = response.json()
except requests.exceptions.HTTPError as e:
    print(f"HTTP error: {e}")
except requests.exceptions.RequestException as e:
    print(f"Request error: {e}")
```

## Next Steps

- Learn about [streaming responses](../core-features/inference-api/completions-and-chat.md)
- Explore [response parameters](../core-features/advanced/response-parameters.md)
- Try [vision models](../core-features/inference-api/vision-analysis.md) for image understanding