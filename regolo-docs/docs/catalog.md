# Regolo Catalog

The **Regolo Catalog** is the official catalog of available models on the 
Regolo AI platform. It provides an up-to-date list of all models that can be 
used with the Regolo API.

## Overview

The catalog is automatically updated every 24 hours by fetching the latest list 
of available models from the Regolo API. This ensures that all applications 
using Regolo models always have access to the most up-to-date model list.

## Accessing the Catalog

### Repository

Visit the official Regolo Model Catalog repository on GitHub:

**🔗 [regolo-ai/regolo-model-catalog](https://github.com/regolo-ai/regolo-model-catalog)**

### List Available Models

You can retrieve the list of available models directly from the Regolo API 
using the `/models` endpoint:

=== "Python"

    ```python
    import requests

    url = "https://api.regolo.ai/models"
    headers = {
        "Authorization": "Bearer YOUR_REGOLO_KEY"
    }

    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        models = response.json()
        print(models)
    else:
        print("Failed to fetch models:", response.status_code, 
              response.text)
    ```

=== "CURL"

    ```bash
    curl -X GET https://api.regolo.ai/models \
        -H "Authorization: Bearer YOUR_REGOLO_KEY"
    ```

## Request a Model

Want to see a specific model on Regolo? Let us know! You can:

* Open an issue on the [catalog repository](https://github.com/regolo-ai/regolo-model-catalog)
* Start a discussion
* Contact us at help@regolo.ai

