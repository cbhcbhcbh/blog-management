package post

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/likexian/gokit/assert"
	"go.uber.org/mock/gomock"

	"blog/internal/blog/biz"
	"blog/internal/blog/biz/post"
	v1 "blog/pkg/api/blog/v1"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func TestPostController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := &v1.CreatePostResponse{PostID: "post-22vtll"}

	mockPostBiz := post.NewMockPostBiz(ctrl)
	mockBiz := biz.NewMockIBiz(ctrl)
	mockPostBiz.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(want, nil).Times(1)
	mockBiz.EXPECT().Posts().AnyTimes().Return(mockPostBiz)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	body := bytes.NewBufferString(`{"title":"blog installation guide","content":"The installation method is coming."}`)
	c.Request, _ = http.NewRequest("POST", "/v1/posts", body)
	c.Request.Header.Set("Content-Type", "application/json")

	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = blw

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
		want   *v1.CreatePostResponse
	}{
		{
			name:   "default",
			fields: fields{b: mockBiz},
			args: args{
				c: c,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &PostController{
				b: tt.fields.b,
			}
			ctrl.Create(tt.args.c)
			var resp v1.CreatePostResponse
			err := json.Unmarshal(blw.body.Bytes(), &resp)
			assert.Nil(t, err)
			assert.Equal(t, resp.PostID, want.PostID)
		})
	}
}
