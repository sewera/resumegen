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
	AccentColor     RGB              `yaml:"accentColor,omitempty"`
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

type RGB struct {
	R uint8 `yaml:"r"`
	G uint8 `yaml:"g"`
	B uint8 `yaml:"b"`
}

func FromYAML(r io.Reader) Resume {
	decoder := yaml.NewDecoder(r)
	var parsed Resume
	err := decoder.Decode(&parsed)
	if err != nil {
		panic(err)
	}
	parsed = verifyColor(parsed)
	return parsed
}

func verifyColor(resume Resume) Resume {
	if resume.AccentColor.R == 0 &&
		resume.AccentColor.G == 0 &&
		resume.AccentColor.B == 0 {
		resume.AccentColor = RGB{R: 255, G: 255, B: 255}
	}
	return resume
}
