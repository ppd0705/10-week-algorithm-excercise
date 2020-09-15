from typing import List, Dict


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        counter: Dict[tuple, List[str]] = {}

        def count_char(s: str) -> tuple:
            base = [0] * 26
            for c in s:
                base[ord(c) - ord('a')] += 1
            return tuple(base)

        for s in strs:
            cc = count_char(s)
            if cc in counter:
                counter[cc].append(s)
            else:
                counter[cc] = [s]
        return list(counter.values())

    def test(self):
        for strs, target in [
            ([""], [[""]]),
            (["a", "b"], [["a"], ["b"]]),
            (["ab", "ba", "bb"], [["ab", "ba"], ["bb"]]),
        ]:
            ans = self.groupAnagrams(strs)
            assert sorted(ans) == sorted(target)

        print("all done")


if __name__ == "__main__":
    Solution().test()
