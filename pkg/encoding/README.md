# encoding
--
    import "github.com/autom8ter/tasks/pkg/encoding"


## Usage

```go
var (
	PBJSONMarshaler = &jsonpb.Marshaler{
		Indent: "  ",
	}
	PBJSONUnmarshaler = &jsonpb.Unmarshaler{
		AllowUnknownFields: false,
	}
)
```
