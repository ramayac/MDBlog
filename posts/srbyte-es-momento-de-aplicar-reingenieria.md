---
title: Es Momento de Aplicar Reingenieria
date: 2009-09-15
author: Robertux
tags: agil, computadoras, desarrollo, codigo, gmail, calidad, programacion, fox, empresa, productividad
post_id: blog-3515952828243908885.post-1715549923484754043
---

[![](http://4.bp.blogspot.com/_jH77WNrMVRA/Sq8nHKoeLqI/AAAAAAAAGDo/0R34Suup-zo/s400/EnConstruccion.jpg)](http://4.bp.blogspot.com/_jH77WNrMVRA/Sq8nHKoeLqI/AAAAAAAAGDo/0R34Suup-zo/s1600-h/EnConstruccion.jpg)"Las aplicaciones requieren constante mantenimiento y actualización. A veces la mejor actualización es su reescritura completa"

Como todos bien sabemos, todas las cosas especialmente los sistemas informáticos tienden a perder utilidad con el tiempo, ya sea debido a que no se adapta a los nuevos avances en hardware, no es compatible con las nuevas plataformas de ejecución o no logra satisfacer las nuevas necesidades del usuario.

De la misma manera como los activos fijos tienen su depreciación y fecha de expiración, cada programa de computadora tiene(o debería tener) definido su período de vida, de manera que los usuarios sepan cuándo sea el momento de reemplazarlos por otros más modernos. Lamentablemente no existe una unidad de medida infalible para saber cuándo un sistema informático ha expirado.

Cuando esto ocurre, es hora de aplicar reingeniería y empezar a rediseñar los sistemas existentes, manteniendo su funcionalidad actual pero utilizando herramientas de desarrollo mas ágiles, técnicas y disciplinas mas ordenadas y frameworks que permiten la extensibilidad del mismo, además de aprovechar para agregar nuevas características que pueda necesitar el usuario.

Si algo ya no sirve, vuélvelo a hacer desde cero.

Si es un programa hecho en Visual Fox Pro 6 que comparte archivos de tablas en una carpeta de red, aunque aún le sea útil al usuario, tú como programador sabes que será un completo dolor de cabeza tratar de consumir web services o transportar datos por Message Queue por lo que en lugar de seguir manteniendo un sistema pasado de moda desarrollado con [código obsoleto](http://www.srbyte.com/2009/04/deprecated-code-codigo-obsoleto.html), es mejor reescribirlo desde cero usando tecnologías que te ahorrarán mucho trabajo en el desarrollo y con capacidad de extenderlo según aparezcan nuevos estándares.

Desarrolla pensando en el futuro.

Como desarrollador puedo estar seguro que los usuarios no siempre saben lo que quieren que haga un sistema y cambian de opinión a medida que el sistema va siendo desarrollado. Por tal motivo, debes tener esto en cuenta a la hora de desarrollar tu sistema y diseñarlo de tal forma que pueda ser adaptable a posibles cambios, que tu sistema no "suponga" ni "imagine" que X o Y proceso se hace de tal manera, que todas las decisiones de negocio sean configurables! Ademas, permite que el programador que retomará tu sistema sea capaz de entenderlo y agregar nuevas funcionalidades que sean requeridas por los usuarios. Como una vez alguien escribió en [Stack Overflow Programming Quotes](http://stackoverflow.com/questions/58640/great-programming-quotes):

> Always
> code as if the guy who ends up maintaining your code will be a violent psychopath who knows
> where you live. (Siempre programa como si el que mantendrá tu código será un violento
> psicópata quien sabe donde vives)
>

-- Rick Osborne

Demuestrale al usuario que algo ha cambiado, y porqué este cambio es para mejora

Como escribí antes, el usuario está conforme y acostumbrado al viejo sistema del año 2000 que le resuelve a medias sus necesidades actuales y posiblemente el cambio que apliques no sea visible en la interfaz sino que solamente en las tecnologías de desarrollo. Aun así, hay que reflejar ese cambio también en la interfaz de usuario aplicando alguna nueva plantilla CSS(en el caso de una aplicación web), agregando nuevos servicios que quizá no eran necesarios pero que reflejen el cambio o agregando una sección de "Nueva versión, nuevas características" a manera de hacerle notar al usuario que algo ha cambiado y poder explicarle cuál es la nueva manera como ahora se realizan los procesos X y Y.

[![](http://4.bp.blogspot.com/_jH77WNrMVRA/Sq8X5TbVcnI/AAAAAAAAGDg/hlskcdMmvS0/s400/gmailnewstuff.JPG)](http://4.bp.blogspot.com/_jH77WNrMVRA/Sq8X5TbVcnI/AAAAAAAAGDg/hlskcdMmvS0/s1600-h/gmailnewstuff.JPG)"Los developers de Gmail siempre han tenido la bondad de notificarnos cuando hay nuevas características disponibles"

Recuerden que en última instancia, nuestro trabajo como desarrolladores es satisfacer los deseos más oscuros y enajenados las necesidades de información de los usuarios de negocio en la empresa, por lo que tampoco es bueno pensar en reescribir todos los sistemas de la empresa solo por estar "in" en tecnologías de desarrollo de software.