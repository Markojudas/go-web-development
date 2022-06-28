<!-- markdownlint-disable -->

# Package: Context

<a href="https://pkg.go.dev/context" target="_blank" rel="noopener noreferrer">Godocs</a><br>
Package context defines the Context type, which carries deadlines, cnacelation signals, and other request-scoped values across API boundaries and between process.

<hr>

`type Context`

```go

type Context interface {
	// Deadline returns the time when work done on behalf of this context
	// should be canceled. Deadline returns ok==false when no deadline is
	// set. Successive calls to Deadline return the same results.
	Deadline() (deadline time.Time, ok bool)

	// Done returns a channel that's closed when work done on behalf of this
	// context should be canceled. Done may return nil if this context can
	// never be canceled. Successive calls to Done return the same value.
	// The close of the Done channel may happen asynchronously,
	// after the cancel function returns.
	//
	// WithCancel arranges for Done to be closed when cancel is called;
	// WithDeadline arranges for Done to be closed when the deadline
	// expires; WithTimeout arranges for Done to be closed when the timeout
	// elapses.
	//
	// Done is provided for use in select statements:
	//
	//  // Stream generates values with DoSomething and sends them to out
	//  // until DoSomething returns an error or ctx.Done is closed.
	//  func Stream(ctx context.Context, out chan<- Value) error {
	//  	for {
	//  		v, err := DoSomething(ctx)
	//  		if err != nil {
	//  			return err
	//  		}
	//  		select {
	//  		case <-ctx.Done():
	//  			return ctx.Err()
	//  		case out <- v:
	//  		}
	//  	}
	//  }
	//
	// See https://blog.golang.org/pipelines for more examples of how to use
	// a Done channel for cancellation.
	Done() <-chan struct{}

	// If Done is not yet closed, Err returns nil.
	// If Done is closed, Err returns a non-nil error explaining why:
	// Canceled if the context was canceled
	// or DeadlineExceeded if the context's deadline passed.
	// After Err returns a non-nil error, successive calls to Err return the same error.
	Err() error

	// Value returns the value associated with this context for key, or nil
	// if no value is associated with key. Successive calls to Value with
	// the same key returns the same result.
	//
	// Use context values only for request-scoped data that transits
	// processes and API boundaries, not for passing optional parameters to
	// functions.
	//
	// A key identifies a specific value in a Context. Functions that wish
	// to store values in Context typically allocate a key in a global
	// variable then use that key as the argument to context.WithValue and
	// Context.Value. A key can be any type that supports equality;
	// packages should define keys as an unexported type to avoid
	// collisions.
	//
	// Packages that define a Context key should provide type-safe accessors
	// for the values stored using that key:
	//
	// 	// Package user defines a User type that's stored in Contexts.
	// 	package user
	//
	// 	import "context"
	//
	// 	// User is the type of value stored in the Contexts.
	// 	type User struct {...}
	//
	// 	// key is an unexported type for keys defined in this package.
	// 	// This prevents collisions with keys defined in other packages.
	// 	type key int
	//
	// 	// userKey is the key for user.User values in Contexts. It is
	// 	// unexported; clients use user.NewContext and user.FromContext
	// 	// instead of using this key directly.
	// 	var userKey key
	//
	// 	// NewContext returns a new Context that carries value u.
	// 	func NewContext(ctx context.Context, u *User) context.Context {
	// 		return context.WithValue(ctx, userKey, u)
	// 	}
	//
	// 	// FromContext returns the User value stored in ctx, if any.
	// 	func FromContext(ctx context.Context) (*User, bool) {
	// 		u, ok := ctx.Value(userKey).(*User)
	// 		return u, ok
	// 	}
	Value(key any) any
}

```

<br>
<hr>

`func WithCancel`

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

WithCancel returns a copy of parent with a new Done channel. The return context's Done channel is closed
when the returned cancel function is called or when the parent context's Done channel is closed,
which ever happenes first.

Canceling this context releases resources assocaited with it, so code should call cancel as soon as the operations
running in this Context complete
<br>
<br>

<hr>

`func Background`

```go
func Background() Context
```

Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline.
It is typically used by the main function, initilization, and tests, and as the top-level Context for incoming
requests.
<br>
<br>

<hr>

`func WithValue`

```go
func WithValue(parent Context, key, val any) Context
```

WithValue returns a copy of parent in which the value associated with key is val.

Use context Values only for request-scoped data that transits processes and APIs, not for passing
optional parameters to functions.

The provided key must be comparable and should not be of type string or any other built-in type
to avoid collisions between packages using context. Users of WithValue should define their own
types for keys. To avoid allocating when assigning to an interface{}, context keys often have
concrete type struct{}. Alternatively, exported context key variables' static type should be
a pointer or interface.
<br>
<br>

<hr>

`func WithDeadline`

```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```

WithDeadline returns a copy of the parent context with the deadline adjusterd to be no later than d.
If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically equivalent
to parent. The returned context's Done channel is closed when the deadline expires, when the returned cnacel
function is called, or when the parent context's Done channel is closed, whichever happens first.

Canceling this context releases resources assocaited with it, so code should call cancel as soon as the operations
running in this Context complete.
<br>
<br>

<hr>

`func WithTimeout`

```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)

Canceling this context releases resources associated with it, so code should call cancel as
soon as the operation running in this Context complete:

<pre><code>func slowOperationWithTimeout(ctx context.Context) (Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()  //releases resources if slowOperation completes before timeout elapses
	return slowOperation(ctx)
}</code></pre>
<br>
<hr>

`type CancelFunc`

```go
type CancelFunc func()
```

A CancelFunc() tells an operation to abandon its work. A CancelFunc does not wait for the work to stop. A CancelFunc may be called by multiple goroutines simultaneously. After the first call, subsequent calls to a CancelFunc do nothing.
