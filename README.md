# restfulandgocql

createdata.go�����s���A�X�y�[�X�ƃe�[�u�����쐬���A�f�[�^���f�[�^�x�[�X�֓������Ă����Ă��������B

���go�t�@�C�����쐬���A���L�̃R�[�h�����s���Ă݂Ă��������B
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
