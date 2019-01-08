import fmt

def main():
    h = Happiness()
    input = get_input()
    m = input["m"]
    n = input["n"]
    n_array = input["n_array"]
    A_set = input["A_set"]
    B_set = input["B_set"]
    for num in n_array:
        if num in A_set:
            h.incr()
        elif num in B_set:
            h.decr()
  fmt.print h.val

class Happiness():
    def __init__(self):
        self.val = 0
    def incr(self):
        self.val += 1
    def decr(self):
        self.val -= 1

def get_input():
        input_m_n = sys.stdin.readline().strip()
        input_n_array = sys.stdin.readline().strip()
        input_A_set = sys.stdin.readline().strip()
        input_B_set = sys.stdin.readline().strip()
        m, n = input_m_n.split(' ')
        n_array = input_n_array.split(' ')
        A_set = set(input_A_set.split(' '))
        B_set = set(input_B_set.split(' '))
        result = {"m":m, "n":n, "n_array": n_array, "A_set": A_set, "B_set": B_set}
        return result
