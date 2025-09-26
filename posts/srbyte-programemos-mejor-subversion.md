---
title: Programemos Mejor: Subversion
date: 2008-03-26
author: Robertux
tags: subversion, programacion, svn, herramienta
post_id: blog-3515952828243908885.post-5173755952151884031
---

[![](http://bp0.blogger.com/_jH77WNrMVRA/R-qoq_a907I/AAAAAAAAAw0/r4Ho663evOs/s320/subversion_logo_hor-468x64.png)](http://bp0.blogger.com/_jH77WNrMVRA/R-qoq_a907I/AAAAAAAAAw0/r4Ho663evOs/s1600-h/subversion_logo_hor-468x64.png)"Logo Oficial de Subversion"

¿Les ha pasado alguna vez que cuando están programando se dan cuenta que las líneas de código que acaban de agregar arruinaron el sistema y desean volver a como lo tenían el día de ayer ya que en ese entonces todavía compilaba, pero ya no se acuerdan qué fue lo último que agregaron para así poder revertirlo?

¿Les ha ocurrido que cada cierto tiempo crean una copia de la carpeta del proyecto en el que trabajan para guardarla como backup y además de que cada copia les abarca más de 10 o 20 MB de espacio en disco, al final no saben si la última versión está en la carpeta "ProyectoUltimo", "ProyectoFinal" o "ProyectoBueno" y les toca comparar las fechas de cada una?

¿Será que cuando trabajan en grupos, cada quién con su copia del proyecto y modificando los archivos que a cada quién le corresponden, al final no saben ni por dónde empezar para unir todos los archivos correctos en un único proyecto para tener la versión final y funcional?

Todas estas situaciones pasan porque no se están utilizando herramientas para el trabajo en grupo y específicamente, para el control de versiones.

Los sistemas de control de versiones son muy populares, más que todo en el mundo del software libre, ya que bajo esta filosofía, los desarrolladores permiten a los demás tener acceso libre (no necesariamente gratuito) al código fuente además de los binarios o ejecutables. Por lo tanto, estos utilizan herramientas que les permiten controlar de una mejor manera el desarrollo y distribución del mismo y cómo este va cambiando a través del tiempo.

¿Qué son los Sistemas de Control de Versiones?

Un sistema de control de versiones se encarga de almacenar de la manera más apropiada los cambios que ocurren sobre un conjunto de archivos en intervalos determinados y centralizados en un repositorio. Como ejemplo de estos sistemas se encuentran CVS (Concurrent Version Control), Visual Source Safe y por supuesto, Subversion. De acá en adelante, todo el contenido se basará específicamente en Subversion.

¿Cómo Funcionan los Sistemas de Control de Versiones?

Se posee un repositorio de svn montado en un servidor local o remoto en el cual se almacenan todas las versiones, desde la inicial. Cada cliente desde su computadora descarga las revisiones que necesite del repositorio y empieza a trabajar en ellas. Cuando ha hecho los cambios, este actualiza el repositorio. Por supuesto, debido a la concurrencia es posible que las versiones de los archivos entren en conflicto si ambos clientes han hecho cambios sobre los mismos archivos. Por ejemplo si dos personas descargan uno o varios archivos y los modifican simultáneamente, el primero en actualizar no tendrá problemas pero el segundo, si ha modificado los mismos archivos, tendrá ciertas dificultades cuando desee actualizar ya que la versión original que el poseía en su computadora ya no coincide con la del servidor.

Esto se resuelve mediante el bloqueo de los archivos por parte de un cliente para que los demás no puedan modificarlo o bien, llegando a un acuerdo entre los que modificaron simultáneamente el archivo para determinar cuál merece ser la última versión. Para ello, Subversion provee herramientas que examinan el contenido de un archivo y lo comparan con otras versiones para remarcar las líneas agregadas, modificadas y/o eliminadas por parte de todos los que han modificado el archivo de manera que sea mas sencillo reconocer los cambios y decidir cuáles deben permanecer.

Tan simple como eso.

Cabe mencionar que Subversion no almacena copias enteras de cada versión almacenada sino que evalúa los cambios realizados entre una versión y otra y solo almacena estos para ahorrar espacio en el repositorio.

Existen diversos clientes de Subversion para poder descargar versiones de un repositorio y actualizarlas en base a nuestros cambios. Para Linux existe la aplicación de línea de comandos llamada [Subversion](http://packages.debian.org/subversion) y para Windows existe una denominada [TortoiseSvn](http://tortoisesvn.tigris.org/), ambas de licencia libre y proporcionadas por los mismos creadores del Subversion: [Tigris](http://www.tigris.org/).

Para más información sobre Subversion, pueden leer de forma gratuita el documento on-line [Control de Versiones con Subversion](http://svnbook.red-bean.com/). El cual también esta disponible a la venta en Amazon.

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte