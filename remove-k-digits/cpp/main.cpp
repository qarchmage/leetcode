#include <iostream>
using namespace std;
class Solution {
public:
    string removeKdigits(string num, int k) {
       string ans = "";                                         // treat ans as a stack in below for loop
       
       for (char c : num) {
           while (ans.length() && ans.back() > c && k) {
               ans.pop_back();                                  // make sure digits in ans are in ascending order
               k--;                                             // remove one char
           }
           
           if (ans.length() || c != '0') { ans.push_back(c); }  // can't have leading '0'
       }
       
       while (ans.length() && k--) { ans.pop_back(); }          // make sure remove k digits in total
       
       return ans.empty() ? "0" : ans;
}
};

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
