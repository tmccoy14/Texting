package sms

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v1"
)

type TestConfig struct {
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Phone    string `yaml:"phone"`
	Carrier  string `yaml:"carrier"`
}

func testSetup(t *testing.T) (SMSClient, TestConfig) {
	buffer, err := ioutil.ReadFile("sms.yml")
	m := make(map[string]TestConfig)
	if err := yaml.Unmarshal(buffer, &m); err != nil {
		t.Error(err)
	}

	config := m["test"]
	client, err := createClient(config.Address, config.Port, config.User, config.Password)
	if err != nil {
		t.Error(err)
	}

	return client, config
}

func TestCreateClient(t *testing.T) {
	testSetup(t)
}

func TestDeliver(t *testing.T) {
	client, config := testSetup(t)

	var message string

	if err := client.Deliver(config.Phone, config.Carrier, message); err != nil {
		t.Error(err)
	}
}
