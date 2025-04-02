# Embedding

The embedding API allows you to get a vector representation of the input to be used from machine learning models or algorithms, leveraging models like `gte-Qwen2`.

## API Call Parameters

* `input`: A string describing the sentence, such as "A white cat resting in Rome."
* `model`: The identifier for the model used in image generation, e.g., "gte-Qwen2."

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
        "prompt": "A white cat resting in Rome",
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

For the exhaustive API's endpoints documentation visit [docs.api.regolo.ai](https://docs.api.regolo.ai).
