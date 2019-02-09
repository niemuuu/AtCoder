// https://atcoder.jp/contests/yahoo-procon2019-qual/tasks/yahoo_procon2019_qual_b

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
    
    char p;
    int c1 = 0;
    int c2 = 0;
    int c3 = 0;
    int c4 = 0;
    while(cin >> p) {
        if (p == '1') ++c1;
        if (p == '2') ++c2;
        if (p == '3') ++c3;
        if (p == '4') ++c4;
    }
    if (c1 == 0 || c2 == 0 || c3 == 0 || c4 == 0) cout << "NO\n";
    else if (c1 >= 3 || c2 >= 3 || c3 >= 3 || c4 >= 3) cout << "NO\n";
    else cout << "YES\n";

    return 0;
}