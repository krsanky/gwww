gwww is a Go WWW framework.

This is a boilerplate site using postrgesql and bootstrap.
(and scss and ...)

It is very modular in design, but not config.  ie. You can easliy 
swap to a different database backend, but you would do it in the code.

It is developed on OpenBSD machines, but could run anywhere. 

---

There is an inner lighter weight framework asking to come out, just chop 
out db, sessions, account etc.  and leave the routing and view and template
stuff

