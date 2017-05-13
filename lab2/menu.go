package main
import 
(
	"fmt"
	"strings"
	"os/exec"
	"os"
	"bufio"
)

var cmd string

func main()  {
	fmt.Println("This system has cmd below:")
	fmt.Println("hello")
	fmt.Println("help")
	fmt.Println("quit")
	fmt.Println("sum")
	fmt.Println("compare")
	fmt.Println("echo")
	fmt.Println("ls")
	fmt.Println("mkdir")
	for true {
		fmt.Printf("\n$:")
        fmt.Scanln(&cmd)
	    if strings.EqualFold(cmd, "hello") {
	    	hello()
	    }else if strings.EqualFold(cmd, "help") {
	    	help()
	    }else if strings.EqualFold(cmd, "quit") {
	    	quit()
	    }else if strings.EqualFold(cmd, "sum") {
	    	sum()
	    }else if strings.EqualFold(cmd, "compare") {
	    	compare()
	    }else if strings.EqualFold(cmd, "echo") {
	    	echo()
	    }else if strings.EqualFold(cmd, "ls") {
	    	ls()
	    }else if strings.EqualFold(cmd, "mkdir") {
	    	mkdir()
	    }else {
	    	fmt.Println("Warning: WORRY COMMAND!!")
	    }
	}
	
}
func hello() {
	fmt.Println("Hello, Wlecome to this menu~")
}
func help() {
	fmt.Println("Need help? I will try my best to help you!")
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