// https://atcoder.jp/contests/yahoo-procon2019-qual/tasks/yahoo_procon2019_qual_c

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
    
    ll k,a,b;
    cin>>k>>a>>b;

    if (a < b-2) {
        ll bis = 1;
        ll m = 0;
        REP(i, k) {
            if (m >= 1) {
                --m;
                bis += b;
                continue;
            }
            if (bis >= a && k-i >= 2) {
                bis -= a;
                ++m;
                continue;
            }
            ++bis;
        }
        cout << bis << "\n";
    } else {
        cout << k+1 << "\n";
    }
    return 0;
}