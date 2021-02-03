// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	client "github.com/amplify-cms/protoc-gen-cobra/client"
	flag "github.com/amplify-cms/protoc-gen-cobra/flag"
	iocodec "github.com/amplify-cms/protoc-gen-cobra/iocodec"
	proto "github.com/golang/protobuf/proto"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	io "io"
)

func CacheClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Cache"),
		Short: "Cache service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_CacheSetCommand(cfg),
		_CacheGetCommand(cfg),
		_CacheMultiSetCommand(cfg),
		_CacheMultiGetCommand(cfg),
	)
	return cmd
}

func _CacheSetCommand(cfg *client.Config) *cobra.Command {
	req := &SetRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Set"),
		Short: "Set RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cache"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cache", "Set"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewCacheClient(cc)
				v := &SetRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Set(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Key, cfg.FlagNamer("Key"), "", "")
	cmd.PersistentFlags().StringVar(&req.Value, cfg.FlagNamer("Value"), "", "")

	return cmd
}

func _CacheGetCommand(cfg *client.Config) *cobra.Command {
	req := &GetRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Get"),
		Short: "Get RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cache"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cache", "Get"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewCacheClient(cc)
				v := &GetRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Get(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Key, cfg.FlagNamer("Key"), "", "")

	return cmd
}

func _CacheMultiSetCommand(cfg *client.Config) *cobra.Command {
	req := &SetRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("MultiSet"),
		Short: "MultiSet RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cache"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cache", "MultiSet"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewCacheClient(cc)
				v := &SetRequest{}

				stm, err := cli.MultiSet(cmd.Context())
				if err != nil {
					return err
				}
				for {
					if err := in(v); err != nil {
						if err == io.EOF {
							_ = stm.CloseSend()
							break
						}
						return err
					}
					proto.Merge(v, req)
					if err = stm.Send(v); err != nil {
						return err
					}
				}

				res, err := stm.CloseAndRecv()
				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Key, cfg.FlagNamer("Key"), "", "")
	cmd.PersistentFlags().StringVar(&req.Value, cfg.FlagNamer("Value"), "", "")

	return cmd
}

func _CacheMultiGetCommand(cfg *client.Config) *cobra.Command {
	req := &GetRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("MultiGet"),
		Short: "MultiGet RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cache"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Cache", "MultiGet"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewCacheClient(cc)
				v := &GetRequest{}

				stm, err := cli.MultiGet(cmd.Context())
				if err != nil {
					return err
				}
				for {
					if err := in(v); err != nil {
						if err == io.EOF {
							_ = stm.CloseSend()
							break
						}
						return err
					}
					proto.Merge(v, req)
					if err = stm.Send(v); err != nil {
						return err
					}
				}

				for {
					res, err := stm.Recv()
					if err != nil {
						if err == io.EOF {
							break
						}
						return err
					}
					if err = out(res); err != nil {
						return err
					}
				}
				return nil

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Key, cfg.FlagNamer("Key"), "", "")

	return cmd
}
