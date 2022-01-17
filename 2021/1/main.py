import os
import time

def firstproblem(input):
    lines = [int(x) for x in input.split('\n')]
    count = 0
    last = lines[0]
    for i in range(0, len(lines)):
        if lines[i] > last:
            count += 1
        last = lines[i]
    return count 

def secondproblem(input):
    lines = [int(x) for x in input.split('\n')]
    count = 0
    sequences = []
    for i in range(0, len(lines)-2):
        sequences.append((lines[i], lines[i+1], lines[i+2]))

    last_sum = sum_sequence(sequences[0])
    for sequence in sequences:
        sum = sum_sequence(sequence)
        if sum > last_sum:
            count += 1
        last_sum = sum
    return count 

def sum_sequence(sequence):
    return sequence[0]+sequence[1]+sequence[2]

if __name__ == "__main__":
    dir_path = os.path.dirname(os.path.realpath(__file__))

    with open(os.path.join(dir_path, 'input'), 'r') as entry:
        payload = entry.read()

    start = time.time()
    r1 = firstproblem(payload)
    runtime_1 = time.time() - start

    start = time.time()
    r2 = secondproblem(payload)
    runtime_2 = time.time() - start

    print(f'Result: {r1} , Runtime: {runtime_1} seconds')
    print(f'Result: {r2} , Runtime: {runtime_2} seconds')
