# Choose Your Language

Regolo AI is fully compatible with the OpenAI API, so you can use any OpenAI-compatible client library. Below are setup instructions for the most popular options.

## Python

### Using Regolo Client (Recommended)

The official Regolo Python client provides a simple, Pythonic interface.

```bash
pip install regolo
```

```python
import regolo

regolo.default_key = "YOUR_API_KEY"
regolo.default_chat_model = "Llama-3.3-70B-Instruct"

# Simple completion
response = regolo.static_chat_completions(
    messages=[{"role": "user", "content": "Hello!"}]
)
print(response)

# Using RegoloClient for chat sessions
client = regolo.RegoloClient()
client.add_prompt_to_chat(role="user", prompt="Tell me a joke")
role, content = client.run_chat()
print(content)
```

### Using OpenAI Client

You can also use the official OpenAI Python library:

```bash
pip install openai
```

```python
from openai import OpenAI

client = OpenAI(
    api_key="YOUR_API_KEY",
    base_url="https://api.regolo.ai/v1"
)

response = client.chat.completions.create(
    model="Llama-3.3-70B-Instruct",
    messages=[{"role": "user", "content": "Hello!"}]
)
print(response.choices[0].message.content)
```

## Node.js

Use the official OpenAI Node.js library:

```bash
npm install openai
```

```javascript
import OpenAI from 'openai';

const client = new OpenAI({
    apiKey: 'YOUR_API_KEY',
    baseURL: 'https://api.regolo.ai/v1'
});

async function main() {
    const response = await client.chat.completions.create({
        model: 'Llama-3.3-70B-Instruct',
        messages: [{ role: 'user', content: 'Hello!' }]
    });
    console.log(response.choices[0].message.content);
}

main();
```

### With Streaming

```javascript
import OpenAI from 'openai';

const client = new OpenAI({
    apiKey: 'YOUR_API_KEY',
    baseURL: 'https://api.regolo.ai/v1'
});

async function streamResponse() {
    const stream = await client.chat.completions.create({
        model: 'Llama-3.3-70B-Instruct',
        messages: [{ role: 'user', content: 'Tell me a story' }],
        stream: true
    });

    for await (const chunk of stream) {
        process.stdout.write(chunk.choices[0]?.delta?.content || '');
    }
}

streamResponse();
```

## cURL

For quick testing or shell scripts, use cURL directly:

```bash
curl -X POST https://api.regolo.ai/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer YOUR_API_KEY" \
    -d '{
        "model": "Llama-3.3-70B-Instruct",
        "messages": [{"role": "user", "content": "Hello!"}]
    }'
```

### With jq for Pretty Output

```bash
curl -s -X POST https://api.regolo.ai/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer YOUR_API_KEY" \
    -d '{
        "model": "Llama-3.3-70B-Instruct",
        "messages": [{"role": "user", "content": "Hello!"}]
    }' | jq '.choices[0].message.content'
```

## Other Languages

Since Regolo AI is OpenAI-compatible, you can use any OpenAI client library by changing the base URL to `https://api.regolo.ai/v1`:

- **Go**: [sashabaranov/go-openai](https://github.com/sashabaranov/go-openai)
- **Ruby**: [alexrudall/ruby-openai](https://github.com/alexrudall/ruby-openai)
- **Java**: [TheoKanning/openai-java](https://github.com/TheoKanning/openai-java)
- **C#**: [betalgo/openai](https://github.com/betalgo/openai)
- **PHP**: [openai-php/client](https://github.com/openai-php/client)

## Environment Variables

For security, store your API key in environment variables:

=== "Linux/macOS"

    ```bash
    export REGOLO_API_KEY="your-api-key"
    ```

=== "Windows (PowerShell)"

    ```powershell
    $env:REGOLO_API_KEY = "your-api-key"
    ```

Then access it in your code:

=== "Python"

    ```python
    import os
    import regolo

    regolo.default_key = os.environ.get("REGOLO_API_KEY")
    ```

=== "Node.js"

    ```javascript
    const client = new OpenAI({
        apiKey: process.env.REGOLO_API_KEY,
        baseURL: 'https://api.regolo.ai/v1'
    });
    ```