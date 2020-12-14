package main
//架構程序入口
import (
	"demo/model"
	"time"
	"net/http"
	"log"
	"runtime"
	"database/sql"

	"demo/controller"

	//a pure Go Postgres driver for the database/sql package
	_ "github.com/lib/pq"
)


func main() {

	//設置最大可同時執行CPU數量，NumCPU查詢本地可用CPU數量
	//func GOMAXPROCS: //https://golang.org/pkg/runtime/#GOMAXPROCS
	runtime.GOMAXPROCS(runtime.NumCPU())

	//設置處理程序
	//func HandleFunc: https://golang.org/pkg/net/http/#HandleFunc
	http.HandleFunc("/v1/getall/", controller.CatGetAll)
	// http.HandleFunc("/v1/getone/", controller.CatGetOne)

	//定義HTTP server 參數
	//type Server: https://golang.org/pkg/net/http/#Server
	s:= &http.Server{
		Addr:			":9090",
		ReadTimeout:	10*time.Second,
		WriteTimeout:	10*time.Second,
	}

	//ListenAndServe監聽 Addr 並呼叫Server處理requests
	//func (*Server) ListenAndServe: https://golang.org/pkg/net/http/#Server
	log.Fatal(s.ListenAndServe())
}

//程序初始化操作: 連接database
func init()  {
	//定義資料庫資訊
	connectStr := "host=localhost" +
		" port=5432" +
		" dbname=demo_db" +
		" user=demo_user" +
		" password='user_password'" +
		" sslmode=disable"

	//開啟指定資料庫連結
	//func Open: https://golang.org/pkg/database/sql/#Open
	var err error = nil
	model.Db, err = sql.Open("postgres", connectStr)
	if err != nil {
		log.Panic(err)
	}
}
