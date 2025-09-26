---
title: Un vistazo al código de Google Chrome (Opinión)
date: 2008-09-14
author: Rodrigo Amaya
tags: open source, opinion, libre, programador, software, codigo, google
post_id: blog-3515952828243908885.post-8311296546419745089
---

![image](https://2.bp.blogspot.com/_ayvorITawE4/SM2-CTy9_UI/AAAAAAAABP4/hHrgVFM8ZcY/s320/Chrome_nav_150x75.jpg)    [El nuevo navegador de Google](https://www.srbyte.com/2008/09/navegador-web-de-google.html) ha logrado llenar la web de opiniones y un sin fin de expectativas:
¿Destronara a IE?, ¿Erradicara a Firefox? y podemos añadir un sin fin de comentarios y "twits" de anuncian el fin, o el principio del fin, para los principales navegadores. De toda eso ya se pueden encontrar mucho en Internet, de lo que nadie habla, es de Google Chrome desde el punto de vista de un Desarrollador de Software. Pues bien, para eso estamos ;)

¿Si Google Chrome es Open Source, donde esta el código? Seamos exactos: NO hay código de Google Chrome. Lo que si esta disponible es el codigo de Google Chromium. Podemos decir con toda propiedad que:

- Google Chrome NO es Google Chromium
- Google Chrome NO es Open Source
- Google Chromium ES Open Source
La primera parada para los incrédulos es revisar la guía de documentación de Chromium en [dev.chromium.org](https://dev.chromium.org/)....

![image](https://1.bp.blogspot.com/_ayvorITawE4/SM2_wG05awI/AAAAAAAABQQ/sr072SSnQ-Q/s320/chormium.png)    
"El logo de Chormium es en
escala de azules."

Así es amigos y amigas, "Chormium" es el nombre de la versión open source de Chrome, y Chrome a su vez, es la versión oficial y estable de Google.

> En la
> vida real Chrome es un producto terminado, afinado y hermoso, mientras que Chromium es solo
> una plasta de código cruda... italic;">¡ouch¡

¿Si es Open Source, es multiplataforma? Eventualmente. El sitio de Chromium explica como descargar el código para Linux, Mac OS X y Windows. Ahora comienzo con las malas noticias: desafortunadamente, no hay una versión de Mac OS X funcional, [según los mismos desarrolladores](https://dev.chromium.org/developers/how-tos/build-instructions-os-x):
> "Esta mas
> cerca del comienzo que del fin".
Es mas, según los [reportes](https://dev.chromium.org/developers/mac-os-x-detailed-status), Chormium ni siquera corre en Mac OS X.

La versión para Linux esta mal, pero no tan acabada como la de Mac OS X. Muchos de ustedes (y yo también) pensaron que un Navegador Web de Google que fuera Open Source seria ganancia automática para Linux.... lo siento, pero no es cierto.

Lo que si es cierto, es que la versión de Chrome para Windows esta diseñada para captar la atención del mercado. Y sera algo así como: "el molde maestro" para las versiones de los demás sistemas. En pocas palabras, no tendremos una versión para Linux hasta que los desarrolladores de Chorme vuelvan el código de [Win32 compatible con Linux](https://dev.chromium.org/developers/how-tos/linux-development).

Puntualizando: no esperen una versión para Mac o Linux en algún momento cercano, simplemente, falta mucho para que eso suceda.

Quiero descargar el código y compilarlo: ¿vale la pena?
![image](https://3.bp.blogspot.com/_ayvorITawE4/SM3BVvq9bVI/AAAAAAAABQY/2PugazfGiTU/s200/subversion_logo-384x332.png)    No, obtener el
código es mas difícil de solo de dar un click y descargar. Por que además de eso, hay que instalar las librerias que "la solucion" demanda, y tambien hay que instalar una serie de scripts que permiten el [acceso al repositorio SVN](https://www.srbyte.com/2008/03/programemos-mejor-subversion.html) de Chromium. Si bien son scripts para la consola, se sienten lo suficientemente familiares para cualquier desarrollador con un poco de experiencia. Lo que implica trabajar con un repositorio SVN, es que se obtiene la ultima versión del código, pero obtenerlo en estos momentos, me parece un poco ridículo, porque Chomium es un proyecto en constante movimiento. Sin mencionar que el código es terriblemente pesado y puede alcanzar cerca de 2.4 GB de código, incluyendo pruebas para depurar WebKit (el renderizador de las paginas web). Solo el código de Chromium mide cerca de 500 MB comprimido. Repito, no, no vale la pena descargar el código e intentar compilarlo. Si desean aprender a compilar proyectos ajenos, comiencen compilando archivos ".tar.gz" pequeños, luego el Kernel de Linux, Mozilla Firefox, y finalmente KDE o GNOME.

Pero si aun así, siguen de necios, [aquí están las instrucciones](https://dev.chromium.org/developers/how-tos/build-instructions-windows).

El código fuente: El código, hasta donde lo vi (si, lo descargue, pero NO lo compile porque al construirlo Chromium en Linux solo muestra la ejecución de pruebas de WebKit), se ve limpio, bien escrito y accesible para el programador casual. Esta bien documentado, y es abundante en comentarios graciosos, que sobran en los proyectos de software libre. ¡Esto es muy bueno! Código fuente bien organizado significa que el proyecto sera atractivo para desarrolladores cansados del código críptico y antiguo de Mozilla, heredado de Netscape Navigator. Es más facil compilar Chromium en Windows, parece que es una simple solucion de Visual Studio 2005... pero puede tardar más de 15 minutos en compilarse.

Notas Finales: Por ser un Beta, no esperen grandes cosas aún de Google Chrome/Chromium. Personalmente no me agrado darme cuenta de la "movida" de Google al dividir su proyecto en una versión "pulida y comercial" y en otra que no lo es tanto. Desde un punto de vista comercial, es sabio y conveniente hacerlo así. Pero por un breve instante, que confieso: fue muuuuy ingenuo de mi parte, quería que Google Chrome fuera en su totalidad Open Source. Si bien Google Chrome es una interesante propuesta es solo
> ...otro navegador web más.
![image](https://2.bp.blogspot.com/_ayvorITawE4/SM2-vr0AjBI/AAAAAAAABQI/VkDhdlHtAJ0/s320/googlechrome_450x257.jpg)    
"Así es, Chrome es solamente
otro navegador web mas."

Ni siquiera Microsoft se ha podido deshacer aún de IE 6, mucho menos Chrome podrá erradicar de la faz de la tierra a Opera, Safari, IE y a Firefox. ¿Pronosticos?, reservado... ¿por qué? Simple porque Chrome es beta, y nadie en su sano juicio demanda, exige o emite pronósticos de software de un beta "0.2"; como se haría con un producto completo. Ahora, como Desarrollador Web bueno, esa es otra historia que se comentará luego ...

¿Y tu, ya descargaste y usaste Google Chrome? ¿Qué te pareció?