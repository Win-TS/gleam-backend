package group

type Role string

const (
	Admin     Role = "creator"
	Moderator Role = "co_leader"
	Member    Role = "member"
)