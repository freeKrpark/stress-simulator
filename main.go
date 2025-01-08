package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"

	"github.com/freeKrpark/stress-simulator/custom"
)

var (
	// TODO: it gonna be set the results
	endpoints = map[string]struct {
		probability int
		Func        func() (*http.Request, error)
	}{
		"login":          {0, custom.Login},
		"API_example_01": {30, custom.Example_01},
		"API_example_02": {10, custom.Example_02},
		"API_example_03": {10, custom.Example_03},
		"API_example_04": {10, custom.Example_04},
		"API_example_05": {10, custom.Example_05},
		"API_example_06": {10, custom.Example_06},
		"API_example_07": {10, custom.Example_07},
		"API_example_08": {10, custom.Example_08},
	}
	numUsers     int           = 100
	testDuration time.Duration = 60 * time.Second
)

type UserSession struct {
	Client *http.Client
}

func performLogin() (*UserSession, error) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar; error is %v", err)
	}

	client := &http.Client{
		Jar: jar,
	}

	req, err := endpoints["login"].Func()
	if err != nil {
		return nil, fmt.Errorf("failed to create login request; error is %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to login; error is %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get 200 status code from login request; statusCode is %d", res.StatusCode)
	}

	log.Println("login success")
	return &UserSession{Client: client}, nil
}

// TODO: calculate duration timein that method
func performRequest(action string, session *UserSession) {
	endpoint, exist := endpoints[action]
	if !exist {
		log.Fatalf("failed to find the action, all actions must be filled.")
	}

	req, err := endpoint.Func()
	if err != nil {
		log.Printf("action %s : failed to create request; error is %v", action, err)
		return
	}

	res, err := session.Client.Do(req)
	if err != nil {
		// it gonna be failure results
		log.Printf("action %s : failed call action's api; error is %v", action, err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		// it gonnabe fail results
		log.Printf("Action: %s, Status Code: %d\n", action, res.StatusCode)
		return
	}
	// it gonnabe success results
	log.Printf("Action: %s, Status Code: %d\n", action, res.StatusCode)

}

func pickAction() string {
	total := 0
	for _, weight := range endpoints {
		total += weight.probability
	}

	r := rand.Intn(total)

	for action, weight := range endpoints {
		if r < weight.probability {
			return action
		}
		r -= weight.probability
	}

	return "login"
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	stop := time.After(testDuration)
	session, err := performLogin()
	if err != nil {
		log.Fatalf("failed to login, it gonna be killed. please check login login; error is %v", err)
	}
	for i := 0; i < numUsers; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case <-stop:
					return
				default:
					action := pickAction()
					performRequest(action, session)
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
				}
			}
		}()
	}

	wg.Wait()
	fmt.Println("stress test is completed")
}
