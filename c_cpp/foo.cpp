#include <iostream>
#include <cstring>
#include <vector>

char* SayHello(char* data, char** mytest) {
	// std::cout << "Hello World\n";
	std::cout << "foo.cpp - data : " << data << "\n";
	std::cout << "foo.cpp - mytest[1] : " << mytest[1] <<"\n";

	std::string returnString = "myReturn ^-^";
	std::vector<char> vchar(returnString.begin(), returnString.end());
	vchar.push_back('\0');
	char* result = &vchar[0];

	return result;
}

// int main() {
// 	SayHello();

// 	return 0;
// }