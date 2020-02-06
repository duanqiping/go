package cal

import "testing"

//编写要给测试用例，去测试addUpper是否正确
func TestGetSub(t *testing.T)  {
	res := getSub(20,10)

	if res != 10 {
		t.Fatalf("getSub(20,10) 期望值为=%v, 实际值为=%v",10,res)
	}

	t.Logf("getSub(20,10) 执行正确……")
}
