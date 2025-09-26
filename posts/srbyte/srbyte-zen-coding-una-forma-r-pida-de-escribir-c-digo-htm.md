---
title: Zen Coding: Una Forma Rápida de Escribir Código HTML
date: 2009-12-09
author: Robertux
tags: wysiwyg, aptana, programacion, zen, html
post_id: blog-3515952828243908885.post-8972450777597994202
---

![image](https://3.bp.blogspot.com/_jH77WNrMVRA/Sx8EzahcduI/AAAAAAAAGII/rnNPT-vXtq0/s400/zen.jpg)    
"Libérate del Stress de
escribir páginas HTML a mano"

Todo desarrollador web alguna vez tuvo que ensuciarse las manos escribiendo código HTML y se habrá dado cuenta de lo inmenso y engorroso que este puede volverse y de la poca eficiencia de los IDEs con capacidades WYSIWYG existentes. Aún teniendo frameworks como [GWT](https://www.srbyte.com/2009/10/trabajando-con-google-web-toolkit-y.html) que nos generan automáticamente todo el HTML y Javascript de nuestro sitio, en muchas ocasiones estos no resuelven nuestras necesidades específicas o "no forman parte del estándar de soluciones viables en nuestra empresa".

Afortunadamente para mi, leyendo [este post](https://www.smashingmagazine.com/2009/11/21/zen-coding-a-new-way-to-write-html-code/) de Smashing Magazine me doy cuenta de la reciente existencia de Zen Coding, un framework que te permite escribir código HTML de una manera bastante rápida y sencilla usando selectores similares a los ya usados en otros frameworks como [prototype](https://www.prototypejs.org/) y [jquery](https://jquery.com/) quienes los usan para hacer búsquedas de elementos HTML pero en este caso, el framework de Zen Coding los utiliza para generar nuestro código HTML reduciendo a su mínima expresion la cantidad de código a escribir para generarlo.

Una vez que te familiarizas con los selectores utilizados, notarás la diferencia de tiempo al momento de crear y editar el código HTML. A continuación un ejemplo de un selector de Zen Coding:

```
html:xt>div#header>h1#logo+ul#nav>li.item-$*5>a

```

y el código HTML que este genera:

```
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
"https://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="https://www.w3.org/1999/xhtml" xml:lang="en"> <head> <title></title> <meta http-equiv="Content-Type" content="text/html;charset=UTF-8" /> </head> <body> <div id="header"> <h1 id="logo"></h1>

<ul id="nav"> <li class="item-1"><a href=""></a></li> <li class="item-2"><a href=""></a></li> <li class="item-3"><a href=""></a></li> <li class="item-4"><a href=""></a></li> <li class="item-5"><a href=""></a></li> </ul> </div> </body> </html>

```

Este framework puede ser incrustado en diferentes IDEs por medio de plugins como por ejemplo Netbeans, Eclipse (y sus derivados), Textmate (mac), WebIDE, Aptana Studio, entre otros. Todos ellos se pueden descargar desde la [página oficial de Zen Coding](https://code.google.com/p/zen-coding/) en Google Code.

[En ese link](https://zen-coding.ru/demo/) puedes probar un demo versión web,
por si no deseas descargar ninguna de sus versiones asociadas a estos IDEs de desarrollo.

"Video demostrativo del
funcionamiento del framework Zen Coding"