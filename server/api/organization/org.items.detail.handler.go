package organization

import (
	"server.com/api/contract"
	"encoding/json"
	"fmt"
	"github.com/Andrew-M-C/go.jsonvalue"
	"net/http"
	"server.com/schema"
)

/*
 * GET: 当前用户为 issuer, 返回该 `item` 的详情及所有参与该 `item` 的学生的进度
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
			return &schema.ServerError{err, "queryItem error", 500}
		}
		res, _ := jsonvalue.Unmarshal(resBytes)
		info, _ := res.GetString("Info")
		var item schema.Item
		err = json.Unmarshal([]byte(info), &item)
		if err != nil {
			return &schema.ServerError{err, "json.Unmarshal error", 500}
		}
		itemBytes, _ := json.Marshal(item)
		_, err = fmt.Fprintf(w, string(itemBytes))
		if err != nil {
			return &schema.ServerError{err, "fmt.Fprintf error", 500}
		}
	}
	return nil
}