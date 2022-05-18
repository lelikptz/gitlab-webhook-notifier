package event

import "fmt"

type MergeRequest struct {
	EventType        string                       `json:"event_type"`
	ObjectAttributes MergeRequestObjectAttributes `json:"object_attributes"`
}

type MergeRequestObjectAttributes struct {
	ID           int64   `json:"id"`
	TargetBranch string  `json:"target_branch"`
	SourceBranch string  `json:"source_branch"`
	Action       string  `json:"action"`
	RequestUrl   string  `json:"url"`
	User         User    `json:"user"`
	Project      Project `json:"project"`
}

type User struct {
	Name string `json:"name"`
}

type Project struct {
	Name   string `json:"name"`
	WebUrl string `json:"web_url"`
}

func (mrp *MergeRequest) GetMessage() string {
	return fmt.Sprintf(
		"Внимание внимание!\nПользователь %v успешно создал новый <a href=\"%v\">merge request</a> из %v в %v в проекте <a href=\"%v\">%v</a>",
		mrp.ObjectAttributes.User.Name,
		mrp.ObjectAttributes.RequestUrl,
		mrp.ObjectAttributes.SourceBranch,
		mrp.ObjectAttributes.TargetBranch,
		mrp.ObjectAttributes.Project.WebUrl,
		mrp.ObjectAttributes.Project.Name,
	)
}
