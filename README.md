# Lab1-CES27

No diretório Process executar 

```
go build Process.go
```

para gerar o executável

No diretório SharedResource executar 
```
go build SharedResource.go
```
para gerar o executável

Em terminais distintos execute 
```
./Process {Process_id} {Process_port} {Process_port_1} {Process_port_2} ... {Process_port_n}
```

Em outro terminal, execute 
```
./SharedResource
```
