package bot

import (
	"context"
	"github.com/google/go-github/github"
)

type contextKey string

const BotContextKey contextKey = "botContext"

// Context encapsulates the fields passed to webhook handlers
type Context struct {
	App     *App
	Payload interface{}
	GitHub  *github.Client
}

// NewContext instantiates a new context
func NewContext(app *App) *Context {
	return &Context{App: app}
}

func FromContext(context context.Context) (*Context, bool) {
	ctx, ok := context.Value(BotContextKey).(*Context)
	return ctx, ok
}
