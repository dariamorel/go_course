package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"project/accounts/models"
	"project/proto"
	"sync"
)

type server struct {
	proto.BankServer
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (s *server) New() *server {
	return &server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

func (s *server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	if len(req.GetName()) == 0 {
		return nil, errors.New("empty name")
	}

	if req.GetAmount() < 0 {
		return nil, errors.New("negative amount")
	}

	s.guard.Lock()

	if _, ok := s.accounts[req.GetName()]; ok {
		s.guard.Unlock()

		return nil, errors.New("account already exists")
	}

	s.accounts[req.GetName()] = &models.Account{
		Name:   req.GetName(),
		Amount: int(req.GetAmount()),
	}

	s.guard.Unlock()

	response := &proto.CreateAccountResponse{Res: "StatusOK"}
	return response, nil
}

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	name := req.GetName()

	s.guard.RLock()

	account, ok := s.accounts[name]

	s.guard.RUnlock()

	if !ok {
		return nil, errors.New("account not found")
	}

	response := &proto.GetAccountResponse{
		Name:   account.Name,
		Amount: int32(account.Amount),
	}
	return response, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {

	if len(req.GetName()) == 0 {
		return nil, errors.New("empty name")
	}

	s.guard.Lock()

	_, ok := s.accounts[req.GetName()]

	if !ok {
		s.guard.Unlock()

		return nil, errors.New("account not found")
	}

	delete(s.accounts, req.GetName())

	s.guard.Unlock()

	response := &proto.DeleteAccountResponse{Res: "account deleted"}

	return response, nil
}

func (s *server) PatchAccount(ctx context.Context, req *proto.PatchAccountRequest) (*proto.PatchAccountResponse, error) {

	if len(req.GetName()) == 0 {
		return nil, errors.New("empty name")
	}

	if req.GetAmount() < 0 {
		return nil, errors.New("negative amount")
	}

	s.guard.Lock()

	account, ok := s.accounts[req.GetName()]

	if !ok {
		s.guard.Unlock()

		return nil, errors.New("account not found")
	}

	account.Amount = int(req.GetAmount())

	s.guard.Unlock()

	response := &proto.PatchAccountResponse{Res: "account amount changed"}

	return response, nil
}

func (s *server) ChangeAccount(ctx context.Context, req *proto.ChangeAccountRequest) (*proto.ChangeAccountResponse, error) {

	if len(req.GetName()) == 0 {
		return nil, errors.New("empty name")
	}
	if len(req.GetNewName()) == 0 {
		return nil, errors.New("empty new name")
	}

	s.guard.Lock()

	account, ok := s.accounts[req.GetName()]

	if !ok {
		s.guard.Unlock()

		return nil, errors.New("account not found")
	}

	if _, ok := s.accounts[req.GetNewName()]; ok {
		s.guard.Unlock()

		return nil, errors.New("account already exists")
	}

	amount := account.Amount
	delete(s.accounts, req.GetName())
	s.accounts[req.GetNewName()] = &models.Account{
		Name:   req.GetNewName(),
		Amount: amount,
	}

	s.guard.Unlock()

	response := &proto.ChangeAccountResponse{Res: "account name changed"}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	newServer := &server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
	proto.RegisterBankServer(s, newServer)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
