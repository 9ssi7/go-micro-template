## Source Folder

This folder contains the source code of the project. The source code is divided into folders according to the architectural structure. Also `app.env` and `main.go` are located under this folder. Detailed information is available in the README inside the folder.


### [`src/app`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/app)

This folder contains the application code. The `app.go` file is located under this folder. `app.go` contains functions such as `Init` and `Serve` and provides an easy-to-manage structure. Detailed information is available in the README inside the folder.

### [`src/config`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/config)

This folder contains the config type of the application. The decode function in the core package will decode according to the type here. Detailed information is available in the README inside the folder.

### [`src/dto`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/dto)

This folder contains the DTOs of the application. DTOs are used to transfer data between layers. DTOs are used when transforming data from client to server and from server to client. Detailed information is available in the README inside the folder.

### [`src/entity`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/entity)

This folder contains the entities of the application. Entities are used to store data in the database. Detailed information is available in the README inside the folder.

### [`src/event`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/event)

This folder contains the events of the application. Events are used to transfer data between microservices. Detailed information is available in the README inside the folder.

### [`src/event_handler`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/event_handler)

This folder contains the event handlers of the application. Event handlers are used to handle events. Detailed information is available in the README inside the folder.

### [`src/event_publisher`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/event_publisher)

This folder contains the event publishers of the application. Event publishers are used to publish events. Detailed information is available in the README inside the folder.

### [`src/internal`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/internal)

This folder contains the internal code of the application. The internal code is the part that does the main work for microservice such as `repo`, `service`, `handler` and `api`. Detailed information is available in the README inside the folder.

### [`src/locales`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/locales)

This folder contains the localization files of the application. Localization files are used to store the text that will be displayed to the user. Detailed information is available in the README inside the folder.

### [`src/mapper`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/mapper)

This folder contains the mappers of the application. Mappers are used to transform data between layers. Detailed information is available in the README inside the folder.

### [`src/app.env`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/app.env)

This file contains the environment variables of the application.

### [`src/main.go`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/main.go)

This file contains the main function of the application.