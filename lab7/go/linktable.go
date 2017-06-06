package main

type LinkTableNode struct {
	pNext *LinkTableNode
}

type LinkTable 	struct {
	pHead *LinkTableNode
	pTail *LinkTableNode
	SumOfNode int
}

func AddLinkTableNode (pLinkTable *LinkTable, pNode *LinkTableNode)  int {
	if pLinkTable == nil || pNode == nil{
		return (-1)
	}
	if pLinkTable.pHead == nil {
		pLinkTable.pHead = pNode
		pLinkTable.pTail = pNode
		pLinkTable.SumOfNode = 1
	}else{
		pLinkTable.pTail.pNext = pNode
		pLinkTable.pTail = pNode
		pLinkTable.SumOfNode ++
		pLinkTable.pTail.pNext = nil
	}

	return 0
}

func DelLinkTableNode (pLinkTable *LinkTable, pNode *LinkTableNode) int {
	if pNode == pLinkTable.pHead {
		pLinkTable.pHead = pLinkTable.pHead.pNext
		pLinkTable.SumOfNode --
		if pLinkTable.SumOfNode == 0 {
			pLinkTable.pTail = nil
		}
		return 0
	}
	tmp := pLinkTable.pHead
	for tmp != nil {
		if tmp.pNext == pNode {
			tmp.pNext = pNode.pNext
			pLinkTable.SumOfNode --
			if pLinkTable.SumOfNode == 0 {
				pLinkTable.pTail = nil
			}
			return 0
		}
		tmp = tmp.pNext
	}
	return -1
}

func GetLinkTableHead ( pLinkTable *LinkTable) *LinkTableNode {
	if pLinkTable ==nil {
		return nil
	}
	return pLinkTable.pHead
}

func GetNextLinkTableNode( pLinkTable *LinkTable, pNode *LinkTableNode) *LinkTableNode{
	if pLinkTable == nil || pNode == nil {
		return nil
	}
	tmp := pLinkTable.pHead
	for tmp != nil {
		if tmp == pNode {
			return tmp.pNext
		}
		tmp = tmp.pNext
	}
	return nil
}

func SearchLinkTableNode(pLinkTable *LinkTable, Conditon func(*LinkTableNode, interface{}) int, args interface{}) *LinkTableNode{
	if pLinkTable ==nil || Conditon == nil{
		return nil
	}
	pNode := pLinkTable.pHead
	for pNode != nil{
		if Conditon(pNode, args) == 0{
			return pNode
		}
		pNode = pNode.pNext
	}
	return nil
}