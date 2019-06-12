#include <iostream>

void SayHello(char* data, char** mytest) {
	// std::cout << "Hello World\n";
	std::cout << data << "\n";

	char* mystr;
	mystr = *mytest;

	std::cout << "%%" << mystr <<"^^\n";
}

// int main() {
// 	SayHello();

// 	return 0;
// }