package main

type Resume struct {
	FullName        string
	Links           [4]Link
	About           string
	Experiences     []Experience
	Educations      []Education
	AdditionalWorks []AdditionalWork
	SkillGroups     []SkillGroup
	CompanyName     string
}

type Link struct {
	Name string
	Href string
}

type Experience struct {
	Position     string
	Period       Period
	Organization string
	Place        string
	Descriptions []string
}

type Period struct {
	Start string
	End   string
}

type Education struct {
	SchoolName string
	Period     Period
	Degree     string
	Place      string
}

type AdditionalWork struct {
	Name         string
	Organization string
	Date         string
}

type SkillGroup struct {
	Category string
	Skills   []string
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
