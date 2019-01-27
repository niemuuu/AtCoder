// https://nikkei2019-qual.contest.atcoder.jp/tasks/nikkei2019_qual_b

#include <bits/stdc++.h>

#define REP(i, n) for(int i = 0; i < n; i++)
#define REPR(i, n) for(int i = n; i >= 0; i--)
#define FOR(i, m, n) for(int i = m; i < n; i++)
#define INF 1e9
#define ALL(v) v.begin(), v.end()

using namespace std;

typedef long long ll;

int n;
string a, b, c;

int main() {
    cin.tie(0);
    ios::sync_with_stdio(false);

    cin >> n >> a >> b >> c;

    string words[3];
    words[0] = a;
    words[1] = b;
    words[2] = c;

    int ch = 0;
    REP(i, n) {
        int m = 0;
        if (words[0][i] == words[1][i]) {
            ++m;
        }
        if (words[0][i] == words[2][i]) {
            ++m;
        }
        if (words[1][i] == words[2][i]) {
            ++m;
        } 

        if (m == 0) {
            ch += 2;
        } else if (m == 1) {
            ch++;
        }
    }

    cout << ch << "\n";
}