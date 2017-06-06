package main

import 
(
	"fmt"
	"unsafe"
)

type DataNode struct{
	next *DataNode
	cmd string
	desc string
	
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

func SearchCondition(ppLinkTableNode *LinkTableNode, args interface{}) int{
	cmd := args
	pNode := (*DataNode)(unsafe.Pointer(ppLinkTableNode))
	if pNode.cmd == cmd{
		return 0
	}
	return (-1)
}
	


func FindCmd(head *LinkTable, cmd string) *DataNode{
	return (*DataNode)(unsafe.Pointer(SearchLinkTableNode(head, SearchCondition, cmd)))
}

func ShowAllCmd(head *LinkTable) int {
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
	*ppLinkTable = new(LinkTable)
	for i = 0; i <= 7; i++ {	
		var pNode = new(DataNode)
		pNode.cmd = cmd1[i]
		pNode.desc = desc[i]
		AddLinkTableNode( *ppLinkTable, (*LinkTableNode) (unsafe.Pointer(pNode)))
	}

}
