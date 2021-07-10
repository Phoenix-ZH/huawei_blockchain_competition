package organization

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
 * POST: 添加 `item`(仅当当前用户为 issuer)
 */
func AddItem(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	if r.Method == "GET" {
		return &schema.ServerError{Error: errors.New("only receive GET request"), Code: 500}
	}
	cookie, _ := r.Cookie("uid")
	if cookie == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		if strings.HasPrefix(cookie.Value, "p") {
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
				Description string `json:"description"`
				Point float64 `json:"point"`
			}
			err = json.Unmarshal(bodyBytes, &body)
			item := schema.Item{
				Id: body.Id,
				Name: body.Name,
				Description: body.Description,
				Issuer: cookie.Value,
				Point: body.Point,
			}
			resBytes, err := contract.Send(item.ToStringArray(), "addItem", contract.Peer0)
			if err != nil {
				return &schema.ServerError{Error: err, Message: "addItem error", Code: 500}
			}
			res, _ := jsonvalue.Unmarshal(resBytes)
			status, _ := res.GetString("Status")
			if status != "SUCCESS" {
				return &schema.ServerError{Error: errors.New("addItem failed"), Code: 500}
			}
		}
	}
	return nil
}