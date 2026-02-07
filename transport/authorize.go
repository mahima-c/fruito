package transport

import (
	"github.com/Mrhb787/hospital-ward-manager/common"
)

var optionsWithoutRouteCheck = []common.ServerOption{
	common.ServerErrorEncoder(common.DefaultErrorEncoder),
}

var optionsWithoutAuth = []common.ServerOption{
	common.ServerErrorEncoder(common.DefaultErrorEncoder),
}
