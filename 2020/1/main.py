import os

def firstproblem(payload):
    lines = {int(x) for x in payload.split('\n') if x}

    for l in lines:
        for l2 in lines:
            if l + l2 == 2020:
                return l * l2

    return -1

def secondproblem(payload):
    lines = {int(x) for x in payload.split('\n') if x}

    for l in lines:
        for l2 in lines:
            for l3 in lines:
                if l + l2 +l3 == 2020:
                    return l * l2 *l3

    return -1


if __name__ == "__main__":
    dir_path = os.path.dirname(os.path.realpath(__file__))

    with open(os.path.join(dir_path, 'input'), 'r') as entry:
        payload = entry.read()

    print(firstproblem(payload))
    print(secondproblem(payload))
