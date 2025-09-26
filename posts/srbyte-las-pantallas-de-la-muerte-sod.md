---
title: Las Pantallas de la Muerte (SOD)
date: 2008-08-07
author: Rodrigo Amaya
tags: linux, windows, mac, ipod, iphone, apple
post_id: blog-3515952828243908885.post-2342286680349823586
---

Tenia rato de no tener una "Pantalla de la Muerte" asechando en mi monitor, y eso me inspiro y motivo a realizar esta entrada. Muchas personas conocen este fenómeno como la famosa "Pantalla Azul", mas, el nombre completo es: Blue Screen of Death (Pantalla Azul de la Muerte, BSOD por sus siglas en ingles). Este terrible fenómeno, sucede cuando el sistema operativo (tipicamente Microsoft Windows, aunque nadie se salva realmente) tiene un error que requiere que todas las operaciones se detengan inmediatamente, por eso mismo también se le dice "Stop Error". Usualmente sucede cuando el error en el sistema es tan fatal, que se requiere que se apague la computadora para prevenir cualquier tipo de daño en el hardware, así que no solo sirven para causar daños psicológicos irreparables, sino que también cuidan tu PC.

En la familia de Windows NT, estos errores suceden cuando los controladores de los dispositivos están mal codificados/programados o cuando el hardware funciona incorrectamente.

[![](http://1.bp.blogspot.com/_ayvorITawE4/SJu6yyvUPYI/AAAAAAAABAc/DwxZZR3MIT4/s400/Windows_XP_BSOD.png)](http://1.bp.blogspot.com/_ayvorITawE4/SJu6yyvUPYI/AAAAAAAABAc/DwxZZR3MIT4/s1600-h/Windows_XP_BSOD.png)"BSOD de Windows NT"

En la familia de Windows 9x (95, 98, 98SE, Me, etc...) eran los DLLs el problema, o bugs en el Kernel:

[![](http://1.bp.blogspot.com/_ayvorITawE4/SJu6ysghdxI/AAAAAAAABAU/mOG-qaLiufI/s400/Windows_9X_BSOD.png)](http://1.bp.blogspot.com/_ayvorITawE4/SJu6ysghdxI/AAAAAAAABAU/mOG-qaLiufI/s1600-h/Windows_9X_BSOD.png)"BSOD de Windows 9x"

Inclusive puede suceder, que si se golpea -fuerte- el CPU, la maquina se reinicia y muestra la BSOD.

Por cierto los "amigos" de Microsoft, han incluido una característica en Windows que permite que la pantalla azul aparezca manualmente. Para habilitarla, con fines "educativos" claro, hay que añadir un valor al registro de Windows (usando regedit.exe) y seguir las instrucciones de [este sitio](http://pcsupport.about.com/od/tipstricks/ht/makebsodxp.htm). Luego, después de reiniciar la maquina, una BSOD aparecerá cuando el usuario presione la tecla SCROLL LOCK (PAUSA INTER) dos veces mientras se sostiene la tecla derecha CTRL. Como ya dije, esta característica es para fines útiles y educativos, sirve para obtener "memory dumps" de la maquina en un estado especifico. También puedes provocar un BSOD al terminar el proceso csrss.exe y winlogon.exe . Aunque si son curiosos, ya habrán notado que con el Task Manager no se puede terminar esos procesos, pero si se puede con software de terceros ;) Y por si fuera poco, también hay un salva pantallas de la famosa pantalla azul, y la pueden descargar [aqui](http://technet.microsoft.com/en-us/sysinternals/bb897558.aspx).

Bien, si dejamos el mundo de Winbugs Windows, veremos que hay muchos dispositivos también tienen esta "característica". La Xbox 360, cuando esta consola experimenta una "fallo general de hardware", tres luces rojas aparecen alrededor del botón de encendido.

[![](http://1.bp.blogspot.com/_ayvorITawE4/SJu6yzLctuI/AAAAAAAABAk/ByMKtlfxTxk/s400/Xbox360-ringofdeath.jpg)](http://1.bp.blogspot.com/_ayvorITawE4/SJu6yzLctuI/AAAAAAAABAk/ByMKtlfxTxk/s1600-h/Xbox360-ringofdeath.jpg)"Xbox 360 anillo de la muerte"

El nintendo DS también tiene una pantalla de la muerte, que sucede cuando se remueve una tarjeta DS o un cartucho de GBA mientras se esta en el menú de selección de Pictochat. El color de la pantalla esta basado en la versión del Firmware del DS:

[![](http://2.bp.blogspot.com/_ayvorITawE4/SJu6ZlRupJI/AAAAAAAAA_s/U9FEKq8tao0/s400/DS_BSoD.JPG)](http://2.bp.blogspot.com/_ayvorITawE4/SJu6ZlRupJI/AAAAAAAAA_s/U9FEKq8tao0/s1600-h/DS_BSoD.JPG)"Pantalla de la muerte del nintendo DS"

Parece que entre la gente que ha comprado el nuevo iPhone, esta apareciendo un nuevo rumor: sobre algun problema del software 2.0 y 2.0.1 del iPhone o de alguna aplicación de maliciosa (para el iPhone) que daña el sistema operativo del mismo, solo que este error... es permanente :-S (hay usuarios que reportan que la característica de restauración no funciona).

[![](http://2.bp.blogspot.com/_ayvorITawE4/SJu6ZoPyGXI/AAAAAAAAA_0/9Mt9eJz0BwQ/s400/iphone-boot-screen.jpg)](http://2.bp.blogspot.com/_ayvorITawE4/SJu6ZoPyGXI/AAAAAAAAA_0/9Mt9eJz0BwQ/s1600-h/iphone-boot-screen.jpg)"Pantalla de la Muerte para el iPhone"

Los productos de Apple tampoco se salvan, antes la mayoría de Mac tenían esta imagen:

[![](http://1.bp.blogspot.com/_ayvorITawE4/SJu6ZxBsVjI/AAAAAAAABAM/jNl5fhL8XrI/s400/Sad_mac.png)](http://1.bp.blogspot.com/_ayvorITawE4/SJu6ZxBsVjI/AAAAAAAABAM/jNl5fhL8XrI/s1600-h/Sad_mac.png)"Pantalla de la Muerte, Sad Mac"

Los primeros iPod, también tienen una especie de "Sad Mac":

[![](http://4.bp.blogspot.com/_ayvorITawE4/SJu6Zze7_0I/AAAAAAAABAE/bhX5gJLt5rM/s400/Sad_iPod.png)](http://4.bp.blogspot.com/_ayvorITawE4/SJu6Zze7_0I/AAAAAAAABAE/bhX5gJLt5rM/s1600-h/Sad_iPod.png)"Sad iPod"

Y desde OS X, las Mac tiene una "elegante" pantalla de la muerte:

[![](http://1.bp.blogspot.com/_ayvorITawE4/SJu6ZtAlHLI/AAAAAAAAA_8/ClH833y2etg/s400/MacOSX_kernel_panic.png)](http://1.bp.blogspot.com/_ayvorITawE4/SJu6ZtAlHLI/AAAAAAAAA_8/ClH833y2etg/s1600-h/MacOSX_kernel_panic.png)"Pantalla de la Muerte de Mac OS X"

También GNU/Linux tiene el temido "Kernel Panic", que en lo personal, solo me ha sucedido dos veces: una por un CD de instalación rayado, y otra por que yo compile mal el kernel. Pero también tiene pantalla de la muerte:

[![](http://4.bp.blogspot.com/_ayvorITawE4/SJu7Zh24zjI/AAAAAAAABAs/0qsPj8K-M_g/s400/Kernel-panic.jpg)](http://4.bp.blogspot.com/_ayvorITawE4/SJu7Zh24zjI/AAAAAAAABAs/0qsPj8K-M_g/s1600-h/Kernel-panic.jpg)"Kernel Panic!"

El problema de un SOD para el usuario, es que sus mensajes de error son algo inexplicables/cripticos/misteriosos y no ayudan en nada a los usuarios a conocer el "porque" del error sucedido, lo que contribuye al pánico y la ansiedad (jaja).

Invito a los lectores a recordar las pantallas de la muerte a lo largo de su vida, ¿En que sistema operativo lo han tenido? ¿Cuantas veces? ¿Que otros aparatos han visto con pantallas de la muerte: cajeros automáticos, reproductores de DVD, el Wii, o en el cable?