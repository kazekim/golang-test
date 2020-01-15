package kazekiminterface

type Test interface {
	Run(str string)
	GetNumber(str string) (*int, error)
}
