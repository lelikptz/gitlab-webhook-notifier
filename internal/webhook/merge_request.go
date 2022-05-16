package webhook

import "fmt"

type MergeRequestPayload struct {
	MergeRequest
}

func NewMergeRequestPayload(request MergeRequest) *MergeRequestPayload {
	return &MergeRequestPayload{request}
}

type MergeRequest struct {
	EventType        string           `json:"event_type"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

type ObjectAttributes struct {
	ID           int64  `json:"id"`
	TargetBranch string `json:"target_branch"`
	SourceBranch string `json:"source_branch"`
	Action       string `json:"action"`
}

func (mrp *MergeRequestPayload) GetMessage() string {
	return fmt.Sprintf(
		"Был создан новый merge request из %v в %v",
		mrp.ObjectAttributes.SourceBranch,
		mrp.ObjectAttributes.TargetBranch,
	)
}
