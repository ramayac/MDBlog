---
title: Depurando CSS y Javascript con Internet Explorer
date: 2008-10-13
author: Robertux
tags: firefox, complementos, diseño, safari, internet, chrome, Opera, addon, navegador, desarrollo, microsoft, herramienta
post_id: blog-3515952828243908885.post-6191280296556684679
---

[![](http://1.bp.blogspot.com/_jH77WNrMVRA/SPL8VhxKCSI/AAAAAAAADcU/ZczDJlYXCIA/s400/postbanner.png)](http://1.bp.blogspot.com/_jH77WNrMVRA/SPL8VhxKCSI/AAAAAAAADcU/ZczDJlYXCIA/s1600-h/postbanner.png) "El lienzo de un diseñador web es la pantalla de su monitor"

Hay muchas y muy variadas herramientas e IDEs que te permiten diseñar/desarrollar tu sitio web de manera que luzca y funcione exactamente como deseas con un par de simples clics pero al final de cuentas, son los navegadores del web quienes tienen la ultima palabra a la hora de decidir como lucirá y se comportara finalmente tu sitio. Estos navegadores, por cierto, son muy quisquillosos respecto a las funcionalidades que puedes implementar ya que ninguno de ellos soporta los estándares del web al 100% sino que utilizan sus propios estándares por lo que posiblemente tu sitio no se vea ni funcione igual en todos los navegadores.

Los desarrolladores del web que utilizan [Mozilla Firefox](http://www.srbyte.com/2008/07/el-guiness-record-de-firefox-3.html) y sus derivados(Iceweasel, Floc, etc) han de estar familiarizados con el muy popular [add-on](http://www.srbyte.com/2008/07/qu-es-un-complementoadd-on-de-firefox.html) llamado Firebug y también con el Web Developer Toolbar, los cuales nos permiten editar el código de nuestro sitio directamente desde el navegador, por lo que sabremos en el momento los resultados de los cambios que realizamos en dicho codigo. De esta manera, podríamos re alinear DIVs que aparecen fuera de lugar, cambiar el color de fondo de un elemento, revisar en que linea de código se produjo un error de Javascript o que datos se han enviado por los HTTP Headers vía Ajax, solo por mencionar algunas de sus utilidades.

[![](http://1.bp.blogspot.com/_jH77WNrMVRA/SPLgWAdkArI/AAAAAAAADcE/vjkhwGdluPE/s400/firebug_srbyte.png)](http://1.bp.blogspot.com/_jH77WNrMVRA/SPLgWAdkArI/AAAAAAAADcE/vjkhwGdluPE/s1600-h/firebug_srbyte.png)"Editando el CSS del sitio del sr byte con Firebug"

Lo lamentable es que no todo mundo utiliza Mozilla Firefox sino que la gran mayoría de usuarios Windows suelen navegar desde Internet Explorer así como la gente de Mac utiliza Safari, otros tantos que se sienten cómodos con [Opera](http://www.blogger.com/www.srbyte.com/2007/09/opera-una-nueva-experiencia.html) y los que recientemente han empezado a probar [Google Chrome](http://www.srbyte.com/2008/09/navegador-web-de-google.html). Esto significa (y lo menciono porque a mi me ocurrió) que si alineamos nuestro sitio perfectamente en Firefox, posiblemente se vea horriblemente mal en los demás navegadores y no tenemos nada que nos permita hacer lo mismo que con Firefox en el resto de navegadores.

Debido a esta situación fue que descubrí que al parecer, Internet Explorer también cuenta con [un conjunto de add-ons](http://www.ieaddons.com/en/) para personalizarlo un poco, aunque no son tan conocidos ni usados como los de Firefox. Entre estos add-ons, casualmente existe uno Developer Tooblar, el cual es muy similar en funcionamiento al Firebug aunque no es tan completo ya que nada mas nos permite inspeccionar el codigo HTML de nuestra pagina y de los elementos de la misma, ademas de mostrarnos el CSS que ha sido aplicado a dichos elementos.

[![](http://2.bp.blogspot.com/_jH77WNrMVRA/SPLkVs1zqnI/AAAAAAAADcM/7o7309Dc8Bo/s400/ie_devtoolbar_srbyte.PNG)](http://2.bp.blogspot.com/_jH77WNrMVRA/SPLkVs1zqnI/AAAAAAAADcM/7o7309Dc8Bo/s1600-h/ie_devtoolbar_srbyte.PNG)"Observando (porque es lo unico que se puede hacer) el codigo HTML y el CSS del sitio del sr byte con IE Developer Toolbar"

Ya que el add-on antes mencionado no nos permite editar el Javascript del sitio, podemos hacer uso del Microsoft Script Debugger para tal efecto. Por cierto, este programa es lanzado cada vez que se encuentra un error en el Javascript de una pagina, una vez que hallas elegido la opcion de utilizarlo como debugger en el cuadro de dialogo que te aparece al momento de producirse el error.

[![](http://4.bp.blogspot.com/_jH77WNrMVRA/SPLajb3eVXI/AAAAAAAADbk/Pa2u2imRmFA/s400/ms_script_debugger.gif)](http://4.bp.blogspot.com/_jH77WNrMVRA/SPLajb3eVXI/AAAAAAAADbk/Pa2u2imRmFA/s1600-h/ms_script_debugger.gif)"Apariencia del editor MS Script Debugger"

Copyleft Rodrigo S. Amaya C. y el staff del Sr. Byte