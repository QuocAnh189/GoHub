package socketio

import (
	"github.com/QuocAnh189/GoBin/logger"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

type Server struct {
	server *socketio.Server
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

var socketConnect = make(map[string]string)

// NewServer creates a new instance of Socket.IO server
func NewServer() (*Server, error) {
	server := socketio.NewServer(&engineio.Options{
		PingInterval: time.Second * 25,
		PingTimeout:  time.Second * 60,
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	// Handle connection event
	server.OnConnect("/", func(s socketio.Conn) error {
		url := s.URL()
		clientID := (&url).Query().Get("id")
		if clientID != "" {
			logger.Info("User connected with ID:", clientID)
			logger.Info("Client connected with ID:", s.ID())
			socketConnect[clientID] = s.ID()
		} else {
			logger.Info("Client connected without ID")
		}
		return nil
	})

	server.OnEvent("/", "follow", func(s socketio.Conn, data map[string]string) {
		followerID := data["follower_id"]
		followeeID := data["followee_id"]
		logger.Info("Follow user")

		server.BroadcastToRoom("/", socketConnect[followeeID], "notify_follow", map[string]string{
			"follower_id": followerID,
		})
	})

	server.OnEvent("/", "invitation", func(s socketio.Conn, inviteeIds []string) {
		logger.Info("Invitation")

		for _, inviteeId := range inviteeIds {
			logger.Info("Invitation User")
			server.BroadcastToRoom("/", socketConnect[inviteeId], "notify_invitation", map[string]string{
				"message": "You have a new invitation",
			})
		}
	})

	// Join Room
	server.OnEvent("/", "join_conversation", func(s socketio.Conn, conversationID string) {
		logger.Info("Client joined conversation:", conversationID)
		s.Join(conversationID)
	})

	// Handle private messages
	server.OnEvent("/", "send_message", func(s socketio.Conn, data map[string]string) {
		senderID := data["sender_id"]
		conversationID := data["conversation_id"]
		message := data["message"]

		logger.Info("Send message from userId: ", senderID, " to conversationId :", conversationID, " with message: ", message)

		// Send message to all clients in the room
		server.BroadcastToRoom("/", conversationID, "receive_message", map[string]string{
			"sender_id": senderID,
			"message":   message,
		})
	})

	//Logout
	server.OnEvent("/", "logout", func(s socketio.Conn) {
		log.Println("Client logged out:", s.ID())
		if err := s.Close(); err != nil {
			logger.Error(err)
		}
	})

	// Handle disconnection event
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		for key, value := range socketConnect {
			if value == s.ID() {
				delete(socketConnect, key)
				break
			}
		}
		log.Println("Disconnected:", s.ID(), "Reason:", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("Socket.IO listen error: %s\n", err)
		}
	}()
	return &Server{server: server}, nil
}

// Run the Socket.IO server on port 9000
func (s *Server) Run(port int) error {
	logger.Info("Socket.IO server is listening on PORT: ", port)
	return http.ListenAndServe(":9000", s.server)
}
