package ml1gongjustests

import (
	"log"
	"testing"
	"xm0jichu/ml0gongjus"
)

func TestMd5JiaMi(t *testing.T) {
	ret := ml0gongjus.Md5JiaMi("123abc")
	log.Println(ret, len(ret))
}
