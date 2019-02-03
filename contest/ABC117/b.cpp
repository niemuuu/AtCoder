// https://atcoder.jp/contests/abc117/tasks/abc117_b

#include <bits/stdc++.h>

#define REP(i, n) for(int i = 0; i < n; i++)
#define REPR(i, n) for(int i = n; i >= 0; i--)
#define FOR(i, m, n) for(int i = m; i < n; i++)
#define INF 1e9
#define ALL(v) v.begin(), v.end()

using namespace std;

typedef long long ll;

int main() {
    cin.tie(0);
    ios::sync_with_stdio(false);

    int n;
    cin >> n;

    int sum = 0;
    int biggest = -1;
    REP(i, n) {
        int v;
        cin >> v;
        sum += v;

        if (v > biggest) {
            biggest = v;
        }
    };

    if (biggest < (sum - biggest)) {
        cout << "Yes" << "\n";
    } else {
        cout << "No" << "\n";
    }
}