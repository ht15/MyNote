## 自旋锁为什么要关闭抢占（preemption）

若P_A持有spin_lock但是时间片结束，P_B的时间片期间可能一直在等待时间片，也就是说一直在空跑；更甚的是P_B有可能是中断处理程序，那么就会发生死锁：中断处理程序一直在spin。

## 可用于多个进程之间

Pthread_spin_init(pthread_spin_t *lock, int pshared) 第二个参数指定能不能被多个进程共享