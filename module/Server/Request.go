package Server

import (
	"GOLONG/module/Resource"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Data struct {
	SearchBar bool
}

func Index(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadFile("./template/" + r.URL.Path[1:] + ".html")
	var t = template.New("index")
	var page string = ""
	var data = map[string]bool{
		"SearchBar": r.URL.Path[1:] == "",
	}

	if err != nil {
		page = Resource.OneLineIfElseStr(r.URL.Path[1:] == "", "./template/main.html", "./template/err404.html")
		t.ParseFiles("./template/index.html", "./template/partial/model.html", page)
	} else {
		t.ParseFiles("./template/index.html", "./template/partial/model.html", "./template/"+r.URL.Path[1:]+".html")
	}

	err = t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
