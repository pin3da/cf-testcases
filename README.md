# cf-testcases

Automatically fetch sample test cases from codeforces.

This code is ugly and slow, please don't use it.

## install

```
cd $GOPATH
go get -u github.com/pin3da/cf-testcases
go install github.com/pin3da/cf-testcases
```

## usage

```
cf-testcase contest_id
```

Example:
This will download all the sample test cases of `Codeforces Beta Round #57 (Div. 2)`

Note that 61 is the number in the URL of the contest.

```
cf-testcase 61
```

----

[Manuel Pineda](https://github.com/pin3da)
