package user

import (
	"fmt"
	"net/http"
	"strings"
)

// UserEndpoint :
type UserEndpoint struct {
	Service UserService
}

// GetUser :
func (ep *UserEndpoint) GetUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	u, err := ep.Service.Get(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintln(w, u.Name)
}
