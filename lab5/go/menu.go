package main
import 
(
	"fmt"
	"os/exec"
	"os"
	"bufio"


)




var head *LinkTable = nil

func main()  {
	InitMenuData(&head)
	//ShowAllCmd(head)
	for true{
		var cmd string
		fmt.Printf("Input a cmd--> ")
		fmt.Scanln(&cmd)
		p := FindCmd(head, cmd)
		if p == nil{
			fmt.Println("WRONG COMMAND!!")
			continue
		}
		fmt.Println(p.cmd+" --- "+p.desc)
		
		handler(p)
		
	}
}

func hello() {
	fmt.Println("Hello, Wlecome to this menu~")
}
func help() {
	ShowAllCmd(head)
}
func quit() {
	os.Exit(1)
}
func sum() {
	a:=0
	b:=0

	fmt.Println("Input two numbers:")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println("Sum of two numbers: ",a+b)
}
func compare() {
	c:=0
	d:=0
	bigger:=0

    fmt.Println("Input two numbers:")
    fmt.Scanf("%d %d", &c, &d)
    if c>d {
    	bigger = c
    } else {
    	bigger = d
    }
    fmt.Println("The bigger one is: ", bigger)
}
func echo() {
	buf := bufio.NewReader(os.Stdin)

	line, err := buf.ReadString('\n')
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(line)
}
func mkdir() {
	fmt.Println("You want make a new dir? Sorry, i have no access to it.")
}
func ls() {
	cm := exec.Command("ls")
    out, err := cm.CombinedOutput()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(out))
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
