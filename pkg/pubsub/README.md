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

#### func (*Nats) Auth

```go
func (n *Nats) Auth(authFunc functions.AuthFunc) nats.AuthTokenHandler
```

#### func (*Nats) NewInbox

```go
func (n *Nats) NewInbox() string
```

#### func (*Nats) ProtoToMsg

```go
func (n *Nats) ProtoToMsg(msg *any.Any, reply string) *nats.Msg
```

#### func (*Nats) ProtoToSubscription

```go
func (n *Nats) ProtoToSubscription(msg *any.Any, queue string) *nats.Subscription
```

#### func (*Nats) Publish

```go
func (n *Nats) Publish(msg *any.Any) error
```

#### func (*Nats) SubscriptionBarrier

```go
func (n *Nats) SubscriptionBarrier(basicFunc functions.BasicFunc) error
```
