---
title: Administrando ancho de banda con Trickle
date: 2007-10-10
author: Rodrigo Amaya
tags: administrar, linux, limitar, trickle, internet, bandwith
post_id: blog-3515952828243908885.post-8103715126442289826
---

[![](http://bp3.blogger.com/_ayvorITawE4/RwzyBFmN2bI/AAAAAAAAAgo/UP4azpHUO_o/s320/work_bandwidth.jpg)](http://bp3.blogger.com/_ayvorITawE4/RwzyBFmN2bI/AAAAAAAAAgo/UP4azpHUO_o/s1600-h/work_bandwidth.jpg)"Estoy seguro de que muchos de nosotros usaríamos esta camisa..."

Como ya saben muchos "linuxeros", en GNU\Linux NO es necesario formatear y re-instalar el sistema operativo para obtener las características mas novedosas. Basta con usar las herramientas que el mismo sistema facilita, como en Debian (y todas sus variantes) por ejemplo, para cambiar a una mejor versión:

> #apt-get dist-upgrade
... claro que faltan muchos pasos (apt-get update, apt-get upgrade, etc...) pero se entiende la idea.

Una de las cosas que mas me gusta hacer, es hacer una actualización del sistema en una terminal, y dejarla como una tarea de fondo. Pero apt-get siempre me "roba" todo el ancho de banda, y me quedo sin poder revisar mi correo, u otras paginas que visito usualmente. A raíz de esto, recordé que hace muuuuucho leí sobre un programa para forzar a que los programas solo usen un determinado porcentaje o limite de la conexión a internet, algo asi como la opcion "--limit-rate" de wget, es decir... un "bandwidth shaper"

Y ese programa del que les hablo es trickle y trickled (daemon de trickle). Trickle se encarga de limitar (pero como tu la limitas, tu administras ;) el uso de ancho de banda de "subida" y "bajada" de una aplicación que le especifiquemos.

[![](http://bp1.blogger.com/_ayvorITawE4/Rwzw-lmN2aI/AAAAAAAAAgg/MIIGtUo8v-o/s320/masanchobanda.jpg)](http://bp1.blogger.com/_ayvorITawE4/Rwzw-lmN2aI/AAAAAAAAAgg/MIIGtUo8v-o/s1600-h/masanchobanda.jpg)"¡Usa más Ancho de Banda!"

Digamos que usted tiene una conexión de 256 Kbps, y necesita actualizar el sistema, pero también quiere leer todos los post de este blog ;) entonces se haría lo siguiente:

> #trickle -u 5 -d 12
> apt-get upgrade
Entonces trickle limitaría "apt-get upgrade" para usar solo 5 Kbps de "subida" y 12 Kbps de "bajada", dejando un aproximado de 13 Kbps libres para leer este blog :-D Y además trickle ser invocado tanto por usuarios como por root. ¡Lo cual lo hace ideal para toda la familia! xD

En lenguaje Geek: trickle es una aplicación del espacio de usuarios, que puede trabajar en forma colaborativa con trickled o en modo solitario (stand alone).

Definitivamente una muy útil aplicación, que puede ser usada para mil cosas más. Visiten la pagina de Trickle para obtener más información: [http://monkey.org/~marius/pages/?page=trickle](http://monkey.org/%7Emarius/pages/?page=trickle)

Saludos!

PD: ¿Cuanto smiley no? :D