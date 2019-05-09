# util
--
    import "github.com/autom8ter/tasks/pkg/util"


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
