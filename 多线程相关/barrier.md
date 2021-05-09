## pthread_barrier_init指定等待的cout

## pthread_barrier_wait

通过返回值是否为PTHREAD_BARRIER_SERIAL_THREAD来选择master

## 使用场景

比如超大列表排序，可以分为多段使用多个线程进行排序，用barrier来等待所有子线程排序完，然后主线程进行merge。