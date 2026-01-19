package isogram

import "strings"

func IsIsogram(word string) bool {
   for i := range word {
           for j:=i+1; j<len(word); j++ {
               if word[j] != byte('-') && word[j] != byte(' ') {
                   
               if strings.ToLower(string(word[i]))==strings.ToLower(string(word[j])) {
                   return false
               }
               }
           }
   }
    return true
}
