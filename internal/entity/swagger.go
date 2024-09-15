package entity

type SwaggerBaseResponseDTO struct {
	AppName string `json:"appName" example:"Customer Miscellaneous API"`
	Build   string `json:"build" example:"1"`
	ID      string `json:"id" example:"16ad78a0-5f8a-4af0-9946-d21656e718b5"`
	Version string `json:"version" example:"1.0.0"`
}

type SwaggerResponseOKDTO struct {
	SwaggerBaseResponseDTO
	Data    any    `json:"data"`
	Message string `json:"message" example:"Success"`
}

type SwaggerResponseCreatedDTO struct {
	SwaggerBaseResponseDTO
	Data    any    `json:"data"`
	Message string `json:"message" example:"Created"`
}

type SwaggerNoContentResponseDTO struct{}

type SwaggerResponseBadRequestDTO struct {
	SwaggerBaseResponseDTO
	//Will return null
	Data    any    `json:"data"`
	Message string `json:"message" example:"Bad Request"`
}

type SwaggerResponseUnauthorizedDTO struct {
	SwaggerBaseResponseDTO
	//Will return null
	Data    any    `json:"data"`
	Message string `json:"message" example:"Unauthorized"`
}

type SwaggerResponseInternalServerErrorDTO struct {
	SwaggerBaseResponseDTO
	//Will return null
	Data    any    `json:"data"`
	Message string `json:"message" example:"Internal Server Error"`
}

type SwaggerResponseForbiddenDTO struct {
	SwaggerBaseResponseDTO
	//Will return null
	Data    any    `json:"data"`
	Message string `json:"message" example:"Forbidden"`
}

type SwaggerResponseNotFoundDTO struct {
	SwaggerBaseResponseDTO
	//Will return null
	Data    any    `json:"data"`
	Message string `json:"message" example:"Not Found"`
}
