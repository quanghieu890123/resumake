package templates

import "text/template"

func Plaintext() *template.Template {
	tmpl, err := template.New("plaintext").Delims("[[", "]]").Parse(plaintext)
	if err != nil {
		panic(err)
	}

	return tmpl
}

var plaintext = `
[[- .Header.Name ]]
[[ .Header.Email ]]
==============================

EDUCATION
==============================
[[- range $eduEntry := .EducationEntries ]]
[[ $eduEntry.Degree ]], [[ $eduEntry.School ]], [[ $eduEntry.TimeSpan.Display ]]
[[ if $eduEntry.GPA ]]GPA: [[ $eduEntry.GPA ]][[ "\n" ]][[ end ]]
[[- end ]]
PROFESSIONAL EXPERIENCE
===============================
[[- range $jobEntry := .JobEntries ]]
[[ $jobEntry.Title ]], [[ $jobEntry.Employer ]], [[ $jobEntry.Location ]], [[ $jobEntry.TimeSpan.Display ]] 
[[- range $bullet := $jobEntry.Bullets ]]
* [[ $bullet ]] 
[[- end ]]
Technologies used: [[ $jobEntry.Skills.Display ]][[ "\n" ]]
[[- end ]]
SKILLS
==============================
Languages: [[ .Languages.Display ]]
Technologies: [[ .Technologies.Display ]]

PROJECTS
==============================
[[- range $project := .Projects ]]
[[ $project.Name ]]
[[ $project.Description ]]
Technologies used: [[ $project.Skills.Display ]][[ "\n" ]]
[[- end ]]`