# Sandbox & Playground

Explore Regolo AI models interactively without writing code.

## Regolo Playground

The Regolo Playground is an interactive web interface where you can:

- **Test different models** - Compare responses across model families
- **Experiment with parameters** - Adjust temperature, max tokens, and more
- **Save conversations** - Keep track of your experiments
- **Export code** - Generate code snippets in your preferred language

[**Open Playground →**](https://dashboard.regolo.ai/playground)

## Available Models

In the playground, you can test all available models:

| Model Family | Example Models | Best For |
|--------------|----------------|----------|
| Llama | Llama-3.3-70B-Instruct | General chat, reasoning |
| Vision | Llama-3.2-11B-Vision-Instruct | Image understanding |
| Embedding | multilingual-e5-large-instruct | Text embeddings |
| Rerank | bge-reranker-v2-m3 | Search result reranking |

See the full [model catalog](../catalog.md) for all available models.

## Playground Features

### System Prompts

Set custom system prompts to control the assistant's behavior:

```
You are a helpful coding assistant specialized in Python.
Provide concise answers with code examples when relevant.
```

### Parameter Tuning

| Parameter | Range | Description |
|-----------|-------|-------------|
| `temperature` | 0.0 - 2.0 | Controls randomness (lower = more focused) |
| `max_tokens` | 1 - 4096+ | Maximum response length |
| `top_p` | 0.0 - 1.0 | Nucleus sampling threshold |
| `frequency_penalty` | -2.0 - 2.0 | Reduces repetition |

### Conversation History

The playground maintains conversation context, allowing you to:

- Build multi-turn conversations
- Test follow-up questions
- Simulate real-world chat scenarios

## API Testing

For direct API testing, check out the interactive API documentation:

[**API Reference →**](https://docs.api.regolo.ai)

The Swagger UI allows you to:

- Browse all available endpoints
- Make authenticated requests
- View request/response schemas
- Download OpenAPI specification

## Tips for Experimentation

1. **Start simple** - Begin with basic prompts before adding complexity
2. **Compare models** - Same prompt, different models, different results
3. **Iterate on prompts** - Small changes can significantly improve outputs
4. **Save successful prompts** - Export configurations for production use
5. **Monitor tokens** - Watch token usage to optimize costs