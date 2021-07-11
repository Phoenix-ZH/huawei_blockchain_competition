package common

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
 * GET: 根据给定的 Hash 值查询对应的证书列表
 */
func QueryCerts(w http.ResponseWriter, r *http.Request) *schema.ServerError {
	query := r.URL.Query()
	hash := query.Get("hash")
	resByte, _ := contract.Query([]string{hash}, "queryCert", contract.Peer0)
	res, _ := jsonvalue.Unmarshal(resByte)
	info, _ := res.GetString("Info")
	var cert schema.Cert
	_ = json.Unmarshal([]byte(info), &cert)
	certBytes, _ := json.Marshal(cert)
	_, err := fmt.Fprintf(w, string(certBytes))
	if err != nil {
		return &schema.ServerError{Error: err, Message: gotrans.Tr("fr", "fpe"), Code: 500}
	}
	return nil
}