# Curso de GO para aplicaciones y web api RESTful

Curso de udemy de diciembre 2020 para el desarrollo y conclusión de un proyecto de Isw y de modo de dar un mejor uso de webservices en Go y posiblemente después con otros frameworks de webservices

## Qué es Go?
* Es un lenguaje de programación opensource de Google, creado en 2009
* Este lenguaje es orientado para ser un lenguaje optimo en tiempo de ejecución para enfoques de desarrollo de aplicaciones para internet (Backends)
* Este lenguaje contiene un tipado estático (tipado fuerte)
* Es un lenguaje compilado para poder ser ejecutado

## Ventajas
* Sencillo de escritura y de lectura con similitudes a lenguaje C
* Soporta miles de conexiones simultaneas (Caso de web / Backend)
* Velocidad y rendimiento optimo
## Para qué siver?
* Crear un script específico para una funcionalidad específica en OS como un demonio en sistema
* Backend / API RESTful para peticiones como servidor
* Sockets para aplicaciones de conexión Cliente-Servidor
## Instalación

__Caso Windows__

Para este caso es necesario que se haga uso de alguna terminal que tenga compatibilidad o comportamiento como Linux el cual puede ser instalado desde la tienda de Windows y puede instalar paquetes de Linux, los cuales tienen una compatibilidad a sistemas Debian (Ubuntu) por si es necesario hacer uso de Go a nivel de Windows visitar la página para bajar el compilador y el lenguaje, mientras que en caso contrario pasar a la siguiente parte que es el caso de Linux.

[Go](https://golang.org/dl/)

__Caso Linux__

* Sistemas basados en RedHat
```bash
  sudo dnf install golang -y
```
* Sistemas basados en Debian o que no tengan repositorio con golang
  * Ir a la pagina principal de golang para bajar la versión deseada [Go](https://golang.org/dl/)
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

Para poder trabajar con este lenguaje de programación es necesario tener conocimiento que nuestros archivos los guardaremos con extensiones de archivo .go los cuales nos permitiran trabajar más tarde

## Hola mundo primer código

Para realizar el sencillo y clásico hola mundo podemos escribir lo siguiente

```go
  package main  /* Paquete principal del lenguaje para el inicio */
  import "ftm"  /* Primer paquete de salida estandar de datos */

  func main(){
    ftm.Printl("Hola mundo desde el curso en Github") /* Entre comillas */
  }
```
