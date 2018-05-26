#include "common.h"

int main()
{
  int listen_fd,conn_fd;
  struct sockaddr_in cli_addr,serv_addr;
  int ret,flags;
  socklen_t cli_len;
  pid_t chld_pid;

  listen_fd = socket(AF_INET,SOCK_STREAM,0);
  if(listen_fd == -1){
    perror("create listen fd");
    exit(1);
  }

  bzero(&serv_addr,sizeof(serv_addr));
  serv_addr.sin_addr.s_addr=htonl(INADDR_ANY);
  serv_addr.sin_family=AF_INET;
  serv_addr.sin_port=htons(SERVER_PORT);

  if((flags=fcntl(listen_fd,F_GETFL,0))<0)
    perror("F_GETEL error");

  flags&=~O_NONBLOCK;

  if(fcntl(listen_fd,F_SETFL,flags)<0)
    perror("F_SETFL error");

  ret=bind(listen_fd,(struct sockaddr*)&serv_addr,sizeof(serv_addr));
  if(ret<0){
    perror("bind server port");
    exit(1);
  }

  listen(listen_fd,LISTENQ);

  while(1)
    {
      cli_len=sizeof(cli_addr);
      printf("The server is waiting for accepting new connection\n");
      conn_fd=accept(listen_fd,(struct sockaddr*)&cli_addr,&cli_len);

      printf("The server create a new connection\n");
      if((chld_pid=fork())==0){
	close(listen_fd);
	printf("child thread doing the service");
	close(conn_fd);
	exit(0);
      }
      close(conn_fd);
    }
}
