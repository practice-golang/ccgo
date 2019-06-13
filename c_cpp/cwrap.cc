#include "foo.hpp"
#include "foo.h"

void DoSay(char* data, char** mytest) {
	// SayHello(data, &data);
	SayHello(data, mytest);
}

// int main() {
// 	DoSay();

// 	return 0;
// }