package data

type Section struct {
	ID     string
	Name   string
	Description string
	NextSectionID string
	PrevSectionID string
	ChildSections []string
	Exercises: []string
	tags: []string
}
