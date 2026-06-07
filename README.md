# fox-contracts

Machine-readable contract schemas, constants, and Go types for the
[Fox instance contract](https://github.com/fox-in-the-box-ai/fox-in-the-box/blob/main/docs/architecture/INSTANCE_CONTRACT.md).

## Contents

- **Go types** — `FoxPrincipal`, response structs for contract endpoints
  (`/health`, `/readyz`, `/version`, `/capabilities`, `/skillset`)
- **Constants** — Fox extension header names, SSE event types, contract versions
- **JSON Schema** — machine-readable schemas for contract endpoint responses
  (`schema/*.schema.json`)

## Usage

```go
import contracts "github.com/fox-in-the-box-ai/fox-contracts"

principal := contracts.FoxPrincipal{
    UserID:   "u-123",
    TenantID: "acme",
    Groups:   []string{"engineering"},
}
```

## License

Apache License 2.0 — see [LICENSE](LICENSE).
