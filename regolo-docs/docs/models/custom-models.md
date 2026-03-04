# Custom Models

Regolo’s **Custom Models** functionality allows you to upload and deploy your own AI models from Hugging Face onto dedicated infrastructure. You maintain full control over the GPU resources, ensuring consistent performance for your specific use cases.

---

## Quick Start Guide

Follow these steps to bring your own model to Regolo:

1. **Add your model**: Click the **"Add model"** button in the top-right corner of your dashboard.
2. **Import from Hugging Face**: Enter the Hugging Face model URL.
3. **Save to Library**: Rename the model if desired and save it to your personal library.
4. **Deploy**: Click **"Deploy model"** from your library list.
5. **Select Instance**: Choose a hardware instance with the appropriate GPU size and count for your model.
6. **Authenticate**: Use an active Regolo API key with **"All models"** permissions to begin making calls.

!!! warning
    A Regolo key with proper authorization for **all models** is required to ensure successful API calls.

---

## Model Support & Specific Endpoints

While you are free to upload any supported model, the API endpoint you use depends on the **Model Family**. You must follow the same URI and body structure used for our regular inference services.

| Model Type | Endpoint Path |
| --- | --- |
| **Text Generation (LLMs)** | `/custom-model/v1/chat/completions` |
| **Image Generation** | `/custom-model/v1/images/generations` |
| **Speech-To-Text** | `/custom-model/v1/audio/transcriptions` |
| **Embeddings** | `/custom-model/v1/embeddings` |
| **Rerankers** | `/custom-model/v1/rerank` |

!!! info
    Ensure you use the specific endpoint corresponding to the model's functionality (e.g., Image, Audio, or Reranker) as detailed in our [Swagger](https://docs.api.regolo.ai/) or on Model Families documentation.

---

## Deployment Requirements & Stability

To ensure a successful deployment and avoid service interruptions, please verify the following before launching:

* **vLLM Compatibility**: For LLMs, ensure your model architecture is supported by vLLM. Check the [vLLM supported models list](https://docs.vllm.ai/en/stable/models/supported_models/).
* **Hardware Matching**: Verify the model's size and minimum VRAM requirements. The instance size must match or exceed these demands to prevent deployment failures.
* **Startup Latency**: Startup times typically range from **a few minutes up to 15 minutes**. During this window, the endpoint will remain unavailable while the weights are loaded.

---

## Custom Inference Example

**Base URL:** `https://api.regolo.ai/custom-model/v1/`

```bash
curl -X POST \
https://api.regolo.ai/custom-model/v1/chat/completions/ \
-H "Authorization: Bearer YOUR-REGOLO-API-KEY" \
-H "Content-Type: application/json" \
-d '{
  "model": "YOUR_CUSTOM_MODEL_NAME",
  "messages": [{"role": "user", "content": "Hello!"}]
}'

```

---

## Pricing & Billing

* **Hourly Usage**: Charges are based on the duration the instance is active.
* **Initial Charge**: You will be billed for the **first hour** immediately upon deployment.
* **Invoicing**: Monthly invoices detailing total hours and charges are available in your dashboard.
