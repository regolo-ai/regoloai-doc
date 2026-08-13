# Embedding

The embedding API allows you to get a vector representation of the input to be used from machine learning models or algorithms, leveraging models like `gte-Qwen2`.

## API Call Parameters

* `input`: A string (or list of strings) to embed, such as "A white cat resting in Rome."
* `model`: The identifier for the embedding model, e.g., `gte-Qwen2` or `Qwen3-Embedding-8B`.
* `dimensions`: The length of the vector to return, e.g. `1024`. Regolo passes this straight to the model with no server-side truncation or renormalization. Whether a reduced dimension is supported — and how MRL truncation is applied — is up to the model.

=== "Using Regolo Client"

    ```python
    import regolo

    regolo.default_key = "<YOUR_REGOLO_KEY>"
    regolo.default_embedder_model = "gte-Qwen2"


    embeddings = regolo.static_embeddings(input_text=["A white cat resting in Rome", "A white cat resting in Paris"])

    print(embeddings)
    ```

=== "Python"

    ```python
    import requests
    import json

    url = 'https://api.regolo.ai/v1/embeddings'
    headers = {
        'Authorization': 'Bearer YOUR_REGOLO_KEY',
        'Content-Type': 'application/json'
    }

    data = {
        "input": "A white cat resting in Rome",
        "model": "gte-Qwen2",
    }

    response = requests.post(url, headers=headers, data=json.dumps(data))

    if response.status_code == 200:
        with open("./embedding.json", 'w') as _file:
            json.dump(response.json(), _file)
    else:
        print("Failed embedding request:", response.status_code, response.text)

    ```

=== "CURL"

    ```bash
    curl -X POST https://api.regolo.ai/v1/embeddings
    -H "Content-Type: application/json"
    -H "Authorization: Bearer YOUR_REGOLO_KEY"
    -d '{
        "model": "gte-Qwen2",
        "input": "The quick brown fox jumps over the lazy dog"
    }'
    ```

## Reduced dimensions

Some embedding models can return shorter vectors directly, so you save storage and speed up similarity search without truncating on the client. Pass `dimensions` and the model does the work.

`dimensions` is forwarded to the model as-is. Regolo does not truncate or renormalize the vector — if the model applies Matryoshka Representation Learning (MRL) and renormalizes, that happens inside the model. If the model does not support a reduced dimension, it ignores the value or returns an error.

=== "Python"

    ```python
    import requests

    url = 'https://api.regolo.ai/v1/embeddings'
    headers = {
        'Authorization': 'Bearer YOUR_REGOLO_KEY',
        'Content-Type': 'application/json'
    }

    data = {
        "model": "Qwen3-Embedding-8B",
        "input": "A white cat resting in Rome",
        "dimensions": 1024
    }

    response = requests.post(url, headers=headers, json=data)
    result = response.json()
    print(len(result["data"][0]["embedding"]))  # 1024
    ```

=== "CURL"

    ```bash
    curl -X POST https://api.regolo.ai/v1/embeddings
    -H "Content-Type: application/json"
    -H "Authorization: Bearer YOUR_REGOLO_KEY"
    -d '{
        "model": "Qwen3-Embedding-8B",
        "input": "A white cat resting in Rome",
        "dimensions": 1024
    }'
    ```

!!! note "Support is model-dependent"
    Not every embedding model accepts `dimensions`. `Qwen3-Embedding-8B` does; `gte-Qwen2` returns its fixed native dimension. Check the model in the [catalog](../catalog.md) before relying on a specific length.

For the exhaustive API's endpoints documentation visit [docs.api.regolo.ai](https://docs.api.regolo.ai).

