class Solution:
    def getHint(self, secret: str, guess: str) -> str:
        secret_counter = [0] * 10
        guess_counter = [0] * 10
        bulls = cows = 0
        for i in range(len(secret)):
            if secret[i] == guess[i]:
                bulls += 1
            else:
                secret_counter[ord(secret[i]) - ord('0')] += 1
                guess_counter[ord(guess[i]) - ord('0')] += 1
        for i in range(len(secret_counter)):
            cows += min(secret_counter[i], guess_counter[i])

        return f"{bulls}A{cows}B"
