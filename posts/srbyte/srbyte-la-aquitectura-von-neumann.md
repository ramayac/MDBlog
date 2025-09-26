---
title: La Aquitectura Von Neumann
date: 2007-12-31
author: Rodrigo Amaya
tags: computadoras, arquitectura
post_id: blog-3515952828243908885.post-8581992345864867649
---

La Arquitectura de Von Neumann (también llamada Eckert-Mauchly) se refiera a un modelo de diseño hardware de computadoras programables, creado por el matemático húngaro John Von Neumann.

![image](https://upload.wikimedia.org/wikipedia/en/d/d6/John_von_Neumann.JPG)    
"Estampilla conmemorativa de John von Neumann."

John von Neumann (Diciembre 28, 1903 - Febrero 8, 1957), fue un matemático que hizo gigantescas contribuciones a un vasto rango de campos que incluyen: física cuántica, análisis funcional, topología, economía, ciencias de las computadoras, análisis numérico, y muchas más.

Este modelo de diseño usa una unidad de procesamiento y una estructura separada para contener instrucciones y datos. Una computadora diseñada de esta forma implementa una [Maquina Universal de Turing](https://srbyte.blogspot.com/2007/12/la-maquina-de-turing.html) ¿ven la conexión?.

![image](https://upload.wikimedia.org/wikipedia/commons/b/bd/Arquitectura_von_Neumann.png)    
"Click en la imagen para verla más
grande."

Como pueden ver a partir de la imagen superior, este modelo de diseño es el abuelo de la mayoría de las computadoras diseñadas en la actualidad.

Pero este modelo de arquitectura no es el único que existe, también hay otros de igual importancia...

La Arquitectura Harvard

La arquitectura Hardvard es otra forma de modelar el hardware en una computadora programable, esta arquitectura separa el almacenamiento y el camino de las señales eléctricas en donde viajan las instrucciones y los datos (data bus). El termino se origina de la computadora Harvard Mark I. Esta computadora poseía una unidad de almacenamiento pequeña contenida en su totalidad en la unidad de procesamiento... es decir, tenia el disco duro dentro de la CPU, imaginemos que era como la memoria cache actual. Así que para cambiar su programa (o grupo de instrucciones) debían de apagar a "Mark", lo cual habrá sido tedioso.

En contraste con una arquitectura de Von Neumann, la arquitectura Harvard es definitivamente la ganadora en velocidad, pero no en conveniencia a la hora de ejecutar grupos de instrucciones distintas. Un ejemplos de computadoras programables que usan la arquitectura Harvard, son los famosos PIC...

![image](https://www.ke4nyv.com/pics5.jpg)    
"Diversos modelos de
PIC's"

Pero no todo es amor y felicidad en la tierra de Von Neumann. La separación que se obtiene entre el CPU y la memoria da como resultado una de las obsesiones mas grandes de la industria del hardware: el famoso cuello de botella. El problema radica, en que la CPU esta continuamente forzada a esperar a que los datos sean transportados desde o hacia la memoria RAM, mas que todo debido a la diferencia entre la taza de transferencia de datos que existen entre la CPU y la RAM.

![image](https://www.labcentrix.com/images/bottleneck_diagram.jpg)    
"Imagen de ejemplo de un cuello de botella"

¿Interesante verdad? ¿Quieres saber más? Lee el articulo "[Lo que no sabias sobre las computadoras](https://srbyte.blogspot.com/2007/12/lo-que-no-sabias-de-las-computadoras.html)
".