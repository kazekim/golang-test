package testgowrap

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/13fd9e8d5b8ef5edc45e6d87c2b7716e7f989b63/templates/validate template

//go:generate gowrap gen -p github.com/kazekim/golang-test/testgowrap -i Test -t https://raw.githubusercontent.com/hexdigest/gowrap/13fd9e8d5b8ef5edc45e6d87c2b7716e7f989b63/templates/validate -o test_mock2.go -v DecoratorName=KimTest -v Name=Kim

// KimTest implements Test interface instrumented with arguments validation
type KimTest struct {
	Test
}

// NewKimTest returns KimTest
func NewKimTest(base Test) KimTest {
	return KimTest{
		Test: base,
	}
}

// GetNumber implements Test
func (_d KimTest) GetNumber(str string) (ip1 *int, err error) {

	if _v, _ok := interface{}(str).(interface{ Validate() error }); _ok {
		if err = _v.Validate(); err != nil {
			return
		}
	}

	return _d.Test.GetNumber(str)
}

// Run implements Test
func (_d KimTest) Run(str string) {
	_d.Test.Run(str)
	return
}
