# Vision models

Vision completions enable the processing of images alongside text, allowing for a wide range of applications such as image description, object recognition, and data extraction from visual content. By sending a combination of text prompts and image URLs, the model can provide insightful responses based on the visual input.

## Vision Completions

You can provide images in two ways:

* **Remote URL**: Supply a publicly accessible URL pointing to the image.
* **Base64 Encoding**: Encode the image in base64 format and pass it in the `image_url` field.

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

