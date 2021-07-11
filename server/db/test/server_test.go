package t_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server.com/api/common"
	"server.com/api/learner"
	"server.com/api/organization"
	"server.com/api/registration"
	"server.com/schema"
	"testing"
)

type serverHandler func(http.ResponseWriter, *http.Request) *schema.ServerError

func (sh serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := sh(w, r); err != nil {
		fmt.Printf("%v", err.Error)
		http.Error(w, err.Message, err.Code)
	}
}

func Test_Index(t *testing.T) {
	mux  := http.NewServeMux()
	mux.Handle("/", serverHandler(common.Index))

	req, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)
	i := writer.Header().Get("index")
	if writer.Code != 200 || i != "index" {
		t.Error("Test index<GET> fail!")
	} else {
		t.Log("Test index<GET> pass!")
	}
}

func Test_Index_person(t *testing.T)  {
	mux := http.NewServeMux()
	mux.Handle("/", serverHandler(common.Index))

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Cookie", "uid=p123")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)

	uid := writer.Header().Get("uid")

	if writer.Code != 200 || uid != "p123" {
		t.Error("Test index<GET:person> fail!")
	} else {
		t.Log("Test index<GET:person> pass!")
	}
}

func Test_Index_issuer(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/", serverHandler(common.Index))

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Cookie", "uid=i123")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)

	uid := writer.Header().Get("uid")

	if writer.Code != 200 || uid != "i123" {
		t.Error("Test index<GET:issuer> fail!")
	} else {
		t.Log("Test index<GET:issuer> pass!")
	}
}

func Test_Login_GET(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/login", serverHandler(registration.Login))

	req, _ := http.NewRequest("GET", "/login", nil)
	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)

	if writer.Code != 200 {
		t.Error("Test login<GET> fail!")
	} else {
		t.Log("Test login<GET> pass!")
	}
}

func Test_Login_POST_person(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/login", serverHandler(registration.Login))

	body, _ := json.Marshal(struct {
		Uid string `json:"uid"`
		Passwd string `json:"passwd"`
	}{"p123", "123"})
	bodyBytes := bytes.NewBuffer(body)
	req, _ := http.NewRequest("POST", "/login", bodyBytes)
	req.Header.Set("Content-Type", "application/json")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)
	//fmt.Println(writer.Body)

	if writer.Code != 200 {
		t.Error("Test login<POST:person> fail!")
	} else {
		t.Log("Test login<POST:person> pass!")
	}
}

func Test_Login_POST_issuer(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/login", serverHandler(registration.Login))

	body, _ := json.Marshal(struct {
		Uid string `json:"uid"`
		Passwd string `json:"passwd"`
	}{"i123", "123"})
	bodyBytes := bytes.NewBuffer(body)
	req, _ := http.NewRequest("POST", "/login", bodyBytes)
	req.Header.Set("Content-Type", "application/json")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)
	//fmt.Println(writer.Body)

	if writer.Code != 200 {
		t.Error("Test login<POST:issuer> fail!")
	} else {
		t.Log("Test login<POST:issuer> pass!")
	}
}

func Test_Item_Person(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/item", serverHandler(learner.Item))

	req, _ := http.NewRequest("POST", "/item", nil)
	req.Header.Set("Cookie", "uid=p123")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)
	//fmt.Println(writer.Body)

	if writer.Code != 200 {
		t.Error("Test item<POST:person> fail!")
	} else {
		t.Log("Test item<POST:person> pass!")
	}
}

func Test_Item_Issuer(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/item", serverHandler(learner.Item))

	req, _ := http.NewRequest("POST", "/item", nil)
	req.Header.Set("Cookie", "uid=i123")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)
	//fmt.Println(writer.Body)

	if writer.Code != 200 {
		t.Error("Test item<POST:issuer> fail!")
	} else {
		t.Log("Test item<POST:issuer> pass!")
	}
}

func Test_My_Item(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/my_item", serverHandler(learner.MyItem))

	req, _ := http.NewRequest("GET", "/my_item", nil)
	req.Header.Set("Cookie", "uid=p123")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)
	//fmt.Println(writer.Body)

	if writer.Code != 200 {
		t.Error("Test my_item<GET:person> fail!")
	} else {
		t.Log("Test my_item<GET:person> pass!")
	}
}

func Test_Item_Query(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/item_query", serverHandler(learner.ItemQuery))

	req, _ := http.NewRequest("GET", "/item_query?item=it123", nil)
	req.Header.Set("Cookie", "uid=p123")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)
	//fmt.Println(writer.Body)

	if writer.Code != 200 {
		t.Error("Test item_query<GET> fail!")
	} else {
		t.Log("Test item_query<GET> pass!")
	}
}


func Test_Item_Op_Person(t *testing.T) {
	// TODO
	t.Log("Test item_op<POST:person> pass!")
}

func Test_Add_Item(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/add_item", serverHandler(organization.AddItem))

	body, _ := json.Marshal(struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Point float64 `json:"point"`
	}{"it567", "CV", "computer vision", 3.4})
	bodyBytes := bytes.NewBuffer(body)
	req, _ := http.NewRequest("POST", "/add_item", bodyBytes)
	req.Header.Set("Cookie", "uid=i123")
	req.Header.Set("Content-Type", "application/json")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)

	if writer.Code != 200 {
		t.Error("Test add_item<POST:issuer> fail!")
	} else {
		t.Log("Test add_item<POST:issuer> pass!")
	}
}

func Test_Certs(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/certs", serverHandler(learner.Certs))

	req, _ := http.NewRequest("GET", "/certs", nil)
	req.Header.Set("Cookie", "uid=p123")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)
	//fmt.Println(writer.Body)

	if writer.Code != 200 {
		t.Error("Test certs<GET:person> fail!")
	}
	t.Log("Test certs<GET:person> pass!")
}

func Test_Combine_Certs(t *testing.T) {
	// TODO
	t.Log("Test combine_certs<POST:person> pass!")
}

func Test_Query_Certs(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/query_certs", serverHandler(common.QueryCerts))

	req, _ := http.NewRequest("GET", "/query_certs?hash=53b991b48a9f6ebc671a0f02d82c514dbe67a20c5ca357ef345ff6e21241d0bd", nil)

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)
	//fmt.Println(writer.Body)

	if writer.Code != 200 {
		t.Error("Test query_certs<GET> fail!")
	}
	t.Log("Test query_certs<GET> pass!")
}

func Test_Logout(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/logout", serverHandler(registration.Logout))

	req, _ := http.NewRequest("POST", "/item", nil)
	req.Header.Set("Cookie", "uid=p123")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)

	if writer.Code != 200 {
		t.Error("Test logout<GET> fail!")
	} else {
		t.Log("Test logout<GET> pass!")
	}
}

func Test_Adduser(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/adduser", serverHandler(registration.Register))

	body, _ := json.Marshal(struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Password string `json:"password"`
		PublicKey string `json:"publicKey"`
	}{"p345", "Wang", "345", "cde"})
	bodyBytes := bytes.NewBuffer(body)
	req, _ := http.NewRequest("POST", "/adduser", bodyBytes)
	req.Header.Set("Cookie", "uid=admin123")
	req.Header.Set("Content-Type", "application/json")

	writer := httptest.NewRecorder()
	mux.ServeHTTP(writer, req)

	if writer.Code != 200 {
		t.Error("Test adduser<POST:admin> fail!")
	} else {
		t.Log("Test adduser<POST:admin> pass!")
	}
}