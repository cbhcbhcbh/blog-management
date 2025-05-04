package post

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/known"
	"blog/internal/pkg/log"
	v1 "blog/pkg/api/blog/v1"
)

func (ctrl *PostController) Create(c *gin.Context) {
	log.C(c).Infow("Create post function called")

	var r v1.CreatePostRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	resp, err := ctrl.b.Posts().Create(c, c.GetString(known.XUsernameKey), &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
