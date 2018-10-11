package viewmodels

import "github.com/WayneShenHH/toolsgo/models"

// SwaggerModels collections of models
// swagger:parameters models
type SwaggerModels struct {
	UserLoginResult models.UserLoginResult
}

// SwaggerViewModels collections of view models
// swagger:parameters Params
type SwaggerViewModels struct {
	APIResult APIResult
}
