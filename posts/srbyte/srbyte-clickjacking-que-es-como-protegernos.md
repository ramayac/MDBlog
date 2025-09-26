---
title: Clickjacking Que es? Como protegernos?
date: 2008-10-08
author: Robertux
tags: firefox, java, complementos, network, safari, internet, chrome, clickjacking, iframe, Opera, addon, navegador, computadoras, seguridad
post_id: blog-3515952828243908885.post-9135386661346630577
---

![image](https://3.bp.blogspot.com/_jH77WNrMVRA/SOwl2YzPRtI/AAAAAAAADa8/y-BTMA8eyLg/s400/spyware.jpg)    
"No es necesario ser
paranoicos para aceptar que nos pueden vigilar y hasta controlar desde las paginas web que visitamos"

Recientemente, dos investigadores de ha.ckers.org [Robert "RSnake" Hansen](https://ha.ckers.org/blog/about) y [Jeremiah Grossman](https://jeremiahgrossman.blogspot.com/2004/11/my-resume.html) [publicaron noticias](https://ha.ckers.org/blog/20080915/clickjacking/) acerca de un problema de vulnerabilidad el cual afecta a todos los navegadores web de la actualidad y que podría monitorear y controlar la actividad de un usuario en un sitio web especifico. Hasta el momento ningún navegador del web se encuentra exento, por lo menos no los que todo mundo utiliza ya que es un problema relativamente reciente, no esta directamente relacionado con javascript y no es tan fácil de solucionar con un simple parche o actualización de los mismos. Los únicos que no se ven afectados por este grave problema son [lynx](https://es.wikipedia.org/wiki/Lynx_%28navegador%29), [Links](https://links.sourceforge.net/), [w3m](https://w3m.sourceforge.net/) y similares.

Ahora dos preguntas vienen a nuestra mente:

En que consiste realmente este problema?

Clickjacking, que se podría traducir como secuestro del clic, se refiere a la capacidad del desarrollador o dueño de un sitio web de tomar el control de los vinculos que nuestro navegador visita, botón, o cualquier elemento de la página sin que ni siquiera uno vea que está sucediendo. Esto lo hacen utilizando un Frame o iFrame con opacidad nula por encima del contenido original de la pagina de manera que el usuario, al tratar de dar clic sobre los elementos de la pagina original, realmente interactúe con los elementos ocultos de este iFrame sin darse cuenta. Este iFrame, a su vez, puede en realidad ser parte de un dominio totalmente diferente al de la pagina original que cubre y tendría la capacidad de hacer requests via Ajax a repositorios de datos de dichos dominios.

Dicho de otra forma, imaginemos un cristal (opacidad 0) con algo debajo, intentamos tocar con el dedo lo que hay debajo, pero tocamos el cristal, esto es lo mismo que sucede cuando ponemos un iFrame con opacidad 0 (invisible) encima de ciertos elementos de nuestra web: el usuario irá a hacer clic en nuestros elementos, pero hará clic dentro del iFrame.

![image](https://3.bp.blogspot.com/_jH77WNrMVRA/SOwmR9EVLQI/AAAAAAAADbE/6rbwdlaYOlQ/s400/clickjacking.png)    

"Imagen que ejemplifica las dos capas invisibles que puede tener un sitio web: la
capa que vemos y la capa con la que realmente interactuamos"

Ahora imaginemos que cargamos en el iFrame una página de noticias, como meneame, fresqui, o digg, lo volvemos transparente, y debajo metemos un botón que diga: haz clic aquí, posicionándolo justo exactamente debajo de donde aparece el botón para votar a la noticia, dentro de la página de noticias. Cuando el usuario haga clic en nuestro botón, que está debajo del iFrame transparente, justo donde está el botón para votar de la web de noticias que hay cargada dentro del iFrame, hará clic en el botón votar de la pagina que hay dentro del iFrame.

Que puedo hacer para protegerme?

Como habíamos mencionado, no es tan fácil como agregar un parche o actualizar a nuestra ultima versión de nuestro navegador. Lo primero seria evitar las visitas a sitios de poca fiabilidad (ustedes saben a cuales me refiero) y si no estamos seguros de adonde nos llevara un link o un botón, dar clic derecho sobre este y revisar sus propiedades.

- Opera: Vayanse a la pagina opera:config y deshabiliten la opción iFrames.
![image](https://4.bp.blogspot.com/_jH77WNrMVRA/SOwiQqaf4eI/AAAAAAAADas/YxB-EONTMXc/s400/ClickJakingOperaConfig.png)    
"Captura de pantalla de
Opera, mostrando la pagina de configuracion donde se deshabilitan los iFrames"

- Firefox: Instalar el plugin ([complemento, add-on o como le quieran llamar](https://www.srbyte.com/2008/07/qu-es-un-complementoadd-on-de-firefox.html)) llamado [NoScript](https://noscript.net/getit) el cual, al igual que el plugin [FlashBlock](https://www.srbyte.com/2008/10/flashblock-complemento-de-firefox.html) que te bloquea los elementos Flash de una pagina, este te permite bloquear una gran cantidad de tipos de contenido de una pagina, entre estos, los scripts, animaciones flash, aplicaciones Java y por supuesto, iFrames.
![image](https://4.bp.blogspot.com/_jH77WNrMVRA/SOwkk3zufcI/AAAAAAAADa0/gt9GEyJiKqA/s400/ClickJackingFirefoxConfig.png)    
"Captura de pantalla de Firefox, mostrando en su
barra de estado la informacion de bloqueos de NoScript y en una ventana emergente, la configuracion del NoScript y donde se deshabilitan los iFrames"

- Internet Explorer Safari y Chrome: Resignense a estar vulnerables, esperen a que los fabricantes encuentren una solución o cambien de navegador.