package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/sheet"
	"reflection_prototype/internal/core/thread"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var subUser = user.User{
	Login:   "test1",
	Name:    "test1",
	Surname: "test1",
	Email:   "test1@test.ru",
	Pwd:     "test1",
}

var subProcess, _ = process.New("p1")
var subThread, _ = thread.New("p1", "t1")
var subSheet = sheet.New("p1", "s1")
var subSheetRow = sheet.NewSheetRow("sh1", time.Now())

var cookie []*http.Cookie

func TestMain(m *testing.M) {
	clear := exec.Command("docker-compose", "down", "-v")
	err := clear.Run()
	if err != nil {
		log.Println(err)
	}
	setup := exec.Command("docker-compose", "up", "-d")
	err = setup.Start()
	if err != nil {
		log.Printf("setup err: %s", err)
	}
	log.Println("Setup testing database")
	time.Sleep(time.Second * 10)
	cmd := exec.Command("./../cmd/server/server")
	err = cmd.Start()
	if err != nil {
		log.Printf("cmd err: %s", err)
		cmd.Process.Kill()
		return
	}

	log.Printf("Server pid: %d", cmd.Process.Pid)
	time.Sleep(time.Second * 3)
	m.Run()
	out, _ := cmd.Output()
	log.Println(string(out))
	shutdown := exec.Command("docker-compose", "down", "-v")
	err = shutdown.Run()
	if err != nil {
		log.Println(err)
	}

	cmd.Process.Kill()
}

func TestRegister(t *testing.T) {

	body, _ := json.Marshal(subUser)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

}

func TestLogin(t *testing.T) {
	body, _ := json.Marshal(subUser)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)
	defer resp.Body.Close()
	cookie = resp.Cookies()
	assert.Equal(t, 200, resp.StatusCode)
}

func TestProcessCreate(t *testing.T) {
	body, _ := json.Marshal(subProcess)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/processes", bytes.NewBuffer(body))
	jwt := cookie[0].Value
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestProcessList(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/list/processes", nil)
	jwt := cookie[0].Value
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestProcessByName(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/processes/"+subProcess.Title, nil)
	jwt := cookie[0].Value
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestThreadCreate(t *testing.T) {
	body, _ := json.Marshal(subThread)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/threads", bytes.NewBuffer(body))
	jwt := cookie[0].Value
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestThreadList(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/list/threads", nil)
	jwt := cookie[0].Value
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestThreadByName(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/processes/"+subThread.Process+"/"+subThread.Title, nil)
	jwt := cookie[0].Value
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestProcessesThreads(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/process/"+subProcess.Title+"/threads", nil)
	jwt := cookie[0].Value
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestSheetCreate(t *testing.T) {
	body, _ := json.Marshal(subSheet)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/processes/"+subSheet.Process+"/sheet", bytes.NewBuffer(body))
	jwt := cookie[0].Value
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestSheetAddContent(t *testing.T) {
	body, _ := json.Marshal(subSheetRow)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/processes/"+subSheet.Process+"/sheet/row", bytes.NewBuffer(body))
	jwt := cookie[0].Value
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)
	require.Nil(t, err)

	c := http.Client{}
	resp, err := c.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
