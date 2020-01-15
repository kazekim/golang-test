package testgowrap

import "github.com/kazekim/golang-test/testgowrap/kazekiminterface"

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using templates/validate template

//go:generate gowrap gen -p github.com/kazekim/golang-test/testgowrap/kazekiminterface -i Test -t templates/validate -o test_mock.struct.go -v DecoratorName=KazekimTest -v Name=Kim

// KazekimTest implements kazekiminterface.Test interface instrumented with arguments validation
type KazekimTest struct {
	kazekiminterface.Test
}

// NewKazekimTest returns KazekimTest
func NewKazekimTest(base kazekiminterface.Test) KazekimTest {
	return KazekimTest{
		Test: base,
	}
}

// GetNumber implements kazekiminterface.Test
func (_d KazekimTest) GetNumber(str string) (ip1 *int, err error) {

	if _v, _ok := interface{}(str).(interface{ Validate() error }); _ok {
		if err = _v.Validate(); err != nil {
			return
		}
	}

	return _d.Test.GetNumber(str)
}

// Run implements kazekiminterface.Test
func (_d KazekimTest) Run(str string) {
	_d.Test.Run(str)
	return
}
