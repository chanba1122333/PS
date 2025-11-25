#include <bits/stdc++.h>
using namespace std;

typedef long long ll;

int main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    string arr[3] = {"U", "O", "S"};
    ll x;
    cin >> x;

    ll idx = (x - 1) % 3;
    cout << arr[idx];

    return 0;
}
