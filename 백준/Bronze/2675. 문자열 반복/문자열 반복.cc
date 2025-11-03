#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
#define FOR(i,a,A) for (ll i = a; i < A; ++i)

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    
    ll a,b;
    string s;
    cin >> a;
    FOR(t, 0, a) {
    cin >> b >> s;
    for (size_t p = 0; p < s.size(); ++p) {
        char ch = s[p];
        for (ll j = 0; j < b; ++j) {
            cout << ch;
            }
        }
    cout << '\n';
    }
    return 0;
}