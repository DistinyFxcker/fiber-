package model

import db "go_web_example/gen/dal/model"

type AdminOutput struct {
}

type GetAdminInput struct {
	Id string `url:"id"`
}

type QueryAdminInput struct {
	KeyWord string
	Ids     []string
	Emails  []string
	State   int

	Limit  int
	Offset int
	Sort   string
	Order  string
}

type QueryAdminOutput struct {
	Total int64
	Count int64
	List  []*db.Admin
}
