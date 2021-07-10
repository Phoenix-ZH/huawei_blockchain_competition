package registration

import (
	"contract"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Andrew-M-C/go.jsonvalue"
	"net/http"
	"schema"
	"strings"
)

/*
 * GET: 返回登录页
 * POST: 处理登录请求 person or issuer
 */
func Login(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	if r.Method == "GET" {
		w.Header().Add("method", "GET")
		_, err := fmt.Fprintf(w, "这里是登录页")
		if err != nil {
			return &schema.ServerError{Error: err, Code: 500}
		}
	} else if r.Method == "POST" {
		bodyBytes := make([]byte, r.ContentLength)
		_, err := r.Body.Read(bodyBytes)
		if err != nil {
			return &schema.ServerError{Error: err, Message: "Read Body error.", Code: 500}
		}
		var body struct{
			Uid string `json:"uid"`
			Passwd string `json:"passwd"`
		}
		err = json.Unmarshal(bodyBytes, &body)
		uid := body.Uid
		passwd := body.Passwd
		var function string
		if strings.HasPrefix(uid, "p") {
			function = "queryPerson"
		} else {
			function = "queryIssuer"
		}
		resBytes, err := contract.Query([]string{uid}, function, contract.Peer0)
		if err != nil || resBytes == nil {
			return &schema.ServerError{Error: err, Message: "The function queryPerson() of chaincode return err", Code: 500}
		} else {
			j, err := jsonvalue.Unmarshal(resBytes)
			if err != nil {
				return &schema.ServerError{Error: err, Message: "jsonvalue.Unmarshal error", Code: 500}
			}
			str, _ := j.GetString("Info")
			var person schema.Person
			_ = json.Unmarshal([]byte(str), &person)
			if passwd != person.Password {
				return &schema.ServerError{Error: errors.New("password not equal"), Message: "Password not equal.", Code: 500}
			}
			personBytes, _ := json.Marshal(person)
			_, err = fmt.Fprintf(w, string(personBytes))
			if err != nil {
				return &schema.ServerError{Error: err, Code: 500}
			}
			cookie := http.Cookie{Name: "uid", Value: person.Id}
			http.SetCookie(w, &cookie)
		}
	}
	return nil
}