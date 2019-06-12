#include "foo.hpp"
#include "foo.h"

void DoSay(char* data, char** mytest) {
	SayHello(data, &data);
}

// int main() {
// 	DoSay();

// 	return 0;
// }