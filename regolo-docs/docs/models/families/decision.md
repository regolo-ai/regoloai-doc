# Decision Models

The Decision Models API lets you classify information, evaluate conditions, and assign scores. You provide the information to evaluate and the questions to answer, together with their allowed options. The API returns structured decisions and probabilities that your application can use directly.

**When to use it** – Route support tickets, classify documents, check whether a condition is met, or assign a priority. For example, a single request can decide which team should handle a customer message, whether it describes an outage, and how urgently it needs attention.

Send requests to **`POST https://api.regolo.ai/v1/systemone`** with a [Regolo API key](../../getting-started/sign-up.md). Choose a decision model from the [Model Catalog](../catalog.md) and pass its ID in `model`. Regolo hosts the model for you.

!!! note "A dedicated decision endpoint"
    Send `state` and `questions` to `/v1/systemone` to obtain decisions over the options you define. Use direct HTTP requests, as shown below, rather than a chat completion call with `messages` or `response_format`.

---

## API Call Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| `model` | `string` (required) | ID of the decision model to use. Check available models and their capabilities in the [Model Catalog](../catalog.md). |
| `state` | String or JSON value (required) | The information the model should evaluate, such as a customer message or a JSON object containing a ticket and relevant context. |
| `questions` | `object` (required) | Questions keyed by an ID you choose, such as `department` or `outage`. Each ID also identifies its answer in the response. |

Each question contains:

| Field | Description |
| --- | --- |
| `type` | Required. One of `choice`, `noul`, or `score`. |
| `instructions` | What the model should decide. Use explicit instructions to make the task clear. |
| `criteria` | The allowed options and their meanings. Required for `choice` and `score`; optional for `noul`. |

### Choose a question type

| Type | Use it to… | `criteria` format | Answer |
| --- | --- | --- | --- |
| `choice` | Select one category. | Object mapping option IDs to descriptions. | `choice`, `confidence`, and `probabilities` for the options. |
| `noul` | Evaluate a yes/no condition. | Optional object with descriptions for `true` and `false`. | `noul`: a number from 0 to 1 representing the probability of **true**. |
| `score` | Rate something on an ordered scale. | List of descriptions, from lowest to highest. | `score`, `confidence`, `legend`, and `probabilities` for the levels. |

For `score`, list positions define the numeric scale, starting at **0**. With three levels, the score ranges from 0 to 2. It is a probability-weighted average, so it can be a decimal rather than a whole number. These types follow the [SystemOne request and response format](https://docs.sglang.io/docs/supported-models/decision_models#system-one-compatible-api).

---

## Example Requests

The examples use **`clef-flash`**, a model in Cloudflare's [CLEF family](https://huggingface.co/Cloudflare/clef). See its [model card](https://huggingface.co/Cloudflare/clef-flash) for background. When using another decision model, set `model` to its catalog ID and check its supported capabilities.

This example evaluates one support ticket with all three question types:

- **`department`** selects the team to handle the ticket.
- **`outage`** checks whether customers are currently unable to complete purchases.
- **`priority`** rates how soon the ticket needs attention.

Set your API key in the terminal before running either example:

```bash
export REGOLO_API_KEY="YOUR_REGOLO_KEY"
```

The Python example uses `requests`, which you can install with `python3 -m pip install requests`. Save the example as `decision.py`, then run `python3 decision.py` from the same terminal where you set the API key.

=== "cURL"

    ```bash
    curl --request POST https://api.regolo.ai/v1/systemone \
      --fail-with-body \
      --max-time 60 \
      --header "Authorization: Bearer $REGOLO_API_KEY" \
      --header "Content-Type: application/json" \
      --data '{
        "model": "clef-flash",
        "state": {
          "ticket": "Our checkout is returning errors. Customers cannot complete purchases. Please help today."
        },
        "questions": {
          "department": {
            "type": "choice",
            "instructions": "Which team should handle this ticket?",
            "criteria": {
              "billing": "Questions about charges, invoices, or refunds.",
              "technical": "Product errors, broken integrations, or service outages.",
              "sales": "Questions about buying a product or choosing a plan."
            }
          },
          "outage": {
            "type": "noul",
            "instructions": "Does the ticket say customers are currently unable to complete purchases?"
          },
          "priority": {
            "type": "score",
            "instructions": "How soon does this ticket need attention?",
            "criteria": [
              "Low: a general question that can wait.",
              "Medium: a problem with a workaround, to address this week.",
              "High: a service is blocked and needs attention today."
            ]
          }
        }
      }'
    ```

=== "Python"

    ```python
    import os

    import requests

    payload = {
        "model": "clef-flash",
        "state": {
            "ticket": "Our checkout is returning errors. Customers cannot complete purchases. Please help today."
        },
        "questions": {
            "department": {
                "type": "choice",
                "instructions": "Which team should handle this ticket?",
                "criteria": {
                    "billing": "Questions about charges, invoices, or refunds.",
                    "technical": "Product errors, broken integrations, or service outages.",
                    "sales": "Questions about buying a product or choosing a plan.",
                },
            },
            "outage": {
                "type": "noul",
                "instructions": "Does the ticket say customers are currently unable to complete purchases?",
            },
            "priority": {
                "type": "score",
                "instructions": "How soon does this ticket need attention?",
                "criteria": [
                    "Low: a general question that can wait.",
                    "Medium: a problem with a workaround, to address this week.",
                    "High: a service is blocked and needs attention today.",
                ],
            },
        },
    }

    response = requests.post(
        "https://api.regolo.ai/v1/systemone",
        headers={"Authorization": f"Bearer {os.environ['REGOLO_API_KEY']}"},
        json=payload,
        timeout=60,
    )
    response.raise_for_status()

    result = response.json()
    answers = result["answers"]
    print("Team:", answers["department"]["choice"])
    print("Team confidence:", answers["department"]["confidence"])
    print("Outage probability:", answers["outage"]["noul"])
    print("Priority (0 to 2):", answers["priority"]["score"])
    ```

## Response

Answers are keyed by the same IDs you sent in `questions`. The following response was returned by Regolo with **HTTP 200** for the cURL request above using `clef-flash`. Exact values can vary between models and requests.

```json
{
  "model": "clef-flash",
  "answers": {
    "department": {
      "type": "choice",
      "choice": "technical",
      "confidence": 0.999999999964877,
      "probabilities": {
        "billing": 1.2256069247280297e-11,
        "technical": 0.9999999999765847,
        "sales": 1.1159278039439907e-11
      },
      "x_label_mass": 0.9699117895992838
    },
    "outage": {
      "type": "noul",
      "noul": 0.9999999999998701,
      "x_label_mass": 0.7104499497215855
    },
    "priority": {
      "type": "score",
      "score": 1.9999999999999405,
      "confidence": 0.9999999999999104,
      "legend": {
        "0": "Low: a general question that can wait.",
        "1": "Medium: a problem with a workaround, to address this week.",
        "2": "High: a service is blocked and needs attention today."
      },
      "probabilities": {
        "0": 1.1663308147704977e-16,
        "1": 5.948063307252533e-14,
        "2": 0.9999999999999405
      },
      "x_label_mass": 0.9985959006521702
    }
  },
  "usage": {
    "input_tokens": 248,
    "output_tokens": 0
  }
}
```

In this response, `technical` is the selected team. The `outage` value is almost **1**, meaning the model assigns a high probability to the condition being true; it is a number, not a Boolean. This evaluates what the ticket reports; checking the service itself requires monitoring data. The priority score is almost **2**, the highest level on the 0–2 scale.

The `usage` object reports token counts: this request used 248 input tokens and generated no output tokens. Additional fields such as `x_label_mass` may appear depending on the serving implementation; the examples read the decision fields directly.

For both `choice` and `score`, `confidence` summarizes certainty in the answer. For `score`, read it separately from the priority value: a high priority and a high confidence describe different things. The serving implementation computes confidence from the option probabilities, so avoid assuming it always equals the largest probability. See the [SystemOne response reference](https://docs.sglang.io/docs/supported-models/decision_models#system-one-compatible-api) for details.

For a `choice` question, the probabilities show how the model distributes its preference across your categories. Use them to inspect close decisions. Your application decides how to act on the result; for example, it can route a ticket or send an uncertain case for review. See Regolo's [support ticket routing guide](https://regolo.ai/route-support-tickets-with-clef-flash-on-regolo/) for a larger example.

!!! tip "Choose thresholds using your own examples"
    Test the decisions against examples you have labeled before choosing a confidence threshold for automatic routing. Confidence is a model signal, not a guarantee that the decision is correct. A threshold such as `0.8` is a starting point to evaluate, not a universal rule.

## Writing Clear Questions

- Keep the information to evaluate in `state`, and describe the decision in `instructions`.
- Give each option a clear description. If two categories overlap, explain which one takes precedence.
- Add an `other` or `needs_review` option to a `choice` question when your application needs a fallback. The model otherwise selects from the categories you provide.
- Use several questions about the same state to obtain separate decisions, such as team, urgency, and topic.
- Keep the state focused on the task. Include relevant context, such as a routing policy, when the decision depends on it.

## Using Decisions in a Workflow

A decision model can select a next step for your application: route a ticket to a team, flag a document for review, or decide which tool should run. Your code then performs that step.

If the workflow also needs a written answer, pass the decision and the relevant context to a [chat model](completions.md). For example, use a decision model to select the support team and priority, then use a chat model to draft a reply.

For JSON text generated by a chat model, see [Structured Outputs](../features/structured-outputs.md).
