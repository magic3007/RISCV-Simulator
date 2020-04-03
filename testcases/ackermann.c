#include <stdlib.h>

int ackermann(int m, int n){
  if (m == 0) return n+1;
  if (n == 0) return ackermann( m - 1, 1 );
  return ackermann( m - 1, ackermann( m, n - 1 ) );
}

int result;

int main(int argc, const char** argv) {

	result = ackermann(3, 2);

    return 0;
}

/*
result:
29
*/
