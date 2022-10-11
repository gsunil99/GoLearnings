package calculator_test
import "testing"
type TestCase struct{
	value int
	actual bool
	expected bool
}
func TestCalculateArmgstrong(t *testing.T){
	testCase := TestCase{
		value:=143
		expected:=true
	}
	testCase.actual = calcuator.CalculateIsArmstrong(testCase.value)

	if testCase.actual != testCase.expected {
		t.Fail();
	}
}