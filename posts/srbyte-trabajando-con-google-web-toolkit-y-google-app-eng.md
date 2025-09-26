---
title: Trabajando con Google Web Toolkit y Google App Engine
date: 2009-10-24
author: Robertux
tags: ruby, java, linux, framework, ajax, python, google, appengine
post_id: blog-3515952828243908885.post-3289853482720812936
---

[![](http://4.bp.blogspot.com/_jH77WNrMVRA/SuJ5sYPGDtI/AAAAAAAAGEE/sN8caM4EpFk/s400/gappengine.gif)](http://4.bp.blogspot.com/_jH77WNrMVRA/SuJ5sYPGDtI/AAAAAAAAGEE/sN8caM4EpFk/s1600-h/gappengine.gif) "Logotipo del framework Google App Engine"

Habiendo salido de mi tesis recientemente, me he interesado en buscar hobbies relacionados con programación y en plasmar muchas ideas de software que vinieron a mi cabeza mientras estaba ocupado haciendo el trabajo de grado. Uno de los inconvenientes que tenemos los desarrolladores Java es que no existen en internet tantas alternativas para hosting de aplicaciones así como la gran variedad disponible para el Stack LAMP (Linux/Apache/Mysql/PHP).

Además, cada programador Java está acostumbrado a trabajar con su propio Stack y es bastante diversa la gama de combinaciones que se pueden armar entre sistema operativo, application server y framework para el entorno visual así como la base de datos y el framework para la persistencia de los datos, esto sin mencionar las versiones de cada una de estas soluciones informáticas. Esto hace aun mas difícil encontrar entornos (pagados o gratuitos) donde alojar nuestras aplicaciones web a menos que nos decidamos a montar nuestro propio server.

Buscando en internet me dí cuenta que Google ofrece una alternativa sencilla y barata (gratuita hasta cierto punto) para alojar aplicaciones desarrolladas con Java: [Google App Engine](http://code.google.com/appengine/).

Este servicio de google te permite desarrollar aplicaciones relativamente sencillas y hostearlas dentro de los servidores que ellos mismos te ofrecen. Eso si, no puede ser cualquier tipo de aplicación y lo mejor sería desarrollar aplicaciones orientadas a esta plataforma de Google, ya que estos ya te ofrecen la alternativa para la persistencia de los datos y los frameworks visuales a utilizar.

La ventaja es que el soporte se ve bastante decente (te brindan un plugin para Eclipse que te hace la mayor parte del trabajo de configuración y el upload de tu aplicación) además que la documentación te explica el framework de una manera bastante sencilla. En la [galería de aplicaciones](http://appgallery.appspot.com/) puedes ver las aplicaciones que otros ya han publicado, notando que algunas pueden ser bastante complejas y de un acabado muy profesional.

Lo único a tener en cuenta es que nada mas puedes subir 10 aplicaciones como máximo, y no puedes borrar una aplicación que ya has subido al server de Google.

Otro tema que me llamó la atención al navegar entre las soluciones de desarrollo de google fue el GWT, [Google Web Toolkit](http://code.google.com/webtoolkit/). Este es un framework visual para apps web que te permite desarrollar aplicaciones sin la necesidad de escribir HTML o Javascript, generándolo todo desde clases Java. Es un framework bastante joven comparado con JSF, Struts o Tapestry pero teniendo el respaldo de google puede llegar lejos, ademas la idea de no escribir HTML será un alivio para muchos desarrolladores web que no saben mucho de estructuración y decoración de sitios web con HTML/CSS y desean una solución rápida y sencilla para sus interfaces. Al igual que con el App Engine Framework, Google apuesta nuevamente por eclipse brindando soporte para la creación de proyectos de GWT mediante asistentes y el uso de un navegador embedido que te permite realizar pruebas de tus aplicaciones sin necesidad de hacer deploy.

[![](http://2.bp.blogspot.com/_jH77WNrMVRA/SuKFrnovl0I/AAAAAAAAGEM/Ri5ZYz6R2lo/s400/sampleGWTapp.png)](http://2.bp.blogspot.com/_jH77WNrMVRA/SuKFrnovl0I/AAAAAAAAGEM/Ri5ZYz6R2lo/s1600-h/sampleGWTapp.png)"Aplicación de ejemplo de uso de GWT, creada a partir de un proyecto web de eclipse con el plugin de GWT y App Engine. En la captura pueden apreciar el browser embedido que incluye el framework de GWT."

Cabe mencionar que GWT puede ser usado nada más para generar el HTML/Javascript de un sitio, usando cualquier otro lenguaje de programación del lado del server. Como ellos bien lo mencionan, este puede ser usado para generar front-ends en aplicaciones Ruby, Python, etc. Además es full Ajax-enabled trabajando de forma transparente con invocaciones a código del lado del servidor hecho con Java y mediante el uso de XML-RPC para otros lenguajes.