# Virenmonitoring SDK

Weekly viral RNA measurements from Basel wastewater for SARS-CoV-2, RSV, and influenza

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Virenmonitoring API

The Virenmonitoring API exposes the wastewater virus surveillance datasets that the Abwasserreinigungsanlage (ARA) Basel publishes through the [Kanton Basel-Stadt open data portal](https://data.bs.ch/). The programme samples wastewater entering Basel's treatment plant and quantifies viral RNA copies for several respiratory pathogens, alongside reported clinical case counts for cross-reference.

What you get from the API:

- Weekly measurements of viral RNA copies per 100,000 residents in the Basel catchment.
- Coverage of SARS-CoV-2, RSV, and influenza A and B.
- Confirmed positive case numbers reported alongside the wastewater signal.
- Dataset metadata (titles, descriptions, update timestamps) for each monitored virus.

The portal is operated on top of OpenDataSoft, so each dataset is reachable through the standard OpenDataSoft records and exports endpoints under `https://data.bs.ch/api`. No authentication is documented for read access; CORS behaviour varies between datasets (the SARS-CoV-2 dataset is reported as CORS-enabled while the influenza A dataset is not).

## Try it

**TypeScript**
```bash
npm install virenmonitoring
```

**Python**
```bash
pip install virenmonitoring-sdk
```

**PHP**
```bash
composer require voxgig/virenmonitoring-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/virenmonitoring-sdk/go
```

**Ruby**
```bash
gem install virenmonitoring-sdk
```

**Lua**
```bash
luarocks install virenmonitoring-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { VirenmonitoringSDK } from 'virenmonitoring'

const client = new VirenmonitoringSDK({})

// List all datasetmetadatas
const datasetmetadatas = await client.DatasetMetadata().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o virenmonitoring-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "virenmonitoring": {
      "command": "/abs/path/to/virenmonitoring-mcp"
    }
  }
}
```

## Entities

The API exposes 2 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **DatasetMetadata** | Catalogue-level information about each wastewater monitoring dataset (title, description, update cadence, fields), served by the portal's OpenDataSoft catalogue endpoints under `/api`. | `/datasets/1.0/100304/` |
| **VirusMonitoring** | The time-series records themselves — weekly viral RNA copies per 100,000 residents plus confirmed case counts for SARS-CoV-2, RSV, and influenza A/B, retrieved from the per-dataset records and exports endpoints. | `/records/1.0/search/` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from virenmonitoring_sdk import VirenmonitoringSDK

client = VirenmonitoringSDK({})

# List all datasetmetadatas
datasetmetadatas, err = client.DatasetMetadata(None).list(None, None)
```

### PHP

```php
<?php
require_once 'virenmonitoring_sdk.php';

$client = new VirenmonitoringSDK([]);

// List all datasetmetadatas
[$datasetmetadatas, $err] = $client->DatasetMetadata(null)->list(null, null);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/virenmonitoring-sdk/go"

client := sdk.NewVirenmonitoringSDK(map[string]any{})

// List all datasetmetadatas
datasetmetadatas, err := client.DatasetMetadata(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "Virenmonitoring_sdk"

client = VirenmonitoringSDK.new({})

# List all datasetmetadatas
datasetmetadatas, err = client.DatasetMetadata(nil).list(nil, nil)
```

### Lua

```lua
local sdk = require("virenmonitoring_sdk")

local client = sdk.new({})

-- List all datasetmetadatas
local datasetmetadatas, err = client:DatasetMetadata(nil):list(nil, nil)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = VirenmonitoringSDK.test()
const result = await client.DatasetMetadata().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = VirenmonitoringSDK.test(None, None)
result, err = client.DatasetMetadata(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = VirenmonitoringSDK::test(null, null);
[$result, $err] = $client->DatasetMetadata(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.DatasetMetadata(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = VirenmonitoringSDK.test(nil, nil)
result, err = client.DatasetMetadata(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:DatasetMetadata(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Virenmonitoring API

- Upstream: [https://data.bs.ch/](https://data.bs.ch/)
- API docs: [https://data.bs.ch/api/](https://data.bs.ch/api/)

- Published as Open Government Data by Kanton Basel-Stadt on the data.bs.ch portal.
- Hosted on OpenDataSoft infrastructure; consult the portal's terms and conditions for reuse and attribution requirements.
- No specific licence string is advertised on the dataset landing page; cite Kanton Basel-Stadt as the source.

---

Generated from the Virenmonitoring API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
