package main

type ProjectCard struct {
	ImgDir             string `json:"imgDir"`
	ProjectName        string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
	ReadMore           bool   `json:"readMore"`
	ReadMoreLink       string `json:"readMoreLink"`
	Github             bool   `json:"github"`
	GithubLink         string `json:"githubLink"`
}
