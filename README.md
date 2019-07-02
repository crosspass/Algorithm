# Algorithm

## ｃ语言实现(算法设计与分析基础)的练习题
  1. 检查变位词
  
  新建一个数组用来保存单词中字母的个数。 单词的字母对应整数作为数组的 index 来保存字母的数量。
  第一次遍历第一个单词，保存其每个字母的数量。然后遍历第二个单词的字母，将其对应的数组中数量进行减少。
  遍历数组判断数组中的字母的数量是否全部为0， 若全部为 0 ， 则 2 个单词为变位词，否则不是变位词。
  
  ``` c
  size = 'z' + 1;
  int array[size];
  
  int check(char * s1, char * s2) { 
    int i = 0;
    for(int i=0; i < size; i++) array[i] = 0;
    while(*s1 != '\0') array[*s1++]++;
    while(*s2 != '\0') array[*s2++]--;
    for(i=0; i < size; i++) if(array[i] != 0) return 0;
    return 1;
   }
  ```
