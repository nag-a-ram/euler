#include <stdio.h>

int odd(int n);

int main(void) 
{  
    int sumsquare = 0;
    int squaresum = 0;
    int history   = 0;
    
    for (int i = 1; i <= 100; i++) {
        squaresum += i;
        history   += odd(i);
        sumsquare += history;
    }

    squaresum *= squaresum;
    printf("square of sum: %d\n", squaresum);
    printf("sum of square: %d\n", sumsquare);
    printf("\ndifference: %d\n", squaresum - sumsquare);
}

int odd(int n)
{
    // gets the nth odd number
    return (n * 2) - 1;
}

// written by @aliazam
