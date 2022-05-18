package webhook

import (
	"encoding/json"
	"fmt"
	"github.com/lelikptz/gitlab-webhook-notifier/internal/operation/webhook/event"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestParser struct {
}

func NewRequestParser() *RequestParser {
	return &RequestParser{}
}

func (rp *RequestParser) GetPayload(r *http.Request) (*Payload, error) {
	mr, err := getEvent(r.Header.Get("X-Gitlab-Event"))
	if err != nil {
		return nil, fmt.Errorf("get event error %s", err)
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	log.Printf(string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &mr.Event)
	if err != nil {
		return nil, err
	}

	if !mr.Event.IsSuccessful() {
		return nil, fmt.Errorf("unknown event structure")
	}

	return NewPayload(mr), nil
}

func getEvent(gitlabEvent string) (*Payload, error) {
	switch gitlabEvent {
	case "Merge Request Hook":
		return NewPayload(&event.MergeRequest{}), nil
	}

	return nil, fmt.Errorf("undefinder event")
}
