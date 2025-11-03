#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
#define FOR(i,a,A) for (ll i = a; i < A; ++i)

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    ll x; 
    cin >> x;

    ll a; 
    cin >> a;
    ll mn = a, mx = a;

    FOR(i, 1, x) {
        cin >> a;
        if (a > mx) mx = a;
        if (a < mn) mn = a;
    }

    cout << mn << ' ' << mx << '\n';
    return 0;
}