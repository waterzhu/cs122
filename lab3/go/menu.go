package main
import 
(
	"fmt"
	"os/exec"
	"os"
	"bufio"
)
	
func main()  {
	cmd1 := DataNode {"help", "Some tips about this system.", nil}
    cmd2 := DataNode{"hello", "Say hello to users,", &cmd1}
    cmd3 := DataNode{"quit", "Quit this system.", &cmd2}
    cmd4 := DataNode{"sum", "Get sum of two integer.", &cmd3}
    cmd5 := DataNode{"compare", "Compare two integer.", &cmd4}
    cmd6 := DataNode{"echo", "Output what you input.",  &cmd5}
    cmd7 := DataNode{"ls", "List files in this dir.",  &cmd6}
    cmd8 := DataNode{"mkdir", "Make a new dir.", &cmd7}
	for true {
		var cmd string
		fmt.Println("Input a cmd--> ")
		fmt.Scanln(&cmd)
		p := FindCmd(&cmd8, cmd)
		if p == nil{
			fmt.Println("WORRY COMMAND!!")
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