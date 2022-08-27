package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func init() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	err := v.RegisterValidation("birthday", birthday)
	if err != nil {
		log.Error("register birthday error", err)
	}
}
