package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/log"
	v1 "blog/pkg/api/blog/v1"
)

func (ctrl *UserController) Update(c *gin.Context) {
	log.C(c).Infow("Update user function called")

	var r v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Users().Update(c, c.Param("name"), &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
