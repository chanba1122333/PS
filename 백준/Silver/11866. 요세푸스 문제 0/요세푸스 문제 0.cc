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
    ll N,K,tmp;
    vll v;
    vll f;
    cin >> N >> K;
    FOR(i,1,N+1){
        v.emplace_back(i);
    }
    
    ll idx = 0;
    while (!v.empty()) {
        idx = (idx + K - 1) % v.size();
        f.emplace_back(v[idx]);
        v.erase(v.begin() + idx);
    }
    
    cout << "<";
    FOR(i, 0, f.size()) {
        cout << f[i];
        if (i + 1 != f.size()){
            cout << ", ";  
        } 
    }
    
    cout << ">\n";
    return 0;
}