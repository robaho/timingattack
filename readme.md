**Overview**

This project is designed to test the variance of making http calls to remote services
in order to determine their susceptibility to [timing attacks](https://en.wikipedia.org/wiki/Timing_attack)

timing_test.go determines the difference in an early-exit string comparison with no characters matching, and the first matching.

main is a Go web app that handles the index url and returns a 401 if the supplied guess is incorrect, and OK if it does.

**Results**

I ran the tests using my local machine, a core i7 3.4 ghz.

For timing_test:

<pre>
goos: darwin
goarch: amd64
pkg: github.com/robaho/timingattack
BenchmarkTimeNone-8   	2000000000	         0.86 ns/op
BenchmarkTimeOne-8    	2000000000	         1.43 ns/op
PASS
</pre>

So, assuming a not possible 0 variance, the difference in is 0.5 nano seconds.

When testing against a local server, with an iteration of 100k tries:

<pre>
avg time  136009 nanos variance 63.623894882347464 nanos
</pre>

It was also tested against a remote server deployed to the Google Cloud, Flex setup.

For these tests, two iteration counts were used, 100, and 5000 (the iterations were capped because:
1) didn't have the time, 2) might of costed some money and 3) any real security would of stopped me detecting a DOS attack far sooner.

For 100 iterations: 
<pre>
avg time  63833616 nanos std 369801.15152876417 nanos
</pre>

For 5000 iterations:
<pre>
avg time  60008116 nanos std 21627.8937485831 nanos
</pre>

**Conclusion**

Statistically, there is a 0% chance of a timing attack due to non-constant string compare in a web application.
The fastest "local" attack has a std with 3 orders of magnitude more than the attack threshold.
For a remote attack, it is more than 6 orders of magnitude, and more importantly, the number of attack requests required would
cause any competent security setup to deny the connection attempt under DOS controls.



 