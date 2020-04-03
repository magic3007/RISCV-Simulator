#include <stdlib.h>

int N, K, M;
int A[3][3], B[3][3], result[3][3];

int main(){
	N = 3;
    K = 3;
    M = 3;

	for(int i = 0; i < N; i++)
		for(int j = 0; j < K; j++)
			A[i][j] = B[i][j]= i * N + j;

	for(int i = 0; i < N; i++)
        for(int k = 0; k < K; k++)
            for(int j = 0; j < M; j++)
                result[i][j] += A[i][k] * B[k][j];


    return 0;
}

/*
result:
15	18	21
42	54	66
69	90	111
*/