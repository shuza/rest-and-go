package main

import (
	"fmt"
	"github.com/smartystreets/assertions"
	"testing"
)

/**
 * :=  created by:  Shuza
 * :=  create date:  11-Apr-2019
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

func TestTmp(t *testing.T) {
	a := assertions.ShouldEqual(2, 2)
	fmt.Println("a : " + a)
}
