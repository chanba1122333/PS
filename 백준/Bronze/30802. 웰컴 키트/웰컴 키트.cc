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
    ll N, T, P;
    vll s(6);
    cin >> N; 
    FOR(i, 0, 6) {
        cin >> s[i]; 
    }
    cin >> T >> P;  

    ll tee = 0;
    FOR(i, 0, 6) {
    tee += (s[i] + T - 1) / T;
    }

    cout << tee << '\n' << N/P << " " << N%P;
    return 0;
}