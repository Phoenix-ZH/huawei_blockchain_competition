package learner

import (
	"server.com/api/contract"
	"encoding/json"
	"fmt"
	"github.com/Andrew-M-C/go.jsonvalue"
	"github.com/bykovme/gotrans"
	"net/http"
	"server.com/schema"
)

/*
 * GET: 返回当前用户(person)的所有证书以及证书组合
 * 证书组合: 用户可以有选择的公开某些证书给某些组织或机构查看
 */
func Certs(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	var personCerts map[string][]string
	personCerts = make(map[string][]string)
	personCerts["p123"] = []string{
		"f892bb6e9ab49df3fd720cef1555f46c772dfdc9391daf39dbae0f38c60b9700",
		"fc131d21c767fea581402f1c9de74fa3614260ac7caa84074e863d5da83a82f1",
	}
	personCerts["p234"] = []string{
		"4e34bf3a3b764a6f9f1943473a811c72cdb76ca141eae166d267c9f25a5ad0ec",
		"53b991b48a9f6ebc671a0f02d82c514dbe67a20c5ca357ef345ff6e21241d0bd",
	}

	cookie, _ := r.Cookie("uid")
	if cookie == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		if certs := personCerts[cookie.Value]; certs != nil {
			resBytes, _ := contract.Query(certs, "queryCerts", contract.Peer0)
			res, _ := jsonvalue.Unmarshal(resBytes)
			info, _ := res.GetString("Info")
			certary, _ := jsonvalue.Unmarshal([]byte(info))
			var certs []schema.Cert
			for v := range certary.IterArray() {
				owener, _ := v.V.GetString("Owner")
				content, _ := v.V.GetString("Content")
				issuer, _ := v.V.GetString("Issuer")
				item, _ := v.V.GetString("Item")
				point, _ := v.V.GetFloat64("Point")
				date, _ := v.V.GetString("Date")
				sig, _ := v.V.GetString("Signature")
				cert := schema.Cert{
					Owner: owener,
					Content: content,
					Issuer: issuer,
					Item: item,
					Point: point,
					Date: date,
					Signature: sig,
				}
				certs = append(certs, cert)
			}
			certsBytes, _ := json.Marshal(certs)
			_, err := fmt.Fprintf(w, string(certsBytes))
			if err != nil {
				return &schema.ServerError{Error: err, Message: gotrans.Tr("fr", "fpe"), Code: 500}
			}
		}
	}
	return nil
}