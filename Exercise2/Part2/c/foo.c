#include <pthread.h>
#include <stdio.h>

int i = 0;


pthread_mutex_t a_mutex = PTHREAD_MUTEX_INITIALIZER; //initialiaztion of mutex

// Note the return type: void*
void* incrementingThreadFunction(){
    pthread_mutex_lock(&a_mutex); //declaration of mutex
    for (int j = 0; j < 1000000; j++) {
	// TODO: sync access to i
	i++;
    }
    pthread_mutex_unlock(&a_mutex);
    return NULL;
    
}

void* decrementingThreadFunction(){
    pthread_mutex_lock(&a_mutex);
    for (int j = 0; j < 1000003; j++) {
	// TODO: sync access to i
	i--;
    }
    pthread_mutex_unlock(&a_mutex);
    return NULL;
}


int main(){
    pthread_t incrementingThread, decrementingThread;
    
    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);
    
    printf("The magic number is: %d\n", i);
    return 0;
}
