import requests

url = "https://api.regolo.ai/v1/models/whisper-large-v3/transcriptions"
headers = {"Authorization": "Bearer $REGOLOAI_API_KEY"}

files = {
            "model": (None, "whisper-1"),
                "file": ("file.mp3", open("file.mp3", "rb"))
                }

response = requests.post(url, headers=headers, files=files)
print(response.json())
