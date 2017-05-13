package main

import (
"fmt"
"strings"
)

type DataNode struct{
	cmd string
	desc string
	next *DataNode
}

func FindCmd(head *DataNode, cmd string) *DataNode{
	var p *DataNode = head
	if(head == nil ||cmd == " "){
		return nil
	}
	for p != nil {
		if (strings.EqualFold(p.cmd, cmd)){
			return p
		}
		p = p.next
	}
	return nil
}

func ShowAllCmd (head *DataNode) int {
	fmt.Println("Menu List:")
	p := head
	for p != nil {
		fmt.Println(p.cmd+" --- "+p.desc)
		p = p.next
	}	
	return 0
}

func handler(datanode *DataNode) {
	if datanode.cmd == "help" {
		help()
	}else if datanode.cmd == "hello" {
		hello()
	}else if datanode.cmd == "quit"{
		quit()
	}else if datanode.cmd == "sum"{
		sum()
	}else if datanode.cmd == "compare"{
		compare()
	}else if datanode.cmd == "echo"{
		echo()
	}else if datanode.cmd == "ls"{
		ls()
	}else if datanode.cmd == "mkdir"{
		mkdir()
	}
}
