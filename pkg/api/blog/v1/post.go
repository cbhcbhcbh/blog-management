package v1

type CreatePostRequest struct {
	Title   string `json:"title" valid:"required,stringlength(1|256)"`
	Content string `json:"content" valid:"required,stringlength(1|10240)"`
}

type CreatePostResponse struct {
	PostID string `json:"postID"`
}

type GetPostResponse PostInfo

type UpdatePostRequest struct {
	Title   *string `json:"title" valid:"stringlength(1|256)"`
	Content *string `json:"content" valid:"stringlength(1|10240)"`
}

type PostInfo struct {
	Username  string `json:"username,omitempty"`
	PostID    string `json:"postID,omitempty"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ListPostRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type ListPostResponse struct {
	TotalCount int64       `json:"totalCount"`
	Posts      []*PostInfo `json:"posts"`
}
