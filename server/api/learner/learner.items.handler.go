package learner

import (
	"server.com/api/contract"
	"encoding/json"
	"fmt"
	"github.com/Andrew-M-C/go.jsonvalue"
	"github.com/bykovme/gotrans"
	"net/http"
	"server.com/schema"
	"strings"
)

/*
 * GET: 返回 返回所有/推荐的 items
 */
func Item(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	items := []string{"it123", "it234", "it345", "it456"}
	var issuerItems map[string][]string
	issuerItems = make(map[string][]string)
	issuerItems["i123"] = []string{"it123", "it234"}
	issuerItems["i234"] = []string{"it345", "it456"}

	cookie, _ := r.Cookie("uid")
	if cookie == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		var args []string
		if strings.HasPrefix(cookie.Value, "p") {
			args = items
		} else {
			args = issuerItems[cookie.Value]
		}
		resBytes, err := contract.Query(args, "queryItems", contract.Peer0)
		if err != nil {
			return &schema.ServerError{Error: err, Message: gotrans.Tr("fr", "cnte"), Code: 500}
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
				return &schema.ServerError{Error: err, Message: gotrans.Tr("fr", "fpe"), Code: 500}
			}
		}
	}
	return nil
}