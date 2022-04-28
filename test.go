package main

import (
	"io"
	"net"
)

func TesthandshakeWithServer(test *ethrTest, conn net.Conn) (err error) {
	/*ethrMsg := createSynMsg(test.testID, test.clientParam)
	err = sendSessionMsg(conn, ethrMsg)
	if err != nil {
		ui.printDbg("Failed to send SYN message to Ethr server. Error: %v", err)
		return
	}
	ethrMsg = recvSessionMsg(conn)
	if ethrMsg.Type != EthrAck {
		ui.printDbg("Failed to receive ACK message from Ethr server. Error: %v", err)
		err = os.ErrInvalid
	}*/
	w, err := conn.Write([]byte{1, 2})
	if err != nil {
		return err
	}
	ui.printDbg("DEBUG:write:%v", w)
	buf := make([]byte, 50)
	r, err := io.ReadFull(conn, buf)
	if err != nil {
		return err
	}
	ui.printDbg("DEBUG:read:%v,%v", r, buf)
	return
}
