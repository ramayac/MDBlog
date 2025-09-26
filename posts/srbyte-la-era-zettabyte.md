---
title: La era Zettabyte...
date: 2009-06-11
author: Rodrigo Amaya
tags: cisco, internet, zettabyte
post_id: blog-3515952828243908885.post-7706356879334824840
---

[![](http://3.bp.blogspot.com/_ayvorITawE4/SjGA9WGaTQI/AAAAAAAACBo/oX17atzc0uQ/s320/Cisco_new_logo-thumb.gif)](http://3.bp.blogspot.com/_ayvorITawE4/SjGA9WGaTQI/AAAAAAAACBo/oX17atzc0uQ/s1600-h/Cisco_new_logo-thumb.gif)Hace poco, [CISCO](http://www33.wolframalpha.com/input/?i=CSCO) publico un [White Paper](http://www.cisco.com/en/US/solutions/collateral/ns341/ns525/ns537/ns705/ns827/white_paper_c11-481360_ns827_Networking_Solutions_White_Paper.html)(documento de reporte o guía) sobre su visión de como se comportara el trafico de Internet en el año 2013. Si lo leen con detenimiento, notaran predicciones de crecimiento exponencial (nada sorprendente), y el uso de "números realmente grandes"(o mejor dicho, unidades de medida) para cuantificar el crecimiento del trafico.

La amplitud de los medios estudiados, va desde la televisión digital hasta los teléfonos celulares. Se espera que el trafico en el 2013 llegue hasta los 56 [exabytes](http://es.wikipedia.org/wiki/Exabyte)al mes (eran 8 exabytes en el 2008). 56 exabytes en simple español equivale al "tamaño" de 10,000 millones de DVD's viajando por Internet cada mes... en pocas palabras una cantidad INCREÍBLE de datos.

Más del 90% de este trafico sera de televisión digital, video bajo demanda y P2P ([BitTorrent](http://www.srbyte.com/2007/03/bittorrent-todo-mundo-ama-bittorrent.html), Emule, Gnutella, Frostwire, etc...). Cisco también predice que el video chat incrementara diez veces más entre 2008 y 2013. Otro incremento masivo, sera en los dispositivos moviles, y con móviles como el [Nokia 5800 XpressMusic](http://www.srbyte.com/2009/06/review-nokia-5800-xpressmusic-13.html), iPhone 3g S y el Palm Pre, tampoco es de extrañar que suceda.

[![](http://2.bp.blogspot.com/_ayvorITawE4/SjF7bW0NCfI/AAAAAAAACBQ/_noLVpXtA_g/s320/nokia-5800-xpressmusic1.jpg)](http://2.bp.blogspot.com/_ayvorITawE4/SjF7bW0NCfI/AAAAAAAACBQ/_noLVpXtA_g/s1600-h/nokia-5800-xpressmusic1.jpg)"Nokia 5800 XpressMusic"

[![](http://1.bp.blogspot.com/_ayvorITawE4/SjF7bXoMYxI/AAAAAAAACBY/_Ia5PLQGdWM/s320/palm-pre-perfecta.jpg)](http://1.bp.blogspot.com/_ayvorITawE4/SjF7bXoMYxI/AAAAAAAACBY/_Ia5PLQGdWM/s1600-h/palm-pre-perfecta.jpg)"Palm Pre"

[![](http://3.bp.blogspot.com/_ayvorITawE4/SjF7bLKjmhI/AAAAAAAACBI/Kesb_6bB_WM/s320/iphone3gs_intro_slide.jpg)](http://3.bp.blogspot.com/_ayvorITawE4/SjF7bLKjmhI/AAAAAAAACBI/Kesb_6bB_WM/s1600-h/iphone3gs_intro_slide.jpg)"iPhone 3G S"

Con teléfonos como los mostrados, no es de extrañarnos que cada año se este duplicando la cantidad de trafico que fluye entre los teléfonos móviles, y tampoco es de extrañarnos que para el 2013, según Cisco sera: 66 veces superior al actual. Con números así, los desarrolladores de software, ya tienen que estar pensando en comenzar a usar el SDK: [Java ME](http://java.sun.com/javame/index.jsp), el [iPhone SDK](http://developer.apple.com/iphone/) y [Symbian SDK](http://developer.symbian.org/); después no digan que no les avise ;)

Otros puntos importantes:

- El trafico IP ANUAL llegara en 2013, a 667 exabytes, 2/3 de un zettabyte.
- Trafico de móviles crecerá 131% por año. Llegando a 2 exabytes al mes.
- Para el 2013, el video comprometera el 64% del trafico en móviles.
- El trafico P2P continuara creciendo, pero sera un "componente menor" en el trafico de Internet en términos porcentuales.
- En 2013, el volumen de videos que se transmitan, sera 700 veces mayor a la capacidad de transmisión (Internet Backbone) que tenia Estados Unidos en 2000.
Cisco ve un (nada extraño) crecimiento en el trafico, que probara verdaderamente la capacidad de la infraestructura existente... cabe mencionar que el estudio esta "ligeramente" enfocado a la nueva serie de productos que Cisco planea lanzar al mercado, bonita estrategia la que usan, pero aparte de ese detalle, el estudio si esta interesante. Se me viene a la mente, que si para el 22 de Diciembre de 2012, no somos destruidos por:

1. La realización de alguna profecía Maya 2. Un planeta con órbita irregular 3. [Skynet](http://en.wikipedia.org/wiki/Skynet_%28Terminator%29#Terminator_Salvation) 4. El [LHC](http://en.wikipedia.org/wiki/LHC) 5. Alguna otra tontería Probablemente si nos matara del aburrimiento, navegar en Internet, eso es, sí la infraestructura que se encarga de brindarnos conectividad, no mejora...

[![](http://1.bp.blogspot.com/_ayvorITawE4/SjGA9Hf18rI/AAAAAAAACBg/Qs821LwkXQU/s320/cisco.jpg)](http://1.bp.blogspot.com/_ayvorITawE4/SjGA9Hf18rI/AAAAAAAACBg/Qs821LwkXQU/s1600-h/cisco.jpg) ¿Que tiene que cambiar, para que en el 2013 no nos destruya la falta de comunicación y las descargas lentas?

1. Implementar IPv6, para soportar más clientes, que los que soporta IPv4. 2. Mejorar la capacidad de los Switchs y Routers. No solo en cantidad de paquetes ruteados/procesados, sino en eficiencia de consumo de energia. 3. Aumentar la capacidad de almacenamiento de los servidores, usando sistemas de archivos de alto rendimiento como: [ZFS](http://es.wikipedia.org/wiki/ZFS_%28sistema_de_archivos%29) ó [XFS](http://es.wikipedia.org/wiki/XFS) (ambos con soporte para volúmenes de hasta 16 exabytes). 4. Aumento de infraestructura que soporte mayor transporte/consumo de ancho de banda en general. A paso firme y continuo se esta llegando al final de la era de Gigabytes y Terabytes, y poco a poco nos movemos a los Exabytes y finalmente al gigantesco [Zettabyte](http://es.wikipedia.org/wiki/Zettabyte). Pero seguramente todos nos estamos preguntando: ¿Que tipo de información se estará produciendo y almacenando?, ¿sera significativa?, ¿sera solo ruido y nada de señal?

¡Danos tus opiniones!

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte