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

    ll n, m, tmp;
    ll a, b;
    cin >> n >> m;
    vll arr(n + 1);
    FOR(i, 1, n + 1)
    {
        arr[i] = i;
    }
    FOR(i, 0, m)
    {
        cin >> a >> b;
        tmp = arr[a];
        arr[a] = arr[b];
        arr[b] = tmp;
    }

    FOR(i, 1, n + 1)
    {
        cout << arr[i];
        if (i < n)
            cout << ' ';
    }
    return 0;
}