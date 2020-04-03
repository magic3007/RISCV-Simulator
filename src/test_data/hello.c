#include <stdio.h>

int a = 3;
int b = 2;
int result;

// result:  6

int main(){
	for(int i = 0; i < a; i++)
		result += b;
	return 0;
}