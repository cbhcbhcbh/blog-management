package post

import (
	"testing"

	"github.com/gin-gonic/gin"

	"blog/internal/blog/biz"
)

func TestPostController_List(t *testing.T) {
	type fields struct {
		b biz.IBiz
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &PostController{
				b: tt.fields.b,
			}
			ctrl.List(tt.args.c)
		})
	}
}
