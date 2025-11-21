#include <bits/stdc++.h>
using namespace std;
typedef long long ll;

ll page[10];

void addCount(ll x, ll digit) {
    while (x > 0) {
        int d = x % 10;
        page[d] += digit;
        x /= 10;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    ll n;
    cin >> n;

    ll start = 1;
    ll end = n;
    ll digit = 1; 
    while (start <= end) {
        while (start % 10 != 0 && start <= end) {
            addCount(start, digit);
            start++;
        }
        if (start > end) break;

        while (end % 10 != 9 && start <= end) {
            addCount(end, digit);
            end--;
        }

        ll blockCount = (end / 10 - start / 10 + 1); 
        for (int d = 0; d <= 9; d++) {
            page[d] += blockCount * digit;
        }

        start /= 10;
        end /= 10;
        digit *= 10;
    }

    for (int d = 0; d <= 9; d++) {
        cout << page[d] << " ";
    }
    return 0;
}
