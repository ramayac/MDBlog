---
title: Skype para Ubuntu de 64bits
date: 2009-04-11
author: Rodrigo Amaya
tags: skype, 64bits, ekiga, ubuntu
post_id: blog-3515952828243908885.post-2361935392325699883
---

Hace un par de semanas, finalmente me canse de estar usando Windows en la laptop del trabajo, así que decidí sacar de su miseria a la pobre maquina, ser infiel al sistema operativo de los gamers y ponerle "otra cosa". La "afortunada" es una Compaq CQ45:

![image](https://1.bp.blogspot.com/_ayvorITawE4/SeCqtek1PKI/AAAAAAAAB74/svXCN-J6Q7I/s320/CQ50.jpg)    
"No es una maravilla, pero
es una buena maquina."

Mi única duda, era si ponerle Debian o Ubuntu de 64bits... recordé las palabras de un blogger Hondureño: "Debian es para Desktops, y Ubuntu es para Laptops." Sinceramente, para mi son lo mismo, pero tenia el disco de Ubuntu a la mano. Para no aburrirlos la laptop ya tiene Ubuntu. Y antes de que se haga un flamewar sobre que distro es mejor, dejen me compartir una idea que siempre he tenido bien clara:
>
> "Sinceramente, tener una u otra distro NO IMPORTA, lo que importa es que la maquina tenga UN
> SISTEMA OPERATIVO DE VERDAD (LINUX, HURD, OPENSOLARIS, BSD) y que este sistema sea
> LIBRE".
¿Sabían que el disco de Ubuntu de 64bits no tiene modo "live CD"? usa la instalación en modo texto, lo que para mi es fenomenal - personalmente prefiero ese tipo de instalaciones - pero hay que aceptarlo, no es muy "User Friendly".

Uno de los percances que tuve con respecto a esta migración (de 32 bits a 64 bits) fue con un programa muy útil que estoy usando mucho últimamente: Skype. Como no me pagan para hacer publicidad sobre Skype, lo mejor que puedo hacer es dejar que la Wikipedia [lo explique](https://en.wikipedia.org/wiki/Skype), y dejar un vinculo a su [pagina web](https://www.skype.com/).

Mi problema es que Skype sólo está disponible en paquetes para arquitectura 32 bits, yo demando que mis programas usen por completo la potencia de mis procesadores. Si alguien más se encontró con este problema, lo mejor que puede hacer es seguir esta sencilla guía para instalar Skype de 32 bits, en Ubuntu de 64 bits:

1. Necesitaremos descargar versiones de 32 bits de algunas librerías para que Skype no nos de ningún error al momento de iniciar. Todo se resume a ejecutar este comando:

> sudo
> apt-get install ia32-libs ia32-libs-gtk ia32-libs-sdl lib32asound2
2. Ahora de clic a este vinculo y descarguen el paquete: [https://www.skype.com/go/getskype-linux-ubuntu-amd64](https://www.skype.com/go/getskype-linux-ubuntu-amd64). Ojo, este NO es un paquete de 64bits, es solo un paquete que tiene las dependencias correctas para Ubuntu 8.04 o mayor (de 64bits).

3. Bien, si falta alguna dependencia, pues hay que resolverla, y despues de eso, solo nos falta Skype. Debido a la variedad de Hardware y servidores de audio, la tarea de configuración de audio y video, queda a discreción del lector, me exonero de "ese asunto".

![image](https://2.bp.blogspot.com/_ayvorITawE4/SeCuv0vAxjI/AAAAAAAAB8A/sQTbI4_0sBc/s320/skypeme_big.png)    
"Ahora ya tienes listo tu
Skype :)"

Bien, espero que esta pequeñisima guía le ayude a algun internauta por ahi, a instalar esta util aplicacion. Cabe mencionar que tambien existen aplicaciones 100% libres, que tiene casi la misma utilidad que Skype, como son: [Ekiga](https://www.gnomemeeting.org/) (antes conocido como GnomeMeeting), despues hablaremos un poco del porque usar Software Libre, en vez de Freeware.