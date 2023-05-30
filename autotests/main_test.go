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
	log.Println("Setting up testing database")
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
	if err != nil {
		log.Println(err)
	}
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

	resp, err := Post("http://localhost:8080/processes", cookie[0].Value, bytes.NewBuffer(body))
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestProcessList(t *testing.T) {
	resp, err := Post("http://localhost:8080/list/processes", cookie[0].Value, nil)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestProcessByName(t *testing.T) {
	url := "http://localhost:8080/processes/" + subProcess.Title
	resp, err := Get(url, cookie[0].Value)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestThreadCreate(t *testing.T) {
	body, _ := json.Marshal(subThread)
	resp, err := Post("http://localhost:8080/threads", cookie[0].Value, bytes.NewBuffer(body))
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestThreadList(t *testing.T) {
	resp, err := Post("http://localhost:8080/list/threads", cookie[0].Value, nil)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestThreadByName(t *testing.T) {
	url := "http://localhost:8080/processes/" + subThread.Process + "/" + subThread.Title
	resp, err := Get(url, cookie[0].Value)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestProcessesThreads(t *testing.T) {
	url := "http://localhost:8080/process/" + subProcess.Title + "/threads"

	resp, err := Get(url, cookie[0].Value)
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestSheetCreate(t *testing.T) {
	url := "http://localhost:8080/processes/" + subSheet.Process + "/sheet"
	body, _ := json.Marshal(subSheet)
	resp, err := Post(url, cookie[0].Value, bytes.NewBuffer(body))
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestSheetAddContent(t *testing.T) {
	url := "http://localhost:8080/processes/" + subSheet.Process + "/sheet/row"
	body, _ := json.Marshal(subSheetRow)
	resp, err := Post(url, cookie[0].Value, bytes.NewBuffer(body))
	require.Nil(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
