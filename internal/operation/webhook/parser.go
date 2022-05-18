package webhook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lelikptz/gitlab-webhook-notifier/internal/operation/webhook/event"
)

type RequestParser struct {
}

func NewRequestParser() *RequestParser {
	return &RequestParser{}
}

func (rp *RequestParser) GetPayload(r *http.Request) (*Payload, error) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	log.Printf(string(bodyBytes))

	var mr event.MergeRequest
	err := json.Unmarshal(bodyBytes, &mr)
	if err != nil {
		return nil, err
	}

	if mr.ObjectAttributes.Action != "open" {
		return nil, fmt.Errorf("unknown action %s", mr.ObjectAttributes.Action)
	}

	return NewPayload(&mr), nil
}
