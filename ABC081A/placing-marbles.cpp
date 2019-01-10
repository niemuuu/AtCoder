/*
ABC081A - Placing Marbles
https://atcoder.jp/contests/abs/tasks/abc081_a
*/

#include <iostream>
#include <string>
using namespace std;
int main() {
    cin.tie(0);
        ios::sync_with_stdio(false);

    string s;
    int n;

    cin >> s;
    for (int i = 0; i < 3; i++)
        if(s[i] == '1') n++;

    cout << n << "\n";
}