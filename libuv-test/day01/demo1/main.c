#include <stdio.h>
#include <stdlib.h>
#include <uv.h>

void HelloWorld(){
    uv_loop_t *loop = malloc(sizeof(uv_loop_t));
    uv_loop_init(loop);

    printf("Now quitting.\n");
    uv_run(loop, UV_RUN_DEFAULT);

    uv_loop_close(loop);
    free(loop);
}

int main(){

    return 0;
}