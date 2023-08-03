package campaign

import (
	"starco/user"
	"time"
)

type Campaign struct {
	ID               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatetAt time.Time
	User             user.User
	CampaignImages []CampaignImage
}

type CampaignImage struct{
	ID int
	CampainId int
	FileName string
	IsPrimary int
	CreatedAt        time.Time
	UpdatetAt time.Time
}