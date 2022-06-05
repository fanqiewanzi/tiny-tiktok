package constants

const (
	IdentityKey      = "id"
	UserTableName    = "sys_user"
	VideoTableName   = "sys_video"
	CommentTableName = "sys_comment"
	FollowTableName  = "sys_follow"
	Bucket           = "bazinga-photograph"
)

const (
	FollowKey       = "follow"
	FansKey         = "fans"
	FavoriteKey     = "favorite"
	FavoriteUserKey = "favorite:user"
)

const (
	Action       = "1"
	ActionCancel = "2"
)

const (
	CPURateLimit float64 = 80.0
	DefaultLimit         = 10
)
