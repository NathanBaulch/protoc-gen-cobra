# protoc-gen-cobra

Cobra command line tool generator for Go gRPC.

[![GoDoc](https://godoc.org/github.com/NathanBaulch/protoc-gen-cobra?status.svg)](https://godoc.org/github.com/NathanBaulch/protoc-gen-cobra)

### What's this?

A plugin for the [protobuf](https://github.com/google/protobuf) compiler protoc that generates Go code using [cobra](https://github.com/spf13/cobra). It is
capable of generating client code for command line tools consistent with your protobuf description.

This proto definition:

```proto
service Bank {
	rpc Deposit(DepositRequest) returns (DepositReply);
}

message DepositRequest {
	string account = 1;
	double amount = 2;
}

message DepositReply {
	string account = 1;
	double balance = 2;
}
```

produces a client that can do:

```
$ ./example bank deposit --account foobar --amount 10
$ echo '{"account":"foobar"}' | ./example bank deposit --amount 10 -f -
$ set BANK_DEPOSIT_ACCOUNT=foobar; ./example bank deposit --amount 10
```

It generates one [cobra.Command](https://godoc.org/github.com/spf13/cobra#Command) per gRPC service (e.g. bank). The service's RPC methods are sub-commands and
share the same command line semantics. They take flags, a request file, or stdin for input, and print the response to the terminal in the specified format. The
client currently supports basic connectivity settings such as TLS on/off, TLS client authentication and so on.

```
$ ./example bank deposit -h
Deposit RPC client

Usage:
  example bank deposit [flags]

Flags:
      --account string   account number of recipient
      --amount float     amount to deposit
  -h, --help             help for deposit

Global Flags:
      --auth-access-token string   authorization access token
      --auth-token-type string     authorization token type (default "Bearer")
      --config string              config file (default is $HOME/.example.yaml)
      --jwt-key string             JWT key
      --jwt-key-file string        JWT key file
  -f, --request-file string        client request file; use "-" for stdin
  -i, --request-format string      request format (json, xml, yaml) (default "json")
  -o, --response-format string     response format (json, prettyjson, prettyxml, xml, yaml) (default "json")
  -s, --server-addr string         server address in the form host:port (default "localhost:8080")
      --timeout duration           client connection timeout (default 10s)
      --tls                        enable TLS
      --tls-ca-cert-file string    CA certificate file
      --tls-cert-file string       client certificate file
      --tls-insecure-skip-verify   INSECURE: skip TLS checks
      --tls-key-file string        client key file
      --tls-server-name string     TLS server name override
```

### Streams

gRPC client and server streams are supported by concatenating multiple requests and responses.
Client stream requests can be passed from a file or stdin with optional whitespace between each document (typically a new line).
Server stream responses are printed one per line using the specified response format.

Example client stream:

```
$ cat req.json
{"key":"hello","value":"world"}
{"key":"foo","value":"bar"}

$ ./example cache multiset -f req.json
```

Example client and server streams:

```
$ echo '{"key":"hello"}{"key":"foo"}' | ./example cache multiget -f -
{"value":"world"}
{"value":"bar"}
```

Idle server streams hang until the server closes the stream or a timeout occurs.

### Custom input/output formats

New input decoders and output encoders can be registered in the host program prior to executing the command.

```go
client.RegisterOutputEncoder("fmtprint", func(w io.Writer) iocodec.Encoder {
	return func(v interface{}) error {
		_, err := fmt.Fprint(w, v)
		return err
	}
})
```

These formats can then be specified using the `-o` or `--response-format` flags or via a `RESPONSE_FORMAT` environment variable.

```
$ ./example bank deposit --account foobar --amount 10 -o fmtprint
$ set RESPONSE_FORMAT=fmtprint; ./example bank deposit --account foobar --amount 10
```

See the [yaml format extension](iocodec/yaml/init.go) for a complete example.

### Custom authentication schemes

Pre-dial hooks can be registered in the host program to support custom authentication schemes.

```go
client.RegisterPreDialer(func(ctx context.Context, opts *[]grpc.DialOption) error {
	if creds, ok := ctx.Value(CredsContextKey).(*Creds); ok {
		*opts = append(*opts, grpc.WithPerRPCCredentials(creds))
	}
	return nil
})
```

See the [jwt authentication extension](auth/jwt/init.go) for a complete example.
