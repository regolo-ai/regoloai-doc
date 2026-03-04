# Fine-Tuning Models

Customize models for your specific use case with fine-tuning.

## Overview

Fine-tuning allows you to:

- Adapt models to your domain
- Improve performance on specific tasks
- Reduce hallucinations and errors
- Lower latency and costs
- Maintain privacy with your data

## When to Fine-Tune

### Ideal Use Cases

✅ Domain-specific language (medical, legal, technical)  
✅ Specific writing style or tone  
✅ Instruction following improvements  
✅ Custom classification tasks  
✅ Reducing model bias  

### Not Recommended

❌ General knowledge tasks (use base model)  
❌ Small datasets (< 100 examples)  
❌ One-off requests  

## Preparing Training Data

### Data Format

Prepare JSONL format with examples:

```json
{"messages": [{"role": "system", "content": "You are a helpful assistant."}, {"role": "user", "content": "What is Python?"}, {"role": "assistant", "content": "Python is a programming language..."}]}
{"messages": [{"role": "system", "content": "You are a helpful assistant."}, {"role": "user", "content": "What is JavaScript?"}, {"role": "assistant", "content": "JavaScript is a scripting language..."}]}
```

### Data Quality

```python
import regolo

# Validate training data
validation = regolo.validate_training_data(
    file_path="training_data.jsonl"
)

print(f"Total examples: {validation.total_examples}")
print(f"Valid examples: {validation.valid_examples}")
print(f"Errors: {validation.errors}")

if validation.valid_examples >= 100:
    print("Data is ready for fine-tuning")
```

### Recommended Guidelines

- **Minimum examples**: 100 (1000+ recommended)
- **Quality over quantity**: High-quality examples outweigh quantity
- **Diversity**: Include various input types
- **Validation**: Reserve 10-20% for validation

## Starting Fine-Tuning

### Basic Fine-Tuning

```python
import regolo

# Start fine-tuning job
finetune_job = regolo.create_finetuning_job(
    model="Llama-3.3-70B-Instruct",
    training_file="training_data.jsonl",
    output_model_name="my-finetuned-model",
    epochs=3
)

print(f"Job ID: {finetune_job.id}")
print(f"Status: {finetune_job.status}")
```

### Advanced Configuration

```python
import regolo

finetune_job = regolo.create_finetuning_job(
    model="Llama-3.3-70B-Instruct",
    training_file="training_data.jsonl",
    validation_file="validation_data.jsonl",
    output_model_name="my-finetuned-model",
    
    # Hyperparameters
    epochs=3,
    batch_size=32,
    learning_rate=2e-5,
    warmup_steps=100,
    
    # Optimization
    lora_rank=16,
    lora_alpha=32
)
```

## Monitoring Fine-Tuning

### Check Status

```python
import regolo

job = regolo.get_finetuning_job(job_id="ft_abc123")

print(f"Status: {job.status}")
print(f"Progress: {job.progress_percent}%")
print(f"Epoch: {job.current_epoch} / {job.epochs}")
print(f"Estimated time remaining: {job.eta_hours} hours")
```

### Monitor Training Metrics

```python
import regolo

metrics = regolo.get_finetuning_metrics(job_id="ft_abc123")

for epoch, epoch_metrics in metrics.items():
    print(f"Epoch {epoch}:")
    print(f"  Loss: {epoch_metrics['loss']:.4f}")
    print(f"  Validation loss: {epoch_metrics['val_loss']:.4f}")
    print(f"  Learning rate: {epoch_metrics['learning_rate']:.2e}")
```

## Using Fine-Tuned Models

### Deploy Fine-Tuned Model

```python
import regolo

# After job completes, use the fine-tuned model
response = regolo.static_chat_completions(
    model="my-finetuned-model",  # Your fine-tuned model
    messages=[{"role": "user", "content": "Your prompt"}]
)

print(response.choices[0].message.content)
```

### A/B Testing

```python
import regolo

# Compare base model vs fine-tuned
test_prompts = [
    "What is machine learning?",
    "Explain neural networks",
    "Define overfitting"
]

for prompt in test_prompts:
    # Base model
    response_base = regolo.static_chat_completions(
        model="Llama-3.3-70B-Instruct",
        messages=[{"role": "user", "content": prompt}]
    )
    
    # Fine-tuned model
    response_ft = regolo.static_chat_completions(
        model="my-finetuned-model",
        messages=[{"role": "user", "content": prompt}]
    )
    
    # Compare results
    print(f"Prompt: {prompt}")
    print(f"Base: {response_base.choices[0].message.content}")
    print(f"Fine-tuned: {response_ft.choices[0].message.content}")
    print()
```

## Cost Analysis

### Fine-Tuning Costs

```
Training cost: Tokens in training data × $0.00001
Storage cost: Model size × $0.10/GB/month
Inference cost: Same as base model
```

### Example

```
Training data: 100K tokens
Training cost: 100K × $0.00001 = $1.00

Model size: 70GB
Monthly storage: 70 × $0.10 = $7.00

Total monthly: ~$8.00
```

## Best Practices

1. **Start Simple** - Begin with default hyperparameters
2. **Monitor Training** - Watch for overfitting
3. **Validate Thoroughly** - Test on validation set
4. **Version Control** - Track which data was used
5. **Document Changes** - Record improvements
6. **Cost Tracking** - Monitor training and storage costs

## Troubleshooting

### Loss Not Decreasing

- Check data quality
- Verify learning rate
- Ensure examples are diverse

### Overfitting

- Increase dropout
- Add more training data
- Reduce epochs

### Memory Issues

- Reduce batch size
- Use LoRA for efficient fine-tuning
- Use smaller base model