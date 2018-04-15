#include <xgboost/test.h>

//gcc -I../include -shared -fPIC -o libtest.so test.c
//g++-7 -std=c++11 -fopenmp -I../include -I../rabit/include -I../dmlc-core/include -L. -ltest -lxgboost -o main main.c

int SayHi(){

    printf("hello world\n");
    return 0;
}

