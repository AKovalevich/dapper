package routes

type Route struct {
	RouteBaseModel
	Structure func() (interface{}, error)
}
