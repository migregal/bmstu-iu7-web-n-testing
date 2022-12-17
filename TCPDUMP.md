## Результат перехвата сетевого трафика посредством tcpdump

### Запуск

```bash
$ sudo tcpdump -i lo0
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on lo0, link-type NULL (BSD loopback), capture size 262144 bytes
```

### Регистрация пользователя

```bash
01:23:07.007161 IP6 localhost.50915 > localhost.scp-config: Flags [S], seq 94495319, win 65535, options [mss 16324,nop,wscale 6,nop,nop,TS val 1303128799 ecr 0,sackOK,eol], length 0
01:23:07.007260 IP6 localhost.scp-config > localhost.50915: Flags [S.], seq 2412201015, ack 94495320, win 65535, options [mss 16324,nop,wscale 6,nop,nop,TS val 468124007 ecr 1303128799,sackOK,eol], length 0
01:23:07.007273 IP6 localhost.50915 > localhost.scp-config: Flags [.], ack 1, win 6371, options [nop,nop,TS val 1303128799 ecr 468124007], length 0
01:23:07.007285 IP6 localhost.scp-config > localhost.50915: Flags [.], ack 1, win 6371, options [nop,nop,TS val 468124007 ecr 1303128799], length 0
01:23:07.007364 IP6 localhost.50915 > localhost.scp-config: Flags [P.], seq 1:262, ack 1, win 6371, options [nop,nop,TS val 1303128799 ecr 468124007], length 261
01:23:07.007379 IP6 localhost.scp-config > localhost.50915: Flags [.], ack 262, win 6367, options [nop,nop,TS val 468124007 ecr 1303128799], length 0
01:23:07.023106 IP6 localhost.scp-config > localhost.50915: Flags [P.], seq 1:117, ack 262, win 6367, options [nop,nop,TS val 468124023 ecr 1303128799], length 116
01:23:07.023149 IP6 localhost.50915 > localhost.scp-config: Flags [.], ack 117, win 6370, options [nop,nop,TS val 1303128815 ecr 468124023], length 0
01:23:07.023278 IP6 localhost.50915 > localhost.scp-config: Flags [P.], seq 262:516, ack 117, win 6370, options [nop,nop,TS val 1303128815 ecr 468124023], length 254
01:23:07.023298 IP6 localhost.scp-config > localhost.50915: Flags [.], ack 516, win 6363, options [nop,nop,TS val 468124023 ecr 1303128815], length 0
01:23:07.032526 IP6 localhost.scp-config > localhost.50915: Flags [P.], seq 117:765, ack 516, win 6363, options [nop,nop,TS val 468124032 ecr 1303128815], length 648
01:23:07.032563 IP6 localhost.50915 > localhost.scp-config: Flags [.], ack 765, win 6359, options [nop,nop,TS val 1303128824 ecr 468124032], length 0
01:23:07.032722 IP6 localhost.50915 > localhost.scp-config: Flags [F.], seq 516, ack 765, win 6359, options [nop,nop,TS val 1303128824 ecr 468124032], length 0
01:23:07.032784 IP6 localhost.scp-config > localhost.50915: Flags [.], ack 517, win 6363, options [nop,nop,TS val 468124032 ecr 1303128824], length 0
01:23:07.034270 IP6 localhost.scp-config > localhost.50915: Flags [F.], seq 765, ack 517, win 6363, options [nop,nop,TS val 468124033 ecr 1303128824], length 0
01:23:07.034314 IP6 localhost.50915 > localhost.scp-config: Flags [.], ack 766, win 6359, options [nop,nop,TS val 1303128825 ecr 468124033], length 0
```

### Авторизация пользователя
```bash
01:23:35.342797 IP6 localhost.50922 > localhost.scp-config: Flags [S], seq 2751566466, win 65535, options [mss 16324,nop,wscale 6,nop,nop,TS val 600925846 ecr 0,sackOK,eol], length 0
01:23:35.342960 IP6 localhost.scp-config > localhost.50922: Flags [S.], seq 1701920377, ack 2751566467, win 65535, options [mss 16324,nop,wscale 6,nop,nop,TS val 1312847669 ecr 600925846,sackOK,eol], length 0
01:23:35.342972 IP6 localhost.50922 > localhost.scp-config: Flags [.], ack 1, win 6371, options [nop,nop,TS val 600925846 ecr 1312847669], length 0
01:23:35.342985 IP6 localhost.scp-config > localhost.50922: Flags [.], ack 1, win 6371, options [nop,nop,TS val 1312847669 ecr 600925846], length 0
01:23:35.343060 IP6 localhost.50922 > localhost.scp-config: Flags [P.], seq 1:255, ack 1, win 6371, options [nop,nop,TS val 600925846 ecr 1312847669], length 254
01:23:35.343084 IP6 localhost.scp-config > localhost.50922: Flags [.], ack 255, win 6367, options [nop,nop,TS val 1312847669 ecr 600925846], length 0
01:23:35.354859 IP6 localhost.scp-config > localhost.50922: Flags [P.], seq 1:649, ack 255, win 6367, options [nop,nop,TS val 1312847681 ecr 600925846], length 648
01:23:35.354886 IP6 localhost.50922 > localhost.scp-config: Flags [.], ack 649, win 6361, options [nop,nop,TS val 600925858 ecr 1312847681], length 0
01:23:35.354989 IP6 localhost.50922 > localhost.scp-config: Flags [F.], seq 255, ack 649, win 6361, options [nop,nop,TS val 600925858 ecr 1312847681], length 0
01:23:35.355006 IP6 localhost.scp-config > localhost.50922: Flags [.], ack 256, win 6367, options [nop,nop,TS val 1312847681 ecr 600925858], length 0
01:23:35.357090 IP6 localhost.scp-config > localhost.50922: Flags [F.], seq 649, ack 256, win 6367, options [nop,nop,TS val 1312847683 ecr 600925858], length 0
01:23:35.357127 IP6 localhost.50922 > localhost.scp-config: Flags [.], ack 650, win 6361, options [nop,nop,TS val 600925860 ecr 1312847683], length 0
```
### Загрузка конфигурации нейронной сети

```bash
01:24:17.994999 IP6 localhost.50942 > localhost.scp-config: Flags [S], seq 200710528, win 65535, options [mss 16324,nop,wscale 6,nop,nop,TS val 3439757629 ecr 0,sackOK,eol], length 0
01:24:17.995170 IP6 localhost.scp-config > localhost.50942: Flags [S.], seq 4149584847, ack 200710529, win 65535, options [mss 16324,nop,wscale 6,nop,nop,TS val 3371621660 ecr 3439757629,sackOK,eol], length 0
01:24:17.995190 IP6 localhost.50942 > localhost.scp-config: Flags [.], ack 1, win 6371, options [nop,nop,TS val 3439757629 ecr 3371621660], length 0
01:24:17.995202 IP6 localhost.scp-config > localhost.50942: Flags [.], ack 1, win 6371, options [nop,nop,TS val 3371621660 ecr 3439757629], length 0
01:24:17.995293 IP6 localhost.50942 > localhost.scp-config: Flags [P.], seq 1:736, ack 1, win 6371, options [nop,nop,TS val 3439757629 ecr 3371621660], length 735
01:24:17.995311 IP6 localhost.scp-config > localhost.50942: Flags [.], ack 736, win 6360, options [nop,nop,TS val 3371621660 ecr 3439757629], length 0
01:24:18.004778 IP6 localhost.50942 > localhost.scp-config: Flags [P.], seq 736:1695, ack 1, win 6371, options [nop,nop,TS val 3439757638 ecr 3371621660], length 959
01:24:18.004844 IP6 localhost.scp-config > localhost.50942: Flags [.], ack 1695, win 6345, options [nop,nop,TS val 3371621669 ecr 3439757638], length 0
01:24:18.037648 IP6 localhost.scp-config > localhost.50942: Flags [P.], seq 1:978, ack 1695, win 6345, options [nop,nop,TS val 3371621701 ecr 3439757638], length 977
01:24:18.037697 IP6 localhost.50942 > localhost.scp-config: Flags [.], ack 978, win 6356, options [nop,nop,TS val 3439757670 ecr 3371621701], length 0
01:24:18.037920 IP6 localhost.50942 > localhost.scp-config: Flags [F.], seq 1695, ack 978, win 6356, options [nop,nop,TS val 3439757670 ecr 3371621701], length 0
01:24:18.037945 IP6 localhost.scp-config > localhost.50942: Flags [.], ack 1696, win 6345, options [nop,nop,TS val 3371621701 ecr 3439757670], length 0
01:24:18.040107 IP6 localhost.scp-config > localhost.50942: Flags [F.], seq 978, ack 1696, win 6345, options [nop,nop,TS val 3371621703 ecr 3439757670], length 0
01:24:18.040150 IP6 localhost.50942 > localhost.scp-config: Flags [.], ack 979, win 6356, options [nop,nop,TS val 3439757672 ecr 3371621703], length 0
```

### Общая статистика

```bash
42 packets captured
42 packets received by filter
0 packets dropped by kernel
```
