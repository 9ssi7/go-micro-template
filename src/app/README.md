## App Folder

Under this folder is the `app.go` file and this is where we manage our application.

Here we make database connections, importing locale and env files, connecting to the message broker, creating the internal folder instances.

> Why didn't we do this on `main.go`?
>
> - Because we want to keep the `main.go` file as clean as possible. Also we can write test codes in our application.
>
> - We make fake links while writing test codes. So we don't really connect to a message broker or database, we mock them.
>
> - Thus, the number of codes we will write in the configuration file of both `main.go` and e2e is reduced.
