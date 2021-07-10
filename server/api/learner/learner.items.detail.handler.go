package learner

import (
	"contract"
	"encoding/json"
	"fmt"
	"github.com/Andrew-M-C/go.jsonvalue"
	"github.com/bykovme/gotrans"
	"net/http"
	"schema"
)

/*
 * GET: 当前用户为 person, 返回该 `item` 的详情及个人进度
 */
func ItemQuery(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	cookie, _ := r.Cookie("uid")
	if cookie == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		query := r.URL.Query()
		id := query.Get("item")
		resBytes, err := contract.Query([]string{id}, "queryItem", contract.Peer0)
		if err != nil {
			return &schema.ServerError{Error: err, Message: gotrans.Tr("fr", "cnte"), Code: 500}
		}
		res, _ := jsonvalue.Unmarshal(resBytes)
		info, _ := res.GetString("Info")
		var item schema.Item
		err = json.Unmarshal([]byte(info), &item)
		if err != nil {
			return &schema.ServerError{Error: err, Message: gotrans.Tr("fr", "jse"), Code: 500}
		}
		itemBytes, _ := json.Marshal(item)
		_, err = fmt.Fprintf(w, string(itemBytes))
		if err != nil {
			return &schema.ServerError{Error: err, Message: gotrans.Tr("fr", "fpe"), Code: 500}
		}
	}
	return nil
}