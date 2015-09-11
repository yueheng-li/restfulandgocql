# restfulandgocql

createdata.goを実行し、スペースとテーブルを作成し、データをデータベースへ導入しておいてください。
package main
import (
	"net/http"
	"log"
	"accenrkt"
)

func main() {

	router := accenrky.NewRouter()


	log.Fatal(http.ListenAndServe(":8080", router))
}
