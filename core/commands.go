package core

const QUEUE_LEN = 1024
const readCounter = 0
const writeCounter = 0

var commands = make([]Signal, QUEUE_LEN)

func WriteCommands() {}

func ReadCommands() {}

func ProcessCommands() {}

/*
	Since I am communicating cross language, I may need to do
	length prefixed json because all implementations may not be
	the same.
*/


/* 
	Goals:
	Write Command through cmd or any other method into socket file
	Read complete commands always.
	
*/