// https://atcoder.jp/contests/abc117/tasks/abc117_a

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

    double T, X;
    cin >> T >> X;
    cout << fixed;
    cout << setprecision(10) << (T/X) << "\n";
}