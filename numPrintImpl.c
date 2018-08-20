#include "numPrinter.h"

void print(int a)
{
	int x = 0;

	while(x<50){
		printf("\n Printing %d for the %d'th time\n", a, x);
		x++;
	}
	printf("Done!");
}
