import os
import time

def firstproblem(input):
    input = input.split('\n')[:-1]
    numbers_to_show = input[0]

    boards = [[],[],[]]
    board = 0
    for line in input[2:]:
        if line == '':
            board += 1
            continue
        boards[board].append([x for x in line.split(' ') if x != ''])

    board_lines = len(boards[0])
    board_columns = len(boards[0][0])

    board_map = {
        '1': 'TODO'
    }

    print(board_lines, board_columns)


if __name__ == "__main__":
    dir_path = os.path.dirname(os.path.realpath(__file__))

    with open(os.path.join(dir_path, 'input'), 'r') as entry:
        payload = entry.read()

    start = time.time()
    r1 = firstproblem(payload)
    runtime_1 = time.time() - start

    # start = time.time()
    # r2 = secondproblem(payload)
    # runtime_2 = time.time() - start

    print(f'Result: {r1} , Runtime: {runtime_1} seconds')
    # print(f'Result: {r2} , Runtime: {runtime_2} seconds')
