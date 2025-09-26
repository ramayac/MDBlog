---
title: ¿Qué es Heurística?
date: 2009-10-14
author: Rodrigo Amaya
tags: algoritmos, heuristica
post_id: blog-3515952828243908885.post-6943288365160740787
---

Si has "escaneado" tu máquina con un antivirus, y pasaste minutos de tu vida mirando la barra de progreso y los comentarios en cada archivo, probablemente viste mas de alguna vez (y con cualquier buen antivirus) la palabra: heurística.

![image](https://4.bp.blogspot.com/_ayvorITawE4/StaLNqjxc3I/AAAAAAAACMY/PGHIFcanHw8/s320/searchradarpc.jpg)    
"Fuera interesante ver el
progreso de escaneo de un ativirus así"

En la mayoría de antivirus, la [heurística](https://es.wikipedia.org/wiki/Heur%C3%ADstica_%28inform%C3%A1tica%29) se refiere a una técnica especifica para reconocer código malicioso ([virus](https://es.wikipedia.org/wiki/Virus_%28inform%C3%A1tica%29), [gusanos](https://es.wikipedia.org/wiki/Gusano_%28inform%C3%A1tica%29), [troyanos](https://es.wikipedia.org/wiki/Caballo_de_Troya_%28inform%C3%A1tica%29), etc.) que este (el antivirus) no posea en su base de datos por diversas razones, como que los programas maliciosos sean nuevos o poco divulgados.

La heurística, como termino general, implica funcionalidades como detección a través de firmas genéricas, reconocimiento del código compilado, desensamblado, desempaquetamiento, entre otros. Recordemos que los antivirus, son como inmensos programas de listas negras, si encuentran algo en memoria o en el disco que sea igual a lo que esta en su base de datos lo eliminan inmediatamente. Y con la heurística, sus capacidades para detectar software malicioso aumenta. La importancia comercial y técnica de la heurística (al menos en los antivirus) radica en el hecho de ser la única defensa automática y reactiva posible frente a la aparición de nuevo software maliciosos de los que no se posea "firmas" o conocimiento alguno (en la base de datos).

> Pero claro, el costo de la seguridad es el desempeño y por esta misma razón mientras
> mas detallado o profundo sea el análisis heurístico de un antivirus, más pobre sera el
> rendimiento del mismo.
La heurística no solo se emplea como palabra complicada para denotar una técnica que todo antivirus decente debería tener, sino que también se emplea en el area de la inteligencia artificial. Muchos algoritmos de "denominada" [inteligencia artificial](https://es.wikipedia.org/wiki/Inteligencia_artificial) son heurísticos por naturaleza o usan reglas heurísticas, un excelente ejemplo es [SpamAssassin](https://spamassassin.apache.org/) que usa una amplia variedad de reglas heurísticas para determinar cuando un email es [spam](https://es.wikipedia.org/wiki/Spam)...

![image](https://upload.wikimedia.org/wikipedia/commons/b/b7/SpamAssassin_logo.png)    
"SpamAssassin = Apache.org + Heuristics"

Cualquiera de las reglas usadas de forma independiente pueden llevar a errores de clasificación, pero cuando se unen múltiples reglas heurísticas, la solución es más robusta y creíble. Esto se llama alta credibilidad en el reconocimiento de patrones.

Ahora bien, no todo es bueno con la heurística... si algo podemos decir de un algoritmo cualquiera, es que este se pueda ejecutar varias veces, y su resultado sea óptimo. Por irónico que parezca, un algoritmo heurístico abandona estos objetivos para obtener una buena solucion, que pudiera convertirse arbitrariamente en una mala solucion, como los famosos [falsos positivos en un antivirus](https://foros.softonic.com/seguridad/antivirus-da-falsos-posivos-47252)...

![image](https://3.bp.blogspot.com/_ayvorITawE4/StaLOCESF3I/AAAAAAAACMg/hpNieJRhsrU/s320/utorrent-malware.png)    
":@ esto es lo mas terrible
del mundo, un falso positivo en un archivo perfectamente sano."

O sino, el algoritmo heurístico se ejecuta razonablemente rápido, pero no hay argumento de que esto siga sucediendo siempre. Finalmente, y con la música de Eurythmics de fondo...

"Eurythmics - Sweet Dreams"

...quiero cerrar esta breve explicación de un tema tan amplio citando al científico [Judea Pearl](https://en.wikipedia.org/wiki/Judea_Pearl): "Heurística, son métodos basados en búsquedas inteligentes de estrategias para resolver problemas computacionales utilizando muchos acercamientos alternativos."

Espero que les sirva,