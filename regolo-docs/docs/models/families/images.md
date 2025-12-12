# Images Generation

The image generation API allows you to create images based on textual descriptions, leveraging models like `Qwen-Image`.

## API Call Parameters

* `prompt`: A string describing the desired image, such as "A white cat resting in Rome."
* `n`: An integer specifying the number of images to generate. Generating more images increases response time, so it's best to keep this number small for faster performance.
* `model`: The identifier for the model used in image generation, e.g., "Qwen-Image."
* `size`: A string defining the dimensions of the images. Supported sizes are "256x256," "512x512," and "1024x1024."

Larger images take longer to generate, so consider using smaller sizes for quicker results.

!!! tip
    If you require larger images, consider using an image upscaler after generation. This can help achieve the desired resolution without increasing the generation time

=== "Using Regolo Client"

    ```python
    import regolo
    from io import BytesIO
    from PIL import Image

    # pip install regolo Pillow

    regolo.default_image_generation_model = "Qwen-Image"
    regolo.default_key = "YOUR_REGOLO_KEY"

    img_bytes = regolo.static_image_create(prompt="A Boat in the sea")[0]

    image = Image.open(BytesIO(img_bytes))

    # Save the Image
    output_path = "generated_image.png"
    image.save(output_path)
    print(f"Image saved to: {output_path}")
    ```

=== "Python"

    ```python
    import requests
    import json
    from PIL import Image
    import io
    import base64

    url = 'https://api.regolo.ai/v1/images/generations'
    headers = {
        'Authorization': 'Bearer YOUR_REGOLO_KEY',
        'Content-Type': 'application/json'
    }

    data = {
        "prompt": "A white cat resting in Rome",
        "n": 2,
        "model": "Qwen-Image",
        "size": "1024x1024"
    }

    response = requests.post(url, headers=headers, data=json.dumps(data))

    if response.status_code == 200:
        response_data = response.json()
        
        for index, item in enumerate(response_data['data']):
            b64_image = item['b64_json']
            image_data = base64.b64decode(b64_image)

            image_stream = io.BytesIO(image_data)
            image = Image.open(image_stream)

            # Save the Image
            output_path = f"generated_image_{index + 1}.png"
            image.save(output_path)
            print(f"Image saved to: {output_path}")
    else:
        print("Failed to generate images:", response.status_code, response.text)

    ```

=== "CURL"

    ```bash
    curl --request POST \
      --url 'https://api.regolo.ai/v1/images/generations' \
      --header 'Authorization: Bearer YOUR_REGOLO_KEY' \
      --header 'Content-Type: application/json' \
      --data '{
        "prompt": "A Boat in the sea",
        "n": 2,
        "model": "Qwen-Image",
        "size": "1024x1024"
    }' | python3 -c "
    import sys
    import json
    import base64
    
    response = json.load(sys.stdin)
    if 'data' in response:
        for index, item in enumerate(response['data']):
            b64_image = item['b64_json']
            image_data = base64.b64decode(b64_image)
            output_path = f'generated_image_{index + 1}.png'
            with open(output_path, 'wb') as f:
                f.write(image_data)
            print(f'Image saved to: {output_path}')
    else:
        print('Failed to generate images:', response)
    "
    ```

For the exhaustive API's endpoints documentation visit [docs.api.regolo.ai](https://docs.api.regolo.ai).

