---
title: ¿Qué es ORM?
date: 2009-09-02
author: Rodrigo Amaya
tags: sql, hibernate, toplink, orm, oracle
post_id: blog-3515952828243908885.post-7958704684122756259
---

Object Relational Mapping, u ORM, O/RM y O/R mapping, es una técnica empleada en la programación, para convertir datos entre sistemas incompatibles, como lo son las bases de datos relacionales y los lenguajes de programación. Esta conversión de datos entre los sistemas crea un efecto una base de datos virtual de objetos, que puede ser usada en el programa (en esa forma).

[![](http://3.bp.blogspot.com/_ayvorITawE4/SqB0NsSSqzI/AAAAAAAACKA/wlMDC_4R0Ls/s320/ORM-Overview.png)](http://3.bp.blogspot.com/_ayvorITawE4/SqB0NsSSqzI/AAAAAAAACKA/wlMDC_4R0Ls/s1600-h/ORM-Overview.png)

Hay implementaciones comerciales y libres disponibles para crear el "mapeo" (mapping) objeto-relación, aunque algunos programadores (o mejor dicho empresas) optan (por ignorancia o espiritu de aventura... o ambas quizas) por crear sus propias herramientas ORM.

[![](http://2.bp.blogspot.com/_ayvorITawE4/SqB0O3gwMMI/AAAAAAAACKQ/3FSryZ57OF8/s320/fig02.jpg)](http://2.bp.blogspot.com/_ayvorITawE4/SqB0O3gwMMI/AAAAAAAACKQ/3FSryZ57OF8/s1600-h/fig02.jpg) Las empresas siempre poseerán una base de datos normalizada, para "ahorrar espacio" (como algunos individuos administrativos lo ven). Para un programador, la tarea de leer estos datos, manipularlos y finalmente modificarlos o eliminarlos pende de un hilo, de acuerdo al grado de ignorancia a la hora de elegir a las herramientas y/o librerías de software (de ORM) empleadas para tales fines.

[![](http://4.bp.blogspot.com/_ayvorITawE4/SqB0PSl4TlI/AAAAAAAACKY/9B4p4sRxVYc/s320/hibernate.gif)](http://4.bp.blogspot.com/_ayvorITawE4/SqB0PSl4TlI/AAAAAAAACKY/9B4p4sRxVYc/s1600-h/hibernate.gif) Una librería de ORM (como [Hibernate](http://es.wikipedia.org/wiki/Hibernate), [Oracle Toplink](http://en.wikipedia.org/wiki/TopLink) o [Linq](http://es.wikipedia.org/wiki/Language_Integrated_Query)) siempre, absolutamente siempre reducirá la cantidad de código, porque habrá algo que permitirá realizar el proceso de mapeo (como el IDE), y se encargara de crear las clases equivalente u homologas con las tablas en la base, además permitirá manejar diversos tipos de relaciones entre las tablas (uno a uno, uno a muchos, etc), reducirá la cantidad de defectos en esta delicada area, y todo esto, para beneficio del programador, que se concentrara más en codificar la lógica del negocio, que en hacer "INSERT", "UPDATE", "DELETE" y "SELECT" en la base. Otra razón por la que una librería ORM reduce la cantidad de código, es porque permite centralizar los procesos de búsqueda de datos en la base, liberándonos de escribir [consultas ad-hoc innecesarias o "quemadas" en el código](http://www.srbyte.com/2009/04/la-importancia-de-los-procedimientos.html). Sin mencionar que, también gestionara el pool de conexiones a la base de datos. Todo para que el programador, no se convierta en un esclavo codificando algo que ya existe, ustedes ya saben que [en una empresa el codigo es el enemigo](http://www.srbyte.com/2008/12/en-una-empresa-el-codigo-es-el-enemigo.html)... y que [de nada sirve estar reinventando](http://www.srbyte.com/2009/03/si-no-estas-usando-un-framework.html) la rueda...

[![](http://4.bp.blogspot.com/_ayvorITawE4/SqB1g5GiTkI/AAAAAAAACKg/GRtX3evcniI/s320/coding_slave_cover.jpg)](http://4.bp.blogspot.com/_ayvorITawE4/SqB1g5GiTkI/AAAAAAAACKg/GRtX3evcniI/s1600-h/coding_slave_cover.jpg)

Queda en claro, que una librería ORM, generara el mapeo de tablas a clases de base de datos ([que esperamos que este BIEN diseñada](http://www.srbyte.com/2008/05/3-reglas-al-trabajar-con-bases-de-datos.html)) de una forma completamente automatizada. Netbeans por ejemplo, posee una excelente integración con JPA usando Oracle TopLink, y genera el código necesario para manipular toda la información de la base, en menos de un minuto... para 42 (cuarenta y dos) tablas.

¿Me pregunto cuanto se podría tardar una persona, haciendo el proceso a "pie"?

Si estas en un proyecto de software, en el que NO te permiten emplear librerías para ORM, eso simplemente refleja la ignorancia de tus inmediatos superiores o de los encargados de tu proyecto. Si ya tienes algo que te asista en el proceso, bien por ti!, pero deberías de estar pensando en emplear herramientas que son prácticamente el estándar de la industria ([Hibernate](http://es.wikipedia.org/wiki/Hibernate)), de comprobado rendimiento ([Oracle Toplink](http://en.wikipedia.org/wiki/TopLink)) y que existen, para que nadie tenga que codificar como esclavo, algo que se puede generar en un par de clics y en no mas de "cien segundos". ORM esta, para facilitar la vida de los programadores, reducir a la mínima expresión un proceso que es terriblemente tedioso, y también, para mejorar y producir mejor software.

[![](http://2.bp.blogspot.com/_ayvorITawE4/SqB0OEFXKFI/AAAAAAAACKI/xG_FclchQCM/s320/bettersoftware.jpeg)](http://2.bp.blogspot.com/_ayvorITawE4/SqB0OEFXKFI/AAAAAAAACKI/xG_FclchQCM/s1600-h/bettersoftware.jpeg) ¿Cuantos de ustedes utilizan tecnologías ORM en su trabajo o en la Universidad para proyectos de software?

Más información sobre ORM en la [Wikipedia](http://es.wikipedia.org/wiki/Mapeo_objeto-relacional).

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte