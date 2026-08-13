# Structured Outputs

Structured outputs force the model to reply in a JSON shape you define, so you can parse the response programmatically without guessing the format. Regolo supports the OpenAI `response_format` parameter and passes it straight through to the underlying engine.

## How it works

Add `response_format` to any `/v1/chat/completions` request. Two shapes are supported:

- **`json_schema`** — constrain the response to a JSON schema you provide. With `strict: true` this is a hard contract: the model must produce an object that validates against your schema.
- **`json_object`** — a lighter constraint that only requires the output to be valid JSON. The shape is up to the model, so guide it from the prompt.

```json
"response_format": {
    "type": "json_schema",
    "json_schema": {
        "name": "city_facts",
        "strict": true,
        "schema": {
            "type": "object",
            "properties": {
                "name": { "type": "string" },
                "population": { "type": "integer" },
                "landlocked": { "type": "boolean" }
            },
            "required": ["name", "population", "landlocked"],
            "additionalProperties": false
        }
    }
}
```

### Example: `json_schema`

=== "Python"

    ```python
    import requests

    api_url = "https://api.regolo.ai/v1/chat/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_REGOLO_KEY",
    }
    data = {
        "model": "Llama-3.3-70B-Instruct",
        "messages": [
            {"role": "user", "content": "Give me key facts about Rome."}
        ],
        "response_format": {
            "type": "json_schema",
            "json_schema": {
                "name": "city_facts",
                "strict": True,
                "schema": {
                    "type": "object",
                    "properties": {
                        "name": {"type": "string"},
                        "population": {"type": "integer"},
                        "landlocked": {"type": "boolean"}
                    },
                    "required": ["name", "population", "landlocked"],
                    "additionalProperties": False
                }
            }
        }
    }

    response = requests.post(api_url, headers=headers, json=data)
    result = response.json()
    print(result["choices"][0]["message"]["content"])
    ```

=== "cURL"

    ```bash
    curl -X POST https://api.regolo.ai/v1/chat/completions \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer YOUR_REGOLO_KEY" \
        -d '{
         "model": "Llama-3.3-70B-Instruct",
         "messages": [
            {"role": "user", "content": "Give me key facts about Rome."}
         ],
         "response_format": {
            "type": "json_schema",
            "json_schema": {
                "name": "city_facts",
                "strict": true,
                "schema": {
                    "type": "object",
                    "properties": {
                        "name": {"type": "string"},
                        "population": {"type": "integer"},
                        "landlocked": {"type": "boolean"}
                    },
                    "required": ["name", "population", "landlocked"],
                    "additionalProperties": false
                }
            }
         }
    }'
    ```

### Example: `json_object`

When you only need valid JSON and don't want to lock the shape, use `json_object`. Tell the model in the prompt what structure to produce — with `json_object` the schema is a suggestion, not a contract.

=== "Python"

    ```python
    data = {
        "model": "Llama-3.3-70B-Instruct",
        "messages": [
            {"role": "user", "content": "Return a JSON object with the name and population of Rome."}
        ],
        "response_format": {"type": "json_object"}
    }
    ```

=== "cURL"

    ```bash
    curl -X POST https://api.regolo.ai/v1/chat/completions \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer YOUR_REGOLO_KEY" \
        -d '{
         "model": "Llama-3.3-70B-Instruct",
         "messages": [
            {"role": "user", "content": "Return a JSON object with the name and population of Rome."}
         ],
         "response_format": {"type": "json_object"}
    }'
    ```

## With reasoning models

Structured outputs and reasoning compose cleanly. When you send `response_format` to a reasoning model such as `gpt-oss-120b`, the model still reasons freely in the `reasoning_content` field — only the final answer in `content` is constrained to your schema.

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Work out the area of a 7x5 rectangle and return it as JSON."}
    ],
    "reasoning_effort": "high",
    "response_format": {
        "type": "json_schema",
        "json_schema": {
            "name": "rectangle_area",
            "strict": True,
            "schema": {
                "type": "object",
                "properties": {
                    "area": {"type": "integer"}
                },
                "required": ["area"],
                "additionalProperties": False
            }
        }
    }
}
```

!!! note "Reasoning is not disabled"
    `response_format` constrains the final answer only. The model still emits reasoning tokens in `reasoning_content`, and those tokens are still billed as output. See [Reasoning](reasoning.md).

!!! tip "The `thinking` parameter"
    If your client sends `thinking: true`, Regolo normalizes it to `reasoning_effort`. It is an alias, not a separate mode — reasoning stays on. See [Reasoning](reasoning.md).

## Notes

- `response_format` is passed through to the underlying engine. Schema enforcement depends on the model; most instruction-tuned models honor `json_schema` strictly, but a few only honor `json_object`.
- Set `strict: true` for a hard contract. Some models ignore `strict` and enforce the schema regardless.
- Keep the prompt and schema aligned. If they contradict each other, the model can still fail to satisfy the schema.

For the exhaustive API endpoint reference visit [docs.api.regolo.ai](https://docs.api.regolo.ai).
