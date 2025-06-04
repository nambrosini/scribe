# Docs

The idea is that the command is run like this:

```
scribe commit
```

and a commit is created with the text generated from the LLM.

## Supported LLMs

I'm not sure how the APIs look like for ChatGPT or other LLMs.
For mistral (on cloud) and all local models run with
ollama, the request looks like this:
```
{
    "model": "mistral-small:latest",
    "messages": [
        {
            "role": "user",
            "content": "What is the best Swiss cheese?"
        }
    ],
    "stream": false
}
```

*Note: the `stream` is required for ollama models so that the response
comes all together or not, not for Mistral.*

### Ollama

Ollama must be running in the background or it won't be able to
connect to it and run the model.

#### URL

Note that the `chat` endpoint has to be called, and not the `completion`.

```
http://localhost:11434/api/chat
```

#### Response

```
{
  "model": "mistral-small:latest",
  "created_at": "2025-06-04T15:22:00.208646Z",
  "message": {
    "role": "assistant",
    "content": "..."
  },
  "done_reason": "stop",
  "done": true,
  "total_duration": 32041457917,
  "load_duration": 30625000,
  "prompt_eval_count": 168,
  "prompt_eval_duration": 2657204959,
  "eval_count": 284,
  "eval_duration": 29349870500
}
```

While the duration, count, etc. might be interesting, they are ignored and only
the "message" field is taken in consideration since it contains the relevant response.

### Mistral

```
https://api.mistral.ai/v1/chat/completions
```

#### Response

```
{
  "id": "6c829d4646094b57b3c2f397de619877",
  "object": "chat.completion",
  "created": 1749057936,
  "model": "mistral-large-latest",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "tool_calls": null,
        "content": "..."
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 10,
    "total_tokens": 447,
    "completion_tokens": 437
  }
}

```
