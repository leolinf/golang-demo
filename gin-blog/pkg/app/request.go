package app

import (
	"gin-blog/pkg/logging"

	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Value)
	}
	return
}
