package graph

import (
	"sync"
	//
	"github.com/golangdaddy/tarantula/log"
)

const (
	SHA3_ID_LENGTH = 15
	MAX_CLASS_NAME = 20
)

type Database struct {

	// logging interface
	Log logging.Logger

	// database name
	Name string

	// database client
	Client Client

	// vertex uid index
	vertexIndex *VertexIndex

	// internal classes
	internalClasses *InternalClasses

	// custom classes
	Classes *Classes

	// class representing user
	UserClass *Class

	// callback channels
	Channels Channels

	MemGraph bool

	sync.RWMutex
}

func NewDatabase(logger logging.Logger, databaseName string, client Client, tools ...interface{}) *Database {

	db := &Database{
		Log: logger,
		Name: databaseName,
		Client: client,
		Channels: Channels{},
	}

	// make sure the client has the db pointer
	db.Client.SetDB(db)

	// todo: remove
	global = db

	// create database

	if !db.Client.CreateDatabase(db.Name) { panic("FAILED TO CREATE DATABASE: "+db.Name) }
	// create database object

	db.vertexIndex = (VertexIndex{}).New()

	db.internalClasses = db.newInternalClasses()

	db.Classes = db.newClasses()

	return db
}

