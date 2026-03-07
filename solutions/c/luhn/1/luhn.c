#include "luhn.h"

bool luhn(const char *num) {
    if (num[0]=='0' && !(num[1])) {
        return false;
    }

    if (num[0] == ' ' && num[1] == '0') return false;

    int end=0;
    int flag=0;
    int sum=0;
    int temp;

    while (num[end] != '\0') end++;

    for (int i=end-1; i>=0; i--) {
        if (num[i] != ' ') {
            if (num[i]>='0' && num[i]<='9') {
                if (flag == 0) {
                    sum += num[i] - '0';
                    flag = 1;
                } else {
                    temp = (num[i] - '0') * 2;
                    if (temp > 9) temp -= 9;
                    sum += temp;
                    flag = 0;
                }
            } else {
                return false;
            }
        }
    }

    if (sum % 10 == 0) {
        return true;
    }
    return false;
}