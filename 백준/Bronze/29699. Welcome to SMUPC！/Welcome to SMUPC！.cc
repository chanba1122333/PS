#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;

    string s = "WelcomeToSMUPC";

    n = (n - 1) % 14;

    cout << s[n];
    return 0;
}
