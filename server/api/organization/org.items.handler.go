package organization

import (
	"contract"
	"encoding/json"
	"fmt"
	"github.com/Andrew-M-C/go.jsonvalue"
	"net/http"
	"schema"
)

/*
 * GET: 返回当前组织(issuer)所发布的所有 items
 */
func IssuerItems(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	var personItems map[string][]string
	personItems = make(map[string][]string)
	personItems["p123"] = []string{"it123", "it345"}
	personItems["p234"] = []string{"it234", "it456"}

	cookie, _ := r.Cookie("uid")
	if cookie == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		items := personItems[cookie.Value]
		resBytes, err := contract.Query(items, "queryItems", contract.Peer0)
		if err != nil {
			return &schema.ServerError{Error: err, Message: "queryItems for person error", Code: 500}
		} else {
			res, _ := jsonvalue.Unmarshal(resBytes)
			info, _ := res.GetString("Info")
			itemary, _ := jsonvalue.Unmarshal([]byte(info))
			var items []schema.Item
			for v := range itemary.IterArray() {
				id, _ := v.V.GetString("Id")
				name, _ := v.V.GetString("Name")
				description, _ := v.V.GetString("Description")
				issuer, _ := v.V.GetString("Issuer")
				point, _ := v.V.GetFloat64("Point")
				item := schema.Item{
					Id: id,
					Name: name,
					Description: description,
					Issuer: issuer,
					Point: point,
				}
				items = append(items, item)
			}
			itemsBytes, _ := json.Marshal(items)
			_, err := fmt.Fprintf(w, string(itemsBytes))
			if err != nil {
				return &schema.ServerError{Error: err, Message: "fmt.Fprintf error", Code: 500}
			}
		}
	}
	return nil
}
