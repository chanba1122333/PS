#include <iostream>
using namespace std;
int main(void) {
    long long n;
    cin >> n;
    if(n%5 == 0 || n%5 == 2) {
        cout << "CY" << '\n';
    }else {
        cout << "SK" << '\n';
    }
    return 0;
}