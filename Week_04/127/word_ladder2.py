from typing import List


class Solution:
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        start_used = [False] * len(wordList)
        end_used = [False] * len(wordList)
        word_index_map = {w: i for i, w in enumerate(wordList)}
        if endWord not in word_index_map:
            return 0
        step = 1
        start_queue = [beginWord]
        end_queue = [endWord]
        end_used[word_index_map[endWord]] = True
        letters = [chr(o) for o in range(ord('a'), ord('z') + 1)]
        while start_queue:
            l = len(start_queue)
            step += 1
            for i in range(l):
                q = start_queue[i]
                q_char = list(q)
                for j in range(len(q)):
                    o = q_char[j]
                    for c in letters:
                        q_char[j] = c
                        s = ''.join(q_char)
                        if s in word_index_map:
                            if start_used[word_index_map[s]]:
                                continue
                            if end_used[word_index_map[s]]:
                                return step + 1
                            start_queue.append(s)
                            start_used[word_index_map[s]] = True
                    q_char[j] = o
            start_queue = start_queue[l:]
            if len(start_queue) > len(end_queue):
                start_queue, end_queue = end_queue, start_queue
                start_used, end_used = end_used, start_used
        return 0
