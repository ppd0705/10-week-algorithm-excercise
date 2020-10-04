from typing import List
from collections import defaultdict


class Solution:
    def findLadders(self, beginWord: str, endWord: str, wordList: List[str]) -> List[List[str]]:
        if endWord not in wordList:
            return []
        buckets = defaultdict(list)

        wordList.append(beginWord)
        n = len(beginWord)
        for w in wordList:
            for i in range(n):
                re = w[:i] + "*" + w[i + 1:]
                buckets[re].append(w)

        level_map = defaultdict(int)
        level_map[beginWord] = 1
        pre_nodes_map = defaultdict(list)

        queue = [beginWord]

        while queue:
            curr = queue.pop(0)
            curr_level = level_map[curr]

            if endWord in level_map and curr_level + 1 > level_map[endWord]:
                # all solutions has found
                break

            for i in range(n):
                re = curr[:i] + "*" + curr[i + 1:]
                for w in buckets[re]:
                    level = level_map[w]
                    if level == 0:
                        # new node
                        level_map[w] = curr_level + 1
                        queue.append(w)
                        pre_nodes_map[w] = [curr]
                    elif level == curr_level + 1:
                        # another solution
                        pre_nodes_map[w].append(curr)
        res = []

        def dfs(node, level, solution):
            solution[level - 1] = node
            if level == 1:
                res.append(solution.copy())
                return
            for pre_node in pre_nodes_map[node]:
                dfs(pre_node, level - 1, solution)

        if endWord in level_map:
            dfs(endWord, level_map[endWord], [""] * level_map[endWord])
        return res
