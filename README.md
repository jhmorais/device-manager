# Device Manager

REST service that supports the management of a device database.

## 💻 Required

* go version 1.18

## 🚀 Installing Device Manager

To install the Device Manager, follow the steps:

Linux, MacOS e Windows:
```
go mod tidy
```

## ☕ Using Device Manager

To use Device Manager, follow the steps:

Create the volume
```
make create-volume
```

Start the containers with docker compose
```
make compose-up
```

Enter inside the container of application
```
make docker-exec
```

Run the rest server
```
make run
```

Now it is possible access the application by the port 8088, follow the exemple to list all devices:
 ```
localhost:8088/devices
 ```

## ☕ Testing Device Manager

To run all the unit tests, follow the step:

```
make test
```

[⬆ back to top](#Device Manager)<br>