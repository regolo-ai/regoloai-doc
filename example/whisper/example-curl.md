```sh
curl https://api.regolo.ai/v1/models/whisper-large-v3/transcriptions \
  -H "Authorization: Bearer $REGOLOAI_API_KEY" \
  -H "Content-Type: multipart/form-data" \
  -F model="whisper-1" \
  -F file="@file.mp3"

