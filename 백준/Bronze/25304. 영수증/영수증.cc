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
    ll x, n, a, b;
    ll sum = 0;
    vll s;
    cin >> x >> n;
    FOR(i, 0, n)
    {
        cin >> a >> b;
        s.emplace_back(a * b);
    }
    for (auto r : s)
    {
        sum += r;
    }
    if (sum == x)
        cout << "Yes";
    else
        cout << "No";
    return 0;
}