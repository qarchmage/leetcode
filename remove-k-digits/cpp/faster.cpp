#include <iostream>
using namespace std;
class Solution {
public:
    string removeKdigits(string num, int k) {
        if (num.length() <= k) return "0";
        string ret;
        for (auto d : num) {
            while (k && ret.length() && ret.back() > d) {
                ret.pop_back();
                --k;
            }
            ret.push_back(d);
        }
        ret.resize(ret.length() - k);
        int beg = 0;
        while (beg < ret.length() - 1 && ret[beg] == '0') ++beg;
        return ret.substr(beg, ret.length() - beg + 1);
    }
};

static auto fastio = [](){
    std::ios::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    return nullptr;
}();

int main() 
{
	Solution f;
	cout << f.removeKdigits("1234567890",3) << "\n";
	cout << f.removeKdigits("1120",3) << "\n";
	cout << f.removeKdigits("17890",3) << "\n";
	cout << f.removeKdigits("3303",2) << "\n"; 
	cout << f.removeKdigits("110",1) << "\n";
	cout << f.removeKdigits("220",1) << "\n";

	return 0;
}
