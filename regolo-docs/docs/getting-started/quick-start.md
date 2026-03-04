# Quick Start (< 5 min)

Get up and running with Regolo AI in under 5 minutes.

## 1. Get Your API Key

Sign up at [dashboard.regolo.ai](https://dashboard.regolo.ai) and create a new API key in the **Virtual Keys** section.

## 2. Install the Client

=== "Python"

    ```bash
    pip install regolo
    ```

=== "Node.js"

    ```bash
    npm install openai
    ```

## 3. Make Your First Call

=== "Python"

    ```python
    import regolo

    regolo.default_key = "YOUR_API_KEY"
    regolo.default_chat_model = "Llama-3.3-70B-Instruct"

    response = regolo.static_chat_completions(
        messages=[{"role": "user", "content": "Hello, Regolo!"}]
    )
    print(response)
    ```

=== "Node.js"

    ```javascript
    import OpenAI from 'openai';

    const client = new OpenAI({
        apiKey: 'YOUR_API_KEY',
        baseURL: 'https://api.regolo.ai/v1'
    });

    const response = await client.chat.completions.create({
        model: 'Llama-3.3-70B-Instruct',
        messages: [{ role: 'user', content: 'Hello, Regolo!' }]
    });
    console.log(response.choices[0].message.content);
    ```

=== "cURL"

    ```bash
    curl -X POST https://api.regolo.ai/v1/chat/completions \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer YOUR_API_KEY" \
        -d '{
            "model": "Llama-3.3-70B-Instruct",
            "messages": [{"role": "user", "content": "Hello, Regolo!"}]
        }'
    ```

## Next Steps

- Explore [available models](../core-features/inference-api/completions-and-chat.md)
- Learn about [response parameters](../core-features/advanced/response-parameters.md)
- Check the [API documentation](https://docs.api.regolo.ai)