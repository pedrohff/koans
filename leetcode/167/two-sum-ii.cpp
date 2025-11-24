#include <iostream>
#include <string>
#include <vector>
#include <sstream>

using namespace std;

/*defiprompt
INPUTS:
[2,7,11,15]
9
OUTPUT:
[1,2]
-*-
INPUTS:
[2,3,4]
6
OUTPUT:
[1,3]
-*-
INPUTS:
[-1,0]
-1
OUTPUT:
[1,2]
-*-
INPUTS:
[5,25,75]
100
OUTPUT:
[2,3]
-*-
INPUTS:
[-10,-8,-2,1,2,5,6]
0
OUTPUT:
[3,5]
-*-
INPUTS:
[2,3]
5
OUTPUT:
[1,2]
*/

std::vector<int> parseStringToArray(std::string input) {
  std::vector<int> result;
  std::stringstream ss(input);
  char ch;
  int num;

  ss >> ch;
  while (ss >> num) { 
    result.push_back(num);
    ss >> ch;
  }
  return result;
}

std::vector<int> twoSum(vector<int>& numbers, int target) {
  int left = 0;
  int right = numbers.size()-1;
  while (true) {
    int first = numbers[left];
    int last = numbers[right];
    int sum = first + last;

    // cout << left << " " << right << " | " << first << " " << last << " | " << sum << "\n";
    if (sum == target) {
      return {left+1, right+1};
    }

    if (sum > target) {
      right--;
      continue;
    } else {
      left++;
      continue;
    }

  }
  return {left, right};
}

int main() {
  std::string arrayStr;
  int target;
  cin >> arrayStr;
  cin >> target;
  
  std::vector<int> numbers = parseStringToArray(arrayStr);
  std::vector<int> result = twoSum(numbers, target);
  if (result.size() != 2) {
    cout << "result's array size is invalid";
    return 1;
  }
  cout << "[" << result[0] << "," << result[1] << "]";
  return 0;
}
