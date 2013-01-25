# coroutinepool
=============
go coroutine pool  


# License
===
author gonghh   
Copyright 2013 gonghh(screscent).  
Under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).  


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
