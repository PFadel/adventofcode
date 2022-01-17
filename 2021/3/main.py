import os
import time

FORWARD = 'forward'
DOWN = 'down'
UP = 'up'

def _calculate_most_common(lines):
    bit_len = len(lines[0])
    most_common = ''
    least_common = ''

    appearences = [0 for x in range(0, bit_len)]
    for line in lines:
        for i in range(0, len(line)):
            if line[i] == '1':
                appearences[i]+=1

    for appearence in appearences:
        if appearence >= len(lines)/2:
            most_common = most_common + '1'
            least_common = least_common + '0'
        else:
            most_common = most_common + '0'
            least_common = least_common + '1'

    return most_common, least_common

def firstproblem(input):
    lines = [x for x in input.split('\n') if x != '']
    most_common, least_common = _calculate_most_common(lines)

    gamma_rate = int(most_common, 2)
    epsilon_rate = int(least_common, 2)

    return gamma_rate * epsilon_rate

def secondproblem(input):
    lines = [x for x in input.split('\n') if x != '']

    bit_len = len(lines[0])
    possible = lines.copy()

    for i in range(0, bit_len):
        most_common, least_common = _calculate_most_common(possible)
        for line in lines:
            if line[i] != most_common[i]:
                possible.remove(line)

        lines = possible.copy()
        if len(possible) == 1:
            break

    oxygen_generator = possible[0]

    lines = [x for x in input.split('\n') if x != '']
    possible = lines.copy()
    for i in range(0, bit_len):
        most_common, least_common = _calculate_most_common(possible)
        for line in lines:
            if line[i] != least_common[i]:
                possible.remove(line)

        lines = possible.copy()
        if len(possible) == 1:
            break

    co2_scrubber = possible[0]

    oxygen_generator = int(oxygen_generator, 2)
    co2_scrubber = int(co2_scrubber, 2)

    return oxygen_generator*co2_scrubber

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
