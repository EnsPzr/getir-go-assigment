<h1>Getir-Go-Assigment</h1>

<h2>Description:</h2>
This project was given as homework by Getir. This is a restful api project. Mongodb, sqlite and go-cache are used inside
the project. Mongodb is used for database. Go-cache and sqlite are used for in-memory database. Project has 5
routes. <br>

<a href="https://getir-go-assigment.herokuapp.com/swagger">Swagger</a>

<h2>Summary of Routes:</h2>
<ul>
<li>
Records: <a href="https://getir-go-assigment.herokuapp.com/records">https://getir-go-assigment.herokuapp.com/records</a><br>
Returns:

````json
{
  "code": 0,
  "msg": "Success",
  "records": [
    {
      "createdAt": "2016-04-02T21:08:18.091Z",
      "key": "xlPSAiou",
      "totalCount": 4017
    },
    ...
  ]
}
````

Code Means: <br>

- 0: Success
- 1: Internal Server Error
- 2: Bad request
- 3: Method Not Allowed

</li>
Records with filter:
<ul>
<li>
Records with start date filter: <a href="https://getir-go-assigment.herokuapp.com/records?startDate=2015-10-14">https://getir-go-assigment.herokuapp.com/records?startDate=2015-10-14</a>
</li>
<li>
Records with end date filter: <a href="https://getir-go-assigment.herokuapp.com/records?endDate=2016-12-27">https://getir-go-assigment.herokuapp.com/records?endDate=2016-12-27</a>
</li>
<li>
Records with start date and end date filter: <a href="https://getir-go-assigment.herokuapp.com/records?startDate=2015-10-14&endDate=2016-12-27">https://getir-go-assigment.herokuapp.com/records?startDate=2015-10-14&endDate=2016-12-27</a>
</li>
<li>
Records with min count filter: <a href="https://getir-go-assigment.herokuapp.com/records?minCount=40">https://getir-go-assigment.herokuapp.com/records?minCount=40</a>
</li>
<li>
Records with max count filter: <a href="https://getir-go-assigment.herokuapp.com/records?maxCount=3000">https://getir-go-assigment.herokuapp.com/records?maxCount=3000</a>
</li>
</ul>
<li>
In memory get: <a href="https://getir-go-assigment.herokuapp.com/in-memory?key=getir">https://getir-go-assigment.herokuapp.com/in-memory?key=getir</a>
</li>
<li>
In memory post: <a href="https://getir-go-assigment.herokuapp.com/in-memory">https://getir-go-assigment.herokuapp.com/in-memory</a><br>
Request body: <br>

````go
type InMemory struct {
Key   string `json:"key"`
Value string `json:"value"`
}
````

</li>
<li>
In memory sqlite get: <a href="https://getir-go-assigment.herokuapp.com/in-memory-sqlite?key=getir">https://getir-go-assigment.herokuapp.com/in-memory?key=getir</a>
</li>
<li>
In memory sqlite post: <a href="https://getir-go-assigment.herokuapp.com/in-memory-sqlite">https://getir-go-assigment.herokuapp.com/in-memory</a><br>
Request body: <br>

````go
type InMemory struct {
Key   string `json:"key"`
Value string `json:"value"`
}
````

</li>
</ul>
<h2>Detail of Routes:</h2>
<ul>
<li>
<a href="https://getir-go-assigment.herokuapp.com/records">https://getir-go-assigment.herokuapp.com/records</a> <br>
This route returns records. Route takes 4 parameters. <br>
<ul>
<li>
startDate=YYYY-MM-DD (optional) <br>
If startdate exist, returns records which creation date is greater than the start date.<br>
Example route: <a href="https://getir-go-assigment.herokuapp.com/records?startDate=2015-10-14">https://getir-go-assigment.herokuapp.com/records?startDate=2015-10-14</a> <br>
</li>
<li>
endDate=YYYY-MM-DD (optional) <br>
If endDate exist, returns records which creation date is smaller than the end date.<br>
Example route: <a href="https://getir-go-assigment.herokuapp.com/records?endDate=2016-12-27">https://getir-go-assigment.herokuapp.com/records?endDate=2016-12-27</a> <br>
</li>
<li>
minCount=number (optional) <br>
If minCount exist, returns records which sum counts is greater than the min count.<br>
Example route: <a href="https://getir-go-assigment.herokuapp.com/records?minCount=40">https://getir-go-assigment.herokuapp.com/records?minCount=40</a> <br>
</li>
<li>
maxCount=number (optional) <br>
If maxCount exist, returns records which sum counts is smaller than the max count.<br>
Example route: <a href="https://getir-go-assigment.herokuapp.com/records?maxCount=3000">https://getir-go-assigment.herokuapp.com/records?maxCount=3000</a> <br>
</li>
</ul>
<br>
</li>
<li>
<a href="https://getir-go-assigment.herokuapp.com/in-memory">https://getir-go-assigment.herokuapp.com/in-memory</a> <br>
This route has two methods. GET and POST.
<ul>
<li>GET<br>
This route gets value from go-cache according to key. <br>
If key is empty, returns 400 status code. <br>
If key is not found, returns 404 status code. <br>
If key is exist, returns 200 status code, key and value. <br>
Example route: <a href="https://getir-go-assigment.herokuapp.com/in-memory?key=getir">https://getir-go-assigment.herokuapp.com/in-memory?key=getir</a> <br>
</li>
<li>POST <br>
This route keeps key and value in go-cache.<br>
Attention! This route has validation method.
Key and value is required!<br>
If body cannot be read, returns 500 status code with error message.<br>
If body cannot be bind to model, returns 500 status code with error message.<br>
If validation condition is not true, returns 400 status code with validation errors.<br>
If operation is succesful, returns 200 status code with request body.<br>
Route :<a href="https://getir-go-assigment.herokuapp.com/in-memory">https://getir-go-assigment.herokuapp.com/in-memory</a> <br>
Request Body:

````go
type InMemory struct {
Key   string `json:"key"`
Value string `json:"value"`
}
````

</li>
<br>
</ul>
</li>
<li>
<a href="https://getir-go-assigment.herokuapp.com/in-memory-sqlite">https://getir-go-assigment.herokuapp.com/in-memory-sqlite</a> <br>
This route has two method. GET and POST.
<ul>
<li>GET<br>
This route gets value from sqlite-cache according to key. <br>
If key is empty, returns 400 status code. <br>
If there is any error when it was taken from cache error, returns 500 status code with error message. <br>
If key is not found, returns 404 status code. <br>
If key is exist, returns 200 status code, key and value. <br>
Example route: <a href="https://getir-go-assigment.herokuapp.com/in-memory-sqlite?key=getir">https://getir-go-assigment.herokuapp.com/in-memory-sqlite?key=getir</a> <br>
</li>
<li>POST<br>
This route keeps key and value in sqlite-cache.<br>
Attention! This route has validation method.
Key and value is required!<br>
If body cannot be read, returns 500 status code with error message. <br>
If body cannot be bind to model, returns 500 status code with error message. <br>
If validation condition not true, returns 400 status code with validation errors. <br>
If operation is successful, returns 200 status code with request body. <br>
Route: <a href="https://getir-go-assigment.herokuapp.com/in-memory-sqlite">https://getir-go-assigment.herokuapp.com/in-memory-sqlite</a> <br>
Request Body:

````go
type InMemory struct {
Key   string `json:"key"`
Value string `json:"value"`
}
````

</li>
</ul>
</li>
</ul>



<h2>Folder Structure</h2>
<ul>
<li>
cache<br>
This folder holds files which contain cache(in-memory) structure and methods.<br>
There are 2 cache(in-memory) structures.<br>
- go-cache <br>
- sqlite
</li>
<li>
cmd<br>
This folder contains main function. <br>
The project starts with this function.<br>
This function calls database connection, caches initialization, route setup and http server starts functions.
</li>
<li>
database<br>
This folder contains mongodb database client, connect and disconnect function.
</li>
<li>
handlers<br>
This folder holds files containing functions to process requests.</li>
<li>
model<br>
This folder holds files containing the models required for the project.
</li>
<li>
router<br>
This folder contains route definition function.</li>
<li>
service<br>
This folder contains database operations functions.
</li>
<li>
test <br>
This folder contains files containing test functions.
</li>
</ul>
