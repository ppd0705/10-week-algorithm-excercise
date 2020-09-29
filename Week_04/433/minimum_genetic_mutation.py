class Solution:
    def minMutation(self, start, end, bank):
        is_used = {b: False for b in bank}

        if end not in bank:
            return -1

        queue = [start]
        count = 0

        mutation_map = {
            "A": "CGT",
            "C": "AGT",
            "G": "CAT",
            "T": "CGA",
        }
        while queue:
            n = len(queue)
            count += 1
            for i in range(n):
                s = queue[i]
                for j in range(len(s)):
                    for c in mutation_map[s[j]]:
                        ss = s[:j] + c + s[j + 1:]
                        if ss == end:
                            return count
                        if ss in is_used and not is_used[ss]:
                            queue.append(ss)
                            is_used[ss] = True
            queue = queue[n:]
        return -1

    def test(self):

        for start, end, bank, target in [
            ("AACCGGTT", "AAACGGTA", ["AACCGGTA", "AACCGCTA", "AAACGGTA"], 2),
            ("AACCGGTT", "AAACGGTA", ["GACCGGTA", "AACCGCTA", "AAACGGTA"], -1),
        ]:
            ans = self.minMutation(start, end, bank)
            assert ans == target, f"target: {target}, ans: {ans}"


if __name__ == "__main__":
    Solution().test()
