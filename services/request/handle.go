package request

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Handle(c *gin.Context, form Handler) (interface{}, error) {
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"err_msg": err.Error()})
		return nil, err
	}

	if err := check(form); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"err_msg": err.Error()})
		return nil, err
	}

	data, err := form.Do(c)
	if err == nil {
		return data, nil
	} else if biz, ok := err.(*BizErr); ok {
		c.JSON(biz.HttpCode, biz)
	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"err_msg": "Not Found"})
	} else {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": "Server Error"})
	}

	return data, err
}

func check(form Handler) error {
	checker, ok := form.(Checker)
	if !ok {
		return nil
	}

	if err := checker.Check(); err != nil {
		return err
	}

	return nil
}

type Handler interface {
	Do(c *gin.Context) (interface{}, error)
}

type Checker interface {
	Check() error
}
