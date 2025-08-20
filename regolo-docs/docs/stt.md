# Speech To Text

The image generation API allows you to create images based on textual descriptions, leveraging models like `FLUX.1-dev`.

## API Call Parameters

* `file`: A binary file in ogg format.
* `model`: The identifier for the model used in image generation, e.g., "faster-whisper-large-v3".
* `language`: A string defining the language, like `english`, `italian` etc.

Consider that this models have a timeout so is better to split the audio in little sequences, like five minutes.

=== "CURL"

    ```bash
    curl --request POST \
      --url 'https://api.regolo.ai/v1/audio/transcriptions' \
      --header 'Authorization: Bearer YOUR_REGOLO_KEY' \
      --header "Content-Type: multipart/form-data"  
      --data '{
        "file": @"/path/of/your/file",
        "model": "faster-whisper-large-v3"
    }

    ```

For the exhaustive API's endpoints documentation visit [docs.api.regolo.ai](https://docs.api.regolo.ai).
