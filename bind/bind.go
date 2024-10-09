package bind

import (
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Bind 参数校验
func Bind(c *gin.Context, d interface{}, bindings ...binding.Binding) error {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = c.ShouldBindUri(d)
		} else {
			err = c.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			err = nil
			continue
		}
		if err != nil {
			return err
		}
	}

	if err1 := vd.Validate(d); err1 != nil {
		return err1
	}

	return nil
}
