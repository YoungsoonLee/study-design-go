package main

type collection interface {
	createIterator() iterator
}

// test 1
