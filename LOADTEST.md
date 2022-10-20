## Загрузка страницы документации
### Без балансировки

```bash
$ k6 run docs-script.js --env HOST="https://..."

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: docs-script.js
     output: -

  scenarios: (100.00%) 1 scenario, 300 max VUs, 1m0s max duration (incl. graceful stop):
           * checkSwggerDocs: 15000.00 iterations/s for 1m0s (maxVUs: 100-300, exec: checkSwaggerDocs)

WARN[0000] Insufficient VUs, reached 300 active VUs and cannot initialize more  executor=constant-arrival-rate scenario=checkSwggerDocs

running (1m00.0s), 000/300 VUs, 31814 complete and 300 interrupted iterations
checkSwggerDocs ✓ [======================================] 300/300 VUs  1m0s  15000 iters/s

     ✓ status is 200

     checks.....................: 100.00% ✓ 31814        ✗ 0
     data_received..............: 127 MB  2.1 MB/s
     data_sent..................: 1.6 MB  26 kB/s
     dropped_iterations.........: 867886  14463.302739/s
     http_req_blocked...........: avg=2.49ms   min=0s     med=0s     max=401ms    p(90)=1µs  p(95)=1µs
     http_req_connecting........: avg=836.57µs min=0s     med=0s     max=138.78ms p(90)=0s   p(95)=0s
     http_req_duration..........: avg=550.82ms min=2.27ms med=5.47ms max=5.01s    p(90)=2s   p(95)=3s
     http_req_failed............: 0.00%   ✓ 0            ✗ 31814
     http_req_receiving.........: avg=72.27µs  min=11µs   med=28µs   max=8.64ms   p(90)=68µs p(95)=205µs
     http_req_sending...........: avg=36.09µs  min=14µs   med=30µs   max=2.37ms   p(90)=52µs p(95)=69µs
     http_req_tls_handshaking...: avg=1.65ms   min=0s     med=0s     max=268.3ms  p(90)=0s   p(95)=0s
     http_req_waiting...........: avg=550.71ms min=2.22ms med=5.35ms max=5.01s    p(90)=2s   p(95)=3s
     http_reqs..................: 31814   530.17967/s
     iteration_duration.........: avg=553.42ms min=2.34ms med=5.58ms max=5.4s     p(90)=2s   p(95)=3s
     iterations.................: 31814   530.17967/s
     vus........................: 300     min=300        max=300
     vus_max....................: 300     min=300        max=300
```

### Балансировка 2-1-1

```bash
$ k6 run docs-script.js --env HOST="https://..."

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: docs-script.js
     output: -

  scenarios: (100.00%) 1 scenario, 300 max VUs, 1m0s max duration (incl. graceful stop):
           * checkSwggerDocs: 15000.00 iterations/s for 1m0s (maxVUs: 100-300, exec: checkSwaggerDocs)

WARN[0000] Insufficient VUs, reached 300 active VUs and cannot initialize more  executor=constant-arrival-rate scenario=checkSwggerDocs

running (1m00.0s), 000/300 VUs, 31787 complete and 300 interrupted iterations
checkSwggerDocs ✓ [======================================] 298/300 VUs  1m0s  15000 iters/s

     ✓ status is 200

     checks.....................: 100.00% ✓ 31787        ✗ 0
     data_received..............: 127 MB  2.1 MB/s
     data_sent..................: 1.6 MB  26 kB/s
     dropped_iterations.........: 867923  14464.617432/s
     http_req_blocked...........: avg=2.41ms   min=0s     med=0s     max=431.31ms p(90)=1µs  p(95)=1µs
     http_req_connecting........: avg=1.14ms   min=0s     med=0s     max=219.57ms p(90)=0s   p(95)=0s
     http_req_duration..........: avg=551.22ms min=2.18ms med=5.29ms max=5.04s    p(90)=2s   p(95)=3s
     http_req_failed............: 0.00%   ✓ 0            ✗ 31787
     http_req_receiving.........: avg=67.94µs  min=10µs   med=22µs   max=11.08ms  p(90)=47µs p(95)=205µs
     http_req_sending...........: avg=28.35µs  min=14µs   med=25µs   max=2.4ms    p(90)=36µs p(95)=44µs
     http_req_tls_handshaking...: avg=1.26ms   min=0s     med=0s     max=217.83ms p(90)=0s   p(95)=0s
     http_req_waiting...........: avg=551.12ms min=2.13ms med=5.19ms max=5.04s    p(90)=2s   p(95)=3s
     http_reqs..................: 31787   529.755283/s
     iteration_duration.........: avg=553.72ms min=2.26ms med=5.37ms max=5.41s    p(90)=2s   p(95)=3s
     iterations.................: 31787   529.755283/s
     vus........................: 300     min=300        max=300
     vus_max....................: 300     min=300        max=300
```

## Регистрация в системе

### Без балансировки

```bash
$ k6 run registration-script.js --env HOST="https://..."

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: registration-script.js
     output: -

  scenarios: (100.00%) 1 scenario, 300 max VUs, 1m0s max duration (incl. graceful stop):
           * registrationFlow: 15000.00 iterations/s for 1m0s (maxVUs: 100-300, exec: registrationFlow)

WARN[0000] Insufficient VUs, reached 300 active VUs and cannot initialize more  executor=constant-arrival-rate scenario=registrationFlow

running (1m00.0s), 000/300 VUs, 13800 complete and 300 interrupted iterations
registrationFlow ✓ [======================================] 300/300 VUs  1m0s  15000 iters/s

     ✗ status is 200
      ↳  83% — ✓ 11466 / ✗ 2334

     checks.....................: 83.08% ✓ 11466        ✗ 2334
     data_received..............: 11 MB  188 kB/s
     data_sent..................: 5.6 MB 94 kB/s
     dropped_iterations.........: 885910 14763.878518/s
     http_req_blocked...........: avg=3.45ms   min=0s      med=0s       max=559.2ms  p(90)=1µs   p(95)=1µs
     http_req_connecting........: avg=166.21µs min=0s      med=0s       max=30.51ms  p(90)=0s    p(95)=0s
     http_req_duration..........: avg=679.57ms min=6.04ms  med=630.24ms max=2.67s    p(90)=1.22s p(95)=1.42s
     http_req_failed............: 8.94%  ✓ 2334         ✗ 23746
     http_req_receiving.........: avg=32.43µs  min=10µs    med=25µs     max=813µs    p(90)=46µs  p(95)=63µs
     http_req_sending...........: avg=43.25µs  min=19µs    med=37µs     max=754µs    p(90)=61µs  p(95)=77µs
     http_req_tls_handshaking...: avg=3.28ms   min=0s      med=0s       max=551.96ms p(90)=0s    p(95)=0s
     http_req_waiting...........: avg=679.49ms min=5.97ms  med=630.16ms max=2.67s    p(90)=1.22s p(95)=1.42s
     http_reqs..................: 26080  434.628745/s
     iteration_duration.........: avg=1.28s    min=57.13ms med=1.24s    max=3.7s     p(90)=2.02s p(95)=2.24s
     iterations.................: 13800  229.979934/s
     vus........................: 300    min=300        max=300
     vus_max....................: 300    min=300        max=300
```

### Балансировка 2-1-1

```bash
$ k6 run registration-script.js --env HOST="https://..."

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: registration-script.js
     output: -

  scenarios: (100.00%) 1 scenario, 300 max VUs, 1m0s max duration (incl. graceful stop):
           * registrationFlow: 15000.00 iterations/s for 1m0s (maxVUs: 100-300, exec: registrationFlow)

WARN[0000] Insufficient VUs, reached 300 active VUs and cannot initialize more  executor=constant-arrival-rate scenario=registrationFlow

running (1m00.0s), 000/300 VUs, 29981 complete and 300 interrupted iterations
registrationFlow ✓ [======================================] 300/300 VUs  1m0s  15000 iters/s

     ✗ status is 200
      ↳  90% — ✓ 27000 / ✗ 2981

     checks.....................: 90.05% ✓ 27000        ✗ 2981
     data_received..............: 25 MB  416 kB/s
     data_sent..................: 12 MB  204 kB/s
     dropped_iterations.........: 869729 14494.221611/s
     http_req_blocked...........: avg=2.11ms   min=0s      med=0s       max=870.48ms p(90)=1µs   p(95)=1µs
     http_req_connecting........: avg=505.5µs  min=0s      med=0s       max=209.45ms p(90)=0s    p(95)=0s
     http_req_duration..........: avg=306.41ms min=5.26ms  med=26.01ms  max=8.39s    p(90)=1.02s p(95)=1.1s
     http_req_failed............: 5.17%  ✓ 2981         ✗ 54652
     http_req_receiving.........: avg=27.29µs  min=9µs     med=20µs     max=1.47ms   p(90)=40µs  p(95)=59µs
     http_req_sending...........: avg=35.67µs  min=18µs    med=31µs     max=3.47ms   p(90)=49µs  p(95)=61µs
     http_req_tls_handshaking...: avg=1.61ms   min=0s      med=0s       max=661.95ms p(90)=0s    p(95)=0s
     http_req_waiting...........: avg=306.35ms min=5.21ms  med=25.94ms  max=8.39s    p(90)=1.02s p(95)=1.1s
     http_reqs..................: 57633  960.466391/s
     iteration_duration.........: avg=592.15ms min=15.35ms med=535.37ms max=9.21s    p(90)=1.23s p(95)=1.9s
     iterations.................: 29981  499.63984/s
     vus........................: 300    min=300        max=300
     vus_max....................: 300    min=300        max=300
```

## Загрузка новой модели в систему

### Без балансировки

```bash
$ k6 run model-uploading-script.js --env HOST="https://..."

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: model-uploading-script.js
     output: -

  scenarios: (100.00%) 1 scenario, 10 max VUs, 1m30s max duration (incl. graceful stop):
           * fullUserFlow: 10 looping VUs for 1m0s (exec: fullUserFlow, gracefulStop: 30s)


running (1m01.0s), 00/10 VUs, 541 complete and 0 interrupted iterations
fullUserFlow ✓ [======================================] 10 VUs  1m0s

     ✓ status is 200
     ✓ token is present
     ✓ id is present

     checks.....................: 100.00% ✓ 2705      ✗ 0
     data_received..............: 1.1 MB  18 kB/s
     data_sent..................: 1.1 MB  19 kB/s
     http_req_blocked...........: avg=318.92µs min=0s       med=1µs     max=73.82ms  p(90)=1µs      p(95)=2µs
     http_req_connecting........: avg=235.68µs min=0s       med=0s      max=51.64ms  p(90)=0s       p(95)=0s
     http_req_duration..........: avg=280.41ms min=4.18ms   med=56.39ms max=738.45ms p(90)=646.4ms  p(95)=662.19ms
     http_req_failed............: 0.00%   ✓ 0         ✗ 2164
     http_req_receiving.........: avg=112.58µs min=12µs     med=58µs    max=3.69ms   p(90)=141.7µs  p(95)=215.54µs
     http_req_sending...........: avg=113.65µs min=22µs     med=96.5µs  max=1.28ms   p(90)=191µs    p(95)=232.84µs
     http_req_tls_handshaking...: avg=81.72µs  min=0s       med=0s      max=23.07ms  p(90)=0s       p(95)=0s
     http_req_waiting...........: avg=280.19ms min=4.03ms   med=56.26ms max=737.6ms  p(90)=646.27ms p(95)=662.1ms
     http_reqs..................: 2164    35.491282/s
     iteration_duration.........: avg=1.12s    min=548.92ms med=1.16s   max=1.38s    p(90)=1.27s    p(95)=1.28s
     iterations.................: 541     8.872821/s
     vus........................: 10      min=10      max=10
     vus_max....................: 10      min=10      max=10
```

### Балансировка 2-1-1

```bash
$ k6 run model-uploading-script.js --env HOST="https://..."

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: model-uploading-script.js
     output: -

  scenarios: (100.00%) 1 scenario, 10 max VUs, 1m30s max duration (incl. graceful stop):
           * fullUserFlow: 10 looping VUs for 1m0s (exec: fullUserFlow, gracefulStop: 30s)


running (1m01.0s), 00/10 VUs, 540 complete and 0 interrupted iterations
fullUserFlow ✓ [======================================] 10 VUs  1m0s

     ✓ status is 200
     ✓ token is present
     ✓ id is present

     checks.....................: 100.00% ✓ 2700      ✗ 0
     data_received..............: 1.1 MB  18 kB/s
     data_sent..................: 1.1 MB  19 kB/s
     http_req_blocked...........: avg=232.18µs min=0s       med=1µs     max=52.68ms  p(90)=1µs      p(95)=1µs
     http_req_connecting........: avg=163.88µs min=0s       med=0s      max=36.71ms  p(90)=0s       p(95)=0s
     http_req_duration..........: avg=279.78ms min=4.2ms    med=73.01ms max=723.58ms p(90)=613.09ms p(95)=629.15ms
     http_req_failed............: 0.00%   ✓ 0         ✗ 2160
     http_req_receiving.........: avg=117.15µs min=10µs     med=53µs    max=9.02ms   p(90)=142.1µs  p(95)=276.49µs
     http_req_sending...........: avg=111.53µs min=21µs     med=90µs    max=3.21ms   p(90)=179.1µs  p(95)=221µs
     http_req_tls_handshaking...: avg=67.13µs  min=0s       med=0s      max=16.23ms  p(90)=0s       p(95)=0s
     http_req_waiting...........: avg=279.55ms min=4.05ms   med=72.86ms max=723.36ms p(90)=612.92ms p(95)=628.83ms
     http_reqs..................: 2160    35.416837/s
     iteration_duration.........: avg=1.12s    min=579.81ms med=1.15s   max=1.37s    p(90)=1.24s    p(95)=1.25s
     iterations.................: 540     8.854209/s
     vus........................: 10      min=10      max=10
     vus_max....................: 10      min=10      max=10
```
