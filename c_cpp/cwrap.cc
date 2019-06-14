#include "foo.hpp"
#include "foo.h"
#include <cstdio>

char* DoSay(char* data, char** mytest) {
	// SayHello(data, &data);
	char* result = SayHello(data, mytest);

	printf("cwrap.cc from foo.cpp return : %s\n", result);

	return result;
}

// int main() {
// 	DoSay();

// 	return 0;
// }