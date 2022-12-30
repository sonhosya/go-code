linux namespace是kernel的一个功能，可以隔离一系列的系统资源，比如PID（Process ID）、User ID、Network等。
Unix 中有一个名为 chroot 的系统调用（通过修改根目录把用户 jail 到一个特定目录下），chroot 提供了一种简单的隔离模式：chroot 内部的文件系统无法访问外部的内容。Linux Namespace 在此基础上，提供了对 UTS（UNIX Time-sharing System）、IPC、Mount、PID、Network、User 等六种隔离机制。
|   Namespace | Description |
| ----------- | ----------- |
|   UTS |   hostname、NIS(Network Information) domai nname   |
|   IPC |   Text    |