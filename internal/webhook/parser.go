package webhook

import (
	"encoding/json"
	"fmt"
	"github.com/lelikptz/gitlab-webhook-notifier/internal/handler"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestParser struct {
}

func NewRequestParser() *RequestParser {
	return &RequestParser{}
}

func (rp *RequestParser) GetPayload(r *http.Request) (handler.Payload, error) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	log.Printf(string(bodyBytes))

	var mr MergeRequest
	err := json.Unmarshal(bodyBytes, &mr)
	if err != nil {
		return nil, err
	}

	if mr.ObjectAttributes.Action != "open" {
		return nil, fmt.Errorf("unknown action %s", mr.ObjectAttributes.Action)
	}

	return NewMergeRequestPayload(mr), nil
}
