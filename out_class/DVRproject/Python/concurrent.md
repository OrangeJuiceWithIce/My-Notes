# 背景
在科研项目中，常常有渲染大量图片的需求

# 概念
[阮一峰:进程和线程的一个简单解释](https://www.bing.com/search?q=%E9%98%AE%E4%B8%80%E5%B3%B0&form=ANNTH1&refig=67b9b10d9c934236a12c1107f26e5ea5&pc=LCTS&adppc=EdgeStart)
1.进程拥有独立的资源  
2.进程下可包含多个线程，线程之间共享资源  
3."互斥锁"(mutual exclusion),防止多个线程同时使用同一份资源  
4.信息量"Semaphore"，规定了一份资源最多可以被几个线程同时使用  

# Thread
## Threading.Thread()
```python
import threading

def worker(name,delay):
    print(f"Worker {name} is running")
    time.sleep(delay)
    print(f"Worker {name} is done")

thread1=threading.Thread(target=worker,arges=("Thread-1",2))
thread2=threading.Thread(target=worker,arges=("Thread-2",4),daemon=True)

# 运行线程
thread1.start()  
thread2.start()
# 等待线程结束
thread1.join()
```

**target:** 线程函数,默认作为thread.run()方法
**name:** 线程名称(default:Thread-N)
**args:** 传递给线程函数的参数
**kwargs:** 传递给线程函数的关键字参数
**daemon:** 是否为守护线程，当主线程结束时，守护线程也会结束
## 继承Thread类并重写run()方法
```python
import threading
class MyThread(threading.Thread):
    def __init__(self,name):
        super().__init__()
        self.name=name

    def run(self):
        print(f"Thread {self.name} is running")
        time.sleep(2)
        print(f"Thread {self.name} is done")
```
## 线程间通信
### Queue
put(item)：将 item 放入队列
get()：从队列中取出并返回一个 item
join()，可以等待队列中的所有元素都被取出
task_done()，每个任务完成后需要显式调用 task_done() 来标记任务完成

## Threading.Lock()
原子操作：一旦开始执行，不能中断
```python
import threading

counter = 0
lock = threading.Lock()

def increment():
    global counter
    for _ in range(100000):
        with lock:  # 加锁
            counter += 1  # 原子操作

thread1 = threading.Thread(target=increment)
thread2 = threading.Thread(target=increment)

thread1.start()
thread2.start()

thread1.join()
thread2.join()

print(f"Final counter value: {counter}")
```
with lock实则进行了lock.acquire()和lock.release()操作
## Threading.RLock()
**情景**：
如果function A加了锁，而它调用的function B(或者递归调用自己)也要加锁，传统的Lock()会造成死锁，而RLock()可以解决这个问题
## Threading.Semaphore()
**情景**：
限制同时访问共享资源的线程数量（例如连接池、线程池）
## Threading.Condition()
线程等待某个条件被满足，条件满足时通知线程继续执行
```python
import threading
# 自带一把RLock锁
condition = threading.Condition()
import threading
import time
import random

# 共享资源
queue = []
max_size = 5
condition = threading.Condition()

def producer():
    while True:
        with condition:
            if len(queue) == max_size:
                print("Queue is full, producer is waiting")
                condition.wait()  # 等待消费者消费，此时释放锁
            item = random.randint(1, 100)
            queue.append(item)
            print(f"Produced {item}, queue: {queue}")
            condition.notify()  # 通知消费者，此时重新获得锁
        time.sleep(random.random())

def consumer():
    while True:
        with condition:
            if not queue:
                print("Queue is empty, consumer is waiting")
                condition.wait()  # 等待生产者生产,此时释放锁
            item = queue.pop(0)
            print(f"Consumed {item}, queue: {queue}")
            condition.notify()  # 通知生产者，此时重新获得锁
        time.sleep(random.random())

producer_thread = threading.Thread(target=producer)
consumer_thread = threading.Thread(target=consumer)

producer_thread.start()
consumer_thread.start()

producer_thread.join()
consumer_thread.join()
```

# Process
## multiprocessing.Process()
```python
class multiprocess.Process(
    group=None, 
    target=None,
    name=None,  # 为None时，默认为Process-N
    args=(),
    kwargs={},
    daemon=None  
)
```
参数和使用都与Thread类差不多
方法：
.start()：启动进程
.join()：等待进程结束
.terminal()：强制终止进程
.is_alive()：判断进程是否存活

### 查看进程信息：
```python
import multiprocessing

process_name=multiprocessing.current_process().name
process_pid = multiprocessing.current_process().pid
parent_pid= multiprocessing.parent_process().pid
```
父进程:子进程的父进程就是创建它的进程

### 进程间通信
#### Queue
put(item)：将 item 放入队列
get()：从队列中取出并返回一个 item
empty()：判断队列是否为空
full()：判断队列是否已满
qsize()：返回队列中的元素数量（近似值）
##### JoinableQueue
支持.join()方法和.task_done()方法
#### Pipe
Pipe(duplex=True)返回两个连接对象，默认情况下，两个连接对象可以双向通信
Pipe(duplex=False)：创建单向通信的管道

send(obj)：发送数据
recv()：接收数据
close()：关闭管道
```python
from multiprocessing import Process, Pipe
import time

def worker(conn):
    while True:
        msg = conn.recv()  # 接收数据
        if msg == "exit":  # 结束信号
            break
        print(f"Worker received: {msg}")
        conn.send(f"Echo: {msg}")  # 发送数据

if __name__ == '__main__':
    parent_conn, child_conn = Pipe()  # 创建管道

    # 创建进程
    p = Process(target=worker, args=(child_conn,))

    # 启动进程
    p.start()

    # 主进程发送数据
    for msg in ["Hello", "World", "exit"]:
        print(f"Main sending: {msg}")
        parent_conn.send(msg)
        if msg != "exit":
            response = parent_conn.recv()  # 接收数据
            print(f"Main received: {response}")

    # 等待子进程完成
    p.join()
```
### 进程池
情景：有时任务多，有时任务少，不应该创建任务数量的进程（一方面数量过多的进程无法同时执行，另一方面创建和删除进程存在开销）
定义一个pool，在里面放上固定数量的进程，如果有任务，就取一个池中的进程来处理任务
```python
multiprocessing.Pool(
    processes=None, # 进程池中进程数量，如果为None，默认使用 os.cpu_count()，即CPU核数。最好小于等于CPU核数
    initializer=None, # 如果不为None，那么每个进程在开始时，都会调用initializer(*initargs)
    initargs=(),     # 初始化时的输入参数
    maxtasksperchild=None # 子进程完成maxtasksperchild个任务后会重启，或者直到主进程结束，有助于释放内存
)
```
方法:
apply(task函数,args参数),发布任务，主线程阻塞
apply_async(task函数,args参数,callback(res),error_callback(error)),发布任务，主线程不阻塞，返回一个对象,用.get()方法获取结果
close()：关闭进程池，不再接受新的任务
join()：等待所有进程结束
>close()必须在join()之前调用

# concurrent.futures
## Executor
ThreadPoolExecutor(max_workers=None)
ProcessPoolExecutor(max_workers=None)
方法:
.submit(task函数，task参数)
.map(task函数，task参数序列,timeout=None,chunksize=1)

## Future
Future实例应该由Executor.submit()创建，由执行器来管理

as_completed 函数接受一个 Future 对象的集合
```python
import concurrent.futures
import time
import random

# 模拟一个耗时任务
def task(n):
    print(f"Task {n} started")
    time.sleep(random.uniform(1, 3))  # 模拟不同时间的任务
    return f"Task {n} completed"

if __name__ == "__main__":
    # 创建一个线程池
    with concurrent.futures.ThreadPoolExecutor(max_workers=5) as executor:
        # 提交多个任务
        futures = [executor.submit(task, i) for i in range(5)]
        
        # 使用 as_completed 动态处理每个完成的任务
        for future in concurrent.futures.as_completed(futures):
            try:
                # 获取任务结果
                result = future.result()
                print(result)
            except Exception as e:
                print(f"Task failed with exception: {e}")
```