package http_util

type Method struct { value string }

var (
	GET    = Method { "GET" }
	POST   = Method { "POST" }
	PUT    = Method { "PUT" }
	PATCH  = Method { "PATCH" }
	DELETE = Method { "DELETE" }
	OPTION = Method { "OPTION" }
)

func (m Method) ToString() string {
	return m.value
}
