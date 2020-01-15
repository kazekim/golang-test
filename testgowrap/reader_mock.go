package testgowrap

import "io"

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/13fd9e8d5b8ef5edc45e6d87c2b7716e7f989b63/templates/validate template

//go:generate gowrap gen -p io -i Reader -t https://raw.githubusercontent.com/hexdigest/gowrap/13fd9e8d5b8ef5edc45e6d87c2b7716e7f989b63/templates/validate -o reader_mock.go

// ReaderWithValidation implements io.Reader interface instrumented with arguments validation
type ReaderWithValidation struct {
	io.Reader
}

// NewReaderWithValidation returns ReaderWithValidation
func NewReaderWithValidation(base io.Reader) ReaderWithValidation {
	return ReaderWithValidation{
		Reader: base,
	}
}

// Read implements io.Reader
func (_d ReaderWithValidation) Read(p []byte) (n int, err error) {

	if _v, _ok := interface{}(p).(interface{ Validate() error }); _ok {
		if err = _v.Validate(); err != nil {
			return
		}
	}

	return _d.Reader.Read(p)
}
