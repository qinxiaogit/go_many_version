#include <stdio.h>
#include "callC.h"

void cHello(){
    printf("hello from C!\n");
}

void printMessage(char *message){
    printf("go send me %s\n",message);
}