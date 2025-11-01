#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
#define FOR(i,a,A) for (ll i = a; i < A; ++i)

int main() {
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    ll s[9];
    set<ll> st;

    FOR(i, 0, 9) {
        cin >> s[i];
        st.insert(s[i]);
    }

    ll maxValue = *st.rbegin();

    ll pos = 0;
    FOR(i, 0, 9) {
        if (s[i] == maxValue) {
            pos = i + 1;
            break;
        }
    }

    cout << maxValue << '\n' << pos << '\n';
    return 0;
}
