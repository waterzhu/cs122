#include <stdio.h>
#include <stdlib.h>
#include "linktable.h"

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
#define CMD_NUM     128

typedef struct DataNode
{
    tLinkTableNode *pNext;
    char* cmd;
    char* desc;
    int   (*handler)();
    struct DataNode *next;
}tDataNode;
tDataNode *FindCmd(tLinkTable *head, char *cmd)
{
    tDataNode* pNode=( tDataNode *) GetLinkTableHead(head);
    while( pNode != NULL)
    {
        if( strcmp(pNode->cmd,cmd) == 0)
        {
            return pNode;
        }
        pNode=( tDataNode *) GetNextLinkTableNode(head,( tLinkTableNode *) pNode);
    }
    return NULL;
}
int ShowAllCmd( tLinkTable * head)
{
    tDataNode * pNode=( tDataNode *)GetLinkTableHead(head);
    while(pNode != NULL)
    {
        printf("%-10s - %s\n", pNode->cmd, pNode->desc);
        pNode =( tDataNode *) GetNextLinkTableNode( head, (tLinkTableNode *) pNode);
    }
    return 0;
}
int InitMenuData( tLinkTable ** ppLinkTable)
{
    * ppLinkTable = CreateLinkTable();
    tDataNode* pNode = ( tDataNode *) malloc( sizeof( tDataNode));
    pNode->cmd = "help";
    pNode->desc = "help command:";
    pNode->handler = help;
    AddLinkTableNode( * ppLinkTable,( tLinkTableNode *) pNode);
    pNode=( tDataNode *) malloc( sizeof( tDataNode));
    pNode->cmd = "hello";
    pNode->desc = "Say hello to users.";
    pNode->handler = hello;
    AddLinkTableNode(*ppLinkTable,(tLinkTableNode *)pNode);
    pNode = (tDataNode*)malloc(sizeof(tDataNode));
    pNode->cmd = "quit";
    pNode->desc = "Quit from Menu";
    pNode->handler = quit;
    AddLinkTableNode( * ppLinkTable,( tLinkTableNode *) pNode);
    pNode = ( tDataNode *) malloc( sizeof( tDataNode));
    pNode->cmd = "sum";
    pNode->desc = "An integer plus another";
    pNode->handler = sum;
    AddLinkTableNode(*ppLinkTable, (tLinkTableNode*) pNode);
    pNode = ( tDataNode *) malloc( sizeof( tDataNode));
    pNode->cmd = "compare";
    pNode->desc = "Get the bigger from two integers.";
    pNode->handler = compare;
    AddLinkTableNode(*ppLinkTable,(tLinkTableNode *)pNode);
    pNode = (tDataNode*)malloc(sizeof(tDataNode));
    pNode->cmd = "ls";
    pNode->desc = "List files in this dir.";
    pNode->handler = ls;
    AddLinkTableNode(*ppLinkTable,(tLinkTableNode *)pNode);
    pNode = (tDataNode*)malloc(sizeof(tDataNode));
    pNode->cmd = "echo";
    pNode->desc = "Output what you input.";
    pNode->handler = echo;
    AddLinkTableNode(*ppLinkTable,(tLinkTableNode *)pNode);
    pNode = (tDataNode*)malloc(sizeof(tDataNode));
    pNode->cmd = "mkdir";
    pNode->desc = "Make a new dir.";
    pNode->handler = mkdir;
    AddLinkTableNode(*ppLinkTable,(tLinkTableNode *)pNode);
    return 0;
}
tLinkTable * head= NULL;


int main()
{
    InitMenuData(&head);
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
