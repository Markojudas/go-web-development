# Exercise #1

<!-- markdownlint-disable -->
<h2>ListenAndServe on port 8080 of localhost</h2>

for the default route "/" Have a func called "foo" which writes to the response "foo ran"

for the route "/dog" have a func called dog which parses a template called dog.gohtml and writes to the response "This is Dog" and also shows a picture of a dog when the template is excuted. Use "Http.ServeFile" to serve the file "dog.jpeg"
