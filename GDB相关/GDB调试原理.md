## 简单思考

对于调试，我们能想到的应该有：

- 调试进程以某种方式可以获取被调试进程的堆栈以及寄存器的信息（查看局部变量、全局变量等操作）
- 调试进程可以影响被调试进程的运行，并且可以等待被调试进程被暂停（打断点等操作）

怎么做到上述两点的就是深究GDB调试的原理了。

## 进程状态

首先，linux进程的状态分为以下几种：

1. D (TASK_UNINTERRUPTIBLE)，不可中断的睡眠状态。
2. R (TASK_RUNNING)，进程执行中。
3. S (TASK_INTERRUPTIBLE)，可中断的睡眠状态。
4. T (TASK_STOPPED)，暂停状态。
5. t (TASK_TRACED)，进程被追踪。
6. w (TASK_PAGING)，进程调页中，2.6以上版本的内核中已经被移除。
7. X (TASK_DEAD – EXIT_DEAD)，退出状态，进程即将被销毁。
8. Z (TASK_DEAD – EXIT_ZOMBIE)，退出状态，进程成为僵尸进程。

被调试进程处于状态5，即**被追踪状态**

## ptrace系统调用

GDB正是通过[ptrace系统调用](https://man7.org/linux/man-pages/man2/ptrace.2.html)来实现调式相关的各种操作的。ptrace原型为：

```c++
long ptrace(enum __ptrace_request request, pid_t pid, void *addr, void *data);
```

有多种request分别对应不同的功能，如：

- PTRACE_TRACEME，表明该进程将被父进程追踪
- PTRACE_ATTACH，attach到指定进程即追踪指定进程
- PTRACE_SINGLESTEP，单步调试

调试大体的原理就是，当被调试程序和调试程序**建立调式关系**后，被追踪程序在接收到信号时（SIGKILL除外）都会被停止，调试程序则会通过wait系统调用获取到被调试程序停止的原因，据此可以判断是否是由于断点引起的程序停止，如果是断点引起的则可以使用多种ptrace请求来检查和编辑被追踪线程。下面看下如何建立调试关系，以及断点原理。

### 建立调试关系

- fork + execve执行被测试程序，子进程在执行execve之前调用ptrace(PTRACE_TRACEME)，从而建立与父进程（debuggerd)的跟踪关系，即dgb **.exe的原理
- attach到被测试程序，调用ptrace(PTRACE_ATTACH, pid...)

### 断点原理

1. 打断点。先将该位置原来的指令保存，然后向该位置写入int 3（*具体是插入int 3还是保存该位置的指令不是很确定，确定的是要实现从断点恢复执行后，下一个指令应该是正确的*）。当执行到int 3时触发软中断，内核给子进程发出SIGTRAP信号，因为子进程已经建立了调试关系，信号会传递给父进程。
2. 断点命中判断：gdb把所有断点位置都存放在一个链表中，命中判定即把被调试程序当前停止的位置和链表中的断点位置进行比较，看是断点产生的信号，还是无关信号。
3. 后一步操作，使用ptrace获取子进程各种状态，或释放这次断点

### pstack

pstack是基于GDB实现的命令行工具，可以显示指定进程每个线程的堆栈快照，便于排查程序异常和性能评估

**另外需要注意一点，只有保留了 debug symbols 的程序才可以 pstack**

