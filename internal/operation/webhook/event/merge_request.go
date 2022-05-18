package event

import "fmt"

type MergeRequest struct {
	EventType        string                       `json:"event_type"`
	ObjectAttributes MergeRequestObjectAttributes `json:"object_attributes"`
	User             User                         `json:"user"`
	Project          Project                      `json:"project"`
}

type MergeRequestObjectAttributes struct {
	ID           int64  `json:"id"`
	TargetBranch string `json:"target_branch"`
	SourceBranch string `json:"source_branch"`
	Action       string `json:"action"`
	RequestUrl   string `json:"url"`
}

type User struct {
	Name string `json:"name"`
}

type Project struct {
	Name   string `json:"name"`
	WebUrl string `json:"web_url"`
}

func (mr *MergeRequest) IsSuccessful() bool {
	return mr.ObjectAttributes.Action == "open"
}

func (mr *MergeRequest) GetMessage() string {
	return fmt.Sprintf(
		"Внимание внимание!\nПользователь %v успешно создал новый <a href=\"%v\">merge request</a> из %v в %v в проекте <a href=\"%v\">%v</a>",
		mr.User.Name,
		mr.ObjectAttributes.RequestUrl,
		mr.ObjectAttributes.SourceBranch,
		mr.ObjectAttributes.TargetBranch,
		mr.Project.WebUrl,
		mr.Project.Name,
	)
}
