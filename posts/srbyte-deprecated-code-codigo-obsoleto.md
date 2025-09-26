---
title: Deprecated Code (Codigo Obsoleto)
date: 2009-04-18
author: Rodrigo Amaya
tags: documentacion, practicas, codigo
post_id: blog-3515952828243908885.post-1093618359359466506
---

Todo programador, se habrá (debería?) de encontrar mas de alguna vez con este termino: Deprecated (Obsoleto).

Y este termino es especialmente frecuente al usar frameworks, que de tan "robustos y seguros", ya se estarán volviendo obsoletos en si mismos. Un excelente lugar donde encontrar código marcado como obsoleto, es utilizando metodos viejoss en un framework "nuevo", por ejemplo, llamar metodos del JDK 1.4 en el JDK 1.6.

Este "atributo" (Obsoleto) que se le da al código fuente, se especifica cuando una clase, función o API ha perdido su importancia, y tiene tan poca importancia que no debería de ser usada en lo absoluto, y probablemente deje de existir en un futuro cercano.

[![](http://2.bp.blogspot.com/_ayvorITawE4/SeeJgLxA_aI/AAAAAAAAB8w/5af8X0SFFMg/s320/deprecated.png)](http://2.bp.blogspot.com/_ayvorITawE4/SeeJgLxA_aI/AAAAAAAAB8w/5af8X0SFFMg/s1600-h/deprecated.png)"Java Doc que muestra una función marcada como 'Deprecated' (Obsoleta)"

La necesidad de marcar una función, clase, API... código en general, como obsoleto, surge del desarrollo natural del código, para adaptarse constantemente a las nuevas tecnologías, mejoras de lógica, etc...

Supongamos que tienes una librería muy popular, y cada 3 meses liberas una versión nueva. Y en cada versión nueva, añades funcionalidad adicional y modificas metodos, y reescribes otros. Como tu librería es muy popular, tienes que mantener cierta compatibilidad con las versiones anteriores, y por esa razón tienes que mantener código viejo... porque tienes que darle tiempo a los developers que migren su código y comiencen a usar las novedades que has añadido a tu código. Y claro, tu no quieres que los developers sigan usando los metodos antiguas, que solo mantienes por motivos de compatibilidad, entonces la solución, es marcar el metodo, clases, o código en general, como obsoleto. ¿Útil no?

[![](http://1.bp.blogspot.com/_ayvorITawE4/SeeJf2h4X8I/AAAAAAAAB8o/xNILZn9x-U4/s320/deprecated-javascript.png)](http://1.bp.blogspot.com/_ayvorITawE4/SeeJf2h4X8I/AAAAAAAAB8o/xNILZn9x-U4/s1600-h/deprecated-javascript.png)"¿Cuantos developers tendrán el 'honor' de marcar como obsoleto el codigo de alguien más?"

Como ven, la actividad de marcar clases, métodos, etc... como obsoleto, resuelve muchos problemas. Las clases existentes que usan el API antiguo continúan trabajando bien, pero el compilador reconoce y avisa sobre los elementos marcados como obsoletos. Además, marcar el código fuente como obsoleto, es una responsabilidad que todo BUEN programador debe asumir, especialmente: cuando se trabaja en equipo, y cuando muchos developers usan el código que tu haces...

[![](http://4.bp.blogspot.com/_ayvorITawE4/SepztXhuXbI/AAAAAAAAB84/BBaaqbTf0iY/s320/take_responsability_card-p137699379057593903q6k5_400.jpg)](http://4.bp.blogspot.com/_ayvorITawE4/SepztXhuXbI/AAAAAAAAB84/BBaaqbTf0iY/s1600-h/take_responsability_card-p137699379057593903q6k5_400.jpg)"Un buen programador, siempre se hace responsable del código que escribe."

En Java la etiqueta @deprecated, marca como obsoleto el código que le sigue, vean el [Java API Documentator Generator (JavaDoc)](http://java.sun.com/j2se/1.5.0/docs/tooldocs/windows/javadoc.html) para más información sobre como usar esta etiqueta en su código. Y claro no solo esta la etiqueta @deprecated, también esta el tipo de "annotation" [Deprecated](http://java.sun.com/j2se/1.5.0/docs/api/java/lang/Deprecated.html). Y finalmente les comparto esta útil página, que tiene información de "[Cuando y Como" marcar como obsoleto el codigo fuente](http://java.sun.com/j2se/1.5.0/docs/guide/javadoc/deprecation/deprecation.html).

Y tu, ¿alguna vez has marcado código fuente como obsoleto?

UPDATE: Por una importante sugerencia de Xtecuan, se renombraron todas las incidencias de la palabra "funcion" por "metodos".

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte