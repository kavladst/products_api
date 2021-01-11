package test

import (
	"net/http"
	"testing"
)

func TestApi_CreateUser(t *testing.T) {
	client := &http.Client{}
	clearTables()

	request := getRequest("POST", "/v1/users", CreateUser1InputBody, nil, "")
	err := checkTaskResponse(client, request, CreateUser1Answer)
	if err != nil {
		t.Error(err)
		return
	}
	request = getRequest("POST", "/v1/users", CreateUser1InputBody, nil, "")
	err = checkTaskResponse(client, request, CreateExitingUserAnswer)
	if err != nil {
		t.Error(err)
		return
	}
}
