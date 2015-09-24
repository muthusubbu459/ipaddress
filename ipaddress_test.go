package main
import(
"testing"
)
func value()bool{
return true
}
func testisip(t *testing.T) {
expected :=value()
actual:=isip("127.0.0.1")
if actual!=expected{
t.Errorf("Test Failed")
}else {
t.Log("first test passed ")
}
//return value
}
