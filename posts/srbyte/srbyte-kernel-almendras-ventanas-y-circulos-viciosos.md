---
title: Kernel, Almendras, Ventanas y Circulos Viciosos...
date: 2007-07-08
author: Rodrigo Amaya
tags: libre, gnu, virtualizacion, linux, software, kernel
post_id: blog-3515952828243908885.post-1172145899024398733
---

En el oscuro y ofuscado mundo de los Sistemas Operativos, un Kernel (también conocido en español como núcleo) es el componente central que se responsabiliza en manejar los recursos del sistema (manejo de memoria por ejemplo) y la comunicación entre el Hardware (parte física de una computadora) y el Software (programas).

Como componente básico de cualquier sistema operativo, un Kernel provee el nivel mas bajo en la capa de abstracción antes de llegar al Hardware... algo mas "abajo" del sistema operativo... son unos y ceros... o cambios voltaje.

La tarea del Kernel es ejecutar procesos (Software, aplicaciones, programas, como le quieran decir) y facilitarles a los mismos la tarea de interacción con el Hardware, mediante la comunicación entre procesos y llamadas al sistema (system calls).

Existen una gran variedad de Kernels: Micro Kernel y Mono Kernel y además esta el ExoKernel, el Kernel Híbrido y el Nano Kernel...

![image](https://bp3.blogger.com/_ayvorITawE4/RpJTN1X49XI/AAAAAAAAAYY/KiFx4OD18RE/s400/kernel.jpg)    
"Macro (Mono), Micro y Exo
Kernel... el concepto"

...el diseño y creación de estos obedece usualmente a la implementación de los "anillos de protección". Estos anillos obedecen a un mecanismo que busca proteger datos y funcionalidad: de fallos y de comportamiento malicioso.

Los dos tipos de Kernel más usados son: Micro Kernel y Mono Kernel... y entre ellos existen un sin fin de implementaciones. Esto es como decir que solo existen Negros y Blancos... mentira, hay morenos, mulatos, cheles, chilucos, etc... Lo mismo es con las diversas implementaciones que existen entre el Micro Kernel y el Mono Kernel.

La forma mas fácil de imaginar un Kernel es como una almendra...

![image](https://bp2.blogger.com/_ayvorITawE4/RpJHDlX49SI/AAAAAAAAAXw/jB-Si-zZrJE/s400/almendras.jpg)    
"Kernel en ingles, es el
centro comestible de una almendra"

Hay que mencionar, que Microsoft decidió omitir esta regla de "anillos de protección" desde Windows 3.x hasta Windows 98. ¿Por qué?... bueno, la comunicación entre los anillos de protección requiere MUCHOS ciclos de procesamiento. Y con Windows siendo un sistema operativo elementalmente visual, Microsoft no tenia mas remedio que ELIMINAR los anillos de protección para mejorar el rendimiento general del sistema. Claro, esto llevo a problemas terribles muy bien conocidos en todos los Windows de la serie 3.x y 9x, como la famosa BSOD (Pantalla Azul).

![image](https://bp0.blogger.com/_ayvorITawE4/RpJHJFX49TI/AAAAAAAAAX4/2C2NNCbJ57Q/s400/kernelwin98.jpg)    
"Así se vería el Kernel de
Windows 98 después de una BSOD"

Ahora la serie de sistemas Windows XP/Vista eliminan (enormemente) este problema implementando en su Kernel los una vez desplazados anillos de protección... que se asemejan mucho con las capas de una cebolla... pero basta de hablar de comida.

El tema del Kernel no sera muy tratado por los usuarios de Windows... pero crean cuando les digo que es un tema muy interesante y elegido como tema de conversación (bueno, ni tanto) por muchos usuarios de GNU\Linux... y no es solo interesante, sino también controversial y [¡extenso!](https://en.wikipedia.org/wiki/Image:Unix-history.svg).

Lo "bonito" del Kernel de Linux, es que como es Open Source, todo mundo sabe como funciona y esto permite algunas características muy interesantes a la hora del famoso proceso de "compilación del Kernel".

Compilar el Kernel en un sistema, no se trata de una tarea exclusiva para "geeks", basta con leer un poco y entender la dinámica. Lo que si aburre, es el TIEMPO que tarda en compilarse, y sumado a esto, tener que re-iniciar la maquina para probar el nuevo kernel, y sino funciona, hacer el proceso de nuevo.

Pues para este circulo vicioso, existe un método que simplifica MUCHO las cosas, y además, sirve de ejemplo perfecto de la flexibilidad que presenta el Software Libre. Este método consiste en la utilización de una maquina virtual para probar nuestro Kernel, y definitivamente NO consiste en copiar el kernel a la maquina virtual y arrancarla o peor aun, compilar el kernel en una maquina virtual, NO NO NO! Es mucho mejor que eso (y menos tedioso), consiste en arrancar la imagen de un disco duro virtual con un Kernel (el que hemos compilado) EXTERNO a esta maquina virtual.

![image](https://bp2.blogger.com/_ayvorITawE4/RpJO_lX49WI/AAAAAAAAAYQ/1BVmDQ3g530/s400/qemu-logo.png)    
"Logo de
QEmu"

Esta fantástica característica parte de [QEMU (QEmu es un emulador de computadoras virtuales)](https://fabrice.bellard.free.fr/qemu/), se llama "Direct Linux Boot":

![image](https://bp2.blogger.com/_ayvorITawE4/RpJLvlX49UI/AAAAAAAAAYA/7fgT_RPSTSM/s400/linuxboot.png)    
"Captura de pantalla de Qemu
Laucher con Direct Linux Boot habilitado"

Direct Linux Boot, permite ahorrar tiempo de una manera increíble, al evitar estar re-iniciando la maquina constantemente y evitando el infame "Kernel Panic".

![image](https://bp2.blogger.com/_ayvorITawE4/RpJMRlX49VI/AAAAAAAAAYI/PYIEc1YPXgw/s400/instant%C3%A1nea1.png)    
"Captura de pantalla de
Direct Linux Boot con Kernel EXTERNO 2.6.21rac Click para hacerlo más grande"

En el transcurso de la semana escribo un tutorial sobre como compilar y probar un Kernel usando QEMU.

Hasta