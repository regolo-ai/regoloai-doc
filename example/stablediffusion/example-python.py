import requests
import os

url = os.getenv('ENDPOINT')
headers = {
            "Content-Type": "application/json",
                "Authorization": f"Bearer {os.getenv('REGOLO_TOKEN')}"
                }
data = { "data": ["Cat playing the piano"] }

response = requests.post(url, headers=headers, json=data)
print(response.json())
