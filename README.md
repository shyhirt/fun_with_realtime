# fun_with_realtime

Instead of using a real data source, implement a goroutine that asynchronously generates random numbers in the range 1 to 10 and pushes them into an input queue.

In the background, run a worker that drains the input queue and recalculates the average values for the last 10 seconds.

The service fun_with_realtime should support an unlimited number of clients/subscribers connected via WebSocket.

Whenever the aggregated data is recalculated, broadcast the updated results to all subscribers.

Additionally, implement a test client/subscriber within the same project—for example, a simple console application.
