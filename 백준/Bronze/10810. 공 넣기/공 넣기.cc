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
    cin.tie(nullptr);

    ll n, m;
    cin >> n >> m;

    vll arr(n + 1, 0);

    ll a, b, c;
    FOR(i, 0, m)
    {
        cin >> a >> b >> c;
        FOR(j, a, b + 1)
        arr[j] = (ll)c;
    }

    FOR(i, 1, n + 1)
    {
        cout << arr[i];
        if (i < n)
            cout << ' ';
    }
    return 0;
}