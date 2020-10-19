package cmd

import (
	"github.com/NathanBaulch/protoc-gen-cobra/example/pb"

	_ "github.com/NathanBaulch/protoc-gen-cobra/auth/jwt"
	_ "github.com/NathanBaulch/protoc-gen-cobra/auth/oauth"
	_ "github.com/NathanBaulch/protoc-gen-cobra/iocodec/yaml"
)

func init() {
	rootCmd.AddCommand(pb.BankClientCommand())
	rootCmd.AddCommand(pb.CacheClientCommand())
	rootCmd.AddCommand(pb.TimerClientCommand())
	rootCmd.AddCommand(pb.NestedClientCommand())
	rootCmd.AddCommand(pb.CRUDClientCommand())
	rootCmd.AddCommand(pb.TypesClientCommand())
	rootCmd.AddCommand(pb.Proto2ClientCommand())
	rootCmd.AddCommand(pb.DeprecatedClientCommand())
}
