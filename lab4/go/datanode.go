package main

import 
(
	"fmt"
	"unsafe"
)

type DataNode struct{
	hehe int
	cmd string
	desc string
	next *DataNode
}

var cmd1 = [8] string {
	"help",
	"hello",
	"quit",
	"sum",
	"compare",
	"echo",
	"ls",
	"mkdir"}
var desc = [8] string {
	"Some tips about this system.",
	"Say hello to users,",
	"Quit this system.",
	"Get sum of two integer.",
	"Compare two integer.",
	"Output what you input.",
	"List files in this dir.",
	"Make a new dir."}

func FindCmd(head *LinkTable, cmd string) *DataNode{
	var p *DataNode = (*DataNode)(unsafe.Pointer(GetLinkTableHead(head)))
	if(head == nil ||cmd == " "){
		return nil
	}
	for p != nil {
		if p.cmd == cmd{
			return p
		}
		p = (*DataNode)(unsafe.Pointer(GetNextLinkTableNode(head, (*LinkTableNode) (unsafe.Pointer(p)))))
	}
	return nil
}

func ShowAllCmd (head *LinkTable) int {
	fmt.Println("Menu List:")
	p := (*DataNode)(unsafe.Pointer(GetLinkTableHead(head)))
	for p != nil {
		//fmt.Printf("%d",p.hehe)
		fmt.Println(p.cmd+" --- "+p.desc)

		p = (*DataNode)(unsafe.Pointer(GetNextLinkTableNode(head, (*LinkTableNode) (unsafe.Pointer(p)))))
			
	}	
	return 0
}


func InitMenuData(ppLinkTable **LinkTable) {
	var i int
	//*ppLinkTable = CreateLinkTable()
	*ppLinkTable = new(LinkTable)
	for i = 0; i <= 7; i++ {	
		var pNode = new(DataNode)
		//fmt.Println(pNode.cmd+" "+pNode.desc)	
		pNode.hehe = 1
		pNode.cmd = cmd1[i]
		pNode.desc = desc[i]
		AddLinkTableNode( *ppLinkTable, (*LinkTableNode) (unsafe.Pointer(pNode)))
		//fmt.Println(pNode.cmd+pNode.desc)
	}

}
