package client

import (
	"time"

	"github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	"github.com/NathanBaulch/protoc-gen-cobra/naming"
)

type Option func(*Config)

func WithServerAddr(addr string) Option {
	return func(c *Config) {
		c.ServerAddr = addr
	}
}

func WithRequestFormat(format string) Option {
	return func(c *Config) {
		c.RequestFormat = format
	}
}

func WithResponseFormat(format string) Option {
	return func(c *Config) {
		c.ResponseFormat = format
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

func WithEnvVars(prefix string) Option {
	return func(c *Config) {
		c.UseEnvVars = true
		c.EnvVarPrefix = prefix
	}
}

func WithCommandNamer(namer naming.Namer) Option {
	return func(c *Config) {
		c.CommandNamer = namer
	}
}

func WithFlagNamer(namer naming.Namer) Option {
	return func(c *Config) {
		c.FlagNamer = namer
	}
}

func WithEnvVarNamer(namer naming.Namer) Option {
	return func(c *Config) {
		c.EnvVarNamer = namer
	}
}

func WithTLSCACertFile(certFile string) Option {
	return func(c *Config) {
		c.TLS = true
		c.CACertFile = certFile
	}
}

func WithTLSCertFile(certFile, keyFile string) Option {
	return func(c *Config) {
		c.TLS = true
		c.CertFile = certFile
		c.KeyFile = keyFile
	}
}

func WithTLSServerName(serverName string) Option {
	return func(c *Config) {
		c.TLS = true
		c.ServerName = serverName
	}
}

func WithFlagBinder(binder FlagBinder) Option {
	return func(c *Config) {
		c.flagBinders = append(c.flagBinders, binder)
	}
}

func WithPreDialer(dialer PreDialer) Option {
	return func(c *Config) {
		c.preDialers = append(c.preDialers, dialer)
	}
}

func WithInputDecoder(format string, maker iocodec.DecoderMaker) Option {
	return func(c *Config) {
		d := make(map[string]iocodec.DecoderMaker, len(c.inDecoders)+1)
		for k, v := range c.inDecoders {
			d[k] = v
		}
		d[format] = maker
		c.inDecoders = d
	}
}

func WithOutputEncoder(format string, maker iocodec.EncoderMaker) Option {
	return func(c *Config) {
		e := make(map[string]iocodec.EncoderMaker, len(c.outEncoders)+1)
		for k, v := range c.outEncoders {
			e[k] = v
		}
		e[format] = maker
		c.outEncoders = e
	}
}
