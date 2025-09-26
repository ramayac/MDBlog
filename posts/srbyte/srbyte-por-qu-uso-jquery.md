---
title: ¿Por qué uso jQuery?
date: 2010-06-24
author: Rodrigo Amaya
tags: javascript, jquery
post_id: blog-3515952828243908885.post-8499818573577577942
---

En estos dorados tiempos, la mayoría de aplicaciones web no pueden ser liberadas a la audiencia mundial, sin usar alguna de las millones de librerías de JavaScript que "pululan" en Internet. En un articulo anterior, se mencionaron los beneficios de usar alguna de las librerías más populares para obtener el máximo de compatibilidad en tu sitio web. Si bien todas estas librerías (Prototype, Scriptaculous, Dojo, jQuery) son muy buenas, en esta ocasión quiero escribir de la librería, que a mi me salva todos los días en el trabajo... jQuery.

![image](https://3.bp.blogspot.com/_ayvorITawE4/SkD8ZuQG_JI/AAAAAAAACDg/33VKLTOS9xM/s320/jquery-logo.gif)     jQuery es uno de
tantos hijos que Prototype tiene regados en Internet, pero no por eso lo vean de menos. Si bien jQuery reusa extensivamente la función $() , también la lleva a niveles que uno no se imagina hasta que trabajas mucho con el. Es más, jQuery básicamente esta basado en el signo $. En jQuery, el $ manda ;) En fin, estas son mis cinco simples razones, por la cual yo prefiero jQuery...

Primera razón: Selectores jQuery soporta muchos "selectores", un selector es una forma de obtener un etiquetas html, div's o elementos CSS, lo que antes tenia que obtener con document.getElementById("idelemento"); ahora lo puedo obtener con $("#identificador") .

Supongamos que tienes un div en tu pagina, asi: <div id="prueba">texto</div>

Y quieres obtener el texto enmarcado en ese div, basta con hacer un $("#prueba").val(); para realizar esta tarea.

### Segunda razón: Atributos

Con jQuery es facil añadir y remover elementos HTML. Con el div del ejemplo anterior, supongamos que deseamos modificar su contenido: $("#prueba").val("nuevo texto");

Y si queremos manipular el tipo de clase CSS al que pertenece... :

$("#prueba").attr("title", "Div de Prueba").addClass("nuevaClaseCSS");

Se puede añadir tantas clases CSS como se deseen. Y si son observadores, habran notado que se hace un uso extensivo del operador ".", eso es porque jQuery usa "objetos", y la concatenacion de varios metodos en jQuery se llama: [Chainability](https://blogs.telerik.com/supportdept/posts/09-02-06/chainability_the_magic_of_jquery.aspx) (Encadenamiento). El encadenamiento, en pocas palabras conciste en que: cada metodo en jQuery, retorna el objeto solicitado, lo que permite que operemos sobre este.

### Tercera razón: Ajax

Con jQuery me olvide de manipular Ajax, para mi eso es cosa del pasado. Usar Ajax es tan facil como escribir 1, 2 y 3...

> $.post class="br0">( class="st0">"https://www.elsitio.com/eltexto",
> function(data class="br0">){
> alert("obtuve el
> siguiente texto: " + data); class="br0">});

Cuarta razón: Documentación y Comunidad

De nada te sirve la mejor libreria de software, si esta no esta documentada, y con jQuery, no solo hay ejemplos de todas sus funciones, tutoriales, presentaciones y ejemplos, sino que tambien hay una comunidad activa, creando constantemente nuevos plugins para extender el uso de jQuery.

![image](https://3.bp.blogspot.com/_ayvorITawE4/SkD8Z2oZCgI/AAAAAAAACDw/FuFLvnTTjXQ/s320/jquery12_colorcharge.png)    

Quinta razón: jQuery UI

jQuery UI, provee una serie de metodos de animacion, interaccion, efectos y controles construidos sobre jQuery. Es facil de usar, aprender e implementar, y ademas, soporta "temas" para una mejor integracion con tu sitio web.

![image](https://3.bp.blogspot.com/_ayvorITawE4/SkD8ZxsxG1I/AAAAAAAACDo/poR6fdeUVlc/s320/jquery_ui_logo.png)