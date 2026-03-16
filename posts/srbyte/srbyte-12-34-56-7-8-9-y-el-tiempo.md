---
title: 12:34:56 7/8/9 y el tiempo...
date: 2009-08-07
author: Rodrigo Amaya
tags: tiempo, linux, unix
post_id: blog-3515952828243908885.post-8334987840556348617
---

![image](https://www.pixelydixel.com/img/2009/08/123456789-400x105.png)  

Este viernes a las 12 del medio día, exactamente, a las 12:34:56 asistiremos por primera y única vez en el siglo a la secuencia perfecta de dígitos 12:34:56 7/8/9.

![image](https://4.bp.blogspot.com/_ayvorITawE4/SnxniE_t7gI/AAAAAAAACIE/os4JZU0LPds/s200/linux-y2k-unix.jpg)    Estas secuencias...
o mejor dicho, curiosidades de juegos de números, siempre me sacan una sonrisa, e inmediatamente me hacen pensar en el famoso bug del milenio, ¿recuerdan al infame Y2K bug? No fue tan malo como muchos creían, más que todo porque el error estaba asociado a como se muestra la informacion, y no con su funcionalidad interna. Al final, todo mundo sobrevivió ese problemita. Pero hay un problema similar y vigente llamado Y2K38, que afecta a los sistemas sistemas de la familia Unix, y este bug es mucho más difícil de resolver. Una solucion practica es migrar a un sistema operativo, que use una representación de tiempo de 64 bits, sin embargo el problema persiste en sistemas de 32 bits. ¿Y quien usa sistemas de 32 bits?... los dispositivos móviles y pequeños reproductores de música.

![image](https://2.bp.blogspot.com/_ayvorITawE4/Snxm6d3aCsI/AAAAAAAACH8/gzoyt8PT_cc/s320/Year_2038_problem.gif)    
Bien, pero seamos realistas, al paso que va la industria, para el 2025 espero que ya todos tengamos más 64 bits, y los dispositivos moviles probablemente también vayan por ese camino. Asi que, por ese lado estamos moderadamente seguros. Pero... el verdadero problema, esta en el software.

El tiempo en la computadora/ordenador...

El tiempo, en las computadoras, es representado por el número de segundos que han transcurrido desde el Unix epoch, es decir desde: 00:00:00 UTC Enero 1 de 1970. Ese numero de segundos transcurridos desde esa fecha se conoce como un "timestamp", bien, muchísimos programas usan [timestamp](https://en.wikipedia.org/wiki/Timestamp) para obtener la representación del tiempo y mostrarnos la fecha actual, la fecha de la ultima modificación de un archivo, etc etc etc, el asunto es que estos mismo programas asumen que el tamaño de ese campo NO cambia (siempre es de 32 bits), entonces un programa de 32 bits, migrado (que se ejecute en modo de compatibilidad) a un sistema de 64 bits, leerá el timestamp correcto (de 64 bits) de manera incorrecta (lo leeria como uno de 32 bits).... ¡Ooops!

La situación es interesante, pero NO es fatal. Ya que no se ve realista seguir usando software de 32 bits en el 2038... al menos para mi, no lo es... pero el ciclo de vida de un software, [puede durar mucho más de lo que esperamos](https://www.srbyte.com/2008/11/el-fin-de-win-311-e-ideas-sobre-la-vida.html) y tal vez, alguien se verá en el problema de brindar soporte a aplicaciones antiguas ([Legacy Applications](https://en.wikipedia.org/wiki/Legacy_system))... Aquellos que lo hagan, no estarán en problemas por brindar soporte a una aplicación caducada, estarán en problemas por trabajar en empresas (o con empresas) que necesite "mantener viva" una Legacy Aplication para seguir funcionado.

![image](https://4.bp.blogspot.com/_ayvorITawE4/Snxm6NK1b2I/AAAAAAAACH0/RJ-In9fm4qM/s320/1000000000seconds.jpg)    
"[Unix Billenium](https://en.wikipedia.org/wiki/Unix_billenium)
"

Otros afectados por la falta de visión, o por el limite de un numero entero sin signo, es Twitter, que aparte de que [ayer fue atacado](https://mashable.com/2009/08/06/denial-of-service-attack/) (junto con FaceBook y LiveJournal) con una denegación de servicio, le llega el [apocalipsis](https://www.twitpocalypse.com/) (otra vez) el 29 de Octubre de 2009. ¿Cuando tendremos que preocuparnos por otro infame bug de tiempo?, usando un timestamp de 64bits (con signo), hasta Diciembre 4 del año 292,277,026,596 ... para esa fecha, ya no tendremos preocupaciones. La mayoria de defectos relacionados con el tiempo (timestamp, date bugs, etc) han sido de caracter cosmetico, y una vez aparentes, pues se pueden resolver bastante rapido.

Estas irregularidades me recuerdan dos cosas:

1. Como le gusta exagerar a la gente los problemas 2. Como nos gusta usar cualquier excusa para divertirnos un rato Me voy a celebrar "12:34:56 7/8/9", apuro el paso y entrego más artículos la otra semana, siento haber dejado tirado el blog, pero estamos apretados con la Tesis, y bueno, hay prioridades, más adelante les cuento como sale todo eso. Saludos!

Si quieren leer más sobre "el tiempo" en los sistemas Unix, pueden encontrar más información