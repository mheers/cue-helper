# Create an ocigateway config

## Prerequisites

- [hof](https://github.com/hofstadter-io/hof)
- [cue](https://github.com/cuelang/cue)

## Usage

```bash
hof mod init cue github.com/mheers/ocigatewayexample
hof mod get github.com/mheers/ocigateway
hof mod vendor cue
touch config.example.cue
```

## config.example.cue

```cue
package ocigatewayexample

import (
	cfg "github.com/mheers/ocigateway"
)

cfg.#ociGateway & {
	gateways: []
}
```

## Generate config

```bash
cue export --out yaml
```
