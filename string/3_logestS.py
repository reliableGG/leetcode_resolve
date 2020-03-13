class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        start = maxlength = 0
        dict = {}
        for i in range(len(s)):
            if s[i] in dict and start <= dict[s[i]]:
                start = dict[s[i]] + 1
            else:
                maxlength = max(maxlength, i - start + 1)
            dict[s[i]] = i
        return maxlength

def stringToString(input):
    return input[1:-1].decode('string_escape')

def intToString(input):
    if input is None:
        input = 0
    return str(input)

def main():
    s = "abcdb"
    s = "tmmzuxt"
    ret = Solution().lengthOfLongestSubstring(s)
    print(ret)

if __name__ == '__main__':
    main()
