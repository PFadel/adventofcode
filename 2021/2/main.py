import os
import time

FORWARD = 'forward'
DOWN = 'down'
UP = 'up' 

def firstproblem(input):
    lines = input.split('\n')
    horizontal = 0
    depth = 0

    for line in lines:
        command = line.split(' ')

        if command[0] == FORWARD:
            horizontal += int(command[1])
        elif command[0] == DOWN:
            depth += int(command[1])
        elif command[0] == UP:
            depth -= int(command[1])
    
    return horizontal*depth

def secondproblem(input):
    lines = input.split('\n')
    horizontal = 0
    depth = 0
    aim = 0

    for line in lines:
        command = line.split(' ')

        if command[0] == FORWARD:
            horizontal += int(command[1])
            depth += aim*int(command[1])
        elif command[0] == DOWN:
            aim += int(command[1])
        elif command[0] == UP:
            aim -= int(command[1])
    
    return horizontal*depth

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
