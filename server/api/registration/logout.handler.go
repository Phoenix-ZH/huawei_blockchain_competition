package registration

import (
	"fmt"
	"net/http"
	"server.com/schema"
)

/*
 * GET: 登出当前帐号
 */
func Logout(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	cookie := http.Cookie{
		Name: "uid",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	_, err := fmt.Fprintf(w, "用户已登出")
	if err != nil {
		return &schema.ServerError{err, "Logout error.", 500}
	}
	return nil
}
