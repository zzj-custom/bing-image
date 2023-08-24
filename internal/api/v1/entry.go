package v1

import "image/internal/api/v1/image"

type ApiGroup struct {
	image.Api
}

var ApiGroupApp = new(ApiGroup)
