#include <stdio.h>
#include <string.h>
#include "linklist.h"

tDataNode* FindCmd(tDataNode* head, char* cmd)
{
    tDataNode *p = head;
    if(head == NULL ||cmd == NULL)
    {
        return NULL;
    }
    while(p != NULL)
    {
        if(strcmp(p->cmd, cmd) == 0)
        {
            return p;
        }
    p = p->next;
    }
    return NULL;
}

int ShowAllCmd(tDataNode* head)
{
    printf("Menu List:\n");
    tDataNode *p = head;
    while(p != NULL)
    {
        printf("%s -  %s\n", p->cmd, p->desc);
        p = p->next;
    }
    return 0;
} 
