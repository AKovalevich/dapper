package service

type Service struct {
	ServiceBaseModel
	Routes func () ([]Route, error)
}
