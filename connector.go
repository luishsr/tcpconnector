package connector

import (
	"log"
	"net"
	"strconv"
)

// represent each connected device
type Party struct {
	address string
	name    string
	role    string
}

type Message struct {
	mtype int
	value string
}

// a list of connected parties
var parties []Party

/*
get connections and register parties
@address caller IP address:port
@name caller identification
@role Server or Client
*/
func Register(address string, name string, role string) string {

	//creates a Party object
	newParty := Party{address, name, role}

	//adds the party to the list of parties
	parties = append(parties, newParty)

	return name + " - Successfuly registered as a " + role
}

/*
routes messages to a party
@from sender address
@to sender address
@message a Message object
*/
func SendMessageTo(message Message, from string, to string) {

	// Open a TCP Session to Server
	c, err := net.Dial("tcp", to)
	if err != nil {
		log.Fatalf("Unable to open TCP Connection: %s", err)
	}
	defer c.Close()

	// Write data
	log.Printf("TCP Session Open")
	b := []byte(strconv.Itoa(message.mtype) + "|" + message.value)
	_, err = c.Write(b)

	if err != nil {
		log.Fatalf("Error from TCP Session: %s", err)
	}
}
