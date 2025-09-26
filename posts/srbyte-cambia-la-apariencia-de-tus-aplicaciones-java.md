---
title: Cambia la apariencia de tus aplicaciones Java
date: 2008-10-02
author: Robertux
tags: java, swing, lookandfeel, nimbus, gui, programacion, mac, substance, nimrod, apple
post_id: blog-3515952828243908885.post-7954172961803422846
---

[![](http://1.bp.blogspot.com/_jH77WNrMVRA/SORVn4D7NqI/AAAAAAAADZk/mSklzdp0Jek/s400/java_beautiful_logo.jpg)](http://1.bp.blogspot.com/_jH77WNrMVRA/SORVn4D7NqI/AAAAAAAADZk/mSklzdp0Jek/s1600-h/java_beautiful_logo.jpg)"Logotipo de Java, incrustado en una taza de cafe con coffee art"

Las aplicaciones de Java son muy características por su capacidad de ser ejecutadas en diversos sistemas operativos sin la necesidad de modificar el código o de incluir algún otro tipo de librerías externas, ya que lo único que se necesita es tener instalada la apropiada maquina virtual de Java en dicho sistema donde desea ejecutarse.

Estas aplicaciones hacen uso de la librería llamada Swing para poder mostrar su entorno gráfico al usuario. El problema es que Swing incluye de forma predeterminada un tema visual conocido como metal, el cual, en mi opinión personal, no es muy agradable para el usuario.

Porque? pues porque al igual que lo metálico, luce demasiado frío y liso, un tanto cuadrado y quizá des actualizado. Me recuerda un poco a los temas visuales que venían incluidos en las primeras versiones de Linux.

[![](http://3.bp.blogspot.com/_jH77WNrMVRA/SORSfGtE6lI/AAAAAAAADZc/gpkMWMS9qrA/s400/java_metal_combined.jpg)](http://3.bp.blogspot.com/_jH77WNrMVRA/SORSfGtE6lI/AAAAAAAADZc/gpkMWMS9qrA/s1600-h/java_metal_combined.jpg)"Netbeans 6 y aplicaciones Java con apariencia Metal, corriendo en tres diferentes sistemas operativos: Debian GNU Linux con Compiz/Emerald, Windows Vista y Mac OS. (Clic para agrandar)"

La ventaja de este tema es que luce igual en todos los sistemas operativos donde se ejecute y, según he leído, lo diseñaron de esta forma para no ocupar tantos recursos de la PC en la interfaz, la desventaja es que no se integra con la apariencia que posea el sistema operativo. Como se puede comprobar en la imagen anterior, esta interfaz no combina con ninguna de las apariencias de los tres principales sistemas operativos en los que se probó, GNU/Linux, Windows Vista y Mac OS.

La ventaja de Swing es que nos permite cambiar esta interfaz por defecto y aplicarles a nuestras aplicaciones, primeramente, el tema especifico de un sistema operativo especifico, de manera que se integre a el y ademas, poder aplicarle librerías definidas por terceras personas, y ademas, su forma de implementación nada mas consiste en agregar una librería y un par de lineas a nuestro código.

Entre las librerías que cambian la apariencia (también llamado Look and Feel) de nuestras aplicaciones Java, estas son mis favoritas:

- Substance:
URL: [https://substance.dev.java.net/](https://substance.dev.java.net/) Nombre: org.jvnet.substance.skin.SubstanceOfficeSilver2007LookAndFeel Descripción: Posee varios temas con diferentes formas y combinaciones de colores. Permite cambiar también el borde de la ventana.

[![](http://1.bp.blogspot.com/_jH77WNrMVRA/SORkAeulVxI/AAAAAAAADZ0/2n-k-vmH6iU/s400/java_substance.jpg)](http://1.bp.blogspot.com/_jH77WNrMVRA/SORkAeulVxI/AAAAAAAADZ0/2n-k-vmH6iU/s1600-h/java_substance.jpg)"Apariencia de una Aplicación Java con Look & Feel Substance"

- Nimbus:
URL: [https://nimbus.dev.java.net/](https://nimbus.dev.java.net/) Nombre: org.jdesktop.swingx.plaf.nimbus.NimbusLookAndFeel Descripción: Apariencia con esquinas redondeadas y relieves. No cambia el borde de la ventana. Colores suaves. No posee temas adicionales.

[![](http://1.bp.blogspot.com/_jH77WNrMVRA/SORkVURFR3I/AAAAAAAADZ8/Z-5aWoPstvM/s400/java_nimbus.jpg)](http://1.bp.blogspot.com/_jH77WNrMVRA/SORkVURFR3I/AAAAAAAADZ8/Z-5aWoPstvM/s1600-h/java_nimbus.jpg)"Apariencia de una Aplicación Java con Look & Feel Nimbus"

- Quaqua
URL: [https://quaqua.dev.java.net/](https://quaqua.dev.java.net/) Nombre: ch.randelshofer.quaqua.QuaquaLookAndFeel Descripción: Apariencia al estilo Mac OS. Único tema incluido y si adapta el borde de la ventana al tema aunque no aplica las esquinas redondeadas en todos los sistemas operativos donde es ejecutado.

[![](http://2.bp.blogspot.com/_jH77WNrMVRA/SORkuqNhF5I/AAAAAAAADaE/qkyPlTAap-M/s400/java_quaqua.jpg)](http://2.bp.blogspot.com/_jH77WNrMVRA/SORkuqNhF5I/AAAAAAAADaE/qkyPlTAap-M/s1600-h/java_quaqua.jpg) "Apariencia de una Aplicación Java con Look & Feel Quaqua"

- Nimrod
URL: [http://personales.ya.com/nimrod/index-en.html](http://personales.ya.com/nimrod/index-en.html) Nombre: com.nilo.plaf.nimrod.NimRODLookAndFeel Descripción: Unico tema pero las combinaciones de colores y transparencias son configurables. No modifica el borde de la ventana. [![](http://3.bp.blogspot.com/_jH77WNrMVRA/SORsuRgTA4I/AAAAAAAADaM/nV0J-WVkmd4/s400/java_nimrod.jpg)](http://3.bp.blogspot.com/_jH77WNrMVRA/SORsuRgTA4I/AAAAAAAADaM/nV0J-WVkmd4/s1600-h/java_nimrod.jpg)"Apariencia de una Aplicacion Java con Look & Feel Nimrod"

Para aplicar cualquiera de estos temas a nuestras aplicaciones, simplemente hay que descargar el .jar del Look & Feel que deseamos y ubicarlo en la carpeta de nuestro proyecto para después agregarlo a la librería del mismo. Luego, en el constructor de cada una de las clases que componen la vista de cada FrameView o JFrame agregar las siguiente lineas:

> JFrame.setDefaultLookAndFeelDecorated(true);
> UIManager.setLookAndFeel("NombreDelLook&Feel");
Por supuesto, estos Look & Feel tienen su coste en cuanto a rendimiento de la aplicación pero considero que valen la pena ya que la apariencia de nuestra aplicacion define en parte su usabilidad.