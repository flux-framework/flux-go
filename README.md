# Flux Core Go Bindings

![img/flux-go-banner.png](img/flux-go-banner.png)

These will provide common interfaces to interacting with Flux in Go.

## Server

This is easiet built outside of the developer environment.

 - [api](cmd/api/api.go) Deploy a flux API service.

```bash
make build
docker run -it  ghcr.io/flux-framework/flux-go:latest
```

Or in the [.devcontainer](.devcontainer) environment:

```bash
make server
```
```console
🦩️ This is the flux service
[GRPCServer] gRPC Listening on [::]:4242
```

## Examples

These are most easiest built in the [developer container environment](.devcontainer).

```bash
make examples
```

 - [submit](example/submit/submitn.go): Submit a job! 
 - [keygen](example/keygen/keygen.go): Use zeromq to generate a curve certificate.
 - [list-jobs](example/list-jobs/list-jobs.go): List jobs example


```console
$ ./bin/flux-submit 
⭐️ Testing flux submit in Go! ⭐️
Submitting a Sleep Job: sleep 10
Flux Future: &{{{}}}

$ flux jobs -a
       JOBID USER     NAME       ST NTASKS NNODES     TIME INFO
  ƒ2y8mkYpjR vscode   sleep       R      1      1   1.688s 94dd63b69bfb


$ ./bin/flux-keygen 
⭐️ Testing flux keygen in Go! ⭐️
Generating to path...
Saving to ./curve.cert
Generating to string...
#   ****  Generated on 2023-04-26 22:54:42 by CZMQ  ****
#   ZeroMQ CURVE **Secret** Certificate
#   DO NOT PROVIDE THIS FILE TO OTHER USERS nor change its permissions.
    
metadata
    name = "curve-cert"
    keygen.hostname = "94dd63b69bfb"
curve
    public-key = "v(4E@IfTOMaW8#)x6jc6}^+.ERAW)IJNzro}w9oX"
    secret-key = "Vg4o}MV}z7SVzs#f^(o8aLZYx6r-29bLhH&Sva7t"


$ ./bin/flux-list-jobs 
⭐️ Testing flux list jobs rpc in Go! ⭐️
[
 {
  "id": 251877870534656,
  "userid": 1000,
  "urgency": 16,
  "priority": 16,
  "t_submit": 1739667371.1322868,
  "state": 64
 }
]
```

## Development

We provide a [Devcontainer environment](.devcontainer) for use in VScode.

## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614
