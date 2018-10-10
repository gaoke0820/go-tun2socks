package tun2socks

import (
	"net"
)

// ConnectionHandler handles connections comming from TUN.
type ConnectionHandler interface {
	// Connect connects the proxy server.
	Connect(conn Connection, target net.Addr) error

	// DidReceive will be called when data arrives from TUN.
	DidReceive(conn Connection, data []byte) error

	// DidSend will be called when sent data has been acknowledged by local clients.
	DidSend(conn Connection, len uint16)

	// DidClose will be called when the connection has been closed.
	DidClose(conn Connection)

	// DidAbort will be called when the connection has been aborted.
	DidAbort(conn Connection)

	// DidReset will be called when the connection has been reseted.
	DidReset(conn Connection)

	// LocalDidClose will be called when local client has close the connection.
	LocalDidClose(conn Connection)
}
