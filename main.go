package main

import (
	"fmt"
	"sort"
)

const (
	SIZE = 10;
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
	for i:=0; i < len(requests); i ++ {
		if requests[i].time <= q {
			requests[i].processed = true;
			requests = append(requests, requests[i]);
			_requests := remove(requests, 0)
			fmt.Println("Update");
			fmt.Println("--------------------------------");
			for i:=0; i < len(_requests); i ++ {
				fmt.Printf("%+v\n", _requests[i]);
			}
			return _requests;
		} else {
			requests = append(requests, requests[i]);
			_requests := remove(requests, 0)
			fmt.Println("--------------------------------");
			for i := 0; i < len(_requests); i++ {
				fmt.Printf("%+v\n", _requests[i]);
			}
			fmt.Println("--------------------------------");
			return _requests;
		}
	}
	return nil;
}

func check (requests []Request, checkResult bool) (bool)  {
	for i:=0; i < len(requests); i ++ {
		if requests[i].processed == false {
			checkResult = false;
		}
	}
	return checkResult;
}

func remove(requests  []Request, i int) []Request {
	copy(requests [i:], requests [i+1:])
	return requests [:len(requests)-1]
}

func main()  {
	requests := []Request{};

	for i:=0; i < SIZE; i ++ {
		var priority int;
		var time int;
		fmt.Print("Priority: ")
		fmt.Scan(&priority);
		fmt.Print("Time: ")
		fmt.Scan(&time);
		_requests := createPQ(requests, priority, time, false);
		requests = _requests;
	}

	sort.SliceStable(requests, func(i, j int) bool {
		return requests[i].priority < requests[j].priority
	})

	q := 10;

	fmt.Println("Queue")

	for i:=0; i < len(requests); i ++ {
		fmt.Printf("%+v\n", requests[i]);
	}

	fmt.Println("Time for order execution: ", q)

	for i:=0; i < len(requests); i++ {
		_requests := processedRequests(requests, q);
		requests = _requests;
	}

	fmt.Println("Result");

	for i := 0; i < len(requests); i++ {
		fmt.Printf("%+v\n", requests[i])
	}

	checkResult := true;
	_checkResult := check(requests, checkResult);
	checkResult = _checkResult;

	if checkResult == false {
		newPQ := requests;
		newPQ2 := []Request{}
		fmt.Println("New queue")

		for i := range newPQ {
			if newPQ[i].processed == false {
				newPQ2 = append(newPQ2, newPQ[i])
			}
		}

		q = q * q;

		for i:=0; i < len(newPQ2); i++ {
			_newPQ2 := processedRequests(newPQ2, q);
			newPQ2 = _newPQ2;
		}

		fmt.Println("Result");

		for i := 0; i < len(newPQ2); i++ {
			fmt.Printf("%+v\n", newPQ2[i])
		}
	} else {
		fmt.Println("Finished")
	}
}
