package main

import (
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server.com/api/common"
	"server.com/api/learner"
	"server.com/api/organization"
	"server.com/api/registration"
	"server.com/schema"
)

type serverHandler func(http.ResponseWriter, *http.Request) *schema.ServerError

func (sh serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := sh(w, r); err != nil {
		fmt.Printf("%v", err.Error)
		http.Error(w, err.Message, err.Code)
	}
}

func main() {
	// 初始化 locales
	if err := gotrans.InitLocales("locales"); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Handle("/login", serverHandler(registration.Login))
	r.Handle("/register", serverHandler(registration.Register))
	r.Handle("/logout", serverHandler(registration.Logout))
	r.Handle("/", serverHandler(common.Index))
	r.Handle("/query_certs", serverHandler(common.QueryCerts))
	r.Handle("/learner/{learner_id}/items", serverHandler(learner.Item))
	r.Handle("/learner/{learner_id}/items/{item_id}", serverHandler(learner.ItemQuery))
	r.Handle("/learner/{learner_id}/my_items", serverHandler(learner.MyItem))
	r.Handle("/learner/{learner_id}/my_items/{item_id}", serverHandler(learner.ItemQuery))
	r.Handle("/learner/{learner_id}/certs", serverHandler(learner.Certs))
	r.Handle("/org/{org_id}/items", serverHandler(organization.IssuerItems))
	r.Handle("/org/{org_id}/items/{item_id}", serverHandler(organization.ItemQuery))
	r.Handle("/org/{org_id}/items/add", serverHandler(organization.AddItem))
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}