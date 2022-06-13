# Exercise #4 & #5
<!-- markdownlint-disable -->

<h2>Building upon the code from the previous problem:</h2>

<p>Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".

Pass the connection of type net.Conn as an argument into this function.

Add "go" in front of the call to "serve" to enable concurrency and multiple connections.</p>

<h2>Exercise 5</h2>

Add code to WRITE to the connection