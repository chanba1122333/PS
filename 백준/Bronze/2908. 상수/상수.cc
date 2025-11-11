#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef vector<ll> vll;
#define FOR(i,a,A) for (ll i = a; i < A; ++i)


ll rev(ll x) {
    ll r = 0;
    while (x > 0) {
        r = r * 10 + (x % 10);
        x /= 10;
    }
    return r;
}

int main(){
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    ll a, b;
    cin >> a >> b;

    ll ra = rev(a);
    ll rb = rev(b);
    if (ra > rb){
        cout << ra;
    } else {
        cout << rb;
    }
    return 0;
}
