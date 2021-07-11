package common

import (
	"fmt"
	"net/http"
	"server.com/schema"
	"strings"
)

/*
 * GET: 返回首页, 以登录或输入Hash查看证书
 */
func Index(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	cookie, _ := r.Cookie("uid")
	if cookie != nil {
		if strings.HasPrefix(cookie.Value, "p") {
			w.Header().Add("uid", cookie.Value)
			_, err := fmt.Fprintf(w, "learner id: %s", cookie.Value)
			if err != nil {
				return &schema.ServerError{Error: err, Code: 500}
			}
		} else if strings.HasPrefix(cookie.Value, "i") {
			w.Header().Add("uid", cookie.Value)
			_, err := fmt.Fprintf(w, "org id: %s", cookie.Value)
			if err != nil {
				return &schema.ServerError{Error: err, Code: 500}
			}
		}
	} else {
		w.Header().Add("index", "index")
		_, err := fmt.Fprintf(w, "homepage")
		if err != nil {
			return &schema.ServerError{Error: err, Code: 500}
		}
	}
	return nil
}