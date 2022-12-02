package handlers

import (
	"crucible/vega/configs"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var httpClient = resty.New()

func HandleSemaphoreLogin(c *fiber.Ctx) error {

	t, err := getUserToken()

	if err != nil {
		return fiber.NewError(fiber.StatusForbidden,
			fmt.Sprintf("failed to login to semaphore: %v", err.Error()),
		)
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
	t, err := getUserToken()

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
	requestURL := fmt.Sprintf("%s/api/project/%d",
		configs.Config.Semaphore.URL,
		1,
	)
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

func getProjectByName(name string) (Project, error) {
	var project Project

	t, err := getUserToken()

	if err != nil {
		log.Println(err)
		return project, err
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		Get(fmt.Sprintf("%s/api/projects", configs.Config.Semaphore.URL))

	if err != nil {
		log.Println(err)
		return project, err
	}

	var projects []Project

	err = json.Unmarshal(resp.Body(), &projects)
	if err != nil {
		log.Println(err)
		return project, err
	}

	for _, p := range projects {
		if p.Name == name {
			return p, nil
		}
	}

	return project, fmt.Errorf("project not found")
}

func HandleSemaphoreCreateProject(c *fiber.Ctx) error {
	var project Project

	err := c.BodyParser(&project)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if project.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing project name")
	}

	t, err := getUserToken()

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	p, err := getProjectByName(project.Name)
	if err == nil {
		return c.SendString(fmt.Sprintf(`{"id": "%d"}`, p.Id))
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		SetBody(
			fmt.Sprintf(
				`{"name": "%s", "alert": %t}`,
				project.Name,
				project.Alert,
			),
		).
		Post(fmt.Sprintf("%s/api/projects", configs.Config.Semaphore.URL))

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	var createdProject Project

	err = json.Unmarshal(resp.Body(), &createdProject)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	err = createNoneSshKey(createdProject.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred while creating 'None' ssh key")
	}

	return c.SendString(fmt.Sprintf(`{"id": "%d"}`, createdProject.Id))
}

func HandleSemaphoreCreateRepository(c *fiber.Ctx) error {
	pId := c.Params("id")
	if pId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing project id")
	}

	pIdIntVal, err := strconv.Atoi(pId)

	var repository Repository

	err = c.BodyParser(&repository)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if repository.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing repository name")
	}

	if repository.GitUrl == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing repository url")
	}

	if repository.GitBranch == "" {
		repository.GitBranch = "main"
	}

	t, err := getUserToken()

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		SetBody(
			fmt.Sprintf(
				`{
						"project_id": %d, 
						"name": %s, 
						"git_url": %s,, 
						"git_branch": %s, 
						"ssh_key_id": %d
						}`,
				pIdIntVal,            // project id
				repository.Name,      // repository name
				repository.GitUrl,    // repository url "github.com/XXX/test-repo.git"
				repository.GitBranch, // repository branch
				0,                    // ssh key id
			),
		).
		Post(fmt.Sprintf("%s/api/project/%d/repositories", configs.Config.Semaphore.URL, pIdIntVal))

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	return c.SendString(fmt.Sprintf("%d - %s", resp.StatusCode(), resp.Body()))
}

func createNoneSshKey(projectId int) error {
	t, err := getUserToken()

	if err != nil {
		return errors.New(fmt.Sprintf("createNoneSshKey() An error occurred while getting token: %s", err.Error()))
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		SetBody(
			fmt.Sprintf(
				`{
						"name": "none",
						"type": "none",
						"project_id": %d
						}`,
				projectId,
			),
		).
		Post(fmt.Sprintf("%s/api/project/%d/keys", configs.Config.Semaphore.URL, projectId))

	if err != nil {
		return errors.New(fmt.Sprintf("createNoneSshKey() An error occurred while calling semaphore API: %s - %s", err.Error(), resp.Body()))
	}

	return nil
}

func HandleSemaphoreCreateSSHKey(c *fiber.Ctx) error {
	var sshKey SshKey

	err := c.BodyParser(&sshKey)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if sshKey.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing ssh key name")
	}

	if sshKey.Value == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing ssh key value")
	}

	t, err := getUserToken()

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		SetBody(
			fmt.Sprintf(
				`{"name": "%s", "value": "%s"}`,
				sshKey.Name,
				sshKey.Value,
			),
		).
		Post(fmt.Sprintf("%s/api/ssh_keys", configs.Config.Semaphore.URL))

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	return c.SendString(fmt.Sprintf("%s", resp.Body()))
}

func HandleSemaphoreRunTaskTemplate(c *fiber.Ctx) error {
	t, err := getUserToken()

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		SetBody(
			fmt.Sprintf(
				`{"project_id": %d, "template_id": %d}`,
				1,
				1,
			),
		).
		Post(fmt.Sprintf("%s/api/project/1/tasks", configs.Config.Semaphore.URL))

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	return c.SendString(fmt.Sprintf("%s", resp.Body()))
}

func getUserToken() (*Token, error) {
	cookies, err := fetchCookies()
	if err != nil {
		return nil, err
	}

	t, err := fetchTokens(cookies)
	if err != nil {
		return nil, err
	}

	if len(*t) == 0 {
		err = generateNewToken(cookies)
		if err != nil {
			return nil, err
		}
		t, err = fetchTokens(cookies)
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

func fetchTokens(c []*http.Cookie) (*[]Token, error) {
	r, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetCookies(c).
		Get(fmt.Sprintf("%s/api/user/tokens", configs.Config.Semaphore.URL))

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

func generateNewToken(c []*http.Cookie) error {
	_, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetCookies(c).
		Post(fmt.Sprintf("%s/api/user/tokens", configs.Config.Semaphore.URL))

	if err != nil {
		return err
	}

	return nil
}

func fetchCookies() ([]*http.Cookie, error) {
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(
			fmt.Sprintf(
				`{"auth": "%s", "password": "%s"}`,
				configs.Config.Semaphore.Username,
				configs.Config.Semaphore.Password,
			),
		).
		Post(fmt.Sprintf("%s/api/auth/login", configs.Config.Semaphore.URL))

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

type Project struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Alert       bool   `json:"alert,omitempty"`
}

type Repository struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	ProjectId   int    `json:"project_id"`
	GitUrl      string `json:"git_url"`
	GitBranch   string `json:"git_branch"`
	SshKeyId    int    `json:"ssh_key_id"`
	LastUpdated string `json:"last_updated"`
}

type SshKey struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Type      string `json:"type"`
	ProjectId int    `json:"project_id"`
}
