package cal

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T)  {
	res := addUpper(10)

	if res != 55 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		t.Fatalf("addUpper(10) 执行错误，期望值=%v 实际值=%v \n",55,res)
	}
	//如果正确
	t.Logf("addUpper(10) 执行正确")
}

func TestHello(t *testing.T) {

	fmt.Println("TestHello被调用..")

}