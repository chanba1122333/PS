#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
#define FOR(i, a, A) for (ll i = a; i < A; ++i)

int main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    string s;
    cin >> s;

    int pos[26];
    FOR (i,0,26){
        pos[i] = -1;
    }
    FOR (i,0,s.size()) {
        ll idx = s[i] - 'a';
        if (pos[idx] == -1) {
            pos[idx] = i;
        }
    }
    
    FOR (i,0,26){
        cout << pos[i] << " ";
    }
    
    return 0;
}
