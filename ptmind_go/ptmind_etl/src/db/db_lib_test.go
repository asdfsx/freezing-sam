package db

import (
	"fmt"
	"strconv"
	"testing"
)

DBADDR := "127.0.0.1:3306"
USER := ""
PASSWD := ""
DBNAME := "ptmind_common"
CHARSET := "utf-8"

func TestGetlogfileconfig(t *testing.T) {
	logfiles := Getlogfileconfig(DBADDR, USER, PASSWD, DBNAME, CHARSET)
	if logfiles {
		t.Log("queue length should be 10")
		t.Fail()
	}
}
