# Function Calling

**Function calling** enables the language model to invoke external functions during a conversation, allowing it to retrieve real-time information or perform operations.

## Supported Models

The following models support function calling:

- `deepseek-r1-70b`
- `gpt-oss-120b`
- `Llama-3.3-70B-Instruct`
- `mistral-small3.2`
- `qwen3-30b`
- `qwen3-coder-30b`

## Web Search Example

Here's a complete example using DuckDuckGo web search with the Regolo API:

```python
from ddgs import DDGS
import logging
import os
import requests
import json

logger = logging.getLogger(__name__)

REGOLO_BASE_URL = "https://api.regolo.ai/v1"
API_KEY = os.getenv("REGOLO_API_KEY", "YOUR_REGOLO_KEY")


def duckduckgo_search(query: str, max_results: int = 10, region: str = "wt-wt") -> str:
    """
    Search the web using DuckDuckGo and return formatted results.
    
    Args:
        query: Search query string
        max_results: Maximum number of results (default: 10)
        region: Region code for localized results (default: "wt-wt" for worldwide)
    
    Returns:
        Formatted markdown string with search results
    """
    try:
        logger.info(f"DUCKDUCKGO: Searching for '{query}' (max_results={max_results})")
        
        with DDGS() as ddgs:
            # ddgs package uses 'query' parameter
            results = list(ddgs.text(query=query, max_results=max_results, region=region))
        
        if not results:
            return "No search results found."
        
        # Format results as markdown
        formatted = "## Search Results\n\n"
        for result in results:
            title = result.get("title", "No title")
            url = result.get("href", "")
            body = result.get("body", "")
            
            formatted += f"[{title}]({url})\n"
            formatted += f"{body}\n\n"
        
        logger.info(f"DUCKDUCKGO: Found {len(results)} results")
        return formatted.strip()
        
    except Exception as e:
        logger.error(f"DUCKDUCKGO: Error: {e}")
        return f"Error performing search: {str(e)}"


tools = [
    {
        "type": "function",
        "function": {
            "name": "duckduckgo_search",
            "description": (
                "Search the web for current information using DuckDuckGo. "
                "Use this when you need up-to-date information, news, facts, "
                "or any information that may change over time."
            ),
            "parameters": {
                "type": "object",
                "properties": {
                    "query": {
                        "type": "string",
                        "description": "The search query string"
                    },
                    "max_results": {
                        "type": "integer",
                        "description": "Maximum number of results to return (1-20)",
                        "minimum": 1,
                        "maximum": 20,
                        "default": 10
                    },
                    "region": {
                        "type": "string",
                        "description": "Region code for localized results (e.g., 'wt-wt' for worldwide, 'us-en' for US English)",
                        "default": "wt-wt"
                    }
                },
                "required": ["query"]
            }
        }
    }
]
```

=== "Non-Streaming"

    === "Clean"

        ```python
        # Create messages list
        messages = [
            {"role": "user", "content": "Find me some news about AI"}
        ]
        
        # Initial request with tools
        headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {API_KEY}"
        }
        
        payload = {
            "model": "gpt-oss-120b",
            "messages": messages,
            "tools": tools,
            "tool_choice": "auto"
        }
        
        response = requests.post(
            f"{REGOLO_BASE_URL}/chat/completions",
            headers=headers,
            json=payload
        )
        
        message = response.json()["choices"][0]["message"]
        
        # Handle function call
        if message.get("tool_calls"):
            for tool_call in message["tool_calls"]:
                if tool_call["function"]["name"] == "duckduckgo_search":
                    args = json.loads(tool_call["function"]["arguments"])
                    search_results = duckduckgo_search(
                        query=args.get("query"),
                        max_results=args.get("max_results", 10),
                        region=args.get("region", "wt-wt")
                    )
                    
                    # Add assistant message with tool_calls and tool response
                    messages.append(message)
                    messages.append({
                        "role": "tool",
                        "tool_call_id": tool_call["id"],
                        "name": tool_call["function"]["name"],
                        "content": search_results
                    })
                    
                    # Make follow-up call with updated messages
                    followup_payload = {
                        "model": "gpt-oss-120b",
                        "messages": messages
                    }
                    
                    followup_response = requests.post(
                        f"{REGOLO_BASE_URL}/chat/completions",
                        headers=headers,
                        json=followup_payload
                    )
                    
                    followup_response.raise_for_status()
                    final_answer = followup_response.json()["choices"][0]["message"]["content"]
                    print(final_answer)
        else:
            print(message.get("content", ""))
        ```

    === "Full"

        ```python
        # Create messages list
        messages = [
            {"role": "user", "content": "Find me some news about AI"}
        ]
        
        # Initial request with tools
        headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {API_KEY}"
        }
        
        payload = {
            "model": "gpt-oss-120b",
            "messages": messages,
            "tools": tools,
            "tool_choice": "auto"
        }
        
        response = requests.post(
            f"{REGOLO_BASE_URL}/chat/completions",
            headers=headers,
            json=payload
        )
        
        # Show full JSON response
        print(response.json())
        
        message = response.json()["choices"][0]["message"]
        
        # Handle function call
        if message.get("tool_calls"):
            for tool_call in message["tool_calls"]:
                if tool_call["function"]["name"] == "duckduckgo_search":
                    args = json.loads(tool_call["function"]["arguments"])
                    search_results = duckduckgo_search(
                        query=args.get("query"),
                        max_results=args.get("max_results", 10),
                        region=args.get("region", "wt-wt")
                    )
                    
                    # Add assistant message with tool_calls and tool response
                    messages.append(message)
                    messages.append({
                        "role": "tool",
                        "tool_call_id": tool_call["id"],
                        "name": tool_call["function"]["name"],
                        "content": search_results
                    })
                    
                    # Make follow-up call with updated messages
                    followup_payload = {
                        "model": "gpt-oss-120b",
                        "messages": messages
                    }
                    
                    followup_response = requests.post(
                        f"{REGOLO_BASE_URL}/chat/completions",
                        headers=headers,
                        json=followup_payload
                    )
                    
                    followup_response.raise_for_status()
                    # Show full JSON response
                    print(followup_response.json())
        else:
            print(message.get("content", ""))
        ```

=== "Streaming"

    === "Clean"

        ```python
        # Create messages list
        messages = [
            {"role": "user", "content": "Find me some news about AI"}
        ]
        
        # Initial request with tools
        headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {API_KEY}"
        }
        
        payload = {
            "model": "gpt-oss-120b",
            "messages": messages,
            "tools": tools,
            "tool_choice": "auto"
        }
        
        response = requests.post(
            f"{REGOLO_BASE_URL}/chat/completions",
            headers=headers,
            json=payload
        )
        
        message = response.json()["choices"][0]["message"]
        
        # Handle function call
        if message.get("tool_calls"):
            for tool_call in message["tool_calls"]:
                if tool_call["function"]["name"] == "duckduckgo_search":
                    args = json.loads(tool_call["function"]["arguments"])
                    search_results = duckduckgo_search(
                        query=args.get("query"),
                        max_results=args.get("max_results", 10),
                        region=args.get("region", "wt-wt")
                    )
                    
                    # Add assistant message with tool_calls and tool response
                    messages.append(message)
                    messages.append({
                        "role": "tool",
                        "tool_call_id": tool_call["id"],
                        "name": tool_call["function"]["name"],
                        "content": search_results
                    })
                    
                    # Make follow-up call with streaming enabled
                    followup_payload = {
                        "model": "gpt-oss-120b",
                        "messages": messages,
                        "stream": True
                    }
                    
                    followup_response = requests.post(
                        f"{REGOLO_BASE_URL}/chat/completions",
                        headers=headers,
                        json=followup_payload,
                        stream=True
                    )
                    
                    # Extract only content from streamed JSON
                    for line in followup_response.iter_lines():
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
        else:
            print(message.get("content", ""))
        ```

    === "Full"

        ```python
        # Create messages list
        messages = [
            {"role": "user", "content": "Find me some news about AI"}
        ]
        
        # Initial request with tools
        headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {API_KEY}"
        }
        
        payload = {
            "model": "gpt-oss-120b",
            "messages": messages,
            "tools": tools,
            "tool_choice": "auto"
        }
        
        response = requests.post(
            f"{REGOLO_BASE_URL}/chat/completions",
            headers=headers,
            json=payload
        )
        
        # Show full JSON response
        print(response.json())
        
        message = response.json()["choices"][0]["message"]
        
        # Handle function call
        if message.get("tool_calls"):
            for tool_call in message["tool_calls"]:
                if tool_call["function"]["name"] == "duckduckgo_search":
                    args = json.loads(tool_call["function"]["arguments"])
                    search_results = duckduckgo_search(
                        query=args.get("query"),
                        max_results=args.get("max_results", 10),
                        region=args.get("region", "wt-wt")
                    )
                    
                    # Add assistant message with tool_calls and tool response
                    messages.append(message)
                    messages.append({
                        "role": "tool",
                        "tool_call_id": tool_call["id"],
                        "name": tool_call["function"]["name"],
                        "content": search_results
                    })
                    
                    # Make follow-up call with streaming enabled
                    followup_payload = {
                        "model": "gpt-oss-120b",
                        "messages": messages,
                        "stream": True
                    }
                    
                    followup_response = requests.post(
                        f"{REGOLO_BASE_URL}/chat/completions",
                        headers=headers,
                        json=followup_payload,
                        stream=True
                    )
                    
                    # Show full JSON response
                    for line in followup_response.iter_lines():
                        if line:
                            print(line.decode('utf-8'))
        else:
            print(message.get("content", ""))
        ```

## Tool Choice

By default, the model will determine when and how many tools to use. You can force specific behavior with the `tool_choice` parameter.

### Auto (Default)

The model decides whether to call zero, one, or multiple functions:

```python
payload = {
    "model": "gpt-oss-120b",
    "messages": messages,
    "tools": tools,
    "tool_choice": "auto"  # Default behavior
}
```

### Required

Force the model to call one or more functions:

```python
payload = {
    "model": "gpt-oss-120b",
    "messages": messages,
    "tools": tools,
    "tool_choice": "required"
}
```

### Forced Function

Force the model to call exactly one specific function:

```python
payload = {
    "model": "gpt-oss-120b",
    "messages": messages,
    "tools": tools,
    "tool_choice": {
        "type": "function",
        "function": {"name": "duckduckgo_search"}
    }
}
```

### Allowed Tools

Restrict the tool calls the model can make to a subset of the tools available. This is useful when you want to make only a subset of tools available across model requests without modifying the list of tools you pass in, maximizing savings from prompt caching.

```python
payload = {
    "model": "gpt-oss-120b",
    "messages": messages,
    "tools": tools,  # All available tools
    "tool_choice": {
        "type": "allowed_tools",
        "mode": "auto",
        "tools": [
            {"type": "function", "name": "duckduckgo_search"}
        ]
    }
}
```

### None

Disable function calling entirely, imitating the behavior of passing no functions:

```python
payload = {
    "model": "gpt-oss-120b",
    "messages": messages,
    "tools": tools,
    "tool_choice": "none"
}
```

## Further Reading

- [OpenAI Function Calling docs](https://platform.openai.com/docs/guides/function-calling)
