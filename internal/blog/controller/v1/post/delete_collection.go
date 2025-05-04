package post

import (
	"github.com/gin-gonic/gin"

	"blog/internal/pkg/core"
	"blog/internal/pkg/known"
	"blog/internal/pkg/log"
)

func (ctrl *PostController) DeleteCollection(c *gin.Context) {
	log.C(c).Infow("Batch delete post function called")

	postIDs := c.QueryArray("postID")
	if err := ctrl.b.Posts().DeleteCollection(c, c.GetString(known.XUsernameKey), postIDs); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
