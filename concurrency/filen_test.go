package concurrency

import "testing"

// Men använd
//  https://github.com/stretchr/testify

func TestVerifyAllBlablaIsCorrect(t *testing.T) {
	result := Run4()

	if result != 23 {
		t.Fail()
	}

}

func BenchmarkRun4(t *testing.B) {
	Run4()
}
