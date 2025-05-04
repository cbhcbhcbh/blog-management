package post

import (
	"github.com/gin-gonic/gin"

	"blog/internal/pkg/core"
	"blog/internal/pkg/known"
	"blog/internal/pkg/log"
)

func (ctrl *PostController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete post function called")

	if err := ctrl.b.Posts().Delete(c, c.GetString(known.XUsernameKey), c.Param("postID")); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
