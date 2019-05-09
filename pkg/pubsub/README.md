# pubsub
--
    import "github.com/autom8ter/tasks/pkg/pubsub"


## Usage

#### type Nats

```go
type Nats struct {
}
```


#### func  NewNats

```go
func NewNats(url string, opts ...nats.Option) (*Nats, error)
```

#### func (*Nats) Publish

```go
func (n *Nats) Publish(msg *any.Any) error
```
