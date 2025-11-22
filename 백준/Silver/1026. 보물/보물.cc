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

    ll n, a, b;
    cin >> n;

    multiset<int> A;
    vector<int> B(n);

    FOR(i, 0, n) {
        cin >> a;
        A.insert(a);
    }
    FOR(i, 0, n) {
        cin >> B[i];
    }

    vector<int> idx(n);
    FOR(i, 0, n) idx[i] = i;

    sort(idx.begin(), idx.end(), [&](int i, int j){
        return B[i] > B[j];
    });

    auto itA = A.begin();

    ll ans = 0;

    FOR(i, 0, n) {
        ans += (*itA) * B[idx[i]];
        itA++;
    }

    cout << ans;

    return 0;
}
