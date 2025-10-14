# Vision models

Vision completions enable the processing of images alongside text, allowing for a wide range of applications such as image description, object recognition, and data extraction from visual content. By sending a combination of text prompts and image URLs, the model can provide insightful responses based on the visual input.

## Vision Completions

You can provide images in two ways:

* [**Remote URL**](#remote-image-url): Supply a publicly accessible URL pointing to the image.
* [**Base64 Encoding**](#base64-encoding): Encode the image in base64 format and pass it in the `image_url` field.

### Remote Image URL

=== "Using Regolo Client"
    ```python
    import regolo

    regolo.default_key = "YOUR_REGOLO_KEY"
    regolo.default_model = "Qwen2.5-VL-32B-Instruct"

    print(regolo.static_chat_completions(messages=[{
        "role": "user",
        "content": [
            {
                "type": "text",
                "text": "What’s in this image?"
            },
            {
                "type": "image_url",
                "image_url": {"url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/de/Colosseo_2020.jpg/960px-Colosseo_2020.jpg"},
                "format": "image/jpeg"
            }
        ]
    }]))
    ```

=== "Python"
    ```python
    import requests

    url = "https://api.regolo.ai/v1/chat/completions"
    payload = {
        "model": "Qwen2.5-VL-32B-Instruct",
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "What’s in this image?"
                    },
                    {
                        "type": "image_url",
                        "image_url": {"url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/de/Colosseo_2020.jpg/960px-Colosseo_2020.jpg"},
                        "format": "image/jpeg"
                    }
                ]
            }
        ]
    }
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_REGOLO_KEY"
    }

    response = requests.post(url, json=payload, headers=headers)
    print(response.json())
    ```


=== "CURL"
    ```bash
    curl -X POST https://api.regolo.ai/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer YOUR_REGOLO_KEY" \
    -d '{
        "model": "Qwen2.5-VL-32B-Instruct",
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "What’s in this image?"
                    },
                    {
                        "type": "image_url",
                        "image_url": {"url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/de/Colosseo_2020.jpg/960px-Colosseo_2020.jpg"},
                        "format": "image/jpeg"
                    }
                ]
            }
        ]
    }'
    ```

### Base64 Encoding

This script demonstrates how to encode a local image as Base64 and send it to a multimodal model (text + image) for analysis.

Replace 'YOUR-API-KEY' with your actual API key before running.

=== "Python"
    ```python
    import base64
    import json
    import requests
    from pathlib import Path

    API_URL = "https://api.regolo.ai/v1/chat/completions"
    API_KEY = "YOUR-API-KEY"
    MODEL = "gemma-3-27b-it"

    IMAGE_PATH = Path("colosseo.jpg")

    if not IMAGE_PATH.exists():
        raise FileNotFoundError(f"Image not found: {IMAGE_PATH.resolve()}")

    with open(IMAGE_PATH, "rb") as f:
        image_bytes = f.read()
    
    image_b64 = base64.b64encode(image_bytes).decode("utf-8")
    
    payload = {
        "model": MODEL,
        "messages": [
            {
                "role": "user",
                "content": [
                    {"type": "text", "text": "What’s in this image?"},
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": f"data:image/jpeg;base64,{image_b64}",
                            "format": "image/jpeg"
                        }
                    }
                ]
            }
        ]
    }
    
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {API_KEY}"
    }
    
    print("Sending request to Regolo AI API...")
    response = requests.post(API_URL, headers=headers, data=json.dumps(payload))
    
    if response.status_code != 200:
        print(f"Error {response.status_code}:")
        print(response.text)
    else:
        result = response.json()
        try:
            content = result["choices"][0]["message"]["content"]
            print("Model response:")
            print(content)
        except Exception:
            print("Unexpected response format:")
            print(response.text)
    ```