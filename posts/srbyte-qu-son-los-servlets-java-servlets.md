---
title: ¿Qué son los Servlets? (Java Servlets)
date: 2008-12-28
author: Rodrigo Amaya
tags: java, programacion
post_id: blog-3515952828243908885.post-8517962772108876376
---

Advertencia, lenguaje geek a continuación:

¿Qué es un Servlet?

De manera fundamental, un Java Servlet (o solo: Servlet) es una aplicación que extiende la funcionalidad de un Servidor. Un Servlet no esta confinado exclusivamente a un servidor web, pero casi siempre se refiere a un Servlet a este contexto. Casi todas las descripciones o referencias de los Servlets, los citan como un reemplazo de los CGI, por eso no es extraño que sea más fácil pensar en un Servlet como un programa Java que realiza funciones de CGI...

"Introducción Java Servlets"

¿Un Reemplazo de los CGI?

Uno de los temas mas intrigantes, citados y curiosos de los Servlets es su aclamada mejora en el rendimiento versus losCGI. Un CGI tradicional es escrito en Perl (script CGI), C (CGI compilado), etc... y su desventaja es que para servir a cada usuario (usualmente un http request) conectado con el servidor donde se aloje el CGI, el sistema tendrá que crear un nuevo proceso de ese mismo CGI... si hay 10,000 usuarios conectados al servidor solicitando utilizar el CGI "Hola Mundo", este iniciara 10,000 procesos para servir a cada usuario.

[![](http://4.bp.blogspot.com/_ayvorITawE4/SVenYuEEpeI/AAAAAAAAB1I/mph7ettKGKU/s320/javaservletur2.jpg)](http://4.bp.blogspot.com/_ayvorITawE4/SVenYuEEpeI/AAAAAAAAB1I/mph7ettKGKU/s1600-h/javaservletur2.jpg)"Para la próxima, añade un libro de Java a la carta que haces a Santa"

El impacto sucede con CGI que se manejan a niveles corporativos, con miles de solicitudes de usuarios por segundo (req/seg), de esa manera se sobrecarga el CPU y llena la memoria RAM a tal grado que se gasta más tiempo administrando los procesos que sirviendo los usuarios, lo que se traduce a una carga innecesaria para el servidor y que puede influir en el rendimiento del mismo.

Un Servlet resuelve este problema, creando un Thread o Hilo en cada request, en vez de iniciar un proceso entero. Así, un proceso de crea para cada Servlet, y para cada usuario que invoca a ese servlet (para cada request) se crea un Hilo.

Usualmente un servlet se une en conjunto con un JSP, en un patrón de codificación lógico llamado: "Modelo 2", o patrón MVC (Model View Controller). Además, como están escritos en Java, los servlets son portables entre servidores (web y de aplicaciones) y entre sistemas operativos.

La API para programar Servlets (Java Servlet API) es una parte "estandar" de la plataforma J2EE.

[![](http://4.bp.blogspot.com/_ayvorITawE4/SVek339_MMI/AAAAAAAAB1A/Vg69jkGx0-8/s320/Wave.png)](http://4.bp.blogspot.com/_ayvorITawE4/SVek339_MMI/AAAAAAAAB1A/Vg69jkGx0-8/s1600-h/Wave.png)La utilidad de conocer que son o como funcionan los Servlets radica en que se utilizan enormemente, tanto en el ámbito empresarial y como en la red. Si estas buscando que tecnología conocer para conseguir trabajo, pues te sugerimos que comiences por aprender a utilizar Servlets y JSP, que es un requerimiento en la mayoría de empresas formales para conseguir trabajo aquí en El Salvador.

Algunos recursos adicionales sobre los Servlets:

- [Sun's servlet tutorial](http://java.sun.com/j2ee/tutorial/1_3-fcs/doc/Servlets.html)
- [Sun's servlet product description](http://java.sun.com/products/servlet)
Saludos!

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte