package model 

type AVLTree struct {
	Root *Node
}

type Node struct {
	Key int 
	Left *Node 
	Right *Node 
}
