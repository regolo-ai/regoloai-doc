# Regolo AI API - Embeddings

This guide provides an example of how to use Regolo AI's embeddings API to generate embeddings from a specified model.

## Endpoint

```
POST https://api.regolo.ai/v1/embeddings
```

## Headers

- `Content-Type: application/json`
- `Authorization: Bearer YOUR_TOKEN`

## Request Body

- `model`: Specifies the embedding model to use. Example: `"Alibaba-NLP/gte-Qwen2-7B-instruct"`. A list of available models can be found [here](https://regolo.ai/models.json).
- `input`: The text string for which the embedding is to be generated.

### Example Request

REGOLO_TOKEN=YOUR_TOKEN

curl https://api.regolo.ai/v1/embeddings \
     -X POST \
     --header "Authorization: Bearer ${REGOLO_TOKEN}" \
     -H 'Content-Type: application/json' \
     -d '{
           "input": "Your text string goes here",
           "model": "Alibaba-NLP/gte-Qwen2-7B-instruct"
         }'
```

## Response

The API will return a JSON object containing the embeddings for the input text.


NOTE: Make sure to replace `YOUR_TOKEN` with your actual API token
