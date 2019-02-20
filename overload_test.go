package httpsrv

import "testing"

func TestCheck(t *testing.T) {
	checker := NewDefaultOverloadChecker(10, "MY_SERVICE")
	err := checker.Check(nil)
	if err != nil {
		t.Errorf("failed,%v", err)
	}
}
