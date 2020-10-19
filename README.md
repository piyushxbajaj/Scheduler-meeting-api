# Meeting Scheduler API

Made with Go Lang, this api can help setup meetings without collisions and huge ease.

## Installation

Install go and run this command in terminal to start the app

```bash
go run app.go
```

## Test

Test GET meeting between start and end time
```bash
Running tool: /usr/bin/go test -benchmem -run=^$ -bench ^(BenchmarktimeMeeting)$
goos: darwin
goarch: intel
BenchmarktimeMeeting-8   	    1346	    4441239 ns/op	   15587 B/op	     120 allocs/op
PASS
ok  	Home/piyush/dev/schduler-meeting-api	3.332s
```

Test GET Participant
```bash
Running tool: /usr/bin/go test -benchmem -run=^$ -bench ^(Benchmarkidparticipant)$
goos: darwin
goarch: intel
Benchmarkidparticipant-8   	    900	    6621239 ns/op	   16667 B/op	     120 allocs/op
PASS
ok  	Home/piyush/dev/schduler-meeting-api	3.332s
```

Test GET Participant
```bash
Running tool: /usr/bin/go test -benchmem -run=^$ -bench ^(Benchmarkparticipant)$
goos: darwin
goarch: intel
Benchmarkparticipant-8   	    1716	    8621239 ns/op	   16787 B/op	     120 allocs/op
PASS
ok  	Home/piyush/dev/schduler-meeting-api	3.332s
```
## Usage

2 features out of other features are displayed here 
POST Create Meeting
```json
{
    "ID":"5",
    "Title":"Piyush",
    "Partipicants":"2",
    "StartTime": "2019-01-22T11:24:39.34Z",
    "EndTime": "2019-01-23T11:24:39.34Z"

}
```
GET Meeting with ID
```json
{"ID":"2","Title":"Lyft budget","Partcipants":0,"StartTime":"2019-01-30T11:24:39.34Z","EndTime":"2019-01-31T11:24:39.34Z","CreationTime":"0001-01-01T00:00:00Z"}
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
