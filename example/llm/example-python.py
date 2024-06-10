import requests

url = "$ENDPOINT"
headers = {
            "Content-Type": "application/json",
                "Accept": "application/json",
                    "Authorization": f"Bearer {os.getenv('REGOLO_TOKEN')}"
                    }
data = {
            "model": "mistralai/Mistral-7B-Instruct-v0.2",
                "messages": [{"role": "user", "content": "Tell me about Rome in a concise manner"}]
                }

response = requests.post(url, headers=headers, json=data)
print(response.json())
