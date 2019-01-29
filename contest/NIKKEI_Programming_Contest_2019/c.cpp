// https://nikkei2019-qual.contest.atcoder.jp/tasks/nikkei2019_qual_c

#include <bits/stdc++.h>

#define REP(i, n) for(int i = 0; i < n; i++)
#define REPR(i, n) for(int i = n; i >= 0; i--)
#define FOR(i, m, n) for(int i = m; i < n; i++)
#define INF 1e9
#define ALL(v) v.begin(), v.end()

using namespace std;

typedef long long ll;
typedef pair<int, int> P;

int main() {
    cin.tie(0);
    ios::sync_with_stdio(false);
    
    int n; cin >> n;
    P pa[n];

    REP(i, n) cin >> pa[i].first >> pa[i].second;
    sort(pa, pa+n, [](P &a, P &b) {
        return a.first + a.second > b.first + b.second;
    });
    
    ll ans = 0;
    REP(i, n) {
        if (i % 2 == 0) {
            ans += pa[i].first;
        } else {
            ans -= pa[i].second;
        }
    }
    cout << ans << "\n";
}