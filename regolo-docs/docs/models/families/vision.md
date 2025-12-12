# Vision models

Vision completions enable the processing of images alongside text, allowing for a wide range of applications such as image description, object recognition, and data extraction from visual content. By sending a combination of text prompts and image URLs, the model can provide insightful responses based on the visual input.

## Note

**This API supports only images, PDF files are not supported.**

## Vision Completions

You can provide images in three ways:

* [**Remote URL**](#remote-image-url): Supply a publicly accessible URL pointing to the image.
* [**Base64 (Image from PATH)**](#base64-image-from-path): Read a local image file, encode it as Base64, and send it to the model.
* [**Base64 Encoding (Remote URL)**](#base64-encoding-remote-url): Download an image from a URL, encode it as Base64, and send it to the model.

### Remote Image URL

This section demonstrates how to use a remote image URL directly with the vision model.

=== "Using Regolo Client"
    ```python
    import regolo

    regolo.default_key = "YOUR_REGOLO_KEY"
    regolo.default_chat_model = "qwen3-vl-32b"

    print(regolo.static_chat_completions(messages=[{
        "role": "user",
        "content": [
            {
                "type": "text",
                "text": "Describe this image in detail."
            },
            {
                "type": "image_url",
                "image_url": {
                    "url": "https://cdn.britannica.com/16/234216-050-C66F8665/beagle-hound-dog.jpg",
                    "format": "image/jpeg"
                }
            }
        ]
    }]))
    ```

=== "Python"
    ```python
    import requests

    url = "https://api.regolo.ai/v1/chat/completions"
    payload = {
        "model": "qwen3-vl-32b",
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "Describe this image in detail."
                    },
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": "https://cdn.britannica.com/16/234216-050-C66F8665/beagle-hound-dog.jpg",
                            "format": "image/jpeg"
                        }
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
        "model": "qwen3-vl-32b",
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "Describe this image in detail."
                    },
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": "https://cdn.britannica.com/16/234216-050-C66F8665/beagle-hound-dog.jpg",
                            "format": "image/jpeg"
                        }
                    }
                ]
            }
        ]
    }'
    ```

### Base64 (Image from PATH)

This section demonstrates how to read a local image file from your filesystem, encode it as Base64, and send it to the vision model.

=== "Using Regolo Client"
    ```python
    import regolo
    import base64
    from pathlib import Path

    regolo.default_key = "YOUR_REGOLO_KEY"
    regolo.default_chat_model = "qwen3-vl-32b"

    IMAGE_PATH = Path("beagle-hound-dog.jpg")

    if not IMAGE_PATH.exists():
        raise FileNotFoundError(f"Image not found: {IMAGE_PATH.resolve()}")

    with open(IMAGE_PATH, "rb") as f:
        image_bytes = f.read()
    
    image_b64 = base64.b64encode(image_bytes).decode("utf-8")

    result = regolo.static_chat_completions(
        messages=[{
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": "Describe this image in detail."
                },
                {
                    "type": "image_url",
                    "image_url": {
                        "url": f"data:image/jpeg;base64,{image_b64}",
                        "format": "image/jpeg"
                    }
                }
            ]
        }],
        max_tokens=4096,  # Important: maintain a wide output token window to avoid issues
        full_output=True
    )

    print("Model response:")
    print(result["choices"][0]["message"]["content"])
    ```

=== "Python"
    ```python
    import base64
    import json
    import requests
    from pathlib import Path

    API_URL = "https://api.regolo.ai/v1/chat/completions"
    API_KEY = "YOUR-API-KEY"
    MODEL = "qwen3-vl-32b"

    IMAGE_PATH = Path("beagle-hound-dog.jpg")

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
                    {"type": "text", "text": "Describe this image in detail."},
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": f"data:image/jpeg;base64,{image_b64}",
                            "format": "image/jpeg"
                        }
                    }
                ]
            }
        ],
        "max_tokens": 4096  # Important: maintain a wide output token window to avoid issues
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

=== "CURL"
    ```bash
    # Encode local image to base64
    IMAGE_B64=$(base64 -w 0 beagle-hound-dog.jpg)

    curl -X POST https://api.regolo.ai/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer YOUR-API-KEY" \
    -d "{
        \"model\": \"qwen3-vl-32b\",
        \"messages\": [
            {
                \"role\": \"user\",
                \"content\": [
                    {
                        \"type\": \"text\",
                        \"text\": \"Describe this image in detail.\"
                    },
                    {
                        \"type\": \"image_url\",
                        \"image_url\": {
                            \"url\": \"data:image/jpeg;base64,$IMAGE_B64\",
                            \"format\": \"image/jpeg\"
                        }
                    }
                ]
            }
        ],
        \"max_tokens\": 4096  # Important: maintain a wide output token window to avoid issues
    }"
    ```

### Base64 Encoding (Remote URL)

This section demonstrates how to download an image from a remote URL, encode it as Base64, and send it to the vision model.

=== "Using Regolo Client"
    ```python
    import regolo
    import base64
    import requests

    regolo.default_key = "YOUR_REGOLO_KEY"
    regolo.default_chat_model = "qwen3-vl-32b"

    # Download image and convert to base64
    image_url = "https://cdn.britannica.com/16/234216-050-C66F8665/beagle-hound-dog.jpg"
    response = requests.get(image_url)
    image_b64 = base64.b64encode(response.content).decode('utf-8')

    result = regolo.static_chat_completions(
        messages=[{
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": "Describe this image in detail."
                },
                {
                    "type": "image_url",
                    "image_url": {
                        "url": f"data:image/jpeg;base64,{image_b64}",
                        "format": "image/jpeg"
                    }
                }
            ]
        }],
        temperature=0.2,
        max_tokens=4096,  # Important: maintain a wide output token window to avoid issues
        full_output=True
    )

    print("Model response:")
    print(result["choices"][0]["message"]["content"])
    ```

=== "Python"
    ```python
    import base64
    import requests
    import json

    # Regolo API credentials configuration
    api_key = "YOUR_API_KEY"
    model = "qwen3-vl-32b"  # Vision-compatible model

    # Download image and convert to base64
    image_url = "https://cdn.britannica.com/16/234216-050-C66F8665/beagle-hound-dog.jpg"
    response = requests.get(image_url)
    image_b64 = base64.b64encode(response.content).decode('utf-8')

    # Direct API call for multimodal messages
    url = "https://api.regolo.ai/v1/chat/completions"
    payload = {
        "model": model,
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "Describe this image in detail."
                    },
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": f"data:image/jpeg;base64,{image_b64}",
                            "format": "image/jpeg"
                        }
                    }
                ]
            }
        ],
        "temperature": 0.2,
        "max_tokens": 4096  # Important: maintain a wide output token window to avoid issues
    }
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {api_key}"
    }

    print("Sending request to Regolo AI API...")
    response = requests.post(url, json=payload, headers=headers)

    if response.status_code == 200:
        result = response.json()
        try:
            content = result["choices"][0]["message"]["content"]
            print("Model response:")
            print(content)
        except KeyError as e:
            print(f"KeyError: {e}")
            print("Response structure:")
            print(json.dumps(result, indent=2))
    else:
        print(f"Error {response.status_code}:")
        print(response.text)
    ```

=== "CURL"
    ```bash
    # First, download the image and encode it to base64
    # This example assumes you have the image URL
    IMAGE_URL="https://cdn.britannica.com/16/234216-050-C66F8665/beagle-hound-dog.jpg"
    IMAGE_B64=$(curl -s "$IMAGE_URL" | base64 -w 0)

    curl -X POST https://api.regolo.ai/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer YOUR_API_KEY" \
    -d "{
        \"model\": \"qwen3-vl-32b\",
        \"messages\": [
            {
                \"role\": \"user\",
                \"content\": [
                    {
                        \"type\": \"text\",
                        \"text\": \"Describe this image in detail.\"
                    },
                    {
                        \"type\": \"image_url\",
                        \"image_url\": {
                            \"url\": \"data:image/jpeg;base64,$IMAGE_B64\",
                            \"format\": \"image/jpeg\"
                        }
                    }
                ]
            }
        ],
        \"temperature\": 0.2,
        \"max_tokens\": 4096  # Important: maintain a wide output token window to avoid issues
    }"
    ```
