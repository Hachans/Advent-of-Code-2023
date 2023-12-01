"""Advent of code day1 solution"""


def read_input(path: str) -> list[str]:
    """Read lines from input file.

    Keyword arguments:
    path -- path to input.* file
    """
    ll = []
    with open(path, "r", encoding="utf-8") as infile:
        for entry in infile:
            ll.append(entry[:-1])
    return ll


def init_dict() -> dict:
    """Initiates value dictionary"""
    values = {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }
    return values


def extract_first(entry: str, vd: dict) -> str:
    """Extracts the first digit or string value in a line

    Keyword arguments:
    entry -- line from input.* file
    vd -- value dictionary
    """
    for i, c in enumerate(entry):
        if c.isdigit():
            return c
        for k, v in vd.items():
            if k in entry[: i + 1]:
                return v
    return ""


def extract_last(entry: str, vd: dict) -> str:
    """Extracts the last digit or string value in a line

    Keyword arguments:
    entry -- line from input.* file
    vd -- value dictionary
    """
    for i, c in enumerate(reversed(entry)):
        if c.isdigit():
            return c
        for k, v in vd.items():
            if k in entry[len(entry) - i - 1 :]:
                return v
    return ""


if __name__ == "__main__":
    lines = read_input("../input.txt")
    value_dict = init_dict()

    TOTAL_SUM = 0
    for line in lines:
        first = extract_first(line, value_dict)
        last = extract_last(line, value_dict)
        num_sum = int(first + last)
        print(f"line: {line}, first: {first}, last: {last}")
        TOTAL_SUM += num_sum
    print(f"Total sum: {TOTAL_SUM}")
