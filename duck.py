class Duck:
    def __init__(self):
        pass

    def Quack(self):
        return "quack"

class Cat:
    def __init__(self):
        pass

    def Quack(self):
        return "meow"

class Dog:
    def __init__(self):
        pass

    def Bark(self):
        return "bark"

def MakeQuack(x):
    print(x.Quack())

duck = Duck()
cat = Cat()
dog = Dog()

MakeQuack(duck)
MakeQuack(cat)
MakeQuack(dog) # Error because dogs do not Quack()
