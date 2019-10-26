def convert(number):
    computed_value = ''

    if number%3 == 0:
        computed_value += "Pling"

    if number%5 == 0:
        computed_value += "Plang"

    if number%7 == 0:
        computed_value += "Plong"

    if len(computed_value) == 0:
        return str(number)

    return computed_value
