package main

import (
	"fmt"

	"github.com/quickfixgo/quickfix"
	log "github.com/sirupsen/logrus"
)

// FixServer implements the main quickfix interface
type FixServer struct{}

func (s *FixServer) OnCreate(sessionID quickfix.SessionID) {
	log.Println("Session created:", sessionID)
}

func (s *FixServer) OnLogon(sessionID quickfix.SessionID) {
	fmt.Println("Session logged on:", sessionID)
}

func (s *FixServer) OnLogout(sessionID quickfix.SessionID) {
	fmt.Println("Session logged out:", sessionID)
}

func (s *FixServer) ToAdmin(message quickfix.Message, sessionID quickfix.SessionID) {
	fmt.Println("Sending admin message to", sessionID, ":", message)
}

func (s *FixServer) ToApp(message quickfix.Message, sessionID quickfix.SessionID) error {
	fmt.Println("Sending app message to", sessionID, ":", message)
	return nil
}

func (s *FixServer) FromAdmin(message quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	fmt.Println("Receiving admin message from", sessionID, ":", message)
	return quickfix.InvalidMessageType()
}

func (s *FixServer) FromApp(message quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	fmt.Println("Receiving app message from", sessionID, ":", message)
	return quickfix.InvalidMessageType()
}

func New() *FixServer {
	server := &FixServer{}
	return server
}
