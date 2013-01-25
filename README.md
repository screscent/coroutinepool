coroutinepool
=============

gocoroutine pool


example
================


//create
coroutinepool.Create(6)

//process
  for i := 0; i < 100; i++ {
		coroutinepool.Run(func() {
			fmt.Println("test", i)
		})
	}
  
  .
  .
  .
//exit
coroutinepool.Exit()
