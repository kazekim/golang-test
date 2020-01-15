package testgowrap

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/13fd9e8d5b8ef5edc45e6d87c2b7716e7f989b63/templates/validate template

//go:generate gowrap gen -p github.com/kazekim/golang-test/testgowrap -i Test -t https://raw.githubusercontent.com/hexdigest/gowrap/13fd9e8d5b8ef5edc45e6d87c2b7716e7f989b63/templates/validate -o test_mock.go

// TestWithValidation implements Test interface instrumented with arguments validation
type TestWithValidation struct {
	Test
}

// NewTestWithValidation returns TestWithValidation
func NewTestWithValidation(base Test) TestWithValidation {
	return TestWithValidation{
		Test: base,
	}
}

// GetNumber implements Test
func (_d TestWithValidation) GetNumber(str string) (ip1 *int, err error) {

	if _v, _ok := interface{}(str).(interface{ Validate() error }); _ok {
		if err = _v.Validate(); err != nil {
			return
		}
	}

	return _d.Test.GetNumber(str)
}

// Run implements Test
func (_d TestWithValidation) Run(str string) {
	_d.Test.Run(str)
	return
}
