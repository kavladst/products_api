package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/kavladst/products_api/internal/app/storage"
)

var db = gorm.DB{}

func init() {
	pointerDB, err := storage.InitDB(
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		panic(err)
	}
	db = *pointerDB
}

func getRequestTask(response *http.Response) (*http.Request, error) {
	answerBody, err := getAnswer(response, 200)
	if err != nil {
		return nil, err
	}
	mapBody := map[string]string{}
	if err := json.Unmarshal([]byte(answerBody), &mapBody); err != nil {
		return nil, err
	}
	return getRequest("GET", "/v1/tasks", nil, mapBody, ""), nil
}

func getRequest(method string, path string, bodyParams map[string]string, queryParams map[string]string, filePath string) *http.Request {
	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)

	if bodyParams != nil {
		for bodyParam, bodyValue := range bodyParams {
			_ = writer.WriteField(bodyParam, bodyValue)
		}
	}

	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		part3, err := writer.CreateFormFile("file", filepath.Base(filePath))
		_, err = io.Copy(part3, file)
		if err != nil {
			panic(err)
		}

	}
	err := writer.Close()
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(method, ApiURL+path, &body)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	if queryParams != nil {
		query := request.URL.Query()
		for queryParam, queryValue := range queryParams {
			query.Add(queryParam, queryValue)
		}
		request.URL.RawQuery = query.Encode()
	}
	return request
}

func checkTaskResponse(client *http.Client, request *http.Request, expectedBody string) error {
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	requestTask, err := getRequestTask(response)
	if err != nil {
		return err
	}
	for i := 0; i < RetriesCount; i++ {
		responseTask, err := client.Do(requestTask)
		if err != nil {
			return err
		}
		answerTaskString, err := getAnswer(responseTask, 200)
		if err != nil {
			return err
		}
		answerTask := map[string]interface{}{}
		if err := json.Unmarshal([]byte(answerTaskString), &answerTask); err != nil {
			return err
		}
		if answerTask["result"] == nil {
			if i == RetriesCount-1 {
				return errors.New("timeout task")
			} else {
				time.Sleep(time.Duration(SecondsWait) * time.Second)
			}
		} else {
			if answerTaskString == expectedBody {
				return nil
			} else {
				return errors.New("Assertion error: \n\tAnswer:\t\t" + answerTaskString + "\n\tExpected:\t" + expectedBody)
			}
		}
	}
	return errors.New("task not completed")
}

func getAnswer(response *http.Response, expectedCode int) (string, error) {
	var errStatus error
	if response.StatusCode != expectedCode {
		errStatus = fmt.Errorf("status: %d\nexpected status: %d", response.StatusCode, expectedCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		if errStatus != nil {
			return "", errStatus
		}
		return "", err
	}
	if errStatus != nil {
		return "", fmt.Errorf("status: %d\nexpected status: %d\nbody: %s", response.StatusCode, expectedCode, string(body))
	}
	return string(body), nil
}

func clearTables() {
	db.Exec("TRUNCATE TABLE products CASCADE")
	db.Exec("TRUNCATE TABLE users CASCADE")
}
