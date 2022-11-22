package handlers

import (
	"crucible/vega/configs"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
)

var httpClient = resty.New()

func HandleSemaphoreLogin(c *fiber.Ctx) error {

	t, err := getUserToken(
		configs.Config.Semaphore.URL,
		configs.Config.Semaphore.Username,
		configs.Config.Semaphore.Password,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}
	return c.SendString(fmt.Sprintf("%v", t))
}

func HandleSemaphorePing(c *fiber.Ctx) error {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%s/api/ping", configs.Config.Semaphore.URL))

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	return c.SendString(fmt.Sprintf("%s", resp.Body()))
}

func HandleSemaphoreGetProjects(c *fiber.Ctx) error {
	t, err := getUserToken(
		configs.Config.Semaphore.URL,
		configs.Config.Semaphore.Username,
		configs.Config.Semaphore.Password,
	)

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		Get(fmt.Sprintf("%s/api/projects", configs.Config.Semaphore.URL))

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	return c.SendString(fmt.Sprintf("%s", resp.Body()))
}

func HandleSemaphoreGetProject(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/api/project/%d", 3000, 1)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return c.SendString(fmt.Sprintf("%s", resBody))
}

func getUserToken(url, username, password string) (*Token, error) {
	cookies, err := fetchCookies(url, username, password)
	if err != nil {
		return nil, err
	}

	t, err := fetchTokens(url, cookies)
	if err != nil {
		return nil, err
	}

	if len(*t) == 0 {
		err = generateNewToken(url, cookies)
		if err != nil {
			return nil, err
		}
		t, err = fetchTokens(url, cookies)
		if err != nil {
			return nil, err
		}
	}

	for _, token := range *t {
		if !token.Expired {
			return &token, nil
		}
	}

	return nil, fmt.Errorf("no valid token found")
}

func fetchTokens(url string, c []*http.Cookie) (*[]Token, error) {
	r, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetCookies(c).
		Get(fmt.Sprintf("%s/api/user/tokens", url))

	if err != nil {
		return nil, err
	}

	var tokens []Token
	err = json.Unmarshal(r.Body(), &tokens)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

func generateNewToken(url string, c []*http.Cookie) error {
	_, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetCookies(c).
		Post(fmt.Sprintf("%s/api/user/tokens", url))

	if err != nil {
		return err
	}

	return nil
}

func fetchCookies(url, username, password string) ([]*http.Cookie, error) {
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(fmt.Sprintf(`{"auth": "%s", "password": "%s"}`, username, password)).
		Post(fmt.Sprintf("%s/api/auth/login", url))

	if err != nil {
		return nil, err
	}

	return resp.Cookies(), nil
}

type Token struct {
	Value     string `json:"id"`
	CreatedAt string `json:"created"`
	Expired   bool   `json:"expired"`
	UserId    int    `json:"user_id"`
}
