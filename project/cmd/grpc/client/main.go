package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"project/proto"
	"time"
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
	conn, err := grpc.NewClient("0.0.0.0:4567", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	switch c.Cmd {
	case "create":
		if err := c.Create(conn); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := c.Get(conn); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "patch":
		if err := c.Patch(conn); err != nil {
			return fmt.Errorf("patch account failed: %w", err)
		}

		return nil
	case "change":
		if err := c.Change(conn); err != nil {
			return fmt.Errorf("change account failed: %w", err)
		}

		return nil
	case "delete":
		if err := c.Delete(conn); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command %s", c.Cmd)
	}
}

func (c *Command) Create(conn *grpc.ClientConn) error {
	cl := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := cl.CreateAccount(ctx, &proto.CreateAccountRequest{Name: c.Name, Amount: int32(c.Amount)})
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Command) Get(conn *grpc.ClientConn) error {
	cl := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cl.GetAccount(ctx, &proto.GetAccountRequest{Name: c.Name})
	if err != nil {
		panic(err)
	}

	fmt.Printf("response account name: %s and amount: %d", res.GetName(), res.GetAmount())

	return nil
}

func (c *Command) Patch(conn *grpc.ClientConn) error {
	cl := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := cl.GetAccount(ctx, &proto.GetAccountRequest{Name: c.Name})
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Command) Change(conn *grpc.ClientConn) error {
	cl := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := cl.GetAccount(ctx, &proto.GetAccountRequest{Name: c.Name})
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Command) Delete(conn *grpc.ClientConn) error {
	cl := proto.NewBankClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := cl.GetAccount(ctx, &proto.GetAccountRequest{Name: c.Name})
	if err != nil {
		panic(err)
	}

	return nil
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
