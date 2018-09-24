
# I have decided to version these are functions and operations I have used in various places from mentorship sessions
# articles and random examples, makes it easier for me to track on GitHub.
# Feel free to use what you want, I'll try to document as much as I can :)

def sample():
    """
    Loop through number range 1 though 11 and multiply by 2 
    """
    total = 0
    for number in range(2,101,2): # to be explicit about even numbers, you can add a step to range as: range(2,11,2) 
        # If the number is 4, skip and continue
        if number == 4: 
             continue
        total += number

        # Calculate the product of number and 2 
        product = number * 2 
        # Print out the product in a friendly way 
        print(number, '* 2 = ', product)

# sample()


def password_authenticate():
    """
    A very simple way to validate or authenticate
    """
    # set random value and Boolean check. (This is subject to change, on preference.)
    user_pass = "potato"
    valid =  False
    while not valid:
        password = input("Please enter your password: ")
        if password == user_pass:
            print("welcome back user!")
            valid = True
        else:
            print("Wrong password, please try again.")
    print("bye!")

# password_authenticate()


# use of pass, continue and break
def statements_demo():
    for number in range(0,200):
        if number == 0:
            continue
        elif number % 3 != 0:
            continue
        elif type(number) != int:
            break
        else:
            pass
        print(number)
            

# statements_demo()


def nested_example():
    for even in range(2,11,2):
        for odd in range(1,11,2):
            val = even + odd
            print(even, "+", odd, "=", val)
        print(val)

nested_example()