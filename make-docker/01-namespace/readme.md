1. linux namespace是kernel的一个功能，可以隔离一系列的系统资源，比如PID（Process ID）、User ID、Network等。
2. Unix 中有一个名为 chroot 的系统调用（通过修改根目录把用户 jail 到一个特定目录下），chroot 提供了一种简单的隔离模式：chroot 内部的文件系统无法访问外部的内容。Linux Namespace 在此基础上，提供了对 UTS（UNIX Time-sharing System）、IPC、Mount、PID、Network、User 等六种隔离机制。

|   Namespace | Description | System Call   |
| ----------- | ----------- | ----------- |
|   UTS |   hostname、domainname   |    CLONE_NEWUTS    |
|   IPC |   隔离System V IPC和POSIX message queues。每一个IPC Namespace都有自己的System V IPC和POSIX message queue。    |   CLONE_NEWIPC    |
|   PID |   隔离进程ID的。同样一个进程在不同的PID Namespace 里可以拥有不同的PIDText    |   CLONE_NEWPID    |
|   Mount   |   隔离各个进程看到的挂载点视图。在不同Namespace的进程中，看到的文件系统层次是不一样的。在Mount Namespace中调用mount（）和umount（）仅仅只会影响当前Namespace内的文件系统，而对全局的文件系统是没有影响的。    |   CLONE_NEWNS    |
|   User    |   隔离用户的用户组ID。也就是说，一个进程的User ID 和Group ID在User Namespace内外可以是不同的    |   CLONE_NEWUSER    |
|   NETWORK |   隔离网络设备、IP地址端口等网络栈的Namespace。Network Namespace可以让每个容器拥有自己独立的（虚拟的）网络设备，而且容器内的应用可以绑定到自己的端口，每个Namespace内的端口都不会互相冲突。在宿主机上搭建网桥后，就能很方便地实现容器之间的通信，而且不同容器上的应用可以使用相同的端口。    |   CLONE_NEWNET    |

