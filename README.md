# Curso de GO para aplicaciones y web api RESTful

Curso de udemy de diciembre 2020 para el desarrollo y conclusión de un proyecto de Isw y de modo de dar un mejor uso de webservices en Go y posiblemente después con otros frameworks de webservices

## Qué es Go?
* Es un lenguaje de programación opensource de Google, creado en 2009
* Este lenguaje es orientado para ser un lenguaje optimo en tiempo de ejecución para enfoques de desarrollo de aplicaciones para internet \(Backends\)
* Este lenguaje contiene un tipado estático \(tipado fuerte\)
* Es un lenguaje compilado para poder ser ejecutado

## Ventajas
* Sencillo de escritura y de lectura con similitudes a lenguaje C
* Soporta miles de conexiones simultaneas \(Caso de web / Backend\)
* Velocidad y rendimiento optimo
## Para qué siver?
* Crear un script específico para una funcionalidad específica en OS como un demonio en sistema
* Backend / API RESTful para peticiones como servidor
* Sockets para aplicaciones de conexión Cliente-Servidor
## Instalación

__Caso Windows__

Para este caso es necesario que se haga uso de alguna terminal que tenga compatibilidad o comportamiento como Linux el cual puede ser instalado desde la tienda de Windows y puede instalar paquetes de Linux, los cuales tienen una compatibilidad a sistemas Debian \(Ubuntu\) por si es necesario hacer uso de Go a nivel de Windows visitar la página para bajar el compilador y el lenguaje, mientras que en caso contrario pasar a la siguiente parte que es el caso de Linux.

[Go]\(https://golang.org/dl/\)

__Caso Linux__

* Sistemas basados en RedHat
```bash
  sudo dnf install golang -y
```
* Sistemas basados en Debian o que no tengan repositorio con golang
  * Ir a la pagina principal de golang para bajar la versión deseada [Go]\(https://golang.org/dl/\)
  * Con ello ejecutaremos lo siguiente
```bash
  # Ejemplo con la ultima version disponible
  wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
  sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
  sudo echo "PATH=$PATH:/usr/local/go/bin" | sudo tee -a /etc/profile
  sudo source /etc/profile

  # Finalmente escribir
  go version
```

Para finalmente si es con uso de Windows el que exista actualmente la terminal de Linux nos permite instalar con las líneas de comando de Debian el lenguaje de Go, pero preferentemente usa Linux...

## Archivos y extensiones de archivo

Para poder trabajar con este lenguaje de programación es necesario tener conocimiento que nuestros archivos los guardaremos con extensiones de archivo .go los cuales nos permitirán trabajar más tarde

## Hola mundo primer código

Para realizar el sencillo y clásico hola mundo podemos escribir lo siguiente

```go
  /* Archivo: hola.go */
  package main  /* Paquete principal del lenguaje para el inicio */
  import "ftm"  /* Primer paquete de salida estandar de datos */

  func main\(\){
    ftm.Printl\("Hola mundo desde el curso en Github"\) /* Entre comillas */
  }
```
Para ejecutar esto simplemente podemos hacer lo siguiente en nuestra consola:
```bash
  go run hola.go
```
En el cual desplegara nuestro código escrito y así podremos trabajar desde la consola donde aparentemente solo se ejecuto nuestro código.
### Varios imports a la vez
Para aplicaciones muy grandes de desarrollo muchas veces necesitaremos más de un import el cual puede ser consumido en nuestra aplicaciones por lo que haremos uso de la siguiente forma para realizar imports
```go
  /* Usa el import ftm second_import y third_import respectivamente */
  import \("ftm"
  "second_import"
  "third_import"\)
```
### Algunos comandos útiles a la hora de desarrollar
#### Ayuda con algunos imports
Así como el comando todo poderoso de Linux _man_ existe un comando específico para desplegar y analizar la documentación de cierto _import_ realizado así como los parámetros que puede recibir, para desplegar esto solo tendremos que escribir:
```bash
  go doc <package_import>
```
En el cual nos desplegara información importante relacionada con nuestro modulo importado y que otras funciones nos pueden servir.
#### Auto-format file
Así como el IDE Netbeans que te permite dar un autoformat a tu archivo fuente de código, también existe una forma rápida y sencilla de dar un formato automático a nuestros archivos de go, solo bastara realizar los siguiente:
```bash
  go fmt <archivo.go>
```
En donde obtendremos el mismo archivo pero con un mejor formato de identado o sangría para una mejor lectura de nuestro código.
#### Compilado de archivos Go
Si bien al realizar _go run_ podemos ver que ejecutamos nuestro código fuente, también podemos compilar lo que deseamos trabajar con el fin de generar un binario ejecutable para nuestra arquitectura deseada solo hará falta escribir lo siguiente:
```bash
  go build <archivo.go>
```
El cual nos generara un archivo binario compatible con nuestra plataforma de desarrollo.
## Variables
Al ser un lenguaje de tipado fuerte, es decir, debemos especificar el tipo de dato con el que vamos a trabajar en nuestros programas, existen diversas formas de declarar variables en go, y de distintos tipos entre los tipos más comunes de variables tenemos:
```go
  int         /* -> Valores enteros */
  bool        /* -> Valores booleanos */
  float32     /* -> Valores reales */
  byte        /* -> Valores de un byte */
  string      /* -> Valores de cadena de caracteres */
```
## Tipos de datos
Tomando en cuenta algunas ideas de distintos sitios tenemos que el tipo de variables puede ser más extendido:

| Tipo | Descripción |
| ---- | ----------- |
| Booleano | Son variables que pueden tomar valores como _Verdadero_ o _Falso_ |
| Numérico | Son variables que pueden tomar valores de tipo entero <img src="https://render.githubusercontent.com/render/math?math=\mathbb{Z}"> o puede tomar valores reales <img src="https://render.githubusercontent.com/render/math?math=\mathbb{R}"> |
| String | Son variables que pueden ser consideradas como una cadena de caracteres que nos permite resguardar valores escritos por el usuario o con los que estemos trabajando |
| Derivados | Son variables que pueden ir desde _Apuntadores_, _Arreglos_, _Estructuras de datos_, _Uniones_, _Funciones_, _Slides_, _Interfaces_, _Mapas_, _Canales_ |

### Subtipos de variables y tipos de dato
En go al igual que en lenguajes como C, C++, Java que necesitan nombrar el tipo de dato de la variable que van a trabajar existen subvariaciones o subtipos de un mismo tipo de dato, para ello esta escrito la siguiente tabla
#### Entero
| Tipo de dato | Valores que puede tomar y bits |
| ------------ | ------------------------------ |
| uint8        | Entero sin signo de 8 bits \( 0 a 255 \) |
| uint16       | Entero sin signo de 16 bits \( 0 a 65535 \) |
| uint32       | Entero sin signo de 32 bits \( 0 a 4294967295 \) |
| uint64       | Entero sin signo de 64 bits \( 0 a 18446744073709551615 \) |
| int8         | Entero con signo de 8 bits \( -128 a 127 \) |
| int16        | Entero con signo de 16 bits \( -32768 a 32767 \) |
| int32        | Entero con signo de 32 bits \( -2147483648 a 2147483647 \) |
| int64        | Entero con signo de 64 bits \( -9223372036854775808 a 9223372036854775807 \) |
#### Reales o tipo flotante
| Tipo de dato | Características |
| ------------ | --------------- |
| float32      | Flotante de 32 bits en la especificación IEEE-754 |
| float64      | Flotante de 64 bits en la especificación IEEE-754 |
| complex64    | Número complejo con parte real e imaginario de 32 bits respectivamente |
| complex128   | Número complejo con parte real e imaginario de 64 bits respectivamente |
