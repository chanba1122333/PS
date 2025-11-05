#include <bits/stdc++.h>
using namespace std;
typedef long long ll;

int main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    ll a, b, c, d;
    cin >> a >> b >> c >> d;

    ll e = c - a;
    ll f = d - b;

    set<ll> s;
    s.insert({a, b, c, d, e, f});

    ll ans = *s.begin();
    cout << ans;

    return 0;
}
