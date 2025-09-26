---
title: Session Tracking
date: 2009-05-02
author: Rodrigo Amaya
tags: srbyte
post_id: blog-3515952828243908885.post-4912755190029257478
---

Supongamos que estas escribiendo haciendo un servlet que implementa sesiones HTTP, al cual se estan conectado tres clientes (navegadores web) diferentes. Para cada cliente, el servlet deberia de ser capaz de determinar a que cliente le pertenece la informacion que se tiene en sesion. La capacidad de un servlet para saber a que cliente le pertenece que información es comúnmente denominada "Session Tracking" (huella o pista de una sesion). Existen tres formas de implementar esta capacidad en un servlet:

- Cookies
- Re-Escritura de URL (URL Rewriting)
- Informacion SSL
Cookies Usar cookies para identificar o distinguir a un cliente, es la forma mas sencilla y comun para mantener la pista de una sesion HTTP, porque no requiere "tecnicas especiales" para funcionar. Cuando un servidor web, o contenedor de aplicaciones, realiza una peticion, el objeto HttpSession es creado y un identificador unico de sesion es generado para el cliente y enviado al navegador como una cookie. En las siguientes peticiones y respuestas que ocurran entre el cliente y el servidor, el navegador web (el cliente) enviara la informacion que contenga la cookie al servidor, y el "Session Manager" usara esta informacion para encontrar el objeto HttpSession asociado con ese cliente. Suena complicado, pero sinceramente no lo es. Mas informacion del cookie monster... de las cookies, aquí.Re-escritura de URL (URL rewriting) Aunque facilón, hay situaciones en que utilizar cookies no funciona. Siguiendo con el juego de suposiciones: te dicen que la aplicación web que estas haciendo, NO debe de usar cookies, que intentes otro método para mantenerle la pista a las sesiones. Eso puede pasar. Y gracias a Dios, existe la "Re-escritura de URL" (URL rewriting), para administrar las sesiones del usuario.

Con la re-escritura de URL, los vinculos que se retoran al navegador web (para las redirecciones) tienen el ID de la sesion adjuntos a el (a ese vinculo). Por ejemplo, para el siguiente vinculo en una pagina web:

```

href="tienda/catalogo"
```
se deberia de re-escribir asi:

```

href="tienda/catalogo;jsessionid=DA32242SSGE2"
```
Cuando el usuario hace clic en ese vinculo, el URL se envia al servidor como parte del ID de la petición del cliente. El servidor o mejor dicho, el contenedor de la aplicacion reconoce la cadena: ;jsessionid=DA32242SSGE2 como el ID de la sesion de ese cliente y la guarda para obtener el objeto HttpSession correcto para ese usuario/cliente.

Nota de buen programador: Por el amor a Dios, no cometan el terrible error de asumir que el ID de la sesion tiene una longitud o contenido exacto. Un ID de sesion será mucho más largo de lo que presento en este ejemplo.Mi problema, con la re-escritura de URL, es que precisa de ciertas técnicas de programación, o guias mejor dicho, que NO todo programador hará. Estos lineamientos son:

- Los servlets que llevan la sesion, deben utilizar "encode URL".
- El punto de entrada a la aplicacion Web deberá ser un JSP o un Servlet.
- Evite usar HTML estatico en la aplicacion.
Usar "Encode URL" en la aplicación:Si el servlet retorna URL's al navegador, o redirecciones, se deben utilizar los metodos: encodeURL() ó encodeRedirectURL() en el codigo del servlet.

```
out.println("<a href="unaurl"&mt;catalogo</a&mt;</a&mt;");
```
Imagina que le retornas a un cliente, entre otras cosas claro, un vinculo asi:

```

out.println("</a&mt;<a href="unaurl;"&mt;catalogo</a&mt;");
```
En este caso usariamos encodeURL(), para retornarle una URL adecuada al cliente:Usar "Encode Redirect URL" en la aplicación:Y si tienes la siguiente declaracion:
```

response.sendRedirect("https://unhost.com/tienda/catalogo");
```
Tendrias que cambiar el servlet para que llame al metodo encodeRedirectURL() antes de enviar la URL al "output stream":
```

response.sendRedirect(response.encodeRedirectURL("https://unhost.com/tienda/catalogo"));
```
Los metodos encodeURL() y encodeRedirectURL() son parte del objeto HttpServletResponse. Estas llamadas verifican que la re-escritura de URL esta configurada antes de enviar el URL codificado (encoded). Si no esta configurada (la re-escritura de URL) esta retorna el URL original.If both cookies and URL rewriting are enabled and response.encodeURL() or encodeRedirectURL() is called, the URL is encoded, even if the browser making the HTTP request processed the session cookie. You can also configure session support to enable protocol switch rewriting. When this option is enabled, the product encodes the URL with the session ID for switching between HTTP and HTTPS protocols. Supply a servlet or JSP file as an entry point The entry point to an application (such as the initial screen presented) may not require the use of sessions. However, if the application in general requires session support (meaning some part of it, such as a servlet, requires session support) then after a session is created, all URLs must be encoded in order to perpetuate the session ID for the servlet (or other application component) requiring the session support. The following example shows how Java code can be embedded within a JSP file:

```

```
Avoid using plain HTML files in the application Note: To use URL rewriting to maintain session state, do not link to parts of your applications from plain HTML files (files with .html or .htm extensions). The restriction is necessary because URL encoding cannot be used in plain HTML files. To maintain state using URL rewriting, every page that the user requests during the session must have code that can be understood by the Java interpreter. If you have plain HTML files in your application (or Web module) or in portions of the site that the user might access during the session, convert the files to JSP files. This impacts the application writer because maintaining sessions with URL rewriting requires that each servlet in the application must use URL encoding for every HREF attribute on tags, as described previously. Sessions are lost if one or more servlets in an application do not call the encodeURL(String url) or encodeRedirectURL(String url) methods. Session tracking with SSL informationNo special programming is required to track sessions with Secure Sockets Layer (SSL) information. To use SSL information, select Enable SSL ID tracking in the Session Management page of the administrative console. Because the SSL session ID is negotiated between the Web browser and HTTP server, this ID cannot survive an HTTP server failure. SSL tracking is supported by the IBM HTTP Server only. You can control the lifetime of an SSL session ID by configuring options in the Web server. For example, in the IBM HTTP Server, set the configuration variable SSLV3TIMEOUT to provide an adequate lifetime for the SSL session ID. An interval that is too short can cause a premature termination of a session. Also, some Web browsers might have their own timers that affect the lifetime of the SSL session ID. These Web browsers may not leave the SSL session ID active long enough to serve as a useful mechanism for session tracking. The internal HTTP Server of WebSphere Application Server - Express also supports SSL tracking.