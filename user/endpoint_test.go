package user_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/podhmo-sandbox/go-mock-sandbox/user"
	"github.com/podhmo-sandbox/go-mock-sandbox/user/vmock"
)

func TestUserServiceEndpoint(t *testing.T) {
	id := "1"
	ep := &user.UserEndpoint{
		Service: &vmock.UserService{
			GetFunc: func(id string) (*user.User, error) {
				return &user.User{ID: id}, nil
			},
		},
	}

	req := httptest.NewRequest("GET", fmt.Sprintf("http://example.com/users/%s", id), nil)
	w := httptest.NewRecorder()

	ep.GetUser(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status is %d, but %d", 200, res.StatusCode)
	}
}
