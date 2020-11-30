package main

import (
	"fmt"
	"sort"
)

type Request struct {
	priority int
	time     int
	processed bool
}

func createPQ ( requests []Request, priority int, time int, processed bool) []Request {
	request := Request{
		priority: priority,
		time: time,
		processed: processed,
	}
	requests = append(requests, request);
	return requests;
}

func processedRequests(requests [] Request, q int) []Request {
	a := true;
	for a {
		for i:=0; i < len(requests); i ++ {
			if requests[i].time < q {
				requests[i].processed = true;
				return requests;
			} else {
				a = false;
				fmt.Println("Remove:= ", requests[i])
				requests[i] = requests[len(requests)-1];

				fmt.Println("---------------------");

				for i := 0; i < len(requests); i++ {
					fmt.Println(requests[i]);
				}

				fmt.Println("---------------------");
				requests = append(requests, requests[i]);
				return requests;
			}
		}
		return nil;
	}
	return nil
}

func main()  {
	requests := []Request{};

	for i:=0; i < 5; i ++ {
		var priority int;
		var time int;
		fmt.Scan(&priority);
		fmt.Scan(&time);
		_requests := createPQ(requests, priority, time, false);
		requests = _requests;
	}

	for i:=0; i < len(requests); i ++ {
		fmt.Println(requests[i])
	}

	sort.SliceStable(requests, func(i, j int) bool {
		return requests[i].priority < requests[j].priority
	})

	q := 10;

	fmt.Println("PQ")

	for i:=0; i < len(requests); i ++ {
		fmt.Println(requests[i]);
	}

	fmt.Println("Время для исполения заявки:= ", q)

	_requests := processedRequests(requests, q);

	for i := 0; i < len(_requests); i++ {
		fmt.Println(_requests[i])
	}
}
