#include <stdio.h>
#include <stdlib.h>
#define SIZE 3

int main() {
	int a[SIZE];
	printf("a\n");
	for (int i = 0; i < SIZE; i++) {
		a[i] = i;
		printf("%d ", a[i]);
	}
	printf("\n");
	int* b = (int*) malloc(SIZE * sizeof(int));
	printf("b\n");
	for (int i = 0; i < SIZE; i++) {
		b[i] = i*2;
		printf("%d ", b[i]);
	}
	printf("\n");
	int c[SIZE][SIZE];
	printf("c\n");
	for (int i = 0; i < SIZE; i++) {
		for (int j = 0; j < SIZE; j++) {
			c[i][j] = i + j;
			printf("%d ", c[i][j]);
		}
		printf("\n");
	}
	int** d = (int**) malloc(sizeof(int*) * SIZE);
	for (int i = 0; i < SIZE; i++) {
		d[i] = (int*) malloc(sizeof(int) * SIZE);
	}
	for (int i = 0; i < SIZE; i++) {
		for (int j = 0; j < SIZE; j++) {
			d[i][j] = i + j;
			printf("%d ", d[i][j]);
		}
		printf("\n");
	}
	int e[SIZE][SIZE][SIZE];
	printf("\n\nE\n");
	for (int i = 0; i < SIZE; i++)
		for (int j = 0; j < SIZE; j++)
			for (int k = 0; k < SIZE; k++)
				e[i][j][k] = i + j + k;
	for (int i = 0; i < SIZE; i++) {
		for (int j = 0; j < SIZE; j++) {
			for (int k = 0; k < SIZE; k++) {
				printf("%d ", e[i][j][k]);
			}
			printf("\n");
		}
		printf("\n");
	}
	return 0;
}
