package captcha

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type capSolverResponse struct {
	ErrorId          int32          `json:"errorId"`
	ErrorCode        string         `json:"errorCode"`
	ErrorDescription string         `json:"errorDescription"`
	TaskId           string         `json:"taskId"`
	Status           string         `json:"status"`
	Solution         map[string]any `json:"solution"`
}

func capSolver(ctx context.Context, apiKey string, taskData map[string]any) (*capSolverResponse, error) {
	uri := "https://api.capsolver.com/createTask"
	res, err := request(ctx, uri, map[string]any{
		"clientKey": apiKey,
		"task":      taskData,
	})
	if err != nil {
		return nil, err
	}
	if res.ErrorId == 1 {
		return nil, errors.New(res.ErrorDescription)
	}

	uri = "https://api.capsolver.com/getTaskResult"
	for {
		res, err = request(ctx, uri, map[string]any{
			"clientKey": apiKey,
			"taskId":    res.TaskId,
		})
		if err != nil {
			return nil, err
		}
		if res.ErrorId == 1 {
			return nil, errors.New(res.ErrorDescription)
		}
		if res.Status == "ready" {
			return res, err
		}
	}
}

func request(ctx context.Context, uri string, payload interface{}) (*capSolverResponse, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", uri, bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	capResponse := &capSolverResponse{}
	err = json.Unmarshal(responseData, capResponse)
	if err != nil {
		return nil, err
	}
	return capResponse, nil
}

func SolveV3() (map[string]any, error) {
	// fmt.Println("solving v3 captcha...")
	apikey := ""
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	res, err := capSolver(ctx, apikey, map[string]any{
		"type":       "ReCaptchaV3EnterpriseTaskProxyLess",
		"websiteURL": "https://www.samsung.com/us/",
		"websiteKey": "6LfFeqMfAAAAAPTHzfiw74_7tDqsXa2Qq4UgZnQJ",
	})
	if err != nil {
		panic(err)
	}
	return res.Solution, nil
}

func SolveV2() (map[string]any, error) {
	fmt.Println("solving v2 captcha...")
	apikey := ""

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	res, err := capSolver(ctx, apikey, map[string]any{
		"type":       "RecaptchaV2EnterpriseTaskProxyless",
		"websiteURL": "https://www.samsung.com/us/",
		"websiteKey": "6Le3CGAUAAAAADgwHP5vnwfsLbIOoxnh07DcaMcq",
	})
	if err != nil {
		panic(err)
	}
	return res.Solution, nil

}
