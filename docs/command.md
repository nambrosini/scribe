# Command

scribe

## Commit

### Options

#### For the text generation

- LLM Model type (mistral, ollama)
- LLM API Url
- LLM API Key (optional, only if model requires it)
- Model name (optional)

#### Prompts

- Prompt template
  - Full
  - Concise
- Template file (optional)
  - If a custom template should be used
- Issue number (optional)
- Type (default: feat) [feat, fix, etc...]

## Config file

```yaml
model:
  modelType: (mistral|ollama)
  url: <rest api url>
  apiKey: <rest api key>
  name: <model to be used>
commit:
  template: (full|concise)
  templateFile: <filename containing the template>
  issue: <the issue to link this commit to>
  type: <commit type> (feat|fix|chore)
```
