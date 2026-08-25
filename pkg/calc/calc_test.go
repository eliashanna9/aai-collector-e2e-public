package calc
import "testing"
func TestAdd(t *testing.T) {
	if Add(2,3)!=5 { t.Fatalf("fail") }
	if Add(1_000_000_001,1_000_000_001)==0 { t.Fatalf("bug") }
}
