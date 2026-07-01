\# meshery-ai



An early-stage AI Adapter for \[Meshery](https://github.com/meshery/meshery), implementing the `adapter.Handler` interface from `meshery-adapter-library`. Enables "Natural Language to Infrastructure" workflows by bridging Meshery Server to LLMs via gRPC.



Built in response to \[meshery/meshery#19092](https://github.com/meshery/meshery/issues/19092).



\## Status

Early scaffold. Currently supports:

\- `ApplyOperation` wired to a local Ollama instance (`OLLAMA\_HOST`, `OLLAMA\_MODEL` env vars, defaults to `localhost:11434` / `llama3`)

\- Standard Meshery adapter gRPC service bootstrap (`main.go`)



\## Not yet implemented

\- `ProcessOAM` (OAM component parsing)

\- Cloud LLM providers (OpenAI, Anthropic) — BYOM support per issue #19092

\- System-prompt context injection from Meshery Model schemas

\- Unit/integration tests



\## Run locally

```powershell

$env:CGO\_ENABLED="0"

go build ./...

./meshery-ai

```

Requires a running Ollama instance (`ollama serve`) with a model pulled (`ollama pull llama3`).



\## Related

\- Issue: https://github.com/meshery/meshery/issues/19092

\- Adapter library: https://github.com/meshery/meshery-adapter-library

