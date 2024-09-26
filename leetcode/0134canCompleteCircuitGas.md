class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        length = len(gas)
        spare = 0
        min_spare = float('inf')
        min_index = 0

        for i in range(length):
            spare += gas[i] - cost[i]
            if spare < min_spare:
                min_spare = spare
                min_index = i

        return -1 if spare < 0 else (min_index + 1) % length