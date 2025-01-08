package custom

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Login() (*http.Request, error) {
	loginData := map[string]string{
		"username": "test_simulator",
		"password": "test_simulator",
	}
	payload, _ := json.Marshal(loginData)
	req, err := http.NewRequest("POST", "https://yourserver.com/api/login", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func Example_01() (*http.Request, error) {
	return nil, nil
}

func Example_02() (*http.Request, error) {
	return nil, nil
}
func Example_03() (*http.Request, error) {
	return nil, nil
}
func Example_04() (*http.Request, error) {
	return nil, nil
}
func Example_05() (*http.Request, error) {
	return nil, nil
}
func Example_06() (*http.Request, error) {
	return nil, nil
}
func Example_07() (*http.Request, error) {
	return nil, nil
}
func Example_08() (*http.Request, error) {
	return nil, nil
}
