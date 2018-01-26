# Python 3.3.3 and 2.7.6
# python fo.py

from threading import Thread
from threading import Lock

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
i = 0
a_mutex=Lock()

def incrementingFunction():
    global i
    global a_mutex
    a_mutex.acquire()
    try:
        for j in range(1000000): # critical section the same as the C module
            i+=1
    finally:
        a_mutex.release() 
                

def decrementingFunction():
    global i
    global a_mutex
    a_mutex.acquire()
    try:
        for j in range(1000000): # critical section the same as the C module
            i-=1
    finally:
        a_mutex.release() 



def main():
    global i

    incrementing = Thread(target = incrementingFunction, args = (),)
    decrementing = Thread(target = decrementingFunction, args = (),)
    
    # TODO: Start both threads
    incrementing.start()
    decrementing.start()
    
    incrementing.join()
    decrementing.join()
    
    print("The magic number is %d" % (i))


main()
