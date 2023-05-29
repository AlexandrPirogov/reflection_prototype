package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"reflection_prototype/internal/core/auth/user"
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
	assert.Equal(t, 200, resp.StatusCode)
}
