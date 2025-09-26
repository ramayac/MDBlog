---
title: Rendimiento Actual vs. Rendimiento Percibido
date: 2008-03-05
author: Rodrigo Amaya
tags: programador, bad, vista, programacion, interesante, microsoft
post_id: blog-3515952828243908885.post-2477388422239409630
---

[![](http://bp1.blogger.com/_ayvorITawE4/R87T-ImBXYI/AAAAAAAAAlw/c09o5W6-EpU/s200/vista-logo-magnified.jpg)](http://bp1.blogger.com/_ayvorITawE4/R87T-ImBXYI/AAAAAAAAAlw/c09o5W6-EpU/s1600-h/vista-logo-magnified.jpg)Para todos ustedes que lamentablemente han usado Windows Vista y Windows XP, se habrán percatado que el rendimiento al momento de copiar archivos en Windows Vista es notablemente peor que en Windows XP. Eso y la horrible tardanza a la hora de actualizar Vista fue una de las cosas que mas note cuando tuve que usarlo. Pero la ironía es que el algoritmo de copia de archivos de Vista es mejor y rinde muchos mas que el del XP...

> Entonces, ¿Por qué se siente taaaaan
> lennnnnntoooooooooooo?
Pues para tener una idea, pueden leer este [fantástico articulo del blog de Mark Russinovich's](http://blogs.technet.com/markrussinovich/archive/2008/02/04/2826167.aspx), que se encarga de hacer una exhaustiva serie de pruebas en el nuevo Vista SP1:

Extracto de Mark's Blog, el Sr. Byte traduce:

> "Quizás la gran desventaja del nuevo algoritmo de copia de
> Vista, y lo que ha causado que muchos usuarios de Vista se quejen, es que para copias que
> involucran un largo grupos de archivos entre 256KB y decenas de MB en tamaño, el rendimiento
> percibido de una copia puede ser significativamente peor que en XP."
Mark puntualiza: la copia de un archivo no es tan fácil como parece al principio. Como con muchas cosas que suceden en la vida - y aquí nos ponemos filosóficos -, la percepción es la realidad: si los usuarios ven que la copia de un archivo es lenta, entonces es lenta. Y en un mundo, en donde "jamas nunca" (o al menos sin [peyote](http://en.wikipedia.org/wiki/Peyote)) se pueden tocar los programas, el usuario final depende de su visión para determinar si algo es o no es más rápido en su PC. A pesar de todas las mejoras en el algoritmo, a pesar de los resultados de copia superiores, el rendimiento de copia de Vista es peor que en XP. ¿Y como ven los usuarios que es lenta o rápida la copia de un archivo? pues mediante la interfaz, osea, la sencilla usual y sobre valorada: Barra de Progreso.

[![](http://bp3.blogger.com/_ayvorITawE4/R87PZomBXXI/AAAAAAAAAlo/tDtcJVgbxgc/s400/progress_bar.gif)](http://bp3.blogger.com/_ayvorITawE4/R87PZomBXXI/AAAAAAAAAlo/tDtcJVgbxgc/s1600-h/progress_bar.gif)"Imagen: Diversas apariencias de barras de progreso."

Ojo, aquí les va un tipo sobre un factor humano secreto, que se omite en el desarrollo de software: El rendimiento percibido es mas importante que el rendimiento actual. ¿Ya se imaginan por donde vamos, verdad? Los elaborados algoritmos de copia no necesariamente ayudan a construir barras de progreso mas rápidos. Pero entender como funciona el cerebro de tus usuarios definitivamente si lo hará, como lo ilustra el articulo [Rethinking the Progress Bar](http://chrisharrison.net/projects/progressbars/ProgBarHarrison.pdf) (Reinventando la barra de progreso):

Extracto de [Rethinking the Progress Bar](http://chrisharrison.net/projects/progressbars/ProgBarHarrison.pdf), el Sr. Byte traduce:

> "Los seres humanos no perciben el paso del
> tiempo en forma linear. Esto, emparejado con el comportamiento irregular de las barras de
> progreso, causan que la percepción humana de la duración del proceso varíe. La comprensión de
> cuales comportamientos perceptibles acortan lo alargan el proceso de duración puede ser usado
> para diseñar una barra de progreso que aparente ser mas rápida, aunque la duración sea la
> misma."
[![](http://bp0.blogger.com/_ayvorITawE4/R87OM4mBXWI/AAAAAAAAAlg/vtCgkIFVNiA/s400/progress-function-graph.png)](http://bp0.blogger.com/_ayvorITawE4/R87OM4mBXWI/AAAAAAAAAlg/vtCgkIFVNiA/s1600-h/progress-function-graph.png)"Gráfico de uso de 8 comportamientos de una barra de progreso y la reacción del usuario, con respecto a cada una"

Aunque todas las barras de progreso del estudio completaban su tarea en el mismo periodo de tiempo, dos características hicieron que los usuarios pensaran que el proceso era mas rápido, ¡aunque en realidad no lo era! :

1. Barra de progreso que se mueve suavemente hasta completarse. 2. Barra de progreso que aumenta su velocidad cuando se aproxima al final. La idea de que el rendimiento es determinado enormemente por la percepción del usuario, en vez del tiempo, puede ser bastante liberador. Pero también puede ser sumamente frustrante. Porque aunque se tenga las partes técnicas bien, con pruebas de rendimiento solidas para respaldarlo, los sutiles factores de percepción humana pueden aún negar el trabajo del mejor programador del mundo, como lo pudieron apreciar los pobres desarrolladores de Vista.

[![](http://www.mlobit.com/word/wp-content/internetdownloadwww.gif)](http://www.mlobit.com/word/wp-content/internetdownloadwww.gif)"Imagen: Descargando Internet"

En pocas palabras, desarrolladores de software, no comentan el mismo error que el equipo de desarrolladores de Vista hizo. Hay que pensar fuera de las pruebas de rendimiento, y comenzar a usar la inteligencia emocional, y comprender al usuario final... por que al final, el software siempre se diseña para el usuario.

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte