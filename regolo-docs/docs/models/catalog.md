# Model Catalog

Regolo keeps every model it serves in a single catalog. You can query it from your code to see which models are available, what each one supports, how much it costs, and what its limits are.

Two endpoints cover this, each good for a different job:

- **`GET /models`** is the lightweight, OpenAI-compatible list. It hands back the model IDs you can pass as `model`. Use it when you only need to know what to put in a request.
- **`GET /model_group/info`** is the full catalog: capabilities, pricing, token limits, descriptions, and the parameters each model accepts. Use it when you need to compare models or surface cost and features in your own app.

## List of model IDs

The simplest call returns the IDs of every model you can use:

=== "Python"

    ```python
    import requests

    response = requests.get("https://api.regolo.ai/models")
    print(response.json())
    ```

=== "cURL"

    ```bash
    curl -X GET https://api.regolo.ai/models
    ```

## Full catalog: `/model_group/info`

This endpoint returns one object per model, inside a `data` array, each describing what the model does and how it's priced. Pass your API key in the request.

=== "Python"

    ```python
    import requests

    url = "https://api.regolo.ai/model_group/info"
    headers = {"Authorization": "Bearer YOUR_REGOLO_KEY"}

    response = requests.get(url, headers=headers)
    models = response.json()["data"]
    print(f"{len(models)} models available")
    ```

=== "cURL"

    ```bash
    curl -X GET https://api.regolo.ai/model_group/info \
      -H "Authorization: Bearer YOUR_REGOLO_KEY"
    ```

### What the response looks like

Each entry is a single object. Stripped down to the parts that matter, one looks like this:

```json
{
  "data": [
    {
      "model_group": "gpt-oss-120b",
      "description": "Open-weight 117B Mixture-of-Experts model by OpenAI. Supports reasoning, tool use, and more.",
      "mode": "chat",
      "providers": ["hosted_vllm"],
      "quantization": "MXFP4",
      "max_input_tokens": 30000,
      "max_output_tokens": 90000,
      "supports_reasoning": true,
      "supports_function_calling": true,
      "supports_vision": false,
      "supports_web_search": true,
      "input_cost_per_token": 0.000001,
      "output_cost_per_token": 0.0000042,
      "supported_openai_params": ["temperature", "max_tokens", "tools", "reasoning_effort", "stream"]
    }
  ]
}
```

### Identity and limits

| Field | Meaning |
| --- | --- |
| `model_group` | The model name. This is the value you send as `model` in your requests. |
| `description` | A short text description of the model. |
| `mode` | What the model does. One of `chat`, `embedding`, `rerank`, `image_generation`, `ocr`, `audio_transcription`. These line up with the [model families](families/completions.md). |
| `providers` | The provider(s) serving the model under the hood. |
| `quantization` | Quantization in use where it matters, e.g. `FP16`, `FP8`, `MXFP4`. |
| `max_input_tokens` | Maximum input tokens per request. |
| `max_output_tokens` | Maximum output tokens per request. |
| `max_tokens` | Overall token budget where the model defines one. |
| `supported_openai_params` | The OpenAI request parameters the model accepts. |

### Capabilities

These boolean flags tell you what a model can do. A field that doesn't apply to a model comes back as `null`.

| Field | `true` when the model... |
| --- | --- |
| `supports_vision` | accepts image inputs |
| `supports_reasoning` | is a reasoning model (see [Reasoning](features/reasoning.md)) |
| `supports_function_calling` | supports tool and function calls |
| `supports_parallel_function_calling` | can run several tool calls in parallel |
| `supports_web_search` | can search the web |
| `supports_url_context` | can pull in context from URLs |
| `supports_video_understanding` | can process video |
| `supports_audio_input` | accepts audio input |

### Pricing

Costs are split across several fields, and each model uses the ones that match how it's billed. Fields that don't apply are `null`.

| Field | Billed for |
| --- | --- |
| `input_cost_per_token` / `output_cost_per_token` | each input and output token, for chat-style models |
| `input_cost_per_second` | each second of audio, for speech-to-text |
| `input_cost_per_pixel` | each pixel of a generated image |
| `input_cost_per_query` | each query, e.g. for rerank |
| `input_cost_per_request` | each request, e.g. for OCR or embeddings |
| `ocr_cost_per_page` | each page processed by OCR |
| `input_cost_per_token_batches` / `output_cost_per_token_batches` | the batch rate, when batch pricing exists |

### Reading the prices

The per-token prices are in EUR per **single token**, so multiply by one million to get the familiar "per 1M tokens" figure. For example, `input_cost_per_token: 0.000001` works out to **$1 per 1M input tokens**.

Non-chat models don't use per-token pricing. A transcription model uses `input_cost_per_second`, an image model uses `input_cost_per_pixel`, and rerank or OCR models use `input_cost_per_query` or `input_cost_per_request`. Reach for whichever field matches the model's `mode`.

### Example: pick a model from code

A common reason to call this endpoint is to choose a model programmatically. Here's how you'd list the chat models that support reasoning, cheapest output first:

```python
import requests

url = "https://api.regolo.ai/model_group/info"
headers = {"Authorization": "Bearer YOUR_REGOLO_KEY"}
models = requests.get(url, headers=headers).json()["data"]

reasoning_chat = [
    m for m in models
    if m["mode"] == "chat" and m.get("supports_reasoning")
]
reasoning_chat.sort(key=lambda m: m["output_cost_per_token"])

for m in reasoning_chat:
    price_per_m = m["output_cost_per_token"] * 1_000_000
    print(f'{m["model_group"]:<20} ${price_per_m:.2f} / 1M out tokens')
```
