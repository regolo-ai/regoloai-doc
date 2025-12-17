# Rerank

The **Rerank** API lets you re‑order a list of documents (or passages) according to their relevance to a given query.  
It is powered by models such as **`Qwen3‑Reranker‑4B`** and returns the top‑N most relevant documents together with a relevance score.

**When to use it** – After you have retrieved a large set of candidate documents (e.g., with BM25, vector search, or another LLM), feed them to the Rerank endpoint to obtain a concise, high‑quality ranking that can be directly presented to users or passed to a downstream LLM for answer generation.

---

## API Call Parameters

!!! tip

    **Important note** – The endpoint expects a **JSON** payload and the total request size (including all documents) must stay lighter as possible. If you have many candidates, consider chunking them and calling the API multiple times.

| Parameter   | Type                           | Description |
|-------------|--------------------------------|-------------|
| **model**   | `string` (required)            | Identifier of the reranker model, e.g., `Qwen3‑Reranker‑4B`. |
| **query**   | `string` (required)            | The user’s question or search query. |
| **documents** | `array[string]` (required)   | List of candidate documents/passages to be reranked. |
| **top_n**   | `integer` (optional) | Number of highest‑scoring documents to return. |

---

=== "CURL"

    ```bash
    curl --request POST \
      --url https://api.regolo.ai/rerank \
      --header 'Authorization: Bearer REGOLO-API-KEY' \
      --header 'Content-Type: application/json' \
      --data '{
        "model": "Qwen3-Reranker-4B",
        "query": "What is the capital of the United States?",
        "documents": [
          "Carson City is the capital city of the American state of Nevada.",
          "The Commonwealth of the Northern Mariana Islands is a group of islands in the Pacific Ocean. Its capital is Saipan.",
          "Washington, D.C. is the capital of the United States.",
          "Capital punishment has existed in the United States since before it was a country."
        ],
        "top_n": 3
    }'
    ```

=== "Python"

    ```python
    import json
    import requests

    API_KEY = "REGOLO-API-KEY"
    ENDPOINT = "https://api.regolo.ai/rerank"

    payload = {
        "model": "Qwen3-Reranker-4B",
        "query": "What is the capital of the United States?",
        "documents": [
            "Carson City is the capital city of the American state of Nevada.",
            "The Commonwealth of the Northern Mariana Islands is a group of islands in the Pacific Ocean. Its capital is Saipan.",
            "Washington, D.C. is the capital of the United States.",
            "Capital punishment has existed in the United States since before it was a country."
        ],
        "top_n": 3
    }

    headers = {
        "Authorization": f"Bearer {API_KEY}",
        "Content-Type": "application/json"
    }

    response = requests.post(ENDPOINT, headers=headers, data=json.dumps(payload))

    if response.ok:
        results = response.json()
        print("Top documents:")
        for doc in results.get("results", []):
            print(f"- score: {doc['score']:.4f}  →  {doc['document']}")
    else:
        print("Error:", response.status_code, response.text)
    ```

## Response

```json
{
  "id": "rerank-906c2c4ec189b5fe",
  "results": [
    {
      "index": 2,
      "relevance_score": 0.9892732501029968,
      "document": {
        "text": "Washington, D.C. is the capital of the United States."
      }
    },
    {
      "index": 3,
      "relevance_score": 0.425626623916626,
      "document": {
        "text": "Capital punishment has existed in the United States since before it was a country."
      }
    },
    {
      "index": 0,
      "relevance_score": 0.4123265105247498,
      "document": {
        "text": "Carson City is the capital city of the American state of Nevada."
      }
    }
  ],
  "meta": null
}
```
