package controller

import (
	"html/template"
	"fmt"
	"net/http"
	"demo/model"
)

var (
	TPL *template.Template
)

func init() {
	TPL = template.Must(template.ParseGlob("./view/tpl/*.html"))
}
//https://colobu.com/2016/10/09/Go-embedded-template-best-practices/
//http://www.prochainsci.com/2017/06/golang-template.html

//HTTP Method Controller
func CatGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
	}
	
	result, err := model.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	fmt.Println(result)
	TPL.ExecuteTemplate(w, "index.html", result)

	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(result)
	
}



