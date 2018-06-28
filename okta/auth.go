package okta

// AuthService handles communication with the Auth data related
// methods of the OKTA API.
type AuthService service

type Auth struct{
	StateToken string `json:"stateToken"`
	ExpiresAt string `json:"expiresAt"`
	Status string `json:"status"`
	SessionToken string `json:"sessionToken"`
	Embedded struct{
		User User `json:"user"`
	}`json:"_embedded"`
	Links struct {
		Cancel struct{
			HREF string `json:"href"`
			Hints struct {
				Allow []string `json:"allow"`
			} `json:"hints"`
		}`json:"cancel"`
	} `json:"_links"`

}

type AuthCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthService) GetSessionToken(username string, password string) (*Auth, *Response, error) {

	creds := AuthCreds{username, password}
	//creds_json, _ := json.Marshal(creds)
	req, err := a.client.NewRequest("POST", "authn", creds)
	req.Header.Del("Authorization")
	if err != nil {
		return nil, nil, err
	}

	auth := new(Auth)

	resp, err := a.client.Do(req, auth)

	if err != nil {
		return nil, resp, err
	}

	return auth, resp, err
}