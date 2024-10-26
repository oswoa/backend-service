package model

type UserAuthority struct {
	AuthorityName string
	CanSearch     bool
	CanCreate     bool
	CanUpdate     bool
	CanDelete     bool
	CanApprove    bool
	CanPullBack   bool
}
