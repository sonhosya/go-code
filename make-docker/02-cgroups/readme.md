# Linux Cgroups
Linux Cgroups（Control Groups）提供了对一组进程及将来子进程的资源限制、控制和统计的能力，这些资源包括CPU、内存、存储、网络等。通过Cgroups，可以方便地限制某个进程的资源占用，并且可以实时地监控进程的监控和统计信息。
## Cgroups中的3个组件
### cgroup
cgroup是对进程分组管理的一种机制，一个cgroup包含一组进程，并可以在这个cgroup上增加Linux subsystem的各种参数配置，将一组进程和一组subsystem的系统参数关联起来。
### subsystem
- 是一组资源控制的模块，一般包含如下几项
- blkio设置对块设备（比如硬盘）输入输出的访问控制。
- cpu设置cgroup中进程的CPU被调度的策略。
- cpuacct可以统计cgroup中进程的CPU占用。
- cpuset 在多核机器上设置cgroup中进程可以使用的CPU和内存（此处内存仅使用于NUMA架构）。
- devices控制cgroup中进程对设备的访问。
- freezer用于挂起（suspend）和恢复（resume） cgroup中的进程。
- memory用于控制cgroup中进程的内存占用。
- net_cls用于将cgroup中进程产生的网络包分类，以便Linux的tc（traffic controller）可以根据分类区分出来自某个cgroup的包并做限流或监控。
- net_prio设置cgroup中进程产生的网络流量的优先级。
### hierarchy 
hierarchy的功能是把一组cgroup串成一个树状的结构，一个这样的树便是一个hierarchy，通过这种树状结构，Cgroups可以做到继承。比如，系统对一组定时的任务进程通过cgroup1限制了CPU的使用率，然后其中有一个定时dump日志的进程还需要限制磁盘IO，为了避免限制了磁盘IO之后影响到其他进程，就可以创建cgroup2，使其继承于cgroup1并限制磁盘的IO，这样cgroup2便继承了cgroup1中对CPU使用率的限制，并且增加了磁盘IO的限制而不影响到cgroup1中的其他进程。
## 三个组件相互的关系
三个组件相互的关系通过上面组件的描述，Cgroups是凭借这三个组件的相互协作实现的。系统在创建了新的hierarchy之后，系统中所有的进程都会加入这个hierarchy的cgroup根节点hierarchy中创建的cgroup都是这个cgroup根节点的子节点。
1. 一个subsystem只能附加到一个hierarchy上面。
2. 一个hierarchy可以附加多个subsystem。
3. 一个进程可以作为多个cgroup的成员，但是这些cgroup必须在不同的hierarchy中。
4. 一个进程fork出子进程时，子进程是和父进程在同一个cgroup中的，也可以根据需要将其移动到其他cgroup中。