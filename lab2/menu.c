#include <stdio.h>
#include <stdlib.h>

void help();
void hello();
void quit();
void sum();
void compare();
void echo();
void ls();
void mkdir();

int main()
{
    char cmd[128];
	printf("This system has the cmd below:\n");
	printf("hello\n");
        printf("help\n");
        printf("quit\n");
        printf("sum\n");
        printf("compare\n");
        printf("echo\n");
        printf("ls\n");
        printf("mkdir\n");

    while(1)
    {
        printf("\n$:");
        scanf("%s", cmd);
        if (strcmp(cmd, "hello") == 0)
        {
            hello();
	}
        else if(strcmp(cmd, "help") == 0)
        {
            help();
        }
        else if(strcmp(cmd, "quit") == 0)
        {
	    quit();
        }
        else if(strcmp(cmd, "sum") == 0)
        {
            sum();
        }
        else if(strcmp(cmd, "compare") == 0)
        {
	    compare();
        }
        else if(strcmp(cmd, "echo") == 0)
        {
            echo();
        }
        else if(strcmp(cmd, "ls") == 0)
        {
            system(cmd);
        }
        else if(strcmp(cmd, "mkdir") == 0)
        {
            mkdir();
        }
        else
        {
            printf("Warning: WORRY COMMAND!\n");
        }
   }
}

void hello()
{
    printf("Hello, Wlecome to this menu~\n");
}

void help()
{
    printf("Need help? I will try my best to help you!");
}

void quit()
{
    exit(0);
}

void sum()
{
    int a = 0;
    int b = 0;

    printf("Input two number:\n");
    scanf("%d, %d", &a, &b);
    printf("a+b = %d\n", a+b);
}

void compare()
{
    int c = 0;
    int d = 0;
    int bigger = 0;

    printf("Input two numbers:\n");
    scanf("%d, %d", &c, &d);
    bigger=(c>d?c:d);
    printf("The bigger one is: %d", bigger);
}

void echo()
{
    char buf[128];

    printf("Hello, strat your performance pls.\n");
    fgetc(stdin);
	fputs(fgets(buf,128,stdin),stdout);
}

void mkdir()
{
    printf("You want make a new dir? Sorry, i have no access to it.\n");
}
