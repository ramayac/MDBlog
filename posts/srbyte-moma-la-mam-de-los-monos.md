---
title: MoMA, la mamá de los monos...
date: 2007-07-20
author: Rodrigo Amaya
tags: libre, gnu, linux, .net, software, mono
post_id: blog-3515952828243908885.post-5401259346683657008
---

[![](http://bp3.blogger.com/_ayvorITawE4/Rp-dLMQ69AI/AAAAAAAAAa4/ydMjyGgRhcA/s400/mono.gif)](http://bp3.blogger.com/_ayvorITawE4/Rp-dLMQ69AI/AAAAAAAAAa4/ydMjyGgRhcA/s1600-h/mono.gif)MoMA o Mono Migration Analyzer, es una herramienta que ayuda a identificar diferentes problemas que se tengan al portar las aplicaciones del Framework .NET al Framework de Mono. Aunque MoMA pueda ayudar a enseñar problemas potenciales, existen muchos factores completos que no pueden ser previstos mediante el uso de una sola herramienta. Pero la utilidad de MoMA es innegable. Si tienes una aplicación hecha en .NET que no pase ninguna prueba de MoMA... entonces ya tienes una idea del trabajo que tendrás que hacer para migrar esa aplicación al Mono Framework.
> Recordemos que la verdadera prueba
> resulta al correr tu aplicación sin problemas en Mono.
Pero para que se acostumbren a codificar bien, y tengan una idea de como usar MoMA, aquí esta una pequeña guia, ¡espero que les sea de utilidad!

nota aclaratoria: las capturas de pantallas ajustadas a menor tamaño no son tan importantes o interesantes como las de tamaño "grande", esto es claro, bajo el criterio del autor. Como siempre, pueden dar click a cualquier imagen para verla en tamaño real.

Guía ejemplo de uso de MoMA:

1.0 Descargando MoMA

Puedes descargar MoMA de [http://www.mono-project.com/Moma](http://www.mono-project.com/Moma) , luego ve a la seccion de descargas "Download" y allí encontraras el vinculo de descarga.

[![](http://bp3.blogger.com/_ayvorITawE4/Rp-FXMQ68zI/AAAAAAAAAZQ/FYAOvDG9oAE/s400/momamono.png)](http://bp3.blogger.com/_ayvorITawE4/Rp-FXMQ68zI/AAAAAAAAAZQ/FYAOvDG9oAE/s1600-h/momamono.png)"Captura de pantalla del sitio web de MoMA con el URL marcado (muy marcado) para descargarlo"

Recuerden que MoMA necesita el framwork .NET 2.0 ó Mono 1.2+, como solo estoy usando GNU\Linux, usare Mono para correr MoMA.

1.1 Ejecutando MoMA

Después de descargar el archivo .Zip de MoMA, extraemos su contenido y podremos ver lo siguiente:

[![](http://bp3.blogger.com/_ayvorITawE4/Rp-J-MQ680I/AAAAAAAAAZY/BJoqgNixcac/s400/momamono-carpeta.png)](http://bp3.blogger.com/_ayvorITawE4/Rp-J-MQ680I/AAAAAAAAAZY/BJoqgNixcac/s1600-h/momamono-carpeta.png)"Captura de pantalla del contenido del archivo .Zip descargado"

Ahi estan los archivos que necesitamos, muchos diran: ah!, son simples ejecutables de windows. Pues noooo, estos archivos "exe" estan compilados en un lenguaje intermedio (como el bytecode de Java) llamado MSIL (MicroSoft Intermedium Languaje). Y las extensiones "exe" y "dll" pues solo son extensiones re utilizadas. Vamos a usar Mono para correr/interpretar/compilar (como le quieran decir) el archivo "MoMA.exe". Abrimos una consola (en mi caso usare Konsole) y nos vamos a la carpeta donde se encuentra "MoMA.exe", ahi ejecutamos el comando:

> mono MoMA.exe
Aquí hay una captura por si se pierden:

[![](http://bp2.blogger.com/_ayvorITawE4/Rp-Na8Q681I/AAAAAAAAAZg/jZ-1t2cHnec/s400/momamono-ejecutando.png)](http://bp2.blogger.com/_ayvorITawE4/Rp-Na8Q681I/AAAAAAAAAZg/jZ-1t2cHnec/s1600-h/momamono-ejecutando.png)"Captura de la consola, usando mono para ejecutar MoMA.exe"

Y si todo sale bien, pues podremos ver la ventana de MoMA en nuestra pantalla:

[![](http://bp0.blogger.com/_ayvorITawE4/Rp-N3cQ682I/AAAAAAAAAZo/AFz4MpjhEHA/s400/moma-ejecutando.png)](http://bp0.blogger.com/_ayvorITawE4/Rp-N3cQ682I/AAAAAAAAAZo/AFz4MpjhEHA/s1600-h/moma-ejecutando.png)"Captura de MoMA en GNU\Linux, usando KDE como Desktop Manager"

1.2 Usando MoMA

Para usar MoMA solo necesitamos ejecutables hechos con el Framework .NET y dar unos cuantos clicks. Yo usare de ejemplo la librería de un amigo para comprobar si puede ser ejecutado o no con Mono. Ustedes pueden usar cualquier otro compilado (el que ustedes quieran) que corra en el Framework .NET 1.0 y 2.0, y que evidentemente no sea una librería especifica a la plataforma o parte de .NET (como System.dll). Esta es la librería CustomDataSource(CDS) Ver.1.0 de Roberto Linares, que puede ser encontrada [aquí](http://ingenieria.uesocc.edu.sv/foro/viewtopic.php?t=77).

[![](http://bp3.blogger.com/_ayvorITawE4/Rp-QxMQ683I/AAAAAAAAAZw/llaOTQvs_LQ/s400/descargandoCDS-bin.rar.png)](http://bp3.blogger.com/_ayvorITawE4/Rp-QxMQ683I/AAAAAAAAAZw/llaOTQvs_LQ/s1600-h/descargandoCDS-bin.rar.png)"Captura de pantalla, descarga de Bin.rar de CDS, foro ingenieria UES-FMOcc"

En el paso 2 (de 4) de MoMA, podemos agregar uno o varios ejecutables hechos con el Framework .NET, asi que despues de descargar los binarios de CDS, y extraerlos en mi escritorio, agregare el archivo "CrypterClassLib.dll" dando click en el boton con signo "+" y luego eligiendo el/los archivo a agregar, asi:

[![](http://bp3.blogger.com/_ayvorITawE4/Rp-RQMQ685I/AAAAAAAAAaA/iN1rMgwlOqw/s400/libroberto-moma.png)](http://bp3.blogger.com/_ayvorITawE4/Rp-RQMQ685I/AAAAAAAAAaA/iN1rMgwlOqw/s1600-h/libroberto-moma.png)"Captura de pantalla, agregando archivos para analizar con MoMA"

[![](http://bp2.blogger.com/_ayvorITawE4/Rp-RU8Q686I/AAAAAAAAAaI/iLUtlZMJHbk/s400/libroberto-moma-selec.png)](http://bp2.blogger.com/_ayvorITawE4/Rp-RU8Q686I/AAAAAAAAAaI/iLUtlZMJHbk/s1600-h/libroberto-moma-selec.png)"Captura de pantalla, ¿Listo para analizar?"

Ya con la lista de archivos que seran analizados, pues solo basta dar click en "Next"; para que MoMA analize el ejecutable y podamos ver un resumen de compatibilidad del archivo con mono. Al final obtendremos un resultado como este:

[![](http://bp0.blogger.com/_ayvorITawE4/Rp-TEcQ687I/AAAAAAAAAaQ/1J0W0KyEMGY/s400/moma-todobien.png)](http://bp0.blogger.com/_ayvorITawE4/Rp-TEcQ687I/AAAAAAAAAaQ/1J0W0KyEMGY/s1600-h/moma-todobien.png)"Captura de Pantalla, ¡Te salvaste Roberto!"

En este paso (3 de 4) podemos ver que en el análisis general, el .dll analizado paso las pruebas con exito. Asi que en teoria, el archivo "CrypterClassLib.dll" puede ejecutarse sin problemas con el Mono Framwork. Vaya Roberto... te salvaste del bochorno publico jaja ;)

Pero... ¿Que sucede cuando TODO sale mal?

[![](http://bp0.blogger.com/_ayvorITawE4/Rp-U_cQ689I/AAAAAAAAAag/hFekSl1mfj8/s400/moma-mal.png)](http://bp0.blogger.com/_ayvorITawE4/Rp-U_cQ689I/AAAAAAAAAag/hFekSl1mfj8/s1600-h/moma-mal.png)"Captura de Pantalla, ¡OH NO!"

Bien, si sucede algo tragico como lo muestra la captura anterior, o peor... pues MoMA genera un util reporte (paso 4 de 4) que muestra de manera general, que se puede hacer para combatir esta incompatibilidad.

[![](http://bp2.blogger.com/_ayvorITawE4/Rp-Vy8Q68-I/AAAAAAAAAao/_26IwS-qLSU/s400/Report.jpg)](http://bp2.blogger.com/_ayvorITawE4/Rp-Vy8Q68-I/AAAAAAAAAao/_26IwS-qLSU/s1600-h/Report.jpg)"Captura reporte de ejemplo"

Con este reporte en mano y teniendo en cuenta ciertos tips de migracion/codificacion seremos capaces de migrar una aplicacion casi (especulativamente) sin problemas de un S.O. a otro.

1.3 Consideraciones finales

No olviden enviar el reporte (sea bueno o malo) de compatibilidad que genera MoMA (MoMA lo envia automáticamente):

[![](http://bp2.blogger.com/_ayvorITawE4/Rp-Xv8Q68_I/AAAAAAAAAaw/2ZT4mDOU638/s400/envio-reporte.png)](http://bp2.blogger.com/_ayvorITawE4/Rp-Xv8Q68_I/AAAAAAAAAaw/2ZT4mDOU638/s1600-h/envio-reporte.png)"Captura de pantalla, envío de reporte a Mono Project"

Esto ayuda para que los desarrolladores de Mono coloquen especial énfasis en los problemas mas comunes de compatibilidad encontrados, es tan fácil como dar un click al botón "Submit Report".

Una guía mas completa y extensa para corregir problemas de interoperabilidad puede ser hallada [aquí](http://www.mono-project.com/Guide:_Porting_Winforms_Applications).

Hasta Luego!

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte