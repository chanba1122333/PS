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
    ll n;
    cin >> n;
    vll score(n);
    ll mx = 0;
    FOR(i,0,n)
    {
        cin >> score[i];
        mx = max(mx, score[i]);
    }
    double sum = 0;
    FOR(i,0,n)
    {
        sum += (double)score[i] / mx * 100;
    }
    cout << sum / n;
    return 0;
}