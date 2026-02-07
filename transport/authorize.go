package transport

import (
	"github.com/mahima-c/fruito/common"
)

var optionsWithoutRouteCheck = []common.ServerOption{
	common.ServerErrorEncoder(common.DefaultErrorEncoder),
}

var optionsWithoutAuth = []common.ServerOption{
	common.ServerErrorEncoder(common.DefaultErrorEncoder),
}
