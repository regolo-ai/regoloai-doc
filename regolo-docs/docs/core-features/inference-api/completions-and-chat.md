# Completions and Chat

## Static Chat Completions

Static chat completions enable a more interactive session by providing conversation-like exchanges, you can send a series of messages. Each message has a role, such as `user`, `assistant` or `system`. The model processes these to continue the conversation naturally. This is useful for applications requiring a back-and-forth dialogue.

=== "Using Regolo Client"

    ```python
    import regolo

    regolo.default_key = "<YOUR_REGOLO_KEY>"
    regolo.default_chat_model = "Llama-3.3-70B-Instruct"

    print(regolo.static_chat_completions(messages=[{"role": "user", "content": "Tell me something about rome"}]))
    ```

=== "Python"
    ```python
    import requests

    api_url = "https://api.regolo.ai/v1/chat/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_REGOLO_KEY"
    }
    data = {
        "model": "Llama-3.3-70B-Instruct",
        "messages": [
            {"role": "user", "content": "Tell me something about Rome."}
        ]
    }

    response = requests.post(api_url, headers=headers, json=data)
    print(response.json())
    ```

=== "CURL"

    ```bash
    curl -X POST https://api.regolo.ai/v1/chat/completions \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer YOUR_REGOLO_KEY" \
        -d '{
         "model": "Llama-3.3-70B-Instruct",
         "messages": [
            {
                "role": "user",
                "content": "Tell me something about Rome."
            }
        ]
    }'
    ```

## Stream Chat Completions

Stream chat completions provide real-time, incremental responses from the model, enabling dynamic interactions and reducing latency. This feature is beneficial for applications that require immediate feedback and continuous conversation flow.

The streaming response is structured as JSON objects sent line by line. Each line typically contains metadata, including fields like `id`, `created`, `model`, and `object`, along with the `choices` array. Within `choices`, there is a `delta` object, which holds the `content` field representing the actual text response from the model. This structure allows applications to parse and process the conversational content as it arrives, ensuring efficient and timely updates to the user interface.

=== "Regolo Client - Clean"

    ```python
    import regolo

    regolo.default_key = "<YOUR_REGOLO_KEY>"
    regolo.default_chat_model = "Llama-3.3-70B-Instruct"

    client = regolo.RegoloClient()
    response = client.run_chat(
        user_prompt="Tell me something about Rome.",
        full_output=True,
        stream=True
    )

    # Process streamed responses - extract only content (words)
    while True:
        try:
            chunk = next(response)
            # Extract content from delta if it's a dict
            if isinstance(chunk, dict):
                content = chunk.get('choices', [{}])[0].get('delta', {}).get('content', '')
                if content:
                    print(content, end="", flush=True)
            else:
                # If it's already a string, print it directly
                print(chunk, end="", flush=True)
        except StopIteration:
            break
    ```

=== "Regolo Client - Full"

    ```python
    import regolo

    regolo.default_key = "<YOUR_REGOLO_KEY>"
    regolo.default_chat_model = "Llama-3.3-70B-Instruct"

    client = regolo.RegoloClient()
    response = client.run_chat(
        user_prompt="Tell me something about Rome.",
        full_output=True,
        stream=True
    )

    # Process streamed responses - show full JSON structure
    while True:
        try:
            chunk = next(response)
            print(chunk)
        except StopIteration:
            break
    ```

=== "Python - Clean"
    ```python
    import requests
    import json

    api_url = "https://api.regolo.ai/v1/chat/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_REGOLO_KEY"
    }
    data = {
        "model": "Llama-3.3-70B-Instruct",
        "messages": [
            {"role": "user", "content": "Tell me something about Rome."}
        ],
        "stream": True
    }

    response = requests.post(api_url, headers=headers, json=data, stream=True)

    # Extract only content from streamed JSON
    for line in response.iter_lines():
        if line:
            decoded_line = line.decode('utf-8')
            if decoded_line.startswith('data: '):
                json_str = decoded_line[6:]
                if json_str.strip() == '[DONE]':
                    break
                try:
                    json_data = json.loads(json_str)
                    if 'choices' in json_data:
                        delta = json_data['choices'][0].get('delta', {})
                        content = delta.get('content', '')
                        if content:
                            print(content, end='', flush=True)
                except json.JSONDecodeError:
                    pass
    ```

=== "Python - Full"
    ```python
    import requests

    api_url = "https://api.regolo.ai/v1/chat/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_REGOLO_KEY"
    }
    data = {
        "model": "Llama-3.3-70B-Instruct",
        "messages": [
            {"role": "user", "content": "Tell me something about Rome."}
        ],
        "stream": True
    }

    response = requests.post(api_url, headers=headers, json=data, stream=True)

    # Show full JSON response
    for line in response.iter_lines():
        if line:
            print(line.decode('utf-8'))
    ```


## Text Static Completions (Deprecated)

!!! warning
    The static text completions are currently deprecated, and Regolo no longer provides any model that supports them natively. They are still listed among the available endpoints only for backward compatibility and internal use, but no public model currently supports them. Use the chat completions instead.

Static completions allow you to generate text responses based on a given prompt using the Regolo API.

=== "Using Regolo Client"

    ```python
    import regolo

    regolo.default_key = "<YOUR_REGOLO_KEY>"
    regolo.default_chat_model = "Llama-3.3-70B-Instruct"

    print(regolo.static_completions(prompt="Tell me something about Rome."))
    ```

=== "Python"
    ```python
    import requests

    api_url = "https://api.regolo.ai/v1/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_REGOLO_KEY"
    }
    data = {
        "model": "Llama-3.3-70B-Instruct",
        "prompt": "Tell me something about Rome.",
        "temperature": 0.7
    }

    response = requests.post(api_url, headers=headers, json=data)
    print(response.json())
    ```

=== "CURL"

    ```bash
    curl -X POST https://api.regolo.ai/v1/completions 
    -H "Content-Type: application/json" 
    -H "Authorization: Bearer YOUR_REGOLO_KEY" 
    -d '{
        "model": "Llama-3.3-70B-Instruct",
        "prompt": "Tell me something about Rome.",
        "temperature": 0.7
    }'
    ```



For the exhaustive API's endpoints documentation visit [docs.api.regolo.ai](https://docs.api.regolo.ai).

