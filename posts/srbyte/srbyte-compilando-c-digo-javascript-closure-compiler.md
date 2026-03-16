---
title: Compilando código JavaScript (Closure Compiler)
date: 2009-11-14
author: Rodrigo Amaya
tags: javascript, compiler, google
post_id: blog-3515952828243908885.post-5276971736983403518
---

Siempre pensando en la oportunidad de mejorar la forma en la que pruebo código JavaScript, buscando información en la re, [encontré esta brillante idea](https://www.west-wind.com/WebLog/posts/10688.aspx). Cuando lo leí vi las puertas del cielo abrirse y luego pensé en usar el compilador de .Net, para realizar [pruebas unitarias](https://www.srbyte.com/2008/12/herramientas-del-programador-unit.html) sobre archivos simples de JavaScript (solo chequear que están bien era suficiente).

El problema es, que aparte de depender de .Net, el ejemplo de compilación de código JavaScript no me funciono en [SharpDeveloper](https://www.icsharpcode.net/OpenSource/SD/) 3.0 ... opte por utilizar Mono, pero lamentablemente el compilador de JavaScript que este posee esta actualmente sin mantenimiento.

De clic en clic fui encontrando a "viejos conocidos" como el proyecto [Rhino](https://www.mozilla.org/rhino/) de Mozilla, sustituto del motor [SpiderMonkey](https://www.mozilla.org/js/spidermonkey/), ambos motores para interpretar, "compilar" y ejecutar JavaScript. Aun asi, estas soluciones no son completamente multiplataformas, claro, son útiles y funcionales, eso NO lo podemos negar, pero aun así, yo queria algo al estilo Java o Python, y vaya que lo encontre y del mejor autor, la susodicha herramienta de las que les quiero hablar hoy, es el: [Closure Compiler de Google](https://code.google.com/intl/es-AR/closure/).

![image](https://code.google.com/intl/es-AR/closure/images/logo128px.png)    

¿Qué es el Closure Compiler? Es un VERDADERO compilador de código JavaScript escrito 100% en Java. Lo más sorprendente es que realmente compila código JavaScript y lo convierte en mejor código JavaScript. Interpreta JavaScript, lo analiza, remueve código muerto, reescribe rutinas y lo que ya esta optimizado lo "minimiza" (remueve espacios en blanco, comentarios, saltos de linea, etc), verifica sintaxis, referencias de variables y tipos, y ademas, avisa de problemas comunes con código JavaScript. Y por si fuera poco, no solo es Software Libre, sino que ademas, existe una extensión de para Firefox, para emplearlo en conjunto con tu código transformado y [Firebug](https://www.srbyte.com/2009/11/javascript-firebug-jsure.html), llamado: [Closure Inspector](https://closure-inspector.googlecode.com/files/closureinspector09.xpi).

¿Como puedo usarlo?

- Podes utilizar el .jar del Closure Compiler como una aplicacion de linea de comandos (funciona perfecto en Windows y en Linux por supuesto).
- Su funcionalidad puede ser utilizada como un servicio Web
- Y como RESTful API.
¿Beneficios sobre otros compiladores de JavaScript?

- Eficiencia (reduce el tamaño del código JavaScript, para que tu aplicación cargue más rápido y se reduzca tu necesidad de ancho de banda).
- Código Libre (Cero Costo)
- Multiplataforma (Java, Java, Java :)
- Chequeo de código para evitar las típicas pesadillas que da JavaScript.
- Es de Google, y Google lo usa en sus productos... ¿necesito decir más?
Espero que esta herramienta les sirva tanto como a mi. Más adelante explico como eliminar las pequeñas "trabas" que da al usarlo desde la consola. Esta es una herramienta que me atrevo a decir que es casi indispensable al trabajar con código JavaScript, lo recomiendo mucho, tanto como JSure, y si tienen que trabar con mucho código de JS, les va a salvar la vida, y ademas les dará una ventaja competitiva en su trabajo.

¡Saludos!