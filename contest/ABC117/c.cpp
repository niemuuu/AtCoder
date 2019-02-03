// https://atcoder.jp/contests/abc117/tasks/abc117_c

#include <bits/stdc++.h>

#define REP(i, n) for(int i = 0; i < n; i++)
#define REPR(i, n) for(int i = n; i >= 0; i--)
#define FOR(i, m, n) for(int i = m; i < n; i++)
#define INF 1e9
#define ALL(v) v.begin(), v.end()

using namespace std;

typedef long long ll;
typedef pair<int, int> diff;

int main() {
    cin.tie(0);
    ios::sync_with_stdio(false);
    /* ファイル読み込み用 */
    // std::ifstream in("stdin");
    // std::cin.rdbuf(in.rdbuf());
    /* 読み込み done */
    
    ll N, M;
    cin >> N >> M;

    if (N >= M) {
        cout << '0' << "\n";
        return 0;
    }

    vector<int> points;
    REP(i, M) {
        int p;
        cin >> p;
        points.push_back(p);
    }

    sort(points.begin(), points.end());

    vector<ll> D;
    REP(i, (int)(points.size() - 1)) {
        ll d = points[i+1] - points[i]; // 差分
        D.push_back(d);
    }

    sort(D.begin(), D.end());

    int result = 0;
    REP(i, M - N) {
        result += D[i];
    };
    cout << result << "\n";
} 