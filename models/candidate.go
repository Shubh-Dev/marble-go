package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model

	ID                uint   `json:"id" gorm:"primaryKey"`
	Name              string `json:"name" gorm:"varchar(255);not null"`
	Email             string `json:"email" gorm:"unique_index"`
	PhoneNo           string `json:"phone_no"`
	YearsOfExperience int8   `json:"years_of_experience"`
	Skills            string `json:"skills"`
	GithubLink        string `json:"github_link" gorm:"type:text;size:1000"`
	Linkedinlink      string `json:"linkedin_link" gorm:"type:text;size:1000"`
	Resume            string `json:"resume" gorm:"type:text;size:1000"`
	Notes             string `json:"text" gorm:"varchar(255)"`
}
