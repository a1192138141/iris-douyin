package datamodels

type VideoInfo struct {
	UserId int        `json:"user_id"` //用户id 
	Title string      `json:"title"` //标题
	Thumbnail string `json:"thumbnail"`  //缩略图
	Video string   `json:"video"`
	PlayNum int    `json:"play_num"`//播放次数
	LikeNum int    `json:"like_num"` //喜欢次数
	ForwardNum int `json:"forward_num"`//转发次数
	CommentNum int `json:"comment_num"`//评论条数
	Avatar string `json:"avatar"` //头像
}
