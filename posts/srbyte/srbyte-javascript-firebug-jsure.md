---
title: JavaScript, FireBug & JSure
date: 2009-11-08
author: Rodrigo Amaya
tags: javascript, jsure
post_id: blog-3515952828243908885.post-1113928188051244522
---

En estas dos ultimas semanas, me vi en la penosa situación, de comprender, alterar y depurar código JavaScript ajeno. JavaScript es un lenguaje script (lo que generalmente implica que es interpretado) usualmente integrado a los navegadores web, para desarrollar paginas web "dinámicas". Se caracteriza por ser un lenguaje basado en prototipos con tipos de datos débil. Inspirado en múltiples lenguajes, fue diseñado con Java (e inherentemente C) en mente.

Todo navegador web que se respete, debe ser capaz de interpretar el código JavaScript que se emplee en una pagina web, pero como se podrán imaginar, [no todos los navegadores interpretan igual el código](https://www.blogger.com/goog_1257714242864)
[JavaScript](https://www.srbyte.com/2009/06/tu-sitio-compatible-en-todos-los.html)
(entre las diferentes versiones de los mismos). Es por eso, que al crear paginas Web modernas e interactivas, hay que utilizar una librería (framework) que cumpla la importante característica de la interoperabilidad entre multiples navegadores (como recordatorio) a esto se le llama: CrossBrowser.

Dojo, [Google Web Toolkit](https://www.srbyte.com/2009/10/trabajando-con-google-web-toolkit-y.html), ICEFaces, jQuery, son algunos ejemplos de librerías y frameworks CrossBrowser, y como pueden ver en el caso de GWT y ICEFaces, a veces ni siquiera se necesita codifica JavaScript.

Un gran, gran problema con JavaScript, radica en el proceso de depuración del mismo. Si bien existen herramientas bien integradas como Aptana, que te permite depurar código JavaScript en su entorno, en otros casos, como en RAD 7.0... (al menos con mi instalación) este IDE no me deja depurar JavaScript.

![image](https://3.bp.blogspot.com/_ayvorITawE4/Svc3PejcMxI/AAAAAAAACN0/HzOBSkO33Ho/s320/frustration.png)    

Supongamos el escenario con un entorno de programación desfasado, en donde no podemos poner puntos de interrupción en el código JavaScript, analizar la pila, o inspeccionar variables... ¿Qué se emplea en este caso tan desahuciado?

FireBug FireBug es un complemento de Firefox creado y diseñado especialmente para desarrolladores y programadores web. Es un paquete de utilidades con el que se puede analizar, editar, monitorear y depurar el código fuente, CSS, HTML y JavaScript de una pagina web en una manera instantánea y "en

caliente" (por ejemplo si se borra una clase CSS, se ve inmediatamente el cambio reflejado en la pagina, eso si: sin modificar el código original).

![image](https://2.bp.blogspot.com/_ayvorITawE4/Svc3KXmPBpI/AAAAAAAACNk/-hrk_3l7DPI/s320/firebug.png)    

[Para Internet Explorer](https://www.srbyte.com/2008/10/depurando-css-y-javascript-con-internet.html) 7, tenemos el DebugBar, que saca de problemas,
pero no es tan bueno como su contraparte (FireBug), e Internet Explorer 8 trae una herramienta integrada similar en funcionamiento a FireBug.

Todo esto es muy bueno, pero existe otro problema. Quizas estas utilizando un archivo JavaScript de 32 KB en el trabajo (eso es cerca de 900 lineas de codigo), y quizas estas pensando: "si es tan grande, probablemente este mal... cambiare algunas cosas para hacerlo más pequeño", existirá un punto en el que todo funciona perfectamente, y dos minutos despues, ya no funciona nada mas. ¿Les ha sucedido esto alguna vez? es sumamente frustrante. En esos dos minutos, hiciste tantos cambios (unos 20 quizas) que no recuerdas que hiciste exactamente para arruinar catastroficamente todo.

Con los archivos JavaScript tradicionalmente existe "la solución" de revertir todos los cambios y dejarlo todo como cuando funcionaba bien... suponiendo que el archivo .js esta versionado, solo perderas un par de horas de tu vida para arreglar todo nuevamente, esperando no arruinar nuevamente todo de manera irreversible. En caso contrario, si el archivo no esta versionado, comenzaras a presionar "deshacer, deshacer, deshacer..." hasta que te des cuenta que esto no sirve de nada. Antes de ponerte verde de la ira, recomiendo para estos casos extremos, el uso de...

JSure JSure es un analizador de sintaxis de JavaScript (probablemente el mejor que hay), de manera que [JSure](https://www.jsure.org/), te puede decir exactamente donde esta el

problema con el código JavaScript que modificaste, y también te muestra la linea del error y el error en si mismo remarcado. Lo mejor de JSure, es que existe una versión para Linux y MacOS X, y para los usuarios de Windows, se puede emplear la versión en linea ([https://www.jsure.org/](https://www.jsure.org/)), que permite verificar pequeños métodos o grandes archivos.

![image](https://2.bp.blogspot.com/_ayvorITawE4/Svc3MKV8EII/AAAAAAAACNs/_ZANt7ysHsg/s320/logo.png)    

Con estas dos fabulosas herramientas, he logrado convertir mi vieja relación de amor/odio con JavaScript, en "más amor y menos odio" ... pero todavía no dejo de odiarlo... en fin, JSure es una herramienta tan indispensable como FireBug, espero que a ustedes también les sirva tanto como a mi.

Más información:
[https://www.jsure.org/](https://www.jsure.org/)
[https://www.tufuncion.com/firebug](https://www.tufuncion.com/firebug)
(Recomendado)
[https://es.wikipedia.org/wiki/JavaScript](https://es.wikipedia.org/wiki/JavaScript)
[https://en.wikipedia.org/wiki/Firebug_%28Firefox_extension%29](https://en.wikipedia.org/wiki/Firebug_%28Firefox_extension%29)
[https://es.wikipedia.org/wiki/Firebug](https://es.wikipedia.org/wiki/Firebug)