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

    string s;
    getline(cin, s);

    if (s.empty())
    {
        cout << 0;
        return 0;
    }

    ll left = 0;
    while (left < s.size() && s[left] == ' ')
    {
        left++;
    }

    if (left == s.size())
    {
        cout << 0;
        return 0;
    }

    ll right = s.size() - 1;
    while (right >= 0 && s[right] == ' ')
    {
        right--;
    }

    string t = s.substr(left, right - left + 1);

    ll cnt = 0;
    FOR(i, 0, (ll)t.size())
    {
        if (t[i] == ' ')
            cnt++;
    }

    cout << cnt + 1;
    return 0;
}
