package ml1gongjustests

import (
	"log"
	"testing"
	"xm0jichu/ml0gongjus"
)

func TestShouJiHaoPiPei(t *testing.T) {
	ret := ml0gongjus.ShouJiHaoPiPei("1112222ff3333")
	log.Println(ret, len(ret))
}
