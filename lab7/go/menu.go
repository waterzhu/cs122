package main

import 
(
	"fmt"
	"unsafe"
	"strings"
	"bufio"
	"os"
)

type DataNode struct{
	next *DataNode
	cmd string
	desc string
	handler func(args int, argv[] string) int
	
}
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
		fmt.Println(p.cmd+" --- "+p.desc)

		p = (*DataNode)(unsafe.Pointer(GetNextLinkTableNode(head, (*LinkTableNode) (unsafe.Pointer(p)))))
			
	}	
	return 0
}

func MenuConfig(cmd string, desc string, handler func(int, []string) int) int {
	var ppLinkTable = &head
	if(head == nil){
	 *ppLinkTable = new(LinkTable)
	}
		var pNode = new(DataNode)
		pNode.cmd = cmd
		pNode.desc = desc
		pNode.handler = handler
		AddLinkTableNode( *ppLinkTable, (*LinkTableNode) (unsafe.Pointer(pNode)))
		return 0
	}

func ExecuteMenu() int {
	for true{
		var argc = 0
		var argv []string
		var cmd string
		var pcmd []string
		fmt.Printf("Input a cmd $ ")
		buf := bufio.NewReader(os.Stdin)
		cmd, err := buf.ReadString('\n')
		if err != nil{
			fmt.Println(err)
		}
		pcmd = strings.Fields(cmd)
		//fmt.Println(len(pcmd))
		for argc < len(pcmd){
			//fmt.Println(pcmd[argc])
			argv = append(argv, pcmd[argc])
			//fmt.Println(argv[argc])
			argc = argc + 1			
		}
		p := (*DataNode)(unsafe.Pointer(SearchLinkTableNode(head, SearchCondition, argv[0])))
		if p == nil{
			fmt.Println("WRONG CMD!!!")
			continue
		}
		if p.handler != nil{
			p.handler(argc,argv)
		}
	}
		return 0
}