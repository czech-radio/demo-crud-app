# demo-crud-app

__Demo application demonstrating how to work with database in Go language.__

## Installation

To build and run this program do:

```shell
go mod tidy
go build
./crud_go
```

If everything works well, it should output:

```shell
&{1 Jan Novak BEZPP 1 0 soustruznik}
&{2 Pepa Dvorak BEZPP 1 0 fyzik}
CREATE SUCCESSFUL
[{1 Jan Novak BEZPP 1 0 soustruznik} {2 Pepa Dvorak BEZPP 1 0 fyzik}]
&{2 Pepa Dvorak BEZPP 1 0 fyzik}
RETURN SUCCESSFUL
[{1 Jan Novak Pirati 1 0 soustruznik} {2 Pepa Dvorak BEZPP 1 0 fyzik}] <nil>
UPDATE SUCCESSFUL
[{2 Pepa Dvorak BEZPP 1 0 fyzik}] <nil>
[] <nil>
DELETE SUCCESSFUL
```
