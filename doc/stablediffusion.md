# Regolo AI API - Text2Image

This guide provides an example of how to use Regolo AI's Text2Image API to generate images from text descriptions using a specified model.

## Endpoint

```
POST https://api.regolo.ai/v1/models/stable-diffusion/generate
```

## Headers

- `Content-Type: application/json`
- `Authorization: Bearer YOUR_TOKEN`

## Request Body

- `data`: An array containing text descriptions for which the images are to be generated.

### Example Request


REGOLO_TOKEN=YOUR_TOKEN

curl -X POST \
     --location https://api.regolo.ai/v1/models/stable-diffusion/generate  \
     --header 'Content-Type: application/json' \
     --header "Authorization: Bearer ${REGOLO_TOKEN}" \
     --data   '{ "data": ["Cat play piano"] }'
```

## Response

The API will return a base64 encoded data of the generated image.
