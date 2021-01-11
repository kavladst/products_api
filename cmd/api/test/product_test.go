package test

import (
	"net/http"
	"testing"
)

func TestApi_UpdateProducts(t *testing.T) {
	client := &http.Client{}
	clearTables()

	request := getRequest("POST", "/v1/users", CreateUser1InputBody, nil, "")
	err := checkTaskResponse(client, request, CreateUser1Answer)
	if err != nil {
		t.Error(err)
		return
	}

	postProducts1Query := map[string]string{"user_id": CreateUser1InputBody["id"]}
	request = getRequest("POST", "/v1/products/update/xlsx", nil, postProducts1Query, Data1Path)
	err = checkTaskResponse(client, request, PostData1Answer)
	if err != nil {
		t.Error(err)
		return
	}

	request = getRequest("GET", "/v1/products", nil, nil, "")
	err = checkTaskResponse(client, request, GetData1Answer)
	if err != nil {
		t.Error(err)
		return
	}

	request = getRequest("POST", "/v1/products/update/xlsx", nil, postProducts1Query, Data12Path)
	err = checkTaskResponse(client, request, PostData12Answer)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestApi_UpdateProductsWithNotExistedUser(t *testing.T) {
	client := &http.Client{}
	clearTables()

	request := getRequest("POST", "/v1/users", CreateUser1InputBody, nil, "")
	err := checkTaskResponse(client, request, CreateUser1Answer)
	if err != nil {
		t.Error(err)
		return
	}

	postProducts1Query := map[string]string{"user_id": "not_exist_id"}
	request = getRequest("POST", "/v1/products/update/xlsx", nil, postProducts1Query, Data1Path)
	err = checkTaskResponse(client, request, PostDataOfNotExistedUser)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestApi_UpdateProductsBadData(t *testing.T) {
	client := &http.Client{}
	clearTables()

	request := getRequest("POST", "/v1/users", CreateUser1InputBody, nil, "")
	err := checkTaskResponse(client, request, CreateUser1Answer)
	if err != nil {
		t.Error(err)
		return
	}

	postProducts1Query := map[string]string{"user_id": CreateUser1InputBody["id"]}
	request = getRequest("POST", "/v1/products/update/xlsx", nil, postProducts1Query, BadDataPath)
	err = checkTaskResponse(client, request, PostBadDataAnswer)
	if err != nil {
		t.Error(err)
		return
	}
}
