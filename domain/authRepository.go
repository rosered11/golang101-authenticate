package domain

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type AuthRepository interface {
	IsAuthorized(string, string, map[string]string) bool
}

type RemoteAuthRepository struct {
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyUrl(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		return false
	} else {
		m := map[string]bool{}
		if err := json.NewDecoder(response.Body).Decode(&m); err != nil {
			return false
		}
		return m["isAuthorized"]
	}
}

/* Sample: /auth/verify?token=aa.bb.cc&routeName=GetTransaction&customer_id=2000&account_id=2000  */
func buildVerifyUrl(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8001", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
