---
title: Cookies, cookies, cookies... (¿Qué es una Cookie?)
date: 2009-05-23
author: Rodrigo Amaya
tags: cookies, practicas, programacion
post_id: blog-3515952828243908885.post-7279915234354871735
---

![image](https://2.bp.blogspot.com/_ayvorITawE4/ShgNiqf9JII/AAAAAAAAB-o/4bapq8QIRoY/s200/cookie-bite-web.jpg)    Las cookies son
pequeños bits de informacion textual, que un servidor web (o un contenedor de aplicaciones) envía a un navegador cliente para identificarlo; el navegador luego retorna esa informacion (esos bits de informacion textual) cuando se visita nuevamente ese sitio web o dominio. Haciendo que el servidor lea informacion que se le había enviado previamente a un cliente, la aplicación web puede proveer a sus visitantes, con unas cuantas conveniencias que se describen a continuación:

Identificar a un usuario durante una sesión de comercio electrónico: Si alguna vez te pusiste a curiosear Amazon.com, o alguna otra tienda en linea, ya sabrás sobre la metáfora del carrito de compras (Shopping Cart) que se puso tan de moda con el "[e-Commerce](https://en.wikipedia.org/wiki/Electronic_commerce)
", en el que el usuario selecciona un ítem, lo añade a su carrito, y sigue comprando, justo como cuando se visita el super mercado. Como las conexiones [HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) se cierran luego de que cada pagina se envía ([HTTP es Stateless](https://www.yafla.com/dennisforbes/-Web-Apps-Suck-Because-HTTP-is-Stateless-/-Web-Apps-Suck-Because-HTTP-is-Stateless-.html)), cuando el usuario selecciona un nuevo ítem para su carrito, ¿como sabe la tienda que el es ese mismo usuario que puso el ítem anterior en su carrito?... simple, con las cookies. Es mas, las cookies son tan útiles que los Java Servlets tienen API especifico para manejarlas.

Evitar autenticarte constantemente: Muchos sitios web grandes, requieren que te registres para utilizar sus servicios (Twitter, SlideShare, Facebook, Yahoo! Mail, etc), seria muuuuuy inconveniente, cada vez que el usuario realiza una acción (cambiar de pagina por ejemplo) preguntar constantemente el usuario y la contraseña a un usuario si ya se autentico al ingresar por primera vez. Las cookies se utilizan como parte de una solución (de baja seguridad) en la cual, se le puede dar informacion de autenticación al usuario (ID, llave de identificación, tiempo que vive la sesión, etc, nunca las contraseñas o números de tarjetas de crédito) luego de que este ingreso correctamente al sistema, luego cuando el usuario realiza una acción en el sitio, esta informacion se envia al servidor y este busca esa identificación para determinar si pertenece a ese usuario y si esta autenticado correctamente.

![image](https://1.bp.blogspot.com/_ayvorITawE4/ShgMqYG1FFI/AAAAAAAAB-Y/J31ve8aaKQg/s320/cookie-monster-20080603-133713.jpg)    
"De vez en cuando es bueno
borrar las cookies."

Personalizar un sitio: iGoogle es un ejemplo perfecto de la personalización de un sitio. iGoogle, entre otras cosas, usa una cookie para "recordar" que widgets le agregaste, cuantas pestañas tienes, etc...

Publicidad personalizada: Un motor de búsqueda, como Google, puede mantener una pista de las preferencias de un usuario a lo largo del tiempo, de esta forma, la publicidad que se le muestra a ese usuario, estará enfocada a sus preferencias de busqueda. Esa es la magia de AdSense, AdWords, y de la [infame cookie de Google](https://www.google-watch.org/cgi-bin/cookie.htm).

![image](https://1.bp.blogspot.com/_ayvorITawE4/ShgMql54LiI/AAAAAAAAB-g/aF4AS4sxhUI/s320/google_cookie.jpg)    
"Un mundo
Google nos vigila."

Proveer características convenientes a los usuarios y valor añadido al dueño del sitio es el propósito de las cookies. Y a pesar de algunos individuos paranoicos, las cookies no son una amenaza seria (por si mismas) a la seguridad de los usuarios finales.

Consideraciones sobre la seguridad de las cookies: Las cookies no se "interpretan" o ejecutan en un sistema, por lo que no se puede insertar virus en ellas para comprometer la integridad de un sistema. Los navegadores generalmente solo aceptan 20 cookies por sitio, además un navegador nunca poseerá más de 300 cookies en total, y cada cookie esta limitada a un tamaño de 4 KB, así que no se pueden usar para llenar el disco duro de un cliente (especialmente con el tamaño de los discos duros modernos).

![image](https://4.bp.blogspot.com/_ayvorITawE4/ShgMp6n6N1I/AAAAAAAAB-I/g8QmyRDGb7g/s320/cookie-monster.jpg)    
"300 cookies... 1
usuario."

Sin embargo, aunque las cookies no presentan una amenaza a la seguridad, si pueden presentar una seria amenaza a la privacidad. Pero más que culpa de la cookie en si, los problemas de privacidad radican en programadores con practicas poco seguras (como meter informacion de tarjetas de crédito en una cookie, ese SI que es un problema).

Dos consejos finales. Por los problemas reales que comprometen la privacidad de los usuarios, estos algunas veces deshabilitan las cookies. Así que si una cookie te da valor agregado a tu sitio, este no debe de depender totalmente de las cookies para funcionar. Y el segundo, es que los programadores de servlets que usan cookies, deben de ser cuidadosos para no almacenar informacion extremadamente importante (como: NUMERO DE TARJETAS DE CRÉDITO!!!) en las cookies.

Para más informacion sobre las cookies, puede visitar la [Wikipedia](https://en.wikipedia.org/wiki/HTTP_cookie) , tambien Java Tips ([usar una cookie de una JSP](https://www.java-tips.org/java-ee-tips/javaserver-pages/using-cookies-from-jsp.html)) y [Cookie Central](https://www.cookiecentral.com/faq/).

¡Saludos!