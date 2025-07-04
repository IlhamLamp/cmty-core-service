package dto

type CoreFilter struct {
	Page          int      `form:"page"`
	Limit         int      `form:"limit"`
	Types         string   `form:"types"`
	Duration      string   `form:"duration"`
	Participation string   `form:"participation"`
	Location      string   `form:"location"`
	Role          string   `form:"role"`
	Experience    string   `form:"experience"`
	Tags          []string `form:"tags[]"`
}

type MemberFilter struct {
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Types      string `form:"types"`
	Role       string `form:"role"`
	Experience string `form:"experience"`
}
