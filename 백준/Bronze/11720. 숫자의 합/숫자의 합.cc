#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef vector<ll> vll;
#define FOR(i,a,A) for (ll i = a; i < A; ++i)

int main(){
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    ll n;
    string s;
    cin >> n >> s;
    ll sum = 0;
    FOR(i, 0, n){
        sum += s[i] - '0';
    }
    cout << sum;
    return 0;
}
