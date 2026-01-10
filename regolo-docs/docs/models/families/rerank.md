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
        --header 'Authorization: Bearer REGOLO_API_KEY' \
        --header 'Content-Type: application/json' \
        --data '{
          "model": "Qwen3-Reranker-4B",
          "query": "Given a web search query, retrieve relevant passages that answer the query\n<Query>: What is the capital of China?",
          "documents": [
            "<Document>: The capital of China is Beijing.",
            "<Document>: Gravity is a force that attracts two bodies towards each other..."
          ]
        }'
    ```

=== "Python"

    ```python
      import requests
      
      api_key = "REGOLO_API_KEY"
      url = "https://api.regolo.ai/rerank"
      
      task = 'Given a web search query, retrieve relevant passages that answer the query'
      queries = ["What is the capital of China?"]
      documents = [
          "The capital of China is Beijing.",
          "Gravity is a force that attracts two bodies towards each other..."
      ]
      
      payload = {
          "model": "Qwen3-Reranker-4B",
          "query": f"{task}\n<Query>: {queries[0]}",
          "documents": [f"{doc}" for doc in documents],
          "top_n": 5
      }
      
      response = requests.post(url, json=payload, headers={"Authorization": f"Bearer {api_key}"})
      
      for res in response.json()['results']:
         print(f"Score: {res['relevance_score']:.4f} | Text: {res['document']['text']}")
    ```

## Response

```json
{
  "id": "rerank-8b37844a3beeecb7",
  "results": [
    {
      "index": 0,
      "relevance_score": 0.8835278153419495,
      "document": {
        "text": "<Document>: The capital of China is Beijing."
      }
    },
    {
      "index": 1,
      "relevance_score": 0.08649543672800064,
      "document": {
        "text": "<Document>: Gravity is a force that attracts two bodies towards each other..."
      }
    }
  ],
  "meta": null
}
```
