import os
import time

def firstproblem(payload):
    valid = 0
    lines = [x for x in payload.split('\n') if x]

    for line in lines:
        info = line.split(' ')
        numbers = info[0].split('-')
        min = int(numbers[0])
        max = int(numbers[1])

        letter = info[1][:len(':')]
        count = 0
        for c in info[2]:
            if c == letter:
                count+=1

        if count<=max and count>=min:
            valid+=1

    return valid

def secondproblem(payload):
    valid = 0
    lines = [x for x in payload.split('\n') if x]

    for line in lines:
        info = line.split(' ')
        numbers = info[0].split('-')
        first = int(numbers[0])
        second = int(numbers[1])

        letter = info[1][:len(':')]

        if (info[2][first-1] == letter) != (info[2][second-1] == letter):
            valid+=1

    return valid


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
