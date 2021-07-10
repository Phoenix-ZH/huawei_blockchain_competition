package registration

import (
	"contract"
	"encoding/json"
	"errors"
	"github.com/Andrew-M-C/go.jsonvalue"
	"net/http"
	"schema"
	"strings"
)

/*
 * POST: 添加用户, 仅网站管理员具有该权限
 */
func Register(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	if r.Method == "GET" {
		return &schema.ServerError{Error: errors.New("only receive POST request"), Code: 500}
	}
	cookie, _ := r.Cookie("uid")
	if cookie == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		if !strings.HasPrefix(cookie.Value, "admin") {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			bodyBytes := make([]byte, r.ContentLength)
			_, err := r.Body.Read(bodyBytes)
			if err != nil {
				return &schema.ServerError{Error: err, Message: "Read Body error.", Code: 500}
			}
			var body struct{
				Id string `json:"id"`
				Name string `json:"name"`
				Password string `json:"password"`
				PublicKey string `json:"publicKey"`
			}
			err = json.Unmarshal(bodyBytes, &body)
			person := schema.Person{
				Id: body.Id,
				Name: body.Name,
				Password: body.Password,
				PublicKey: body.PublicKey,
			}
			resBytes, err := contract.Send(person.ToStringArray(), "addPerson", contract.Peer0)
			if err != nil {
				return &schema.ServerError{Error: err, Message: "addPerson error", Code: 500}
			}
			res, _ := jsonvalue.Unmarshal(resBytes)
			status, _ := res.GetString("Status")
			if status != "SUCCESS" {
				return &schema.ServerError{Error: errors.New("addPerson failed"), Code: 500}
			}
		}
	}
	return nil
}