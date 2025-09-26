---
title: 3 reglas al trabajar con Bases de Datos
date: 2008-05-14
author: Rodrigo Amaya
tags: programador, mysql, subversion, programacion, base de datos
post_id: blog-3515952828243908885.post-9182436290436764077
---

Aquí presento algunas reglas para trabajar con bases de datos:

Regla 1: Nunca uses un solo servidor de bases de datos para todo el trabajo de desarrollo (exacto, nada de bases de datos centralizadas).La conveniencia de trabajar con un servidor centralizado de bases de datos es tentadora. Todos los desarrolladores se conectan a una sola base de datos que pueden probar y cambiar. Este servidor funcionara como el Anillo Único de Sauron... y todos los cambios que se den, se reflejaran inmediatamente a todos los miembros del equipo de desarrollo. Es más, es tan convincente la idea de usar un servidor centralizado, que este la gente lo usa como repositorio de datos de prueba.

![image](https://bp1.blogger.com/_ayvorITawE4/SCsVN8-FvLI/AAAAAAAAAtM/4iEZNRs4-xY/s400/gollum.jpg)    

"Imagen: Gollum con el Anillo"

Pero, como muchas conveniencias en el desarrollo de software (y así de engañoso y malévolo como el Anillo Único), usar una sola base de datos para todo el equipo, funciona como un pozo de brea... si, como en el que murieron tantos dinosaurios :P.

![image](https://bp1.blogger.com/_ayvorITawE4/SCsVM8-FvKI/AAAAAAAAAtE/yvbJ1g22jYQ/s400/tar.jpg)    
"Imagen: Pozo de Brea"

Todos los desarrolladores están tentados a cambiar los tipos de datos en algún momento, y los mas probable es que así sea, pero el cambio que un desarrollador puede hacer en la base de datos... veamos que sucede, cuando todo sale mal... estos son los pasos al pozo de brea:

1. Un desarrollador modifica un tipo de dato en la base de datos (o peor, una tabla entera... o varias). 2. El cambio probablemente provocara un error en mi código (ya no puedo compilar mi parte). 3. Esto implica que tengo que hacer cambios en mi código para que todo funciona como antes. 4. O tengo que esperar a que el desarrollador que modifico la base de datos haga un "Commit" de su código para que altere el mio... y ver si funciona con lo que el hizo. 5. Pueda que su "Commit" altere o no mi código. 6. Si no lo altera, y ahora YO tendré que ver que $%&&/$ altero el tipo, para que mi código funcione (¿cuanto tiempo perdí ya?). 7. Es posible también que el tipo diga después: "bueno, me equivoque, solo estaba probando...". 8. Ahora hay que regresar a la versión anterior del software, perder horas de cambios y ajustes, y realizar el proceso, cada vez que alguien disponga hacer el paso 1. Ademas desarrollar con una base de datos remota es LENTOOOOOOO...

![image](https://lh3.ggpht.com/Ramayac/SChunc-FvII/AAAAAAAAAs0/ltbMIa8a3HY/nerd1.jpg?imgmax=400)    
"Imagen: Maldito
desarrollador, ¡cambiaste todas las tablas sin avisar!"

Regla 2: Siempre ten una sola y respetada fuente para el esquema de tu base de datos (alguien se encarga de mantener y liberar la base de datos para todos).

Idealmente, esta única fuente sera tu control de versiones ( ver la regla #3). Considere el siguiente entre dos desarrolladores, Pedro y Juan:

Pedro: Es tiempo para comenzar a probar la aplicación, y depurar los errores. ¿Copiamos la base de datos de la maquina de José, o usamos la base de datos de Luis?

Juan: Ummmmmmmm, no recuerdo cual es la base de datos mas reciente (u oficial, o actualizada)

Pedro: Shit... (se pone a llorar).

Todo mundo debería de saber donde esta el ultimo esquema de la base de datos, y también es el derecho de los desarrolladores tener una experiencia sin problema para obtenerla y usarla en su maquina de trabajo. Yo desarrollador quiero caminar a mi estación de trabajo, obtener la ultima versión de la base de datos por medio de cvs ó svn, y trabajar transparentemente con ella, sin preocuparme por nada más que hacer mi trabajo.

Regla 3: Siempre versiona tu base de datos.

Hay muchas maneras de versionar una base de datos, pero el objetivo principal es propagar los cambios, probar y liberar la base de datos en una manera controlada y consistente. Un segundo objetivo es poder regresar y recrear la base de datos de algún momento anterior. Este segundo objetivo es particularmente importante si se produce software para clientes. Si alguien tiene un bug (error) en la versión 1.00.3.4 de tu aplicación (y tu vas por la versión 2.00.5.0) tienes que ser capaz de recrear la aplicación y su base de datos para esa versión, para poder darle soporte técnico al cliente o a los clientes que tengan esa versión.

> Respeta
> estas reglas, y tu desarrollo de software sera muchísimo mas sencillo y menos
> traumatico.

¡Saludos!