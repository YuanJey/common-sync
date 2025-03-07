package third_sso

var ThirdSSOInstance ThirdSSO

type ThirdSSO struct{}

func (t *ThirdSSO) GetAuthorizationURL(redirectURI string, state string) string {
	//TODO implement me
	panic("implement me")
}

func (t *ThirdSSO) ExchangeCodeForToken(code string) (string, error) {
	//TODO implement me
	panic("implement me")
}

// GetUserInfo return union_id
func (t *ThirdSSO) GetUserInfo(accessToken string) (string, error) {
	//TODO implement me
	panic("implement me")
}
