package handlers

import (
	"crucible/vega/configs"
	"crucible/vega/models"
	"crucible/vega/queries"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
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
		log.Errorln(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	return c.SendString(fmt.Sprintf("%s", resp.Body()))
}

func HandleSemaphoreGetProjects(c *fiber.Ctx) error {
	t, err := getUserToken()

	if err != nil {
		log.Errorln(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		Get(fmt.Sprintf("%s/api/projects", configs.Config.Semaphore.URL))

	if err != nil {
		log.Errorln(err)
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

func getProjectByName(name string) (models.Project, error) {
	var project models.Project

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

	var projects []models.Project

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

func getRepositoryByName(projectId int, name string) (models.Repository, error) {
	var repository models.Repository

	t, err := getUserToken()

	if err != nil {
		log.Println(err)
		return repository, err
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		Get(fmt.Sprintf("%s/api/project/%d/repositories", configs.Config.Semaphore.URL, projectId))

	if err != nil {
		log.Println(err)
		return repository, err
	}

	var repositories []models.Repository

	err = json.Unmarshal(resp.Body(), &repositories)
	if err != nil {
		log.Println(err)
		return repository, err
	}

	for _, r := range repositories {
		if r.Name == name {
			return r, nil
		}
	}

	return repository, fmt.Errorf("repository not found")
}

func HandleSemaphoreCreateProject(c *fiber.Ctx) error {
	var project models.Project

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

	var createdProject models.Project

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

	var repository models.Repository

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
		log.Errorln(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	r, err := getRepositoryByName(pIdIntVal, repository.Name)
	if err == nil {
		return c.SendString(fmt.Sprintf(`{"id": "%d"}`, r.Id))
	}

	sshKey, err := getSSHKeyByName(pIdIntVal, "none")

	if err != nil {
		log.Errorln(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	log.Infoln(sshKey)

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		SetBody(
			fmt.Sprintf(
				`{
						"project_id": %d, 
						"name": "%s", 
						"git_url": "%s", 
						"git_branch": "%s", 
						"ssh_key_id": %d
						}`,
				pIdIntVal,            // project id
				repository.Name,      // repository name
				repository.GitUrl,    // repository url "github.com/XXX/test-repo.git"
				repository.GitBranch, // repository branch
				sshKey.Id,            // ssh key id
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

func getSSHKeyByName(projectId int, name string) (models.SshKey, error) {
	var sshKey models.SshKey

	t, err := getUserToken()

	if err != nil {
		return sshKey, errors.New(fmt.Sprintf("getSSHKeyByName() An error occurred while getting token: %s", err.Error()))
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		Get(fmt.Sprintf("%s/api/project/%d/keys", configs.Config.Semaphore.URL, projectId))

	if err != nil {
		return sshKey, errors.New(fmt.Sprintf("getSSHKeyByName() An error occurred while calling semaphore API: %s - %s", err.Error(), resp.Body()))
	}

	var sshKeys []models.SshKey

	err = json.Unmarshal(resp.Body(), &sshKeys)
	if err != nil {
		return sshKey, errors.New(fmt.Sprintf("getSSHKeyByName() An error occurred while unmarshalling response: %s", err.Error()))
	}

	for _, k := range sshKeys {
		if k.Name == name {
			return k, nil
		}
	}

	return sshKey, errors.New(fmt.Sprintf("getSSHKeyByName() SSH key with name '%s' not found", name))
}

func getInventoryByName(projectId int, name string) (models.Inventory, error) {
	var inventory models.Inventory

	t, err := getUserToken()

	if err != nil {
		return inventory, errors.New(fmt.Sprintf("getInventoryByName() An error occurred while getting token: %s", err.Error()))
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		Get(fmt.Sprintf("%s/api/project/%d/inventory", configs.Config.Semaphore.URL, projectId))

	if err != nil {
		return inventory, errors.New(fmt.Sprintf("getInventoryByName() An error occurred while calling semaphore API: %s - %s", err.Error(), resp.Body()))
	}

	var inventories []models.Inventory

	err = json.Unmarshal(resp.Body(), &inventories)
	if err != nil {
		return inventory, errors.New(fmt.Sprintf("getInventoryByName() An error occurred while unmarshalling response: %s", err.Error()))
	}

	for _, i := range inventories {
		if i.Name == name {
			return i, nil
		}
	}

	return inventory, errors.New(fmt.Sprintf("getInventoryByName() Inventory with name '%s' not found", name))
}

func HandleSemaphoreCreateSSHKey(c *fiber.Ctx) error {
	var sshKey models.SshKey

	err := c.BodyParser(&sshKey)
	if err != nil {
		log.Errorln(err)
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

func HandleSemaphoreAddInventory(c *fiber.Ctx) error {
	pId := c.Params("id")
	if pId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing project id")
	}

	pIdIntVal, err := strconv.Atoi(pId)

	var inventory models.Inventory

	err = c.BodyParser(&inventory)
	if err != nil {
		log.Errorln(err)
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	log.Infoln(inventory)

	if inventory.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing inventory name")
	}

	if inventory.Value == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing inventory value")
	}

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	existingInv, err := getInventoryByName(pIdIntVal, inventory.Name)

	if err == nil {
		log.Infof("Inventory with name '%s' already exists", existingInv.Name)
		return c.SendString(fmt.Sprintf(`{"id":%d}`, existingInv.Id))
	}

	sshKey, err := getSSHKeyByName(pIdIntVal, "none")

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	inventory.SshKeyId = sshKey.Id

	inventory.Type = "static"

	err = queries.Database.AddInventory(c.Context(), pIdIntVal, inventory)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	inv, err := getInventoryByName(pIdIntVal, inventory.Name)

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	return c.SendString(fmt.Sprintf(`{"id":%d}`, inv.Id))
}

func HandleSemaphoreCreateEnvironment(c *fiber.Ctx) error {
	pId := c.Params("id")
	if pId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing project id")
	}

	pIdIntVal, err := strconv.Atoi(pId)

	var environment models.Environment

	err = c.BodyParser(&environment)
	if err != nil {
		log.Errorln(err)
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if environment.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing environment name")
	}

	if environment.ProjectId == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "missing inventory id")
	}

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	existingEnv, err := getEnvironmentByName(pIdIntVal, environment.Name)

	if err == nil {
		return c.SendString(fmt.Sprintf(`{"id": %d}`, existingEnv.Id))
	}

	t, err := getUserToken()

	if err != nil {
		log.Errorln(err)
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
							"name": "%s", 
							"project_id": %d,
							"json": "{}"
						}`,
				environment.Name,
				pIdIntVal,
			),
		).
		Post(fmt.Sprintf("%s/api/project/%d/environment", configs.Config.Semaphore.URL, pIdIntVal))

	if err != nil {
		log.Errorln(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	nEnv, err := getEnvironmentByName(pIdIntVal, environment.Name)

	if err != nil {
		log.Errorln(err, resp.Body())
	}

	return c.SendString(fmt.Sprintf(`{"id": %d}`, nEnv.Id))
}

func getEnvironmentByName(pId int, name string) (models.Environment, error) {
	var environment models.Environment

	t, err := getUserToken()

	if err != nil {
		return environment, errors.New(fmt.Sprintf("getEnvironmentByName() An error occurred while getting token: %s", err.Error()))
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		Get(fmt.Sprintf("%s/api/project/%d/environment", configs.Config.Semaphore.URL, pId))

	if err != nil {
		log.Errorln(err)
		return environment, err
	}

	var environments []models.Environment

	err = json.Unmarshal(resp.Body(), &environments)

	if err != nil {
		log.Errorln(err)
		return environment, err
	}

	for _, env := range environments {
		if env.Name == name {
			return env, nil
		}
	}

	return environment, errors.New(fmt.Sprintf("Environment with name '%s' not found", name))
}

func HandleSemaphoreCreateTemplate(c *fiber.Ctx) error {
	pId := c.Params("id")
	if pId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing project id")
	}

	pIdIntVal, err := strconv.Atoi(pId)

	var template models.Template

	err = c.BodyParser(&template)
	if err != nil {
		log.Errorln(err)
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if template.Alias == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing template alias")
	}

	if template.ProjectId == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "missing project id")
	}

	if template.RepositoryId == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "missing repository id")
	}

	if template.InventoryId == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "missing inventory id")
	}

	if template.EnvironmentId == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "missing environment id")
	}

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	existingTemplate, err := getTemplateByName(pIdIntVal, template.Alias)

	if err == nil {
		return c.SendString(fmt.Sprintf("Template with name '%s' already exists", existingTemplate.Alias))
	}

	t, err := getUserToken()

	if err != nil {
		log.Errorln(err)
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
						  "inventory_id": %d,
						  "repository_id": %d,
						  "environment_id": %d,
						  "alias": "%s",
						  "playbook": "%s",
						  "description": "%s"
						}`,
				pIdIntVal,
				template.InventoryId,
				template.RepositoryId,
				template.EnvironmentId,
				template.Alias,
				template.Playbook,
				template.Description,
			),
		).
		Post(fmt.Sprintf("%s/api/project/%d/template", configs.Config.Semaphore.URL, pIdIntVal))

	if err != nil {
		log.Errorln(err)
		return fiber.NewError(fiber.StatusInternalServerError, "An error occurred")
	}

	return c.SendString(fmt.Sprintf("%s", resp.Body()))
}

func getTemplateByName(pId int, name string) (models.Template, error) {
	var template models.Template

	t, err := getUserToken()

	if err != nil {
		return template, errors.New(fmt.Sprintf("getTemplateByName() An error occurred while getting token: %s", err.Error()))
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", t.Value)).
		Get(fmt.Sprintf("%s/api/project/%d/template", configs.Config.Semaphore.URL, pId))

	if err != nil {
		log.Errorln(err)
		return template, err
	}

	var templates []models.Template

	err = json.Unmarshal(resp.Body(), &templates)

	if err != nil {
		log.Errorln(err)
		return template, err
	}

	for _, t := range templates {
		if t.Alias == name {
			return t, nil
		}
	}

	return template, errors.New(fmt.Sprintf("Template with name '%s' not found", name))
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
