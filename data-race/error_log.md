# race_test.go
```log
==================
WARNING: DATA RACE
Write at 0x00c000028250 by goroutine 7:
  command-line-arguments.(*RPC).compute()
      /Users/yong/Documents/GitHub/learn-go/data-race/race_test.go:16 +0x44
  command-line-arguments.TestRace.gowrap1()
      /Users/yong/Documents/GitHub/learn-go/data-race/race_test.go:27 +0x20

Previous read at 0x00c000028250 by goroutine 6:
  command-line-arguments.TestRace()
      /Users/yong/Documents/GitHub/learn-go/data-race/race_test.go:28 +0x120
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1792 +0x180
  testing.(*T).Run.gowrap1()
      /usr/local/go/src/testing/testing.go:1851 +0x40

Goroutine 7 (running) created at:
  command-line-arguments.TestRace()
      /Users/yong/Documents/GitHub/learn-go/data-race/race_test.go:27 +0x114
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1792 +0x180
  testing.(*T).Run.gowrap1()
      /usr/local/go/src/testing/testing.go:1851 +0x40

Goroutine 6 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1851 +0x684
  testing.runTests.func1()
      /usr/local/go/src/testing/testing.go:2279 +0x7c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1792 +0x180
  testing.runTests()
      /usr/local/go/src/testing/testing.go:2277 +0x77c
  testing.(*M).Run()
      /usr/local/go/src/testing/testing.go:2142 +0xb68
  main.main()
      _testmain.go:45 +0x110
==================
RPC computation complete, result: 42, version: 1
--- FAIL: TestRace (1.00s)
    testing.go:1490: race detected during execution of test
FAIL
FAIL    command-line-arguments  1.356s
FAIL
```

# race2_test.go
```log
==================
WARNING: DATA RACE
Write at 0x00c0000d8078 by goroutine 7:
  command-line-arguments.(*RPCWithMutex).compute()
      /Users/yong/Documents/GitHub/learn-go/data-race/race2_test.go:18 +0x44
  command-line-arguments.TestRaceRPCWithMutex.gowrap1()
      /Users/yong/Documents/GitHub/learn-go/data-race/race2_test.go:29 +0x20

Previous read at 0x00c0000d8078 by goroutine 6:
  command-line-arguments.TestRaceRPCWithMutex()
      /Users/yong/Documents/GitHub/learn-go/data-race/race2_test.go:30 +0x124
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1792 +0x180
  testing.(*T).Run.gowrap1()
      /usr/local/go/src/testing/testing.go:1851 +0x40

Goroutine 7 (running) created at:
  command-line-arguments.TestRaceRPCWithMutex()
      /Users/yong/Documents/GitHub/learn-go/data-race/race2_test.go:29 +0x118
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1792 +0x180
  testing.(*T).Run.gowrap1()
      /usr/local/go/src/testing/testing.go:1851 +0x40

Goroutine 6 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1851 +0x684
  testing.runTests.func1()
      /usr/local/go/src/testing/testing.go:2279 +0x7c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1792 +0x180
  testing.runTests()
      /usr/local/go/src/testing/testing.go:2277 +0x77c
  testing.(*M).Run()
      /usr/local/go/src/testing/testing.go:2142 +0xb68
  main.main()
      _testmain.go:45 +0x110
==================
RPC computation complete, result: 42, version: 1
--- FAIL: TestRaceRPCWithMutex (1.00s)
    testing.go:1490: race detected during execution of test
FAIL
FAIL    command-line-arguments  1.158s
FAIL
```

# race3_test.go
```log
RPC computation complete, result: 0
PASS
==================
WARNING: DATA RACE
Write at 0x00c000118210 by goroutine 7:
  command-line-arguments.(*RPCNoWait).compute()
      /Users/yong/Documents/GitHub/learn-go/data-race/race3_test.go:16 +0x44
  command-line-arguments.TestRaceRPCNoWait.gowrap1()
      /Users/yong/Documents/GitHub/learn-go/data-race/race3_test.go:27 +0x20

Previous read at 0x00c000118210 by goroutine 6:
  command-line-arguments.TestRaceRPCNoWait()
      /Users/yong/Documents/GitHub/learn-go/data-race/race3_test.go:31 +0x114
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1792 +0x180
  testing.(*T).Run.gowrap1()
      /usr/local/go/src/testing/testing.go:1851 +0x40

Goroutine 7 (running) created at:
  command-line-arguments.TestRaceRPCNoWait()
      /Users/yong/Documents/GitHub/learn-go/data-race/race3_test.go:27 +0x10c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1792 +0x180
  testing.(*T).Run.gowrap1()
      /usr/local/go/src/testing/testing.go:1851 +0x40

Goroutine 6 (finished) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1851 +0x684
  testing.runTests.func1()
      /usr/local/go/src/testing/testing.go:2279 +0x7c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1792 +0x180
  testing.runTests()
      /usr/local/go/src/testing/testing.go:2277 +0x77c
  testing.(*M).Run()
      /usr/local/go/src/testing/testing.go:2142 +0xb68
  main.main()
      _testmain.go:45 +0x110
==================
Found 1 data race(s)
FAIL    command-line-arguments  1.381s
FAIL
```