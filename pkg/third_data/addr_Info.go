package third_data

type AddrInfo struct {
	Url         string
	Method      string
	ContentType string
	Headers     map[string]string
	Type        int
	Req         map[string]interface{}
	Auth        bool
	TokenName   string
}

func (a *AddrInfo) SetAuthToken(token string) {
	for s := range a.Headers {
		if s == a.TokenName {
			a.Headers[s] = token
		}
	}
	for s := range a.Req {
		if s == a.TokenName {
			a.Req[s] = token
		}
	}
}
