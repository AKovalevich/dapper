package routes

// Service base model definition, including fields `Name`, which could be embedded in your models
type RouteBaseModel struct {
	Pattern string
	Version int
}
