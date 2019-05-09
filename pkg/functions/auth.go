package functions

import "github.com/nats-io/nats"

type AuthFunc func() string

func (a AuthFunc) AsNatsHandler() nats.AuthTokenHandler {
	return nats.AuthTokenHandler(a)
}
