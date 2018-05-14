## Max Profit Trade Dates Problem

### Abstract
This service exposes an API that takes in a stock symbol (e.g. AAPL, GOOG) and and the number of days to look into the past and
returns a single buy and sell date that yields the max profit over the last N number of days.
The service uses REST APIs exposed by https://www.alphavantage.co to retrieve historical stock prices, but only considers the high price for the day.

### Dependencies
* Docker 17.12 or above [The program has been tested only on this docker version]
* curl

### Start the service
* Unarchive the problem.solving.tar.gz
* Traverse to problem.solving folder
* Run following command to build and run unit tests in the all the packages:
```bash
$ docker build . -t max.profit
```
* Run following command:
```bash
$ docker run --rm  -p 8006:8006 -it max.profit
```

### Invoke the API
* In a separate terminal window run the following command:
```bash
$ curl -X GET "http://127.0.0.1:8006/stock/GOOGL/days/180"
```

### Important Note
* The stock history retrieved from www.alphavantage.co provides the history excluding the weekend dates.
  Hence the number of days you use for looking up the history will have slightly different start date and end date.

* Also, www.alphavantage.co has only 2 distinct mode for data sizes for retrieval.
  'compact' data size returns only 100 days worth of history, while 'full' returns 20 years worth of data.
  Due to time constraints, other sources of stock history couldn't be evaluated and
  the performance hit due to retrieval of 20 years of data is a known issue.

* The service cannot handle concurrent API requests at this point.