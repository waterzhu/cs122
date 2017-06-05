package main
import 
(
	"fmt"
	"os/exec"
	"os"
	"strconv"
	//"bufio"
	"strings"


)

var head *LinkTable = nil

func main(){
		MenuConfig("help", "Some tips about this system.", help)
		MenuConfig("hello", "Say hello to users.", hello)
		MenuConfig("quit", "Quit this system.", quit)
		MenuConfig("sum", "Get sum of two integer.", sum)
		MenuConfig("compare", "Compare two integer.", compare)
		MenuConfig("echo", "Output what you input.", echo)
		MenuConfig("ls", "List files in this dir.", ls)
		MenuConfig("mkdir", "ake a new dir.", mkdir)
		ExecuteMenu()
	
}

func hello(argc int, argv []string) int {
	fmt.Println("Hello, Wlecome to this menu~")
	return 0
}
func help(argc int, argv []string) int{
	ShowAllCmd(head)
	return 0
}
func quit(argc int, argv []string) int {
	os.Exit(1)
	return 0
}
func sum(argc int, argv []string) int {
	/*a:=0
	b:=0

	fmt.Println("Input two numbers:")
	fmt.Scanf("%d %d", &a, &b)*/
	a,error := strconv.Atoi(argv[1])
	if error != nil{
	fmt.Println("字符串转换成整数失败")
	}
	b,error := strconv.Atoi(argv[2])
	if error != nil{
	fmt.Println("字符串转换成整数失败")
	}
	fmt.Println("Sum of two numbers: ",a+b)
	return 0
}
func compare(argc int, argv []string) int{
	var bigger int
	/*c:=0
	d:=0
    fmt.Println("Input two numbers:")
    fmt.Scanf("%d %d", &c, &d)
    */
    c,error := strconv.Atoi(argv[1])
	if error != nil{
	fmt.Println("字符串转换成整数失败")
	}
	d,error := strconv.Atoi(argv[2])
	if error != nil{
	fmt.Println("字符串转换成整数失败")
	}
	if c>d {
    	bigger = c
    } else {
    	bigger = d
    }
    fmt.Println("The bigger one is: ", bigger)
    return 0
}
func echo(argc int, argv []string) int {
	s := argv[1:len(argv)]
	fmt.Println(strings.Join(s, " "))
	/*buf := bufio.NewReader(os.Stdin)

	line, err := buf.ReadString('\n')
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(line)*/
	return 0
}
func mkdir(argc int, argv []string) int {
	fmt.Println("You want make a new dir? Sorry, i have no access to it.")
	return 0
}
func ls(argc int, argv []string) int {
	cm := exec.Command("ls")
    out, err := cm.CombinedOutput()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf(string(out))
    return 0
}