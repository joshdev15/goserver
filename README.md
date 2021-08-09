# Heroku 

Esta rama representa una version de prueba de este software y mi repositorio Curriculum-Vitae para el despliegue en 
la plataforma Heroku

# GoServer

Servidor de archivos estaticos configurable hecho en Go

El servidor se puede correr directamente desde el codigo o compilarlo para usarlo casi en cualquier lugar.

Como el lenguaje Go lo permite, es posible compilar un ejecutable para Linux, MAC, y Windows, pudiendo correr
el mismo desde la terminal de su ordenador.

## Como ejecutar?

Primeramente es necesario clonar el proyecto a su ordenador.

Luego para ejecutar un servicio dirigirse a la ubicacion donde se encuentra el proyecto, por ejemplo:

```bash
$ cd /home/user/goserver
```

### Desde el Codigo

Si desea ejecutar el servidor desde el codigo solo es
necesario ejecutar el siguiente comando:

```bash
$ go run main.go
```

### Desde un Ejecutable

Si por el contrario necesita el ejecutable para no tener que buscar constantemente el codigo,
debe construir el programa con el siguiente comando:

```bash
$ go build main.go
```

Esta accion generará un ejecutable, el formato depende del sistema operativo, en Linux y MAC será un binario,
en Windows un archivo exe.

Una vez que se tiene el ejecutable solo es necesario correrlo en donde lo desee, por ejemplo: (Linux)

```bash
$ ./main
```

## Parametros

Bien sea compilado o desde el codigo, el servidor permite la personalizacion del puerto y el folder que se desea servir.

Los comandos son los siguientes:

### Puerto

Se puede configurar el puerto con el siguiente comando:

__Desde el Codigo:__ (Linux)

```bash
$ go run main.go -port 9000
```

__Desde el Ejecutable:__ (Linux)

```bash
$ ./main -port 9000
```

### Ruta / Directorio

Se puede configurar el directorio o ruta que se quiere servir con el siguiente comando:

__Desde el Codigo:__ (Linux)

```bash
$ go run main.go -path ./myCustomFolder
```

__Desde el Ejecutable:__ (Linux)

```bash
$ ./main -path ./myCustomFolder
```

### Ambos Comandos

Tambien es posible ejecutar ambos comandos a la vez

```bash
$ ./main -port 7000 -path ./myCustomFolder
```

## Recomendaciones

Se recomienda generar el ejecutable y colocarlo en o cerca de la carpeta que se desea servir
para que la el uso sea mas facil, rapido y efectivo.

## Adicionales

En el codigo se encuentra una carpeta de prueba llamada dist, que contiene un archivo, "index.html",
que sirve para probar que el servidor funciona correctamente.
