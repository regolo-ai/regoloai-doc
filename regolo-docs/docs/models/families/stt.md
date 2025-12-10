# Speech To Text

The Speech to Text (STT) API enables you to extract and transcribe text from audio files using models such as `faster-whisper-large-v3`.  
**We recommend using audio chunks of less than 2 minutes to prevent hallucinations and duplicate transcriptions.**

## API Call Parameters

- `file`: A binary audio file in OGG format.
- `model`: The identifier for the model used for transcription, e.g., `faster-whisper-large-v3`.
- `language`: A two-letter ISO language code specifying the language of the audio, such as `en` (English), `it` (Italian), etc.

### Important Note
The models have a timeout limit. It is recommended to split audio files into smaller segments, such as five-minute clips, to ensure optimal performance.

## Example Requests

=== "Using Regolo Client"

    ```python
    import regolo
    from pathlib import Path

    # Regolo configuration
    regolo.default_key = "YOUR_REGOLO_KEY"
    regolo.default_audio_transcription_model = "faster-whisper-large-v3"

    # Audio file to transcribe
    AUDIO_FILE = "/path/to/your/audio"
    OUTPUT_FILE = "/path/to/output/transcription.txt"

    # Transcribe the file
    transcript = regolo.static_audio_transcription(file=AUDIO_FILE)

    # Save the transcription
    output_path = Path(OUTPUT_FILE)
    output_path.parent.mkdir(parents=True, exist_ok=True)

    with open(output_path, "w", encoding="utf-8") as f:
        f.write(transcript)

    print(f"Transcription saved to: {OUTPUT_FILE}")
    ```

=== "Python Client"

    ```python
    import openai
    from pathlib import Path

    # OpenAI client configuration
    openai.api_key = "YOUR_REGOLO_KEY"
    openai.base_url = "https://api.regolo.ai/v1/"

    # Audio file to transcribe
    AUDIO_FILE = "/path/to/your/audio"
    OUTPUT_FILE = "/path/to/output/transcription.txt"

    # Transcribe the file
    with open(AUDIO_FILE, "rb") as audio_file:
        transcript = openai.audio.transcriptions.create(
            model="faster-whisper-large-v3",
            file=audio_file,
            language="en",
            response_format="text"
        )

    # Save the transcription
    output_path = Path(OUTPUT_FILE)
    output_path.parent.mkdir(parents=True, exist_ok=True)

    with open(output_path, "w", encoding="utf-8") as f:
        f.write(transcript)

    print(f"Transcription saved to: {OUTPUT_FILE}")
    ```

=== "CURL"

    ```bash
    curl --request POST \
      --url 'https://api.regolo.ai/v1/audio/transcriptions' \
      --header 'Authorization: Bearer YOUR_REGOLO_KEY' \
      -F "file=@/path/to/your/audio" \
      -F "model=faster-whisper-large-v3"
    ```

### Example Implementation

For a practical example of how to use this API, you can refer to the [Telegram Transcriber GitHub Repository](https://github.com/regolo-ai/TelegramTranscriber/). This repository provides a complete implementation for transcribing audio messages from Telegram using the Speech to Text API.

For the exhaustive API's endpoints documentation visit [docs.api.regolo.ai](https://docs.api.regolo.ai).

