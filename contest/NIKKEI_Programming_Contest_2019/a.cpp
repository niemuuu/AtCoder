// https://nikkei2019-qual.contest.atcoder.jp/tasks/nikkei2019_qual_a

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

    int n, a, b;
    cin >> n >> a >> b;

    int ma = min(a, b);
    int mi = (a+b) - n;
    if (mi < 0) {
        mi = 0;
    }

    cout << ma << ' ' << mi << "\n";
}