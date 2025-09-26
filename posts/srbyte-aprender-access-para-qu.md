---
title: Aprender ACCESS ¿para qué?...
date: 2008-08-24
author: Rodrigo Amaya
tags: libre, desarrollo, microsoft, base de datos
post_id: blog-3515952828243908885.post-5235056442898263809
---

Realmente no entiendo, ¿para que enseñan Access en las Universidades?. Creo sinceramente (y si me equivoco, pueden dejar los comentarios para intentar convencerme) que la única forma de que una persona tenga una experiencia "agradable" con Access, es que seas un novato absoluto con las bases de datos o un usuario común y silvestre.

> Access es a SQL lo
> que VBA es para VB.NET
Muchos desarrolladores (Salvadoreños al menos) estarán pensando en que Access es "bueno" para sistemas sencillos que necesitan bases de datos pequeñas. La respuesta a eso, es NO. Permitanme explicarles: Microsoft Access utiliza un modulo de acceso a datos llamado Microsoft JET. Las siglas JET significan: Joint Engine Technology, el sucesor de esta tecnología fue primero MSDE (Microsoft Desktop Engine) y este, a su vez, fue sucedido por SQL Server 2005 Express y recientemente en el SQL Server 2005 Compact Edition.

[![](http://1.bp.blogspot.com/_ayvorITawE4/SLLRZnumzlI/AAAAAAAABIw/FWHCXzyz570/s320/200px-MicrosoftJet.gif)](http://1.bp.blogspot.com/_ayvorITawE4/SLLRZnumzlI/AAAAAAAABIw/FWHCXzyz570/s1600-h/200px-MicrosoftJet.gif)"Guía del programador de JET"

Desde un punto de vista tecnológico: [JET es considerado como una tecnologia desfazada por Microsoft](http://msdn.microsoft.com/en-us/library/ms810810.aspx#mdac%20technologies%20road%20map%20old_topic9). Es más, JET ya no se distribuye con la ultima versión de MDAC (Microsoft Data Access Components). Sin embargo, JET sigue siendo el motor de bases de datos de Microsoft Access 2007 y lo seguirá hasta el fin de los tiempos (o hasta que decidan cambiarlo). ¿Comprenden ahora en problema?

> ¿Cual es la "enseñanza" de utilizar un programa con una base
> tecnológica desfasada?
Sin mencionar el común problema, que las bases de datos tienen un limite de 2GB de capacidad. Tal parece que algunas personas desean enseñar los principios fundamentales de la creación de bases de datos relacionales (normalización, integridad referencial, etc...) utilizando Microsoft Access... es ridículo.

[![](http://3.bp.blogspot.com/_ayvorITawE4/SLLRZzTfnZI/AAAAAAAABJA/UEfZQDqtY_s/s320/WhySoftwareSucks.jpg)](http://3.bp.blogspot.com/_ayvorITawE4/SLLRZzTfnZI/AAAAAAAABJA/UEfZQDqtY_s/s1600-h/WhySoftwareSucks.jpg)"¿Por qué enseñan software que apesta?"

¿Como se soluciona este problema? Utilizando herramientas reales de desarrollo. En vez de gastar dinero en licencias (si es que lo gastan) de Microsoft Office 2007, se recomienda usar [SQL Server 2005 Express](http://www.microsoft.com/sql/editions/express/default.mspx) para este propósito. Y si lo que les gusta de Access, es que pueden elaborar bases de datos de forma gráfica, entonces usen: SQL Server 2005 Management Studio Express, que les permitir hacer exactamente lo mismo. Ambas herramientas son "freeware", se pueden descargar sin cargo alguno, se pueden redistribuir sin problemas de licencias y, por experiencia personal, es lo suficientemente sencillo para que cualquiera lo use. Y si están pensando en bases de datos portables, rápidas y sin limites de tamaño, piensen en alternativas libres como: [SQLite](http://www.sqlite.org/) y [Apache Derby](http://db.apache.org/derby/quick_start.html).

[![](http://3.bp.blogspot.com/_ayvorITawE4/SLLRZuN8P2I/AAAAAAAABI4/-js5AEs2CMM/s320/SQLite.gif)](http://3.bp.blogspot.com/_ayvorITawE4/SLLRZuN8P2I/AAAAAAAABI4/-js5AEs2CMM/s1600-h/SQLite.gif)"SQLite es la solución a los problemas de bases de datos portables"

Yo use Access con VB 6.0, y desde entonces me propuse no contaminar mi mente nunca más con ese tipo de "soluciones informáticas", y claro, estoy arrepentido de haberlo usado, jaja. Y tu: ¿alguna vez utilizaste Access para desarrollar Software?

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte