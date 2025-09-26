---
title: Pruebas de Uso (Usability Test)
date: 2009-03-04
author: Rodrigo Amaya
tags: practicas, util, desarrollo, pruebas, usabilidad
post_id: blog-3515952828243908885.post-3466889532297221442
---

Supongamos que estas trabajando en un software, por los últimos... cuatro o cinco meses. Y tienes que entregar un producto terminado en dos semanas (tu deadline es en dos semanas). Resuelves bugs, tu código esta cubierto (Code Coverage), tu software pasa las pruebas unitarias (Unit Testing), y además, pasa las pruebas de deploy ("Deployar" como escuche mencionar alguna vez) y de pronto, algo realmente importante surge en tu mente, algo que tienes ahi en tu cráneo desde el principio...

> ¿Es usable mi
> programa?
Rayos!, esa, si que es una pregunta muy importante, no solo porque tienes que saber:

> ¿Quien usa mi programa? o ¿Cuanto van a usar mi
> programa?
... para ver si realmente lo que hiciste es útil. No, no solo es por eso. Es importante, debido a la simple y sencilla razón, de que el cliente o usuario final, es la razón, propósito y motivo ulterior por la que estas desarrollando software, y TU software (o el software que estas desarrollando) DEBE (si, DEBE) facilitar la vida del usuario final de la aplicación.

Que este sea el mantra de esta semana:

> "Mi aplicación deberá facilitar la vida del usuario
> final."
La manera para asegurarte de que este mantra se cumpla lo máximo posible es utilizando las [Pruebas de Uso](http://en.wikipedia.org/wiki/Usability_testing). En el sentido técnico una Prueba de Uso consiste en realizar una prueba de "caja negra" a tu software. Es decir, observar a las personas que van a usar el producto, REALMENTE usándolo para descubrir errores y ÁREAS de mejora.

[![](http://3.bp.blogspot.com/_ayvorITawE4/SbB842QJbyI/AAAAAAAAB58/nNfPj9A50LY/s320/eyetracking_corporate_site_about_us.png)](http://3.bp.blogspot.com/_ayvorITawE4/SbB842QJbyI/AAAAAAAAB58/nNfPj9A50LY/s1600-h/eyetracking_corporate_site_about_us.png)"Google realiza pruebas de uso, para conocer que áreas de una pagina son las que mas son 'visitadas' por el cursor (esto también es conocido como HeatMap)"

Suele ser de caja negra, porque no importa lo que hace en el fondo el sistema, sino la forma. Generalmente una prueba de uso, medirá que tan bien responden los usuarios en cuatro áreas: eficiencia, exactitud, recuerdo o familiaridad (recall), y respuesta emocional...
- Rendimiento -- ¿Cuanto tiempo y cuantos pasos se requieren para que una persona realice una tarea básica? (Ordenar Productos, Buscar, etc).
- Exactitud -- ¿Cuantos errores realizaron las personas? (Fueron errores fatales, recuperables, etc).
- Recuerdo-- ¿Que tan bien se recuerda el uso del programa, después de largos periodos de no utilizarlo?
- Respuesta Emocional -- ¿Como se siente una persona al completar una tarea? ¿Estresada, Confiada, Emocionada?
[![](http://1.bp.blogspot.com/_ayvorITawE4/SbB85fg5RTI/AAAAAAAAB6E/d0x_cYeYKNI/s320/usability_testing_absplit.jpg)](http://1.bp.blogspot.com/_ayvorITawE4/SbB85fg5RTI/AAAAAAAAB6E/d0x_cYeYKNI/s1600-h/usability_testing_absplit.jpg)"Es necesario conocer que piensa el usuario final del software que usara"

¿Fácil no? ¿Ves a donde vamos con la implementar de una Prueba de Uso? Con estas pruebas, nos estamos asegurando que:

1. No estas perdiendo tu tiempo ([Codificas lo que realmente se necesita](http://www.srbyte.com/2008/12/en-una-empresa-el-codigo-es-el-enemigo.html)) 2. Facilitas la vida del usuario final (Diseñar buenas GUI) 3. Tienes una retro-alimentación constante (Alimentas el proceso de desarrollo con criticas constructivas y mantienes los pies en la tierra) 4. Se mantienes una participación ACTIVA y proactiva del usuario y los developers (o el equipo de desarrollo). Y definitivamente, es algo que quieres estar haciendo cada quince días o un mes como máximo, no es necesario que cien usuario prueben constantemente el software, uno siempre es mejor que ninguno. A continuación, les describo la mas simple Prueba de Usabilidad que se puedan imaginar:

- Eliges un usuario final (operativo, gerencia, financiero, etc)
- Se aísla al usuario, y se le presenta el software.
- Se le asigna una o un par de tareas a realizar en el sistema.
- Se observa al usuario resolviendo la tarea.
- Se pregunta al usuario que piensa de la tarea (¿cuando esta en el proceso?, ¿como cree que puede mejorarlo?)
- Cuando el usuario termina la prueba (si la termina sin que le de un ataque cardíaco, por la frustración), se le pregunta lo siguiente: ¿Como se siente con respecto a la tarea? ¿Que recuerda del proceso? etc...
- Repita el proceso, con los usuarios aislados (siempre es mejor, probar a un solo usuario, que a varios para las Pruebas de Uso) y evalué en base a los resultados de las pruebas anteriores.
En el [blog de Robin Good](http://www.masternewmedia.org/es/2007/05/02/pruebas_de_usabilidad_analisis_del.htm), se puede encontrar en detalle una introducción al método "A/B Split Testing", les presento un extracto:

##

> ## "A/B
> split testing
> style="font-size:85%;">Por Lisa Halabi

> A partir de cualquier estudio de usabilidad de
> sitio web, generalmente se encuentran un número de href="http://www.masternewmedia.org/es/2006/12/13/pruebas_de_usabilidad_de_sitios.htm">problemas
> de usabilidad. A menudo puede suscitarse un debate dentro de una organización para
> hallar la solución a cada problema, sin que nadie conozca realmente la solución óptima.
> En vez de dejar que
> prevalezca la opinión de la persona que grita más style="font-size:85%;">, una mejor opción puede ser probar dos soluciones en un entorno en
> vivo. Aquella que tenga el mejor rendimiento es claramente la solución superior. ¡Bienvenido
> al A/B split testing! href="http://www.masternewmedia.org/es/2007/05/02/pruebas_de_usabilidad_analisis_del.htm">...
> click aquí para leer mas."
>

Aquí la clave (resaltada en negritas) es no dejar que prevalezca la opinión del que grita mas, sino la mejor opción, pero para el usuario.

[![](http://4.bp.blogspot.com/_ayvorITawE4/SbB9Vh8ntZI/AAAAAAAAB6M/dlKMuqkiKb8/s320/steveDontMakeMeThink.jpg)](http://4.bp.blogspot.com/_ayvorITawE4/SbB9Vh8ntZI/AAAAAAAAB6M/dlKMuqkiKb8/s1600-h/steveDontMakeMeThink.jpg) Steve Krug en su libro [Don't Make Me Think!](http://www.amazon.com/Think-Common-Sense-Approach-Usability/dp/0789723107), en el capitulo nueve, plantea una serie de aclaraciones sobre las Pruebas de Uso, que toda persona que este encargada de crear software debe de leer, pueden encontrar el capitulo 9 (en Ingles) de la primera edición de este genial libro [dando clic aquí.](http://sensible.com/Downloads/DMMTchapter09_for_personal_use_only.pdf)

Espero, sinceramente, que la próxima vez que estén en un proyecto de software, que se preocupen constantemente, en la opinión del usuario final, final, final (el que realmente usara tu software todos los dias, para hacer su vida mas fácil). Para que de esta forma, realmente estés produciendo software enfocado al uso o a la "usabilidad".

Saludos!

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte