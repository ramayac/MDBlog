---
title: Singleton
date: 2009-04-26
author: Rodrigo Amaya
tags: srbyte
post_id: blog-3515952828243908885.post-3548982160569510892
---

Hace más de una semana, se encontraron en la empresa en la que estoy trabajando, con un problema de seguridad bastante particular. La solución (una de cuatro propuestas) que actualmente estoy implementando incluye: clase, con un comportamiento muy especial. Esta clase implementa un patron de diseño de software (de los que muy gustosamente escribiré mañana) llamado: Singleton.

Un Singleton, en poquísimas palabras, es:
"Una clase diseñada de tal forma, que solo exista una instancia de
esta en memoria."

Eso es "todo". Un Singleton, una vez en memoria, ya no podrá ser instanciado nuevamente hasta que: explícitamente se destruya (ya sea declarando su destrucción, o cerrando la aplicación que lo creo... y/o reiniciando la computadora/servidor, obvio, no es un poltergeist.

Para los que leen por primera vez sobre el Singleton, debo de aclarar algo antes de que ocurra alguna catástrofe de conceptos:

Un Singleton NO es una "variable objeto de acceso global glorificado".

En lo absoluto. Es más bien, una solución tan particular y especifica, que su simplicidad y conceptualizacion misma, puede resultar en un engaño que se pagara con terribles problemas a la hora de codificar.

REM this: Using this technique solves the issue of global state because "there is no global state at all. Every object only has references to what it needs directly! No passing around of objects which are not directly needed by the code. Dependencies are obvious since each object only asks for what it needs."

Using this technique solves the issue of memory management. All classes will be created and used when needed. No more Singletons sitting around and taking up memory while doing nothing at all.

Using this technique you can extend any of the classes you desire. Static methods are nowhere. No-sir.

https://srbyte.pastebin.com/f7b181adb