#include <stdio.h>
#include <math.h>

int div(int n);
int tri(int n);

int main(void)
// driver function
{
	int now = 1;

	while (div(tri(now)) < 500)
		now++;

	printf("the answer is %d", tri(now));
}

int div(int n)
// returns number of divisors of n
{
	int out = 0;
	int i =   1;

	while (i <= sqrt(n)) {
		if (n % i == 0) {
		if (n / i != i) {
			out += 2;
		} else { out++; }
		}
		i++;
	}

	return out;
}

int tri(int n)
// returns the n-th triangle number
{
	return (n * (n+1))/2;
}

// written by @aliazam
