package post

import (
	"github.com/gin-gonic/gin"

	"blog/internal/pkg/core"
	"blog/internal/pkg/errno"
	"blog/internal/pkg/known"
	"blog/internal/pkg/log"
	v1 "blog/pkg/api/blog/v1"
)

func (ctrl *PostController) List(c *gin.Context) {
	log.C(c).Infow("List post function called.")

	var r v1.ListPostRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	resp, err := ctrl.b.Posts().List(c, c.GetString(known.XUsernameKey), r.Offset, r.Limit)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
