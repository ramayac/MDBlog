---
title: ¿Esta versionada tu Base de Datos?
date: 2008-05-14
author: Rodrigo Amaya
tags: subversion, programacion, svn
post_id: blog-3515952828243908885.post-1925728383852479214
---

Atencion: esta entrada es para programadores... y geeks en general.

Como mencionaba [Robertux](https://www.blogger.com/profile/15615123126956711175) en [una entrada anterior](https://srbyte.blogspot.com/2008/03/programemos-mejor-subversion.html):

> """¿Les ha
> pasado alguna vez que cuando están programando se dan cuenta que las líneas de código que
> acaban de agregar arruinaron el sistema y desean volver a como lo tenían el día de ayer ya que
> en ese entonces todavía compilaba, pero ya no se acuerdan qué fue lo último que agregaron para
> así poder revertirlo?
> ¿Les ha ocurrido que cada cierto tiempo crean
> una copia de la carpeta del proyecto en el que trabajan para guardarla como backup y además de
> que cada copia les abarca más de 10 o 20 MB de espacio en disco, al final no saben si la
> última versión está en la carpeta "ProyectoUltimo", "ProyectoFinal" o "ProyectoBueno" y les
> toca comparar las fechas de cada una?
> ¿Será que cuando trabajan en
> grupos, cada quién con su copia del proyecto y modificando los archivos que a cada quién le
> corresponden, al final no saben ni por dónde empezar para unir todos los archivos correctos en
> un único proyecto para tener la versión final y funcional?
> Todas estas
> situaciones pasan porque no se están utilizando herramientas para el trabajo en grupo y
> específicamente, para el control de versiones."""

Y lo mismo podríamos decir de las Bases de Datos. Así que, developers, con esta idea en mente les pregunto:
> ¿Esta tu base de datos bajo control
> de version (cvs ó svn)?
...la respuesta a esta pregunta debería de ser (en los casos que lo amerite) SI. ¿Por que? simplemente por que:
> ¡la base de datos es una de las partes más criticas de cualquier
> aplicación! (la base de datos es tan parte de la aplicación, como el código y los modelos
> dentro del software)
y en la etapa de desarrollo, con varias personas trabajando en un proyecto, es muy probable que se cometan errores como los que mencionábamos al principio... o peores. ¿Y como versiono una base de datos?, bien, aquí hay un pequeño ejemplo:

Escenario de trabajo:

> Gestor de Base de Datos: style="font-style: italic;">MySQL.
> Repositorio SVN: style="font-style: italic;">Google Code.
> Herramientas a usar
> (multiplataforma): SVN y MySQL
> (dump).
Para versionar una base de datos en mysql, basta con versionar el dump de la base de datos. Y el proceso de versionado ("Commit" y "Update" de cambios) y restauración de la BD se realiza con dos sencillos script (estos scripts pueden estar en una carpeta que se llame SQL y ser parte del proyecto que contenga esos scripts), uno script sera para realizar el "Commit" y otro para el "Update".

La logica del script de "Commit" es la siguiente: 1. Despues de realizar cambios significativos en la base de datos... 2. Llama a "mysqldump", y realiza un respaldo de la base de datos (con su contenido) en un archivo sql, por ejemplo:

> mysqldump
> --single-transaction -hlocalhost -uROOT -pTOOR BASEDEDATOS >
> basededatos.sql

3. Luego realiza un "Commit" del archivo, asi:

> svn commit -m "Dump de base de datos versionado"
> basededatos.sql

La logica del script de "Update" es la siguiente: 1. Realiza un "Update" invocando a svn...

> svn update
2. Restauramos el dump actualizado obtenido a mysql:

> mysql.exe -uROOT -pTOOR
> BASEDEDATOS < basededatos.sql

No hay que preocuparse por la información de Login para el svn, ya que si la carpeta en la que se invoca el script de commit o update esta agregada al repositorio, svn crea una carpeta llamada ".svn" que contiene la información de login (y otras cosas). Como les mencione, esa es la lógica, si yo hago un cambio a la base de datos, hago el commit, y si hay cambios (o antes de una sesión de trabajo) hago el update.

Si esta entrada te pareció útil, [aquí hay algunas reglas para trabajar con bases de datos que deberías de leer.](https://srbyte.blogspot.com/2008/05/3-reglas-al-trabajar-con-bases-de-datos.html)

Espero que esto les sirva para facilitar su trabajo de desarrollo de software, ¡Saludos!