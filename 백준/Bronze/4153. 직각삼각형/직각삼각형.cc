#include <bits/stdc++.h>

using namespace std;
typedef long long ll;
typedef unsigned long long ull;
typedef pair<ll, ll> pll;
typedef pair<ull, ull> pull;
typedef const ll cll;
typedef queue<ll> qll;
typedef queue<pll> qpll;
typedef priority_queue<ll> pqll;
typedef priority_queue<pll> pqpll;
typedef vector<ll> vll;
typedef vector<pll> vpll;
typedef vector<vll> vvll;
typedef vector<vpll> vvpll;
#define FOR(i, a, A) for (ll i = a; i < A; ++i)
#define IFOR(i, a, A) for (ll i = a; i >= A; --i)

int main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);
    ll a,b,c;

    while (true) {
        cin >> a >> b >> c;

        if (a==0 && b==0 && c==0) {
            break;
        }

        ll x = a, y = b, z = c;

        if (y < x) swap(x,y);
        if (z < y) swap(y,z);
        if (y < x) swap(x,y);

        vll v = {a, b, c};
        sort(v.begin(), v.end());
        if (v[0]*v[0] + v[1]*v[1] == v[2]*v[2]){
        cout << "right\n";
        }else{
        cout << "wrong\n";
        }
    }
    return 0;
}