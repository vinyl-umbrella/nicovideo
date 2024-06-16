package domain

type Video struct {
	Vid           int    `json:"vid"`
	Title         string `json:"title"`
	OwnerNickname string `json:"ownerNickname"`
	Description   string `json:"description"`
	Duration      int    `json:"duration"`
	RegisteredAt  string `json:"registeredAt"`
	ViewCount     int    `json:"viewCount"`
	CommentCount  int    `json:"commentCount"`
	MylistCount   int    `json:"mylistCount"`
	LikeCount     int    `json:"likeCount"`
}
