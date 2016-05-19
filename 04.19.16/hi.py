class Dog(object):
    def speak(self):
        return "Arf! Arf!"

class Cat(object):
    def speak(self):
        return "Meow!"

def makeNoise(thing):
    print(thing.speak())

if __name__ == '__main__':
    dog = Dog()
    cat = Cat()
    makeNoise(dog)
    makeNoise(cat)
