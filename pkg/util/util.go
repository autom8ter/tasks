//go:generate godocdown -o README.md

package util

import "github.com/gogo/protobuf/jsonpb"

var (
	PBJSONMarshaler = &jsonpb.Marshaler{
		Indent: "  ",
	}
	PBJSONUnmarshaler = &jsonpb.Unmarshaler{
		AllowUnknownFields: false,
	}
)
