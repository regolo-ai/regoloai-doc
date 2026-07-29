# Reasoning

Reasoning models work through a problem in steps before they answer. While they think they emit reasoning tokens, and Regolo returns that reasoning so you can follow how the model reached its conclusion.

## Overview

A reasoning model spends extra effort planning, weighing alternatives, and checking its own work before it replies. That makes these models a good fit for math, code, multi-step logic, and any task where the path to the answer matters as much as the answer itself.

Reasoning is a property of the model, not a switch you flip. You get it by picking a reasoning-capable model such as `gpt-oss-120b`, and you tune how hard it thinks with `reasoning_effort`.

!!! note "Reasoning tokens are billed"
    The model's internal thinking uses reasoning tokens on top of the normal input and output. They take up space in the context window and are billed as output tokens, so deeper reasoning costs more.

## How it works

Send a regular chat completion request to a reasoning model. The reasoning comes back in the `reasoning_content` field of the assistant message, kept separate from the final answer in `content`.

```python
import requests

api_url = "https://api.regolo.ai/v1/chat/completions"
headers = {
    "Content-Type": "application/json",
    "Authorization": "Bearer YOUR_REGOLO_KEY",
}
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "What color was Napoleon's white horse?"}
    ],
}

response = requests.post(api_url, headers=headers, json=data)
result = response.json()

# Pull out the reasoning and the final answer
message = result["choices"][0]["message"]
reasoning = message.get("reasoning_content", "")
answer = message.get("content", "")

if reasoning:
    print("=== Reasoning ===")
    print(reasoning)
    print("\n=== Final Answer ===")

print(answer)
```

There is no special flag to turn reasoning on. Choose a reasoning model and the reasoning comes back in the response.

## Controlling effort

Use `reasoning_effort` to decide how much the model should think. Lower effort is faster and cheaper; higher effort spends more tokens to dig into harder problems.

The values most models agree on are:

| Effort | When it fits |
| --- | --- |
| `low` | Quick tasks where speed matters more than depth |
| `medium` | The balanced default for most workloads |
| `high` | Hard problems that reward deeper thinking |

```python
data = {
    "model": "gpt-oss-120b",
    "messages": [
        {"role": "user", "content": "Solve this step by step: What is 15% of 240?"}
    ],
    "reasoning_effort": "high",  # low, medium, or high
}
```

### It varies from model to model

Reasoning parameters are not consistent across models. Regolo normalizes requests where it can, so a model may quietly accept a value that its own documentation never mentions, but the behavior you get is still specific to that model. A few cases you'll run into:

- **Switching reasoning off.** Some models accept `reasoning_effort: "none"` and skip reasoning to answer right away, even when the official model docs don't list that option.
- **Going beyond `high`.** A few models support deeper thinking with values such as `xhigh` or `max`.
- **No scale at all.** Other reasoning models always reason the same way and simply ignore `reasoning_effort`.

So before you depend on any of this, check the model's official documentation for which reasoning parameters it truly supports. An unsupported value may be rejected outright, or it may just do nothing.

!!! tip "Verify it in the Playground"
    You don't have to guess. Every reasoning parameter in the [Regolo Playground](../../getting-started/playground.md) is tested, and the supported options are laid out model by model. If you need to know whether a model lets you switch reasoning off or push past `high`, start there.

!!! note "The usual parameters still apply"
    Standard parameters like `temperature`, `max_tokens`, `top_p`, `frequency_penalty`, and `presence_penalty` keep working alongside reasoning. See [Response Parameters](response-parameters.md) for the details.

## When reasoning earns its cost

- **Hard problems**: math, logic, and multi-step questions where the route to the answer matters
- **Code**: debugging, refactoring, and planning a change before you write it
- **Trust and debugging**: reading the reasoning helps you spot where the model went wrong
- **Agents and tool use**: reasoning models tend to plan tool calls more reliably
