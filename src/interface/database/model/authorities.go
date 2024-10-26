package model

import "time"

type Authorities struct {
	AuthorityName string
	CanSearch     bool
	CanCreate     bool
	CanUpdate     bool
	CanDelete     bool
	CanApprove    bool
	CanPullBack   bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
