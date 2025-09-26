---
title: Hilos y Sincronización...
date: 2009-10-26
author: Rodrigo Amaya
tags: threads, programacion, hilos
post_id: blog-3515952828243908885.post-3802782752421137828
---

En el trabajo, existen un buen numero de componentes que fueron ideados para ser reusables o "genéricos", el diseño de estos a veces funciona y otras... digamos que dejan mucho que desear. Pero existe un componente que resulta interesante (más nada practico en lo personal, a pesar de que se esta usando), que terminan empleando todas las aplicaciones, y en pocas palabras, se encarga de determinar la configuración adecuada a emplear de acuerdo al ambiente (ruta e IP) en el que se encuentra, nos referiremos a este componente con el nombre ficticio de: "ConfigUtility". Para usar este componente en X aplicación, se instancia un objeto "ReadConfiguration", y se especifica de que aplicativo queremos la información, luego el ReadConfiguration, leerá de un archivo XML la configuración adecuada a la aplicación especificada.

Existen dos problemas fundamentales con el ConfigUtility... El primero, es que cada vez que se realiza una instancia de este objeto, se realiza la lectura y conversión de un archivo XML a un objeto. El segundo es que en un segundo, un proceso puede invocar hasta seis veces el ConfigUtility, lo que se convierte en seis lecturas del mismo archivo... y si son 100 usuarios los que usan ese proceso en un determinado instante (algo muy probable), entonces son 600 lecturas en un segundo de ese archivo.

Se me ocurrió mejorar los tiempos de carga del archivo realizan un "[cache](http://en.wikipedia.org/wiki/Cache)" del objeto en memoria, y modificar este hasta que el archivo XML cambie en el disco. Una buena idea cuando se prueba con una sola persona, pero que se convierte en un caos con muchas usuarios.

[![](http://1.bp.blogspot.com/_ayvorITawE4/SuUNCTHQvmI/AAAAAAAACNE/C-X8UEXltno/s320/threads.jpg)](http://1.bp.blogspot.com/_ayvorITawE4/SuUNCTHQvmI/AAAAAAAACNE/C-X8UEXltno/s1600-h/threads.jpg)

¿La razón del fracaso? los hilos de ejecución ([Threads](http://www.javaworld.com/javaworld/jw-04-1996/jw-04-threads.html)).

Un hilo de ejecución, es una característica en los procesadores que permite a una aplicación realizar varias tareas a la vez (concurrentemente), especialmente en los ambientes multi-núcleo modernos que se vuelven más y mas comunes. Los distintos hilos de ejecución comparten una serie de recursos tales como el espacio de memoria, los archivos abiertos, situación de autenticación, etc. Así una aplicación puede llevar a cabo distintas tareas simultáneamente por ejemplo: Writer de OpenOffice.org y Word, emplean hilos para ejecutar el proceso del corrector ortográfico mientras se escribe.

A la sumatoria de los hilos de ejecución y los recursos que estos comparten se les conoce como: [proceso](http://en.wikipedia.org/wiki/Process_%28computing%29) (Lo que ves cuando ejecutas Top en el bash, o el TaskManager en Windows). El hecho de que los hilos de ejecución de un mismo proceso compartan los recursos hace que cualquiera de estos hilos pueda modificar éstos. Cuando un hilo modifica un dato en la memoria, los otros hilos pueden acceder a ese dato modificado inmediatamente. En el ConfigUtility (antes de la dichosa mejora) cada objeto pertenece a un hilo, y este realiza la lectura del archivo. Si existen muchas lecturas al archivo, cada hilo correspondiente se espera a que el otro hilo suelte el archivo para poder leerlo completamente.

En el nuevo ConfigUtility, como la información del archivo en cuestión se mantiene en memoria como objeto (y esta información se comparte entre todas las instancias del objeto), la única vez que esta información pueda cambiar, es cuando el archivo XML real sea modificado. Sin este cambio la lectura del archivo era de 1.6 segundos, con el cambio solo la primera lectura tiene ese tiempo y las lecturas/invocaciones consecutivas retornan en un tiempo de 234 milisegundos.

[![](http://2.bp.blogspot.com/_ayvorITawE4/SuUN2fwNJOI/AAAAAAAACNU/aubU9HuHolA/s320/runforrestrun.jpg)](http://2.bp.blogspot.com/_ayvorITawE4/SuUN2fwNJOI/AAAAAAAACNU/aubU9HuHolA/s1600-h/runforrestrun.jpg)

"Run ... new instance of ReadConfiguration, Run!"

Pero los hilos comenzaron a molestarse, ya que en determinados momentos y circunstancias todos querían verificar si el archivo había cambiado, al mismo tiempo (lo que genera un error de concurrencia) y todos querían emplear el objeto en memoria (objeto que podía contener información corrupta, porque la lectura y escritura a este no era Atómica). Claro, acá falta explicar muchas cosas sobre como es el código en cuestión, [atomicidad](http://es.wikipedia.org/wiki/Atomicidad), [concurrencia](http://en.wikipedia.org/wiki/Concurrency_%28computer_science%29), locks, mutex y demás, pero se entiende la idea (bien) general (quizás más adelante escriba un poco sobre lo mencionado).

[![](http://2.bp.blogspot.com/_ayvorITawE4/SuUNGQUN_BI/AAAAAAAACNM/Dgem8LKcuqs/s320/578px-Dining_philosophers.png)](http://2.bp.blogspot.com/_ayvorITawE4/SuUNGQUN_BI/AAAAAAAACNM/Dgem8LKcuqs/s1600-h/578px-Dining_philosophers.png)

"El dilema de los Filosofos (Concurrencia)"

¿La solución? Sincronizar los hilos. Al menos la solución sencilla de implementar, consiste en emplear la palabra reservada de Java: Synchronize, para asegurarnos que un bloque de código (o un método completo) sea "[Thread Safe](http://en.wikipedia.org/wiki/Thread_Safe)", es decir, que en ese preciso bloque, los hilos harán "fila" para usarlo.

La lección:

Programar con la idea de la concurrencia en mente...  no solo cuando haces algo nuevo, sino que también cuando modificas algo existente. Hay que recordar, que en un ambiente con muchos clientes conectados que pueden emplear el mismo método en un determinado instante, puede suceder que treinta hilos corren salvajemente a utilizar un recurso critico, lo que podría resultar en un pequeño infierno de dudas, incertidumbres y defectos extraños, todo porque tu lógica no es "Thread Safe".

¡No se olviden de sincronizar sus hilos! Saludos :)