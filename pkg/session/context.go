package session

import "context"

type SessionKey struct{}
type TokenKey struct{}

func FromContext(ctx context.Context) (*Session, bool) {
	session, ok := ctx.Value(SessionKey{}).(*Session)
	if !ok {
		return nil, ok
	}
	return session, false
}

func IntoContext(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, SessionKey{}, session)
}
