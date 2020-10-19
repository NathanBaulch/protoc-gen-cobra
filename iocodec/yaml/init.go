package yaml

import (
	"io"

	"gopkg.in/yaml.v3"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
	"github.com/NathanBaulch/protoc-gen-cobra/iocodec"
)

func init() {
	client.RegisterInputDecoder("yaml", decoderMaker)
	client.RegisterOutputEncoder("yaml", encoderMaker)
}

func decoderMaker(r io.Reader) iocodec.Decoder {
	return iocodec.DecodeKnownTypes(yaml.NewDecoder(r).Decode)
}

func encoderMaker(w io.Writer) iocodec.Encoder {
	return iocodec.EncodeKnownTypes(yaml.NewEncoder(w).Encode)
}
