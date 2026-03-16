---
title: Número de Versión en el Software
date: 2007-02-16
author: Rodrigo Amaya
tags: libre, gnu, linux, software, microsoft
post_id: blog-3515952828243908885.post-3385006549121719633
---

Si algo sabe hacer Microsoft bien, es llegar a todo el mundo. Respiremos profundo y aceptar la realidad. ¿Que tiene que ver esto con las versiones de los programas? Pues cuando Microsoft cambio el nombre de Windows 4.0 a Windows 95, miles rasgaron sus vestiduras y muchos dijeron: ¿Cómo sabrá el usuario a que versión de Windows actualizarse?

![image](https://bp3.blogger.com/_ayvorITawE4/RdaAu6SA9-I/AAAAAAAAAHQ/BubUZxd1mhw/s400/grito.jpg)    
"El Grito - Edvard Münch, 1893"
Parece ridículo ahora, pero fue un riesgo significativo en su momento, y para ser francos este cambio de política tiene mucho sentido ¿no? ¿Por que tiene sentido? fácil:

- Al usuario común no le importa las versiones de los programas (no, realmente no le importan, solo quiere tener lo mas reciente y funcional/productivo)
- Un modelo de numeración basado en años es fácil de entender que un modelo de numeración tradicional.
Un problema frecuente con mucho Software Libre es que muchos desarrolladores sigue los lineamientos de numeración de versiones del Kernel de Linux, por ejemplo la versión de numeración de un Kernel Linux, son tres números separados por un punto:

> style="font-size:130%;"> 2.4.19
este número se descompone así:

- Número Mayor: Representa cambios muy significativos en el Software.
- Número Menor: Números impares aquí significan versión de desarrollo, y un numero par significa versión de producción.
- Micro Número / Parche: Significa numero de entullezcan o re-compilación o que arreglaron un bug o alguna otra cosa.
Esa numeración tiene lógica para algo tan delicado como el desarrollo de un Kernel de un Sistema Operativo. Pero los desarrolladores de software libre van mas lejos cuando llaman a sus productos así. Una de las mejores ideas que se ha visto en el software libre son los nombres claves para las diferentes distribuciones de GNU/Linux; Ubuntu Hoary, por ejemplo, es mejor que Ubuntu 5.04 para el usuario (Aunque no mucho para el usuario de habla hispana, pero se entiende la idea). La forma normal de numerar un programa en .NET/Mono es:

> (Versión mayor).(Versión
> Menor).(Numero de revisión).(Numero de construcción)

Pero como dije antes, al usuario todos esos números no le importan y es mas, muchas veces:

- Lo confunden
- No saben para que sirve
- No le dan importancia
Y es que la numeración de las versiones se usa en caso de emergencia o actualización de Software, para la etapa de desarrollo o para darle seguimiento en general pero no para el nombre comercial de un software enfocado para el usuario final!. Por eso Microsoft Office 2007 se lee mas "bonito" que Microsoft Office 11.8125.8122.00 (verdad?). Lo mismo sucede con el Windows Messenger cuando cambio de 7.5 a 8.0 ¿se fijaron que ahora el se llama Windows Live Messenger? Pero los números siguen ahí:

![image](https://bp0.blogger.com/_ayvorITawE4/RdZ1xKSA99I/AAAAAAAAAHI/IHkaYidDPxw/s400/messenger.jpg)    
Al parecer nadie se salva de agregar la fecha de alguna u otra forma en la numeración de sus programas/productos... ni siquiera el caricaturista Ruz:

![image](https://bp3.blogger.com/_ayvorITawE4/RdaCX6SA9_I/AAAAAAAAAHY/ZufbAfpnUYs/s200/ruzhoy.JPG)    ...pero tomando un
caso mejor adaptado a nuestra realidad tecnológica, veamos el de Microsoft Office, su versión de construcción se dividen en 4 dígito con la fecha codificada... no es tan difícil como se lee, la lógica funciona así:

- Toma el año en el que el proyecto comienza. Para Microsoft Office "12" fue el 2003.
- Llama a Enero de ese año "Mes 1".
- Los primeros dos dígito del numero son los números de los meses desde el "Mes 1"
- Los últimos dos dígito son el día de ese mes.
Entonces en un programa de Microsoft Office con una construcción 2413, siguiendo el algoritmo anterior (suponiendo siempre que el proyecto comienza en el 2003): Tomo los primeros dos dígito: el "Mes 12" = Diciembre 2003; por lo tanto el "Mes 24" = Diciembre 2004, y los dos siguientes números son el día 13 de diciembre.

La fecha entonces es: 13 de Diciembre de 2004.

De todos estos datos ¿Cuales son importantes para el usuario?, ¿Y cuantos para los desarrolladores? La idea general de todo esto es confundir lo menos posible al usuario. Siempre que sea posible es mejor usar datos más simples para llamar un programa que utilizar versiones numéricas en su nombre y también variar un poco es bueno, por eso Windows Vista no se llama Windows 2005, esto es aplicable particularmente a los nombres enfocados para la comercialización de nuestro Software/producto.

Hasta la