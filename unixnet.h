#ifndef __UNIXNET_H
#define __UNIXNET_H

#include <sys/types.h>
#include <sys/socket.h>
#include <sys/time.h>
#include <time.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <errno.h>
#include <fcntl.h>
#include <netdb.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/un.h>

#include <sys/select.h>
#include <poll.h>
#include <string.h>
#include <sys/ioctl.h>
#include <pthread.h>

#include <sys/ipc.h>
#include <sys/sem.h>
#include <sys/shm.h>
#include <sys/msg.h>

#define SERVER_PORT 9999
#define LISTENQ 10
#define MAXLINE 1024

#endif __UNIXNET_H
