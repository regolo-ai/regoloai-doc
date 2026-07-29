# Playground

Explore Regolo AI models interactively, without writing a single line of code.

Use the Playground to experiment with models and parameters before wiring them into your application.

The Regolo Playground is an interactive web interface. From here you can:

- **Test different models** and compare how each one answers across families
- **Play with parameters** like temperature and max tokens to see what changes
- **Copy the code** for a request in the language you prefer

[**Open Playground →**](https://dashboard.regolo.ai/playground)

## Playground Features

### System Prompts

Set a custom system prompt to steer the assistant's behavior:

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

## API Testing

Prefer to test the API directly? Head over to the interactive API documentation:

[**API Reference →**](https://docs.api.regolo.ai)

## Tips for Experimentation

1. **Start simple.** Begin with basic prompts before piling on complexity.
2. **Compare models.** Run the same prompt across a few of them and watch how the answers differ.
3. **Iterate on your prompts.** Small tweaks often lead to noticeably better outputs.
4. **Save what works.** Copy the configurations that perform well, so they're ready when you move to production.
5. **Keep an eye on tokens.** Watching token usage is the easiest way to keep costs under control.
