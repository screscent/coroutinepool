# coroutinepool
=============
go coroutine pool
author gonghh
Copyright 2013 gonghh(screscent). All rights reserved.



## example
==============================


* create

```go
coroutinepool.Create(6)

//or

p:=coroutinepool.NewPool()
p.Create(6)
```
* run

```go
  for i := 0; i < 100; i++ {
		coroutinepool.Run(func() {
			fmt.Println("test", i)
		})
  }

//or
  for i := 0; i < 100; i++ {
		p.Run(func() {
			fmt.Println("test", i)
		})
  }
  
```
* exit

```go
coroutinepool.Exit()

//or
p.Exit()
```
