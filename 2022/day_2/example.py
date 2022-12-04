def read_input(input_path: str) -> list[list[str]]:
    with open(input_path, 'r') as input_file:
        rounds = [line.rstrip().split() for line in input_file.readlines()]
    return rounds


def first_part(rounds_input_list: list[list[str]]) -> int:
    my_sign_score_dictionary = {'X': 1, 'Y': 2, 'Z': 3}
    my_defeater_dictionary = {'X': 'B', 'Y': 'C', 'Z': 'A'}
    draw_dictionary = {'X': 'A', 'Y': 'B', 'Z': 'C'}

    draw_bonus = 3
    winning_bonus = 6

    total_score = 0
    for current_round in rounds_input_list:
        opponent_sign = current_round[0]
        my_sign = current_round[1]
        current_round_score = my_sign_score_dictionary[my_sign]
        if opponent_sign == draw_dictionary[my_sign]:
            current_round_score += draw_bonus
        elif opponent_sign != my_defeater_dictionary[my_sign]:
            current_round_score += winning_bonus
        total_score += current_round_score
        print(opponent_sign, "vs", my_sign, "=", current_round_score)
    return total_score


def second_part(rounds_input_list: list[list[str]]):
    opponent_sign_score_dictionary = {'A': 1, 'B': 2, 'C': 3}
    scores_list = [1, 2, 3]

    draw_bonus = 3
    winning_bonus = 6

    total_score = 0
    for current_round in rounds_input_list:
        opponent_sign = current_round[0]
        my_indication = current_round[1]
        current_round_score = 0
        if my_indication == 'Y':  # Draw
            current_round_score += draw_bonus
            current_round_score += opponent_sign_score_dictionary[opponent_sign]
        elif my_indication == 'Z':
            current_round_score += winning_bonus
            my_sign_score_index = opponent_sign_score_dictionary[opponent_sign] % 3
            current_round_score += scores_list[my_sign_score_index]
        else:
            my_sign_score_index = (opponent_sign_score_dictionary[opponent_sign] + 1) % 3
            current_round_score += scores_list[my_sign_score_index]
        total_score += current_round_score
    return total_score


def main():
    rounds_input_list = read_input('input.txt')
    first_part(rounds_input_list)
    # print(f'Result for the first part : {first_part(rounds_input_list)}')
    # print(f'Result for the second part : {second_part(rounds_input_list)}')


if __name__ == '__main__':
    main()