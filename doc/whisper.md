# Regolo AI API - Audio to Text

This guide provides an example of how to use Regolo AI's Audio to Text API to transcribe audio files into text using the Whisper model.

## Endpoint

```
POST https://api.regolo.ai/v1/models/whisper-large-v3/transcriptions
```

## Headers

- `Content-Type: multipart/form-data`
- `Authorization: Bearer YOUR_TOKEN`

## Request Body

- `model`: Specifies the transcription model to use. Example: `"whisper-1"`.
- `file`: The audio file to be transcribed, provided as a file upload.

### Example Request

ENDPOINT=https://api.regolo.ai/v1/models/whisper-large-v3/transcriptions

curl --location https://api.regolo.ai/v1/models/whisper-large-v3/transcriptions  \
     --header "Authorization: Bearer ${REGOLO_TOKEN}" \
     -H "Content-Type: multipart/form-data" \
     -F model="whisper-1" \
     -F file="@demo.mp3"
```

## Response

The API will return a JSON object containing the transcribed text from the audio file.

