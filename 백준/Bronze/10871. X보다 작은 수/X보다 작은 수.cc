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

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    ll N, X, a;
    cin >> N >> X;
    bool first = true;
    for (ll i = 0; i < N; ++i) {
        cin >> a;
        if (a < X) {
            if (!first) cout << ' ';
            cout << a;
            first = false;
        }
    }
    cout << '\n';
    return 0;
}