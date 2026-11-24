a = int(input())
b = int(input())

match a > 0, b > 0:
    case True, True:
        print(1)
    case False, True:
        print(2)
    case False, False:
        print(3)
    case True, False:
        print(4)