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

func RemoveIndex(requests []Request, index int) []Request {
	return append(requests[:index], requests[index+1:]...)
}

func RemoveCopy(requests  []Request, i int) []Request {
	copy(requests [i:], requests [i+1:])
	return requests [:len(requests)-1]
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


		for i:=0; i < len(requests); i ++ {
			if requests[i].time < q {
				requests[i].processed = true;

				requests = append(requests, requests[i]);
				_requests := RemoveCopy(requests, 0)
				//return requests;
				fmt.Println("Updated");
				for i:=0; i < len(_requests); i ++ {
					fmt.Printf("%+v\n", _requests[i]);
				}
				return _requests;
			} else {

				fmt.Println("Remove:= ", requests[i])
				requests = append(requests, requests[i]);
				_requests := RemoveCopy(requests, 0)

				fmt.Println("---------------------");

				for i := 0; i < len(_requests); i++ {
					fmt.Printf("%+v\n", _requests[i]);
				}

				fmt.Println("---------------------");
				//requests = append(requests, requests[i]);
				//processedRequests(requests, q);
				return _requests;
			}
		}
		return nil;
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

	//for i:=0; i < len(requests); i ++ {
	//	fmt.Println(requests[i])
	//}

	sort.SliceStable(requests, func(i, j int) bool {
		return requests[i].priority < requests[j].priority
	})

	q := 10;

	fmt.Println("PQ")

	for i:=0; i < len(requests); i ++ {
		fmt.Println(requests[i]);
	}

	fmt.Println("Время для исполения заявки:= ", q)

	for i:=0; i < len(requests); i++ {
		_requests := processedRequests(requests, q);
		//fmt.Println("result: ", _requests);
		requests = _requests;
	}

	for i := 0; i < len(requests); i++ {
		fmt.Println(requests[i])
	}
}
