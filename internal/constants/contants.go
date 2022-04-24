package constants

type ctxKey int8
type MdKey string

const (
	CtxKeyUser     ctxKey = iota
	MdKeyUserEmail MdKey  = "email"
	MdKeyAuthToken MdKey  = "token"
)
