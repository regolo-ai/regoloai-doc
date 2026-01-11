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

## Technical Explanation of Prompt Tags

To achieve optimal performance with the **Qwen3-Reranker** model, the input strings must follow a specific prompt template. This structure allows the underlying Cross-Encoder to differentiate between instructions, user intent, and the content to be ranked.

| Tag                                                      | Purpose | Description |
|:---------------------------------------------------------| :--- | :--- |
| <code style="white-space:nowrap">&lt;Instruct&gt;</code> | **Task Specification** | Defines the context of the retrieval (e.g., "retrieve relevant passages"). This guides the model's focus based on the specific use case. |
| <code style="white-space:nowrap">&lt;Query&gt;</code>    | **User Intent** | Marks the beginning of the actual question or search term. This helps the model isolate the core information the user is looking for. |
| <code style="white-space:nowrap">&lt;Document&gt;</code> | **Content Boundary** | Identifies each candidate passage. By explicitly tagging documents, the model better understands where one passage ends and another begins during the scoring process. |

#### Why these tags are necessary
Modern LLM-based rerankers are trained using these semantic markers to improve **zero-shot accuracy**. Omitting these tags or using inconsistent formatting between the query and the documents can lead to sub-optimal relevance scores, as the model might fail to distinguish between the instruction and the data.

---

=== "CURL"

    ```bash
      curl --request POST \
        --url https://api.regolo.ai/rerank \
        --header 'Authorization: Bearer REGOLO_API_KEY' \
        --header 'Content-Type: application/json' \
        --data '{
          "model": "Qwen3-Reranker-4B",
          "query": "<Instruct>: Given a web search query, retrieve relevant passages that answer the query\n<Query>: What is the capital of China?",
          "documents": [
            "<Document>: The capital of China is Beijing.",
            "<Document>: Gravity is a force that attracts two bodies towards each other..."
          ],
          "top_n": 5
        }'
    ```

=== "Python"

    ```python
    import requests
    
    api_key = "REGOLO_API_KEY"
    url = "https://api.regolo.ai/rerank"
    
    task = "Given a web search query, retrieve relevant passages that answer the query"
    query_text = "What is the capital of China?"
    
    documents = [
        "The capital of China is Beijing.",
        "Gravity is a force that attracts two bodies towards each other..."
    ]
    
    payload = {
        "model": "Qwen3-Reranker-4B",
        "query": f"<Instruct>: {task}\n<Query>: {query_text}",
        "documents": [f"<Document>: {doc}" for doc in documents],
        "top_n": 5
    }
    
    response = requests.post(
        url,
        json=payload,
        headers={"Authorization": f"Bearer {api_key}"}
    )
    
    results = response.json().get('results', [])
    for res in results:
        score = res['relevance_score']
        clean_text = res['document']['text'].replace("<Document>: ", "")
        print(f"Score: {score:.4f} | Text: {clean_text}")
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