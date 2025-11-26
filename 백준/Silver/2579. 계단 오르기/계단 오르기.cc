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

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n;
    cin >> n;

    vll stair(n + 1);
    FOR(i, 1, n + 1)
    {
        cin >> stair[i];
    }

    vll dp(n + 1, 0);

    if (n >= 1)
    {
        dp[1] = stair[1];
    }
    if (n >= 2)
    {
        dp[2] = stair[1] + stair[2];
    }
    if (n >= 3)
    {
        dp[3] = max(stair[1] + stair[3], stair[2] + stair[3]);
    }

    FOR(i, 4, n + 1)
    {
        dp[i] = max(
            dp[i - 2] + stair[i],
            dp[i - 3] + stair[i - 1] + stair[i]);
    }

    cout << dp[n] << '\n';
    return 0;
}
