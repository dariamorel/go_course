package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"project/accounts/dto"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	Amount  int
}

func (c *Command) Do() error {
	switch c.Cmd {
	case "create":
		if err := c.Create(); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := c.Get(); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "patch":
		if err := c.Patch(); err != nil {
			return fmt.Errorf("patch account failed: %w", err)
		}

		return nil
	case "change":
		if err := c.Change(); err != nil {
			return fmt.Errorf("change account failed: %w", err)
		}

		return nil
	case "delete":
		if err := c.Delete(); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command %s", c.Cmd)
	}
}

func (c *Command) Create() error {
	request := dto.CreateAccountRequest{
		Name:   c.Name,
		Amount: c.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", c.Host, c.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (c *Command) Get() error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", c.Host, c.Port, c.Name),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)

	return nil
}

func (c *Command) Patch() error {
	request := dto.PatchAccountRequest{
		Name:   c.Name,
		Amount: c.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/patch", c.Host, c.Port),
		"application/json",
		bytes.NewReader(data),
	)

	if err != nil {
		return fmt.Errorf("http patch failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (c *Command) Change() error {
	request := dto.ChangeAccountRequest{
		Name:    c.Name,
		NewName: c.NewName,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/change", c.Host, c.Port),
		"application/json",
		bytes.NewReader(data),
	)

	if err != nil {
		return fmt.Errorf("http change failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (c *Command) Delete() error {
	request := dto.DeleteAccountRequest{
		Name: c.Name,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/delete", c.Host, c.Port),
		"application/json",
		bytes.NewReader(data),
	)

	if err != nil {
		return fmt.Errorf("http delete failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	newNameVal := flag.String("new_name", "", "new name of account")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NewName: *newNameVal,
		Amount:  *amountVal,
	}

	if err := cmd.Do(); err != nil {
		panic(err)
	}
}
