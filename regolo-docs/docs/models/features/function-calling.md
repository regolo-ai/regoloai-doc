# Function Calling

Function calling allows models to interact with external tools and APIs by requesting specific function invocations based on the conversation context.

## Overview

Function calling enables models to:
- Request specific functions to be called
- Provide structured arguments for those functions
- Integrate with external APIs and tools
- Create more interactive and capable applications

## Usage

To use function calling, define your functions and pass them to the model in the request.

=== "Using Regolo Client"

    ```python
    import regolo

    regolo.default_key = "<YOUR_REGOLO_KEY>"
    regolo.default_chat_model = "Llama-3.3-70B-Instruct"

    functions = [
        {
            "name": "get_weather",
            "description": "Get the current weather in a given location",
            "parameters": {
                "type": "object",
                "properties": {
                    "location": {
                        "type": "string",
                        "description": "The city and state, e.g. San Francisco, CA"
                    },
                    "unit": {
                        "type": "string",
                        "enum": ["celsius", "fahrenheit"],
                        "description": "The unit for temperature"
                    }
                },
                "required": ["location"]
            }
        }
    ]

    response = regolo.static_chat_completions(
        messages=[{"role": "user", "content": "What's the weather like in Rome?"}],
        functions=functions
    )
    
    print(response)
    ```

=== "Python"

    ```python
    import requests

    api_url = "https://api.regolo.ai/v1/chat/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer YOUR_REGOLO_KEY"
    }
    
    functions = [
        {
            "name": "get_weather",
            "description": "Get the current weather in a given location",
            "parameters": {
                "type": "object",
                "properties": {
                    "location": {
                        "type": "string",
                        "description": "The city and state, e.g. San Francisco, CA"
                    },
                    "unit": {
                        "type": "string",
                        "enum": ["celsius", "fahrenheit"],
                        "description": "The unit for temperature"
                    }
                },
                "required": ["location"]
            }
        }
    ]
    
    data = {
        "model": "Llama-3.3-70B-Instruct",
        "messages": [
            {"role": "user", "content": "What's the weather like in Rome?"}
        ],
        "functions": functions
    }

    response = requests.post(api_url, headers=headers, json=data)
    print(response.json())
    ```

## Function Response Handling

After the model requests a function call, you should execute the function and send the result back to the model.

```python
# Example: Handle function call response
if response["choices"][0]["message"].get("function_call"):
    function_name = response["choices"][0]["message"]["function_call"]["name"]
    function_args = json.loads(response["choices"][0]["message"]["function_call"]["arguments"])
    
    # Execute your function here
    function_result = execute_function(function_name, function_args)
    
    # Send result back to model
    messages.append(response["choices"][0]["message"])
    messages.append({
        "role": "function",
        "name": function_name,
        "content": json.dumps(function_result)
    })
```

## Benefits

- **Tool Integration**: Connect models to external APIs and services
- **Structured Output**: Get structured data instead of free-form text
- **Extensibility**: Add new capabilities without retraining models
- **Reliability**: Models can request specific, well-defined operations

