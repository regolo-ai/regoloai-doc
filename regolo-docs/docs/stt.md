# Speech To Text

The Speech to Text (STT) API enables you to extract and transcribe text from audio files using models such as `faster-whisper-large-v3`.
**We recommend using audio chunks of less than 2 minutes to prevent hallucinations and duplicate transcriptions.**

## API Call Parameters

- `file`: A binary audio file in OGG format.
- `model`: The identifier for the model used for transcription, e.g., `faster-whisper-large-v3`.
- `language`: A string specifying the language of the audio, such as `english`, `italian`, etc.

### Important Note
The models have a timeout limit. It is recommended to split audio files into smaller segments, such as five-minute clips, to ensure optimal performance.

## Example CURL Request

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

### Example Implementation

For a practical example of how to use this API, you can refer to the [Telegram Transcriber GitHub Repository](https://github.com/regolo-ai/TelegramTranscriber/). This repository provides a complete implementation for transcribing audio messages from Telegram using the Speech to Text API.

For the exhaustive API's endpoints documentation visit [docs.api.regolo.ai](https://docs.api.regolo.ai).
