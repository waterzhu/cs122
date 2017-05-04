#include <stdio.h>
#include <stdlib.h>
#include "linklist.h"

void help();
void hello();
void quit();
void sum();
void compare();
void echo();
void ls();
void mkdir();

#define CMD_MAX_LEN 128
#define DESC_LEN    1024
#define CMD_NUM     10

static tDataNode head[] = 
{
    {"help", "Some tips about this system.", help, &head[1]},
    {"hello", "Say hello to users,", hello, &head[2]},
    {"quit", "Quit this system.", quit, &head[3]},
    {"sum", "Get sum of two integer.", sum, &head[4]},
    {"compare", "Compare two integer.", compare, &head[5]},
    {"echo", "Output what you input.", echo, &head[6]},
    {"ls", "List files in this dir.", ls, &head[7]},
    {"mkdir", "Make a new dir.", mkdir, NULL}
};

int main()
{
    while(1)
    {
        char cmd[CMD_MAX_LEN];
        printf("Input a cmd: ");
        scanf("%s", cmd);
        tDataNode *p = FindCmd(head,cmd);
        if(p == NULL)
        {
            printf("WORRY COMMAND!\n");
            continue;
        }
        printf("%s - %s\n",p->cmd, p->desc);
        if(p->handler != NULL)
        {
            p->handler();
        }
    }
    return 0;
}

void help()
{
    ShowAllCmd(head);
}

void hello()
{ 
    printf("Hello ,Welcome to this menu~\n");
}

void quit()
{
    exit(0);
}

void sum()
{
    int a = 0;
    int b = 0;
 
    printf("Input two integers:\n");
    scanf("%d, %d", &a, &b);
    printf("a+b = %d\n", a+b);
}

void compare()
{
    int c = 0;
    int d = 0;
    int bigger = 0;

    printf("Input two integers:\n");
    scanf("%d, %d", &c, &d);
    bigger = c>d?c:d;
    printf("The bigger one is : %d\n", bigger);
}

void echo()
{
    char buf[128];
    printf("Start your performance:\n");
    fgetc(stdin);
    fputs(fgets(buf, 128, stdin), stdout);
}

void mkdir()
{
    printf("Make a new dir.\n");
}

void ls()
{
    system("ls");
}
