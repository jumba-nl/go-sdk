package model

import (
	"log"

	"gopkg.in/mgo.v2"
	"strings"
	"crypto/tls"
	"time"
	"net"
)

var conn *mgo.Session
var D *mgo.Database

var MongoHost string
var MongoDB string

var DialInfo *mgo.DialInfo

func MgoParseURI(url string) (string, *mgo.DialInfo, error) {
	isSSL := strings.Index(url, "ssl=true") > -1
	// Remove ssl option because it is unsupported by mgo ParseURL
	url = strings.Replace(url, "ssl=true", "", 1)

	// Remove other options that are unsupported by mgo ParseURL
	url = strings.Replace(url, "retryWrites=true", "", 1)

	dialInfo, err := mgo.ParseURL(url)

	if err != nil {
		return url, nil, err
	}

	if isSSL {
		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true

		dialInfo.Timeout = time.Second * 10
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	return url, dialInfo, err
}


func ConnectDB(host, db string) error {
	MongoHost = host
	MongoDB = db

	var err error

	MongoHost, DialInfo, err = MgoParseURI(host)
	if err != nil {
		return err
	}

	conn, err = mgo.DialWithInfo(DialInfo)
	if err != nil {
		return err
	}
	D = conn.DB(MongoDB)

	return nil
}

func WithSession(s *mgo.Session) {
	if len(s.LiveServers()) > 0 {
		MongoHost = s.LiveServers()[0]
	}
	MongoDB = s.DB("").Name
	D = s.DB("")
	conn = s
}

func C() *mgo.Session {
	err := conn.Ping()
	if err != nil {
		var err error
		conn, err = mgo.DialWithInfo(DialInfo)
		if err != nil {
			log.Println("Error connecting to DB", err)
			return conn
		}

		D = conn.DB(MongoDB)
	}

	return conn
}

func DB() *mgo.Database {
	err := conn.Ping()
	if err != nil {
		var err error
		conn, err = mgo.DialWithInfo(DialInfo)
		if err != nil {
			log.Println("Error connecting to DB", err)
			return D
		}

		D = conn.DB(MongoDB)
	}

	return D
}

func CloseDB() {
	conn.Close()
}
