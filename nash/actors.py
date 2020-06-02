import nashpy as nash
import numpy as np


A = np.array([[0, -1, 1], [1, 0, -1], [-1, 1, 0]])
rockPaperScissors= nash.Game(A)
rockPaperScissors
print(rockPaperScissors)
sigma_r = [0, 0, 1]
sigma_c = [0, 1, 0]
print(rockPaperScissors[sigma_r, sigma_c])
# calculating nash equilibria 
equilibrium = rockPaperScissors.support_enumeration()
print(list(equilibrium))

# with random seed 
iterations = 100
np.random.seed(10)
count = rockPaperScissors.fictitious_play(iterations = iterations)
for i, j in count:
    print(i, j)