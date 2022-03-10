package entity

type FollowRequest struct {
	FollowedID string `json:"followed_id"`
}

type UnfollowRequest struct {
	UnfollowedID string `json:"unfollowed_id"`
}
