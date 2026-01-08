# OCR

DeepSeek OCR is a powerful optical character recognition model that enables accurate text extraction from images and documents. It supports various modes including native and dynamic resolutions, making it suitable for different use cases from simple OCR to complex document parsing.

## Deepseek-OCR Usage with Regolo API

Use DeepSeek OCR on the Regolo platform with model name: `"deepseek-ocr"`

### Python Examples

=== "Remote Image URL"

    ```python
    import requests

    url = "https://api.regolo.ai/v1/chat/completions"
    payload = {
        "model": "deepseek-ocr",
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "Convert the document to markdown."
                    },
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": "https://cdn.britannica.com/86/22386-050-51E63D13/Silicon-silicon-symbol-square-Si-properties-some.jpg",
                            "format": "image/png"
                        }
                    }
                ]
            }
        ],
        "skip_special_tokens": False
    }
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_REGOLO_KEY"
    }

    response = requests.post(url, json=payload, headers=headers)
    print(response.json())
    ```

=== "Base64 (Image from PATH)"

    ```python
    import base64
    import requests
    from pathlib import Path

    API_URL = "https://api.regolo.ai/v1/chat/completions"
    API_KEY = "YOUR-API-KEY"
    MODEL = "deepseek-ocr"

    IMAGE_PATH = Path("document.png")

    with open(IMAGE_PATH, "rb") as f:
        image_bytes = f.read()

    image_b64 = base64.b64encode(image_bytes).decode("utf-8")

    payload = {
        "model": MODEL,
        "messages": [
            {
                "role": "user",
                "content": [
                    {"type": "text", "text": "Convert the document to markdown."},
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": f"data:image/png;base64,{image_b64}",
                            "format": "image/png"
                        }
                    }
                ]
            }
        ],
        "max_tokens": 4096,
        "skip_special_tokens": False
    }

    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {API_KEY}"
    }

    response = requests.post(API_URL, headers=headers, json=payload)
    result = response.json()
    content = result["choices"][0]["message"]["content"]
    print(content)
    ```

=== "Base64 Encoding (Remote URL)"

    ```python
    import base64
    import requests

    api_key = "YOUR_API_KEY"
    model = "deepseek-ocr"

    image_url = "https://example.com/document.png"
    response = requests.get(image_url)
    image_b64 = base64.b64encode(response.content).decode('utf-8')

    url = "https://api.regolo.ai/v1/chat/completions"
    payload = {
        "model": model,
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "Convert the document to markdown."
                    },
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": f"data:image/png;base64,{image_b64}",
                            "format": "image/png"
                        }
                    }
                ]
            }
        ],
        "max_tokens": 4096,
        "skip_special_tokens": False
    }
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {api_key}"
    }

    response = requests.post(url, json=payload, headers=headers)
    result = response.json()
    content = result["choices"][0]["message"]["content"]
    print(content)
    ```

!!! note "Note on skip_special_tokens"
    The `skip_special_tokens` parameter controls whether special tokens (like `<|grounding|>`) are included in the response:
    - **`skip_special_tokens=False`**: Keeps special tokens (default). Use when you need document structure info.
    - **`skip_special_tokens=True`**: Removes special tokens. Use for clean text output only.

## Prompt Examples

```python

DeepSeek OCR supports various prompt formats for different use cases:

# Convert the document contents to markdown format
<|grounding|>Convert the document to markdown.

# Perform text recognition on this image
<|grounding|>OCR this image.

# Extract all text without layout consideration
Free OCR.

# Parse any figures or tables in the document
Parse the figure.

# Provide a detailed description of the image content
Describe this image in detail.

# Locate the position of <|ref|>xxxx<|/ref|> in the image
Locate <|ref|>xxxx<|/ref|> in the image.
```

## PDF OCR Reader

This example demonstrates how to extract text from a PDF document by converting each page to an image and processing it through the OCR API. All extracted text is aggregated and saved to a markdown file.

```python
import base64
import requests
import fitz
from pathlib import Path

API_URL = "https://api.regolo.ai/v1/chat/completions"
API_KEY = "YOUR-API-KEY"
MODEL = "deepseek-ocr"

PDF_PATH = Path("document.pdf")
OUTPUT_PATH = PDF_PATH.with_suffix(".md")

doc = fitz.open(PDF_PATH)
all_text = []

for page_num in range(len(doc)):
    page = doc[page_num]
    pix = page.get_pixmap(matrix=fitz.Matrix(2, 2))
    img_bytes = pix.tobytes("png")
    
    image_b64 = base64.b64encode(img_bytes).decode("utf-8")
    
    payload = {
        "model": MODEL,
        "messages": [
            {
                "role": "user",
                "content": [
                    {"type": "text", "text": "Convert the document to markdown."},
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": f"data:image/png;base64,{image_b64}",
                            "format": "image/png"
                        }
                    }
                ]
            }
        ],
        "max_tokens": 4096,
        "skip_special_tokens": False
    }
    
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {API_KEY}"
    }
    
    response = requests.post(API_URL, headers=headers, json=payload)
    result = response.json()
    content = result["choices"][0]["message"]["content"]
    
    all_text.append(f"\n\n--- Page {page_num + 1} ---\n\n")
    all_text.append(content)

doc.close()

with open(OUTPUT_PATH, "w", encoding="utf-8") as f:
    f.write("".join(all_text))

print(f"Completed : file saved in \"{OUTPUT_PATH}\"")
```

## Resources

**Deepseek-OCR Links:**
- [GitHub Repository](https://github.com/deepseek-ai/DeepSeek-OCR)
- [Hugging Face](https://huggingface.co/deepseek-ai/DeepSeek-OCR)
- [Arxiv Paper](https://arxiv.org/abs/2510.18234)

For the exhaustive API's endpoints documentation visit [docs.api.regolo.ai](https://docs.api.regolo.ai).
