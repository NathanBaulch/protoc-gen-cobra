package cmd

import (
	"go.amplifyedge.org/protoc-gen-cobra/example/pb"

	_ "go.amplifyedge.org/protoc-gen-cobra/auth/jwt"
	_ "go.amplifyedge.org/protoc-gen-cobra/auth/oauth"
	_ "go.amplifyedge.org/protoc-gen-cobra/iocodec/yaml"
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
	rootCmd.AddCommand(pb.OneofClientCommand())
}
