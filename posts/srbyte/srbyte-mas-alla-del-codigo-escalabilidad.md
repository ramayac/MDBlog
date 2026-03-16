---
title: Mas Alla del Codigo: Escalabilidad
date: 2009-01-20
author: Robertux
tags: elecciones, politica, computadoras, cloud, desarrollo, codigo, el salvador, programacion, twitter, empresa
post_id: blog-3515952828243908885.post-4883452891109540459
---

![image](https://1.bp.blogspot.com/_jH77WNrMVRA/SXOl-nRV4gI/AAAAAAAAFoQ/1Fz8NfsCHww/s400/scalability.jpg)     "Tu sistema debe ser capaz
de crecer junto con tu negocio, para soportar la demanda del mercado"

El año pasado, [Fafamonge](https://www.fafamonge.com/) me preguntó una vez cuánta influencia creía yo que podría tener el Internet y los blogs en las elecciones que estamos celebrando, a lo cual yo respondí que relativamente poca o casi nula, considerando que somos un país demasiado conservador con poca cultura tecnológica, en comparación con otros hermanos latinoamericanos como Chile o Brasil pero ahora que se celebraron las elecciones de diputados y alcaldes, quizá mi perspectiva podría variar un poco.

Me he dado cuenta que los encargados de IT del tribunal supremo electoral compartían mi idea y nunca esperaron la saturación en el tráfico que ha tenido el sitio web de dicho organismo durante las elecciones y al parecer, lo mismo podría suceder con el sistema utilizado para el conteo de votos y transmisión de resultados. Lo mismo nos ocurría en la universidad cuando realizabamos exámenes en línea o al momento de realizar inscripciones en línea. Los servidores donde la aplicación web estaba alojada colapsaba por el exceso de tráfico.

![image](https://3.bp.blogspot.com/_jH77WNrMVRA/SXOoemrnFrI/AAAAAAAAFoY/PUzrzV7QWPg/s400/fail_whale.gif)    
"El ejemplo más claro de falta de escalabilidad:
Twitter en sus inicios"

Para evitar estos inconvenientes, es necesario tomar en cuenta la cantidad de usuarios que harán uso de nuestras aplicaciones al momento de desarrollarlas y publicarlas, además de considerar si tu sistema será capaz de crecer para adaptarse a estas nuevas necesidades de rendimiento.

Lo más común es agregar mas recursos al servidor como mas RAM o un CPU más rapido. Otros optan por el [Cloud Computing](https://en.wikipedia.org/wiki/Cloud_computing), de manera que varias computadoras actuen como una sola y la suma de sus recursos sea suficiente para mejorar el rendimiento de la aplicación lo cual en parte si ayuda. Todo lo anterior es tarea de arquitectos de hardware/software y sysadmins pero nosotros como desarrolladores también tenemos que realizar nuestro aporte pensando en aplicaciones que consuman solamente los recursos necesarios y capaces de crecer y adaptarse a las nuevas necesidades cuando estas aparezcan, de manera que no sea necesario reescribir por completo el código fuente.

Como desarrollador, recomiendo hacer código lo mas genérico posible, no hagas código para solucionar problemas específicos en el momento que estos surgen, haz código que resuelva las necesidades de información actuales y las que podrían aparecer en un futuro y si no fuera capaz de resolverlas, que por lo menos sea lo suficientemente flexible para que se pueda adaptar sin perder mucho tiempo en este proceso. Por supuesto, esto no lo puedes decidir si ya estas en la fase de desarrollo, esto debes hablarlo con los involucrados en el proyecto una vez que este ha nacido, para que se reserve el suficiente tiempo para esta tarea.

Dentro de pocos meses se llevarán a cabo también las elecciones presidenciales. Crees que volveremos a ver otra failwhale en el sitio del TSE y similares?