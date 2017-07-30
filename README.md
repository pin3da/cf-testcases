# cf-testcases

Automatically fetch sample test cases from codeforces.

## install

```
cd $GOPATH
go get -u github.com/pin3da/cf-testcases
go install github.com/pin3da/cf-testcases
```

## usage

```
cf-testcase contest_id [problem_letter]
```

Example:
This will download all the sample test cases of `Codeforces Beta Round #57 (Div. 2)`

Note that 61 is the number in the URL of the contest.

```
cf-testcase 61
```

Download the cases for the problem D

```
cf-testcase 61 d
```

----

[Manuel Pineda](https://github.com/pin3da)
