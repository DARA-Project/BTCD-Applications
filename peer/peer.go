/*
Implements the solution to assignment 2 for UBC CS 416 2016 W2.

Usage:
$ go run peer.go [numPeers] [peerID] [peersFile] [server ip:port]

Example:
$ go run peer.go [numPeers] [peerID] [peersFile] [server ip:port]

*/

package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	// TODO
	"log"
	"strconv"
	"bufio"
	"bytes"
	"encoding/gob"
)

// Resources
var res []Resource

// Invocations remaining
var numInvocationsLeft int

// Resource server type.
type RServer int

// Request that peer sends in call to RServer.InitSession
// Return value is int
type Init struct {
	NumPeers int
	IPaddr   string // Must be set to ""
}

// Request that peer sends in call to RServer.GetResource
type ResourceRequest struct {
	SessionID int
	IPaddr    string // Must be set to ""
}

// Response that the server returns from RServer.GetResource
type Resource struct {
	Resource     string
	PeerID       int
	NumRemaining int
}

// Check errors from function calls
func check(err error) {
	if err != nil {
		log.Fatal("error: ", err)
	}
}

// Main workhorse method. 
func main() {
	args := os.Args[1:]

	// Missing command line args.
	if len(args) != 4 {
		fmt.Println("Usage: go run peer.go [numPeers] [peerID] [peersFile] [server ip:port]")
		return
	}

	// TODO

	numPeers, errInt := strconv.Atoi(args[0])
	check(errInt)

	peerID, errPID := strconv.Atoi(args[1])
	check(errPID)

	fp, errFile := os.Open(args[2])
	check(errFile)

	peersMap := make(map[int]string)
	scan := bufio.NewScanner(fp)
	i := 1

	for scan.Scan() {
		peersMap[i] = scan.Text()
		i++
	}

	check(scan.Err())

	fp.Close()

	init := Init{NumPeers: numPeers, IPaddr: ""}

	conn, errDial := rpc.Dial("tcp", args[3])
	defer conn.Close()
	check(errDial)
	var sessionId int

	// Initialization session
	if peerID == 1 {
		errCallInit := conn.Call("RServer.InitSession", init, &sessionId)
		check(errCallInit)
	}

	connections := make(map[int]*net.TCPConn)
	tcpAddrHost, errTCP := net.ResolveTCPAddr("tcp", peersMap[peerID])
	check(errTCP)
	l, errConn := net.ListenTCP("tcp", tcpAddrHost)
	check(errConn)

	for j := 1; j < peerID; j++ {
		for {
			lCon, errCon := l.AcceptTCP()
			if errCon != nil {
				continue;
			}

			defer lCon.Close()

			receivePeer := make(chan int, 1)
			ReceiveSessionID(lCon, receivePeer)
			recvd := <- receivePeer
			connections[recvd] = lCon
			break;
		}
	}

	// peerID works but not peerID + 1
	for i := peerID + 1; i <= numPeers; i++ { 
		for {
			tcpAddr, err1 := net.ResolveTCPAddr("tcp", peersMap[i])
			if err1 != nil {
				continue;
			}
			connect, err2 := net.DialTCP("tcp", nil, tcpAddr)
			if err2 != nil {
				continue;
			}
			connections[i] = connect
			defer connect.Close()
			
			writeAddr := make(chan int, 1)
			writeAddr <- peerID
			SendSessionID(connect, writeAddr)
			break;
		}
	}

	if peerID == 1 {
		for i := peerID + 1; i <= numPeers; i++ {
			writeMsg := make(chan int, 1)
			writeMsg <- sessionId
			SendSessionID(connections[i], writeMsg)
		}
	}

	if peerID == 2 {
		readMsg := make(chan int, 1)
		ReceiveSessionID(connections[1], readMsg)
		val := <- readMsg
		sessionId = val
	}

	// Arbitrary initial choice for number of resources	
	numResourceRemaining := 1

	if peerID == 1 {
		for ;numResourceRemaining > 0; {
			var response Resource
			initResource := ResourceRequest{SessionID: sessionId, IPaddr: ""}
			errRes := conn.Call("RServer.GetResource", initResource, &response)
			check(errRes)
			numResourceRemaining = response.NumRemaining
			
			if response.PeerID == 1 {
				res = append(res, response)
			}

			for i := peerID + 1; i <= numPeers; i++ {
				writeMsg := make(chan Resource, 1)
				writeMsg <- response
				SendResource(connections[i], writeMsg)
			}

			if numResourceRemaining == 0 {
				break;
			}

			newChan := make(chan Resource, 1)
			ReceiveResource(connections[2], newChan)
			resourceReceived := <- newChan

			if resourceReceived.PeerID == peerID {
		 		res = append(res, resourceReceived)
		 	}

			numResourceRemaining = resourceReceived.NumRemaining

			if numResourceRemaining == 0 {
				break;
			}
		}
	} else if peerID == 2 {
		 for ;numResourceRemaining > 0; {
		 	receiveResource := make(chan Resource, 1)
		 	ReceiveResource(connections[1], receiveResource)
		 	recvd := <- receiveResource
		 	if recvd.PeerID == peerID {
		 		res = append(res, recvd)
		 	}

		 	if recvd.NumRemaining == 0 {
		 		break;
		 	}

		 	var response Resource
		 	initResource := ResourceRequest{SessionID: sessionId, IPaddr: ""}
			errRes := conn.Call("RServer.GetResource", initResource, &response)
			check(errRes)
			
			if response.PeerID == 2 {
				res = append(res, response)
			} 

			for i := 1; i <= numPeers; i++ {
				if i != peerID {
					writeMsg := make(chan Resource, 1)
					writeMsg <- response
					SendResource(connections[i], writeMsg)
				}
			}

		 	numResourceRemaining = response.NumRemaining

		 	if numResourceRemaining == 0 {
		 		break;
		 	}
		}
	} else {
		for ;numResourceRemaining > 0; {
			receiveResource := make(chan Resource, 1)
		 	ReceiveResource(connections[1], receiveResource)
		 	recvd := <- receiveResource
		 	if recvd.PeerID == peerID {
		 		res = append(res, recvd)
		 	}

		 	if recvd.NumRemaining == 0 {
		 		break;
		 	}

		 	receiveResource2 := make(chan Resource, 1)
		 	ReceiveResource(connections[2], receiveResource2)
		 	recvd2 := <- receiveResource2
		 	if recvd2.PeerID == peerID {
		 		res = append(res, recvd2)
		 	}

		 	if recvd2.NumRemaining == 0 {
		 		break;
		 	}
		 }
	}

	for _, element := range res {
		fmt.Printf("Resource: %s\n", element.Resource)
	}
}



	

	/* Add broadcast to communicate with peers based on session id
		func broadcast() {
			for _, ipport := range peers {
				conn.send(seshid)
			}
		}
	*/

func SendSessionID(conn *net.TCPConn, writeMsg chan int) {
	for true {
		var network bytes.Buffer
		enc := gob.NewEncoder(&network)
		message := <- writeMsg
		err := enc.Encode(message)
		check (err)
		conn.Write(network.Bytes())
		break;
	}
}

func ReceiveSessionID(conn *net.TCPConn, readMsg chan int) {
	for true {
		var message int
		var network bytes.Buffer
		buf := make([]byte, 1024)
		dec := gob.NewDecoder(&network)
		n, err := conn.Read(buf)
		if (err != nil) {
			continue;
		}
		network.Write(buf[0:n])
		err = dec.Decode(&message)
		if (err != nil) {
			continue;
		}
		readMsg <- message
		break;
	}
}

func SendResource(conn *net.TCPConn, writeMsg chan Resource) {
	for true {
		var network bytes.Buffer
		enc := gob.NewEncoder(&network)
		message := <- writeMsg
		err := enc.Encode(message)
		check (err)
		conn.Write(network.Bytes())
		break;
	}
}

func ReceiveResource(conn *net.TCPConn, readMsg chan Resource) {
	for true {
		var message Resource
		var network bytes.Buffer
		buf := make([]byte, 1024)
		dec := gob.NewDecoder(&network)
		n, err := conn.Read(buf)
		if (err != nil) {
			continue;
		}
		network.Write(buf[0:n])
		err = dec.Decode(&message)
		if (err != nil) {
			continue;
		}
		readMsg <- message
		break;
	}
}