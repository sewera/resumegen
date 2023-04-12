package main

import (
	"gopkg.in/yaml.v3"
	"io"
)

type Resume struct {
	FullName        string           `yaml:"fullName"`
	Links           [4]Link          `yaml:"links"`
	About           string           `yaml:"about"`
	Experiences     []Experience     `yaml:"experience"`
	Educations      []Education      `yaml:"education"`
	AdditionalWorks []AdditionalWork `yaml:"additionalWorks"`
	SkillGroups     []SkillGroup     `yaml:"skillGroups"`
	CompanyName     string           `yaml:"companyName"`
}

type Link struct {
	Name string `yaml:"name"`
	Href string `yaml:"href"`
}

type Experience struct {
	Position     string   `yaml:"position"`
	Period       Period   `yaml:"period"`
	Organization string   `yaml:"organization"`
	Place        string   `yaml:"place"`
	Descriptions []string `yaml:"descriptions"`
}

type Period struct {
	Start string `yaml:"start"`
	End   string `yaml:"end"`
}

type Education struct {
	SchoolName string `yaml:"schoolName"`
	Period     Period `yaml:"period"`
	Degree     string `yaml:"degree"`
	Place      string `yaml:"place"`
}

type AdditionalWork struct {
	Name         string `yaml:"name"`
	Organization string `yaml:"organization"`
	Date         string `yaml:"date"`
}

type SkillGroup struct {
	Category string   `yaml:"category"`
	Skills   []string `yaml:"skills"`
}

var Example = Resume{
	FullName: "Full Name",
	Links: [4]Link{
		{
			Name: "link 1",
			Href: "link1",
		},
		{
			Name: "link 2",
			Href: "link2",
		},
		{
			Name: "link 3",
			Href: "link3",
		},
		{
			Name: "link 4",
			Href: "link4",
		},
	},
	About: "Lorem ipsum",
	Experiences: []Experience{
		{
			Position: "Position",
			Period: Period{
				Start: "Start",
				End:   "End",
			},
			Organization: "Org",
			Place:        "Place",
			Descriptions: []string{
				"Desc",
			},
		},
	},
	Educations: []Education{
		{
			SchoolName: "School",
			Period: Period{
				Start: "Start",
				End:   "End",
			},
			Degree: "Degree",
			Place:  "Place",
		},
	},
	AdditionalWorks: []AdditionalWork{
		{
			Name:         "Additional Work",
			Organization: "Org",
			Date:         "Date",
		},
	},
	SkillGroups: []SkillGroup{
		{
			Category: "Category",
			Skills: []string{
				"Skill 1",
				"Skill 2",
			},
		},
	},
	CompanyName: "Company Name",
}

func FromYAML(r io.Reader) (parsed Resume) {
	decoder := yaml.NewDecoder(r)
	err := decoder.Decode(&parsed)
	if err != nil {
		panic(err)
	}
	return
}
