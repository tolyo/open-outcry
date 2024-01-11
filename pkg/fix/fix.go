package fix

import (
	"fmt"

	"github.com/quickfixgo/fix40/neworderlist"
	"github.com/quickfixgo/fix42/marketdatarequest"
	"github.com/quickfixgo/fix42/newordersingle"
	"github.com/quickfixgo/fix42/ordercancelrequest"

	"github.com/quickfixgo/quickfix"
	log "github.com/sirupsen/logrus"
)

// FixServer implements the main quickfix interface
type Server struct {
	*quickfix.MessageRouter
}

func New() *Server {
	server := &Server{}
	server.AddRoute(marketdatarequest.Route(server.NewMarketDataReq))
	server.AddRoute(newordersingle.Route(server.NewOrder))
	server.AddRoute(ordercancelrequest.Route(server.CancelOrder))
	server.AddRoute(neworderlist.Route(server.NewOrderList))
	return server
}

func (s *Server) OnCreate(sessionID quickfix.SessionID) {
	log.Println("Session created:", sessionID)
}

func (s *Server) OnLogon(sessionID quickfix.SessionID) {
	fmt.Println("Session logged on:", sessionID)
}

func (s *Server) OnLogout(sessionID quickfix.SessionID) {
	fmt.Println("Session logged out:", sessionID)
}

func (s *Server) ToAdmin(message quickfix.Message, sessionID quickfix.SessionID) {
	fmt.Println("Sending admin message to", sessionID, ":", message)
}

func (s *Server) ToApp(message quickfix.Message, sessionID quickfix.SessionID) error {
	fmt.Println("Sending app message to", sessionID, ":", message)
	return nil
}

func (s *Server) FromAdmin(message quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	fmt.Println("Receiving admin message from", sessionID, ":", message)
	return quickfix.InvalidMessageType()
}

func (s *Server) FromApp(message quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	fmt.Println("Receiving app message from", sessionID, ":", message)
	return quickfix.InvalidMessageType()
}

func (s *Server) NewOrder(msg newordersingle.NewOrderSingle, id quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

func (s *Server) CancelOrder(msg ordercancelrequest.OrderCancelRequest, id quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

func (s *Server) NewOrderList(msg neworderlist.NewOrderList, id quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

func (s *Server) NewMarketDataReq(msg marketdatarequest.MarketDataRequest, id quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}
