---
title: Guía de Construcción de un Kernel Personalizado.
date: 2007-07-12
author: Rodrigo Amaya
tags: linux, tutorial, kernel, guia
post_id: blog-3515952828243908885.post-4817630226343999299
---

tipo: guía/tutorial.

1.0 Construyendo un Kernel Personalizado, a la manera tradicional (o como un Vikingo Leñador).

Esta parte especifica describe el proceso de creación de un Kernel personalizado (o que se volverá así a medida que veas que es un proceso para nada complicado). Abarca desde la descarga de un Kernel (donde encontrarlo), usar un parche en nuestro Kernel, el software necesario para compilar el Kernel, la Configuración y la creación de un "initram" (no se preocupen, ya llegaremos a eso). Todos los pasos se realizan suponiendo que el lector tiene Debian Etch instalado como sistema operativo y GRUB como bootloader (gestor de arranque) instalado. Al final de esta guía/tutorial, se espera que termines con un Kernel (por lo menos ligeramente) personalizado...

[![](http://bp1.blogger.com/_ayvorITawE4/RpkGVMQ68yI/AAAAAAAAAZI/VwAcVCuhLkU/s400/iconos-tux.jpg)](http://bp1.blogger.com/_ayvorITawE4/RpkGVMQ68yI/AAAAAAAAAZI/VwAcVCuhLkU/s1600-h/iconos-tux.jpg)"La belleza del Kernel de Linux, esta en su personalización"

1. 1 Instalando los paquetes requeridos para la compilación del Kernel.

Primero actualizamos la lista de paquetes instalables en el sistema:

> apt-get
> update
Luego instalamos los paquetes

- kernel-package
- libncurses5-dev
- fakeroot
- wget
- bzip2
- build-essential
...con el siguiente comando:

> apt-get install kernel-package
> libncurses5-dev fakeroot wget bzip2 build-essential
1.2 Descargando el código fuente del Kernel

[![](http://bp0.blogger.com/_ayvorITawE4/Rpj6oMQ68uI/AAAAAAAAAYo/SVlNzB3eyeA/s400/korg10yr.gif)](http://bp0.blogger.com/_ayvorITawE4/Rpj6oMQ68uI/AAAAAAAAAYo/SVlNzB3eyeA/s1600-h/korg10yr.gif)"10 años de Kernel.org"

En el sitio [http://kernel.org](http://kernel.org/) se publica el codigo del kernel mas reciente, versiones anteriores del mismo y también parches respectivos. Al momento de hacer esta guía, el Kernel esta en su version estable 2.6.22.1 (este es el Kernel que sera usado). Pero antes de descargarlo vamonos al directorio "/usr/src":

> cd /usr/src
y ahora:

> style="font-family:verdana;">wget
> http://kernel.org/pub/linux/kernel/v2.6/linux-2.6.22.1.tar.bz2
O lo puedes descargar con Firefox o Kget o d4x... cualquiera que sea tu gestor de descargas, lo importante es que el archivo este en "/usr/src", para no perdernos. Ahora vamos a desempaquetar el Kernel y realizar un "acceso directo" (no me tiren tomates) o "symlink" al directorio creado por el proceso de desempaquetado del Kernel:

> style="font-family:verdana;">tar xvf linux-2.6.22.1.tar.bz2
> ln -s linux-2.66.22.1 linux
> cd linux
Ahora estamos en "/usr/src/linux", compruebalo con el comando "pwd".

1.3 Aplicando un parche al código del Kernel

Más tarde pongo como parchar un Kernel, como estamos usando la versión 2.66.22.1 estable, no hay necesidad (aun de parcharlo), cuando salga el primer parche, pongo el proceso.

1.4 Configurando el Kernel

Este es el proceso que mas interacción exige del usuario en la compilación del Kernel. Y quizás la mas temida también. Este proceso consiste en decirle al proceso de compilación, que módulos crear y que no, que hardware soportara o no nuestro Kernel, que características especiales tendrá, etc. Es decir, aquí es donde lo personalizamos, lo afinamos, e integramos con el hardware de nuestra PC. Este proceso genera un archivo llamado ".config", [nótese que es un dotfile o archivo oculto](http://en.wikipedia.org/wiki/Hidden_file#Unix_and_Unix-like).

Como consejo, es buena costumbre usar la configuración del Kernel que tienes funcionando en este momento, para tener una solida base para la personalizacion de mismo. Este paso es opcional; para usar la configuración de tu Kernel actual (y funcional), basta con copiarlo a la carpeta /usr/src/linux con el nombre .config:

> style="font-family:verdana;">cp /boot/config-`uname -r`
> ./.config
uname es un programa de consola que imprime la información del sistema, uname con el argumento "-r" imprime la versión del Kernel actual. El comando anterior, suponiendo que tienes un Kernel 2.6.18-4-686, seria interpretado de la siguiente forma:

> cp /boot/config-2.6.18-4-686
> /usr/src/linux/.config
Ahora podemos comenzar con la configuración personalizada con:

> make
> menuconfig
make se encargara construir el programa menuconfig y mostrarlo, menuconfig lista de opciones (módulos, características, etc) disponibles para el Kernel y los guardada en ".config" (lo lee y lo sobre escribe con las personalizaciones hechas). Existen otras opciones para modificar el archivo ".config" (con un editor de texto por ejemplo ja ja) por ejemplo con:

- xconfig, que usa un front end con Qt
- gconfig, que usa un front end con GTK
Pero menuconfig es la manera de hacerlo este trabajo como un vikingo leñador, y no requiere mayor cosa que los libncurses5, asi que ¡adelante!

[![](http://bp1.blogger.com/_ayvorITawE4/Rpj6ecQ68tI/AAAAAAAAAYg/Yqf9Rdofeqc/s400/makemenuconfig.png)](http://bp1.blogger.com/_ayvorITawE4/Rpj6ecQ68tI/AAAAAAAAAYg/Yqf9Rdofeqc/s1600-h/makemenuconfig.png)"Captura de pantalla de menuconfig"

Ahora vamos a seleccionar la opción: Load an Alternate Configuration File.

[![](http://bp3.blogger.com/_ayvorITawE4/Rpj7E8Q68vI/AAAAAAAAAYw/Mf08kgHqKnQ/s400/loadalter.png)](http://bp3.blogger.com/_ayvorITawE4/Rpj7E8Q68vI/AAAAAAAAAYw/Mf08kgHqKnQ/s1600-h/loadalter.png)"Captura de pantalla, selección de la opción Load an Alter Configuration File"

Y nos aseguramos de que este el nombre del archivo ".config":

[![](http://bp2.blogger.com/_ayvorITawE4/Rpj7MsQ68wI/AAAAAAAAAY4/fHHK0s2oA_Q/s400/loadconfig.png)](http://bp2.blogger.com/_ayvorITawE4/Rpj7MsQ68wI/AAAAAAAAAY4/fHHK0s2oA_Q/s1600-h/loadconfig.png)"Captura de pantalla, nombre del archivo de configuración"

Cuando salgas, si has cambiado alguna configuración y no has salvado el ".config", te aparecerá un "cuadro de dialogo" para salvar el archivo. Hay que aclarar, que el proceso de compilación del Kernel espera que exista un archivo con el nombre ".config", si este no existe, entonces no podremos continuar.

En resumen...

> style="font-weight: bold;">Tiene que existir un archivo llamado ".config" en la carpeta del
> código fuente del Kernel ¡¡¿ok?!?
1.5 Construyendo e Instalando el nuevo Kernel

Los siguientes comandos son para construir el Kernel (recuerda que make es el que hacer la magia):

> make all
> make modules_install
> make
> install
Ahora seamos pacientes, este proceso puede tomar desde algunos minutos, hasta un par de horas. Esto depende de la configuración que hiciste (si habilitaste todas las opciones como un perfecto demente) y la velocidad de tu procesador. Si deseas saber que otras opciones puedes usar para el make, puedes hacer

> make
> help
en "/usr/src/linux", la carpeta del código del Kernel, claro.

1.6 Proceso posterior a la instalación

El nuevo Kernel ahora esta instalado en tu sistema, pero necesita un "ramdisk", si no lo tiene, lo mas probable es que ¡el sistema no arranque!

[![](http://bp0.blogger.com/_ayvorITawE4/Rpj77MQ68xI/AAAAAAAAAZA/06qgK1kKzKA/s400/kernelpanic.png)](http://bp0.blogger.com/_ayvorITawE4/Rpj77MQ68xI/AAAAAAAAAZA/06qgK1kKzKA/s1600-h/kernelpanic.png)"Oh no!... Kernel Panic"

Además necesitamos actualizar el GRUB para que muestre el nuevo Kernel. Bien, lo primero es hacer el "ramdisk", asumiendo que tenemos el Kernel :

> style="font-family:verdana;">depmod 2.6.22.1
Este comando realiza un mapa de las dependencias de los módulos que usa tu Kernel, basándonos en este mapa podemos crear un ramdisk que tenga todos los modulosa necesarios que usa nuestro Kernel. Para mas información de depmod ... [RFTM](http://es.wikipedia.org/wiki/RTFM).

Suponiendo que no tienes instalado mkinitramfs, evidentemente habrá que instalarlo...

> apt-get install
> initramfs-tools
Y ahora que tenemos instalado mkinitramfs, hagamos nuestro ramdisk con:

> style="font-family:verdana;">mkinitramfs -o /boot/initrd.img-2.6.22.1
> 2.6.22.1
Esto genera el ramdisk con respecto a los módulos específicos compilados para nuestro Kernel. Ahora actualizaremos GRUB con:

> update-grub
Esto detectara el nuevo Kernel y ramdisk creado y actualizara el archivo "/boot/grub/menu.lst" por nosotros.

1.7 Últimos pasos...

Reiniciar el sistema:

> shutdown -r
> now
En el menú del grub, selecciona tu nuevo Kernel. Si todo sale bien, ya esta funcionando tu PC con tu nuevo Kernel, ¡felicidades! Podes verificar si "realmente" estas usando tu nuevo Kernel con:

> uname -r
Si tu sistema no arranca, entonces puedes re iniciar la computadora, y seleccionar el Kernel antiguo (y funcional) que tenias, para modificar los cambios.

1.8 Comentarios Finales

Si bien la etapa de configuración/compilación/re-inicar/arrancar etc... parece que no puede ser más corta, existe una interesante, educativa y "mejor" de probar el Kernel recién compilado sin re-iniciar la PC. Esta es la forma del Vikingo Inteligente (por decirle de alguna forma) que examinaremos y explicaremos en algunos días.

1.9 Agradecimientos

Gracias a [Falko Timme por su guía de compilación de Kernel](http://www.howtoforge.com/kernel_compilation_debian_etch_p2?s=2da22972170ef10fb4fdb8a327b463a4&), su guía fue usada como base para la realización de este documento, y gracias a ti por leer esta guía. Se esperan sugerencias y cualquier comentario es bienvenido.

Saludos y hasta la próxima!