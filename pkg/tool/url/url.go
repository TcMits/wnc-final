package url

import (
	"net/url"

	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

func CloneURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	u2 := new(url.URL)
	*u2 = *u
	if u.User != nil {
		u2.User = new(url.Userinfo)
		*u2.User = *u.User
	}
	return u2
}

func JoinQueryString(u *string, queries map[string]string) (*string, error) {
	uHelper, err := url.Parse(*u)
	if err != nil {
		return nil, err
	}
	q := uHelper.Query()
	for k, v := range queries {
		q.Set(k, v)
	}
	uHelper.RawQuery = q.Encode()
	return generic.GetPointer(uHelper.String()), nil
}
