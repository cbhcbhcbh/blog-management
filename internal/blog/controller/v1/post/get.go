package post

import (
	"github.com/gin-gonic/gin"

	"blog/internal/pkg/core"
	"blog/internal/pkg/known"
	"blog/internal/pkg/log"
)

func (ctrl *PostController) Get(c *gin.Context) {
	log.C(c).Infow("Get post function called")

	post, err := ctrl.b.Posts().Get(c, c.GetString(known.XUsernameKey), c.Param("postID"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, post)
}
