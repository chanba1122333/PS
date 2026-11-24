a = int(input())

match a:
    case _ if a >= 90:
        print("A")
    case _ if a >= 80:
        print("B")
    case _ if a >= 70:
        print("C")
    case _ if a >= 60:
        print("D")
    case _:
        print("F")