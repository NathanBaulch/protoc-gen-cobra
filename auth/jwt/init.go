package jwt

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
	"github.com/NathanBaulch/protoc-gen-cobra/naming"
)

var Config = &config{}

type config struct {
	Key     string
	KeyFile string
}

func init() {
	client.RegisterFlagBinder(func(fs *pflag.FlagSet, namer naming.Namer) {
		fs.StringVar(&Config.Key, namer("JWT Key"), Config.Key, "JWT key")
		fs.StringVar(&Config.KeyFile, namer("JWT KeyFile"), Config.KeyFile, "JWT key file")
	})

	client.RegisterPreDialer(func(_ context.Context, opts *[]grpc.DialOption) error {
		cfg := Config

		if cfg.Key != "" {
			cred, err := oauth.NewJWTAccessFromKey([]byte(cfg.Key))
			if err != nil {
				return fmt.Errorf("jwt key: %v", err)
			}
			*opts = append(*opts, grpc.WithPerRPCCredentials(cred))
		}
		if cfg.KeyFile != "" {
			cred, err := oauth.NewJWTAccessFromFile(cfg.KeyFile)
			if err != nil {
				return fmt.Errorf("jwt key file: %v", err)
			}
			*opts = append(*opts, grpc.WithPerRPCCredentials(cred))
		}

		return nil
	})
}
