# Completions and Chat

## Static Completions

Static completions allow you to generate text responses based on a given prompt using the Regolo API.

=== "Using Regolo Client"

    ```python
    import regolo

    regolo.default_key = "<YOUR_REGOLO_KEY>"
    regolo.default_model = "Llama-3.3-70B-Instruct"

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

## Static Chat Completions

Static chat completions enable a more interactive session by providing conversation-like exchanges, you can send a series of messages. Each message has a role, such as `user`, `assistant` or `system`. The model processes these to continue the conversation naturally. This is useful for applications requiring a back-and-forth dialogue.

=== "Using Regolo Client"

    ```python
    import regolo

    regolo.default_key = "<YOUR_REGOLO_KEY>"
    regolo.default_model = "Llama-3.3-70B-Instruct"

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
    curl -X POST https://api.regolo.ai/v1/chat/completions 
    -H "Content-Type: application/json" 
    -H "Authorization: Bearer YOUR_REGOLO_KEY" 
    -d '{
        "model": "gpt-4o",
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

=== "Using Regolo Client"

    ```python
    import regolo

    regolo.default_key = "<YOUR_REGOLO_KEY>"
    regolo.default_model = "Llama-3.3-70B-Instruct"

    client = regolo.RegoloClient()
    response = client.run_chat(user_prompt="Tell me something about Rome.", full_output=True, stream=True)


    while True:
        try:
            print(next(response))
        except StopIteration:
            break
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
        ],
        "stream": True
    }

    response = requests.post(api_url, headers=headers, json=data, stream=True)

    for line in response.iter_lines():
        if line:
            print(line.decode('utf-8'))
    ```

!!! tip

    For the exhaustive API's endpoints documentation visit [docs.api.regolo.ai](https://docs.api.regolo.ai).