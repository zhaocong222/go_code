package main

import "testing"

//编写测试用例，去测试addUpper是否正确
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("执行错误")
	}

	//正确，输出日志
	t.Logf("执行正确")
}
