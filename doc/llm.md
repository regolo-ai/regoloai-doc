# Regolo AI API - Chat Completions

This guide provides an example of how to use Regolo AI's chat completions API to generate text responses from a specified model.

## Endpoint

```
POST https://api.regolo.ai/v1/chat/completions
```

## Headers

- `Content-Type: application/json`
- `Accept: application/json`
- `Authorization: Bearer YOUR_TOKEN`

## Request Body

- `model`: Specifies the model to use. Example: `"llama3.1:70b-instruct-q8_0"`. A list of available models can be found [here](https://regolo.ai/models.json).
- `messages`: An array of message objects, each containing:
- `role`: The role of the message sender. Use `"system"` for system instructions, `"user"` for user inputs, and `"assistant"` for the model's responses.
- `content`: The text input or instruction for the model to process.

### Example Request

REGOLO_TOKEN=YOUR_TOKEN

curl --location https://api.regolo.ai/v1/chat/completions  \
     --header 'Content-Type: application/json' \
     --header 'Accept: application/json' \
     --header "Authorization: Bearer ${REGOLO_TOKEN}" \
     --data   '{
               "model": "llama3.1:70b-instruct-q8_0",
               "messages": [
                 {
                   "role": "system",
                   "content": "You are a helpful assistant."
                 },
                 {
                   "role": "user",
                   "content": "Tell me about Rome"
                 }
               ]
              }'


NOTE: Make sure to replace `YOUR_TOKEN` with your actual API token
