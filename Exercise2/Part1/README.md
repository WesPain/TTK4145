# Mutex and Channel basics

### What is an atomic operation?
> Atomic operations are operations which are indivisible and can’t be interrupted. It’s an operation which a processor can simultaneously read a location and write it in the same bus operation, the operation HAS to be completed.

### What is a semaphore?
> A good metaphor for semaphore is a bouncer at a night club where the semaphore limits the number of threads(processes) using a resource. It appears as a single integer. The semaphore signals the threads telling them if they can use a resource or not.

### What is a mutex?
> Mutually exclusive so in a way a mutex is also a way to limit accessibility to threads. The thread who has the mutex is the only thread which can operate on that said resource at a time. 

### What is the difference between a mutex and a binary semaphore?
> A mutex is a locking mechanism while a semaphore is a signalling mechanism.

### What is a critical section?
> A critical section is a section ensuring that multiple threads doesn’t use the same resources at the same time. The section is protected from unexpected behaviour. 

### What is the difference between race conditions and data races?
> A race condition is when the timing or the order of events affects a programs correctness. While a data race happens when there are two memory accesses in a program which targets the same location, are performed by two concurrent threads and the operation isn’t a reading one, and they are asynchronous.

### List some advantages of using message passing over lock-based synchronization primitives.
> The advantages of message passing over lock-bases sync are that it is easier to build massively parallel hardware, and are more tolerant to message passing programming models.  

### List some advantages of using lock-based synchronization primitives over message passing.
> These can be better in single servers because it could be easier to attain higher performance
