# crud_go
working with database in go lang


## running

To run this program do:

```
go mod tidy
go build
./crud_go
```
and if everything works well, it should output:

```
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
