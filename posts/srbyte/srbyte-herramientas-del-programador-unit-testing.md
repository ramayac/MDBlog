---
title: Herramientas del Programador: Unit Testing
date: 2008-12-11
author: Robertux
tags: .net, desarrollo, software, lenguaje, codigo, tecnicas, programacion, herramienta
post_id: blog-3515952828243908885.post-4202301501316594182
---

![image](https://1.bp.blogspot.com/_jH77WNrMVRA/SUAtVVrVQnI/AAAAAAAAFNg/6xIdHcm541w/s400/sourcecode.png)    
"Jajajaja Nunca te librarás de mi,
programador" Bugs. Todo programador los conoce. Desde el primer hola mundo que escribe en C y se da cuenta que no le compiló porque le hacía falta el punto y coma al final de la sentencia, el programador se da cuenta que los programas que escriba nunca serán perfectos y siempre habrá que arreglar los diferentes errores que comunmente se cometen cuando estamos desarrollando alguna aplicación.

Algunos consideran que el proceso de desarrollo de una aplicación se distribuye regularmente en un 20% de tiempo invertido en la codificación contra un 80% invertido en la depuración del código escrito. Esto nos lleva a pensar que corregimos mas de lo que producimos, lo cual no es muy efectivo que digamos a la hora del desarrollo. Para ello, existen técnicas que nos permiten prevenir la ocurrencia de Bugs o errores que comúnmente se cometen a la hora de la programación y entre estas, las más popular es el uso de Unit Testing.

Unit Testing. Es una de las etapas o técnicas que conforman el proceso de programación extrema, la cual consiste en llevar al extremo las habilidades de un programador y el tiempo invertido de manera que se involucre al cliente lo mas posible en el proceso de desarrollo y se entregue lo mas pronto posible. Unit tests son una serie de porciones de código que se escriben con la finalidad de probar y asegurar el correcto funcionamiento de los módulos y clases que conforman tu aplicación. Estas porciones de código, distribuidas en métodos dentro de clases, se ejecutan y devuelven un estado para saber si el test fue pasado con éxito o si este falló. Con los unit tests puedes probar que tu código funciona cuando debe funcionar y falla cuando debe fallar (aunque estas fallas, obviamente estarán controladas, quizá, mediante excepciones).

Por ejemplo, si tienes el siguiente método escrito en java:

> public
> class Comparador {
> public int getMayor(int[] lista){
> int mayor =
> lista[0];
> for(int i=1; imayor)
> mayor = lista[i];
> }
> return mayor;
> }
Podrias construir un Unit Test que se encargue de invocar a ese método pasándole parámetros de muestra y comparando (mediante un Assert) los resultados obtenidos con resultados esperados. Acá es donde el test te avisa si estos son idénticos (el Test pasa) o difieren (El test falla). Asi de simple es como funciona un Unit Test.

> italic;">@Test
> public void
> testGetMayor() {
> Comparador t
> = new Comparador();
> int
> esperado = 5;
> int obtenido =
> t.getMayor(new int[]{1, 2, 3, 4, 2, 3, 5});
> style="font-style: italic;"> assertEquals(esperado, obtenido);
> }

Si el valor esperado es diferente al valor obtenido, el método assertEquals nos avisará de ello mediante un mensaje.

Las ventajas de probar nuestros métodos con Unit Testing es que podemos saber justo después de codificarlos si ellos funcionan como es debido o no, además también podremos conservar los tests y pasarlos cada vez que modifiquemos el código para asegurarnos que los nuevos cambios realizados no afectaron el funcionamiento original de nuestro método y sigue devolviendo los valores esperados.

Con la experiencia, un programador aprende a conocer todas las posibles fallas que se podrían producir en una porción de código y prevenirlas antes que ocurran. Usando Unit Testing se puede comprobar que dichas fallas fueron mitigadas correctamente. Como recomendación, cada programador debe aprender a ser pesimista con su código, a pensar siempre lo peor de manera que su aplicación esté preparada para ello. Buscar todas las posibles fallas que pueda tener el código, sin importar qué tan ridículas, absurdas o imposibles estas sean. Sino, recuerden esa vez que estaban refinando su aplicación unas horas antes de su defensa final y por arte de magia, a última hora todo dejó de funcionar por culpa de un error que pasaron por alto y no creyeron que fuera a afectar tanto el proyecto.

![image](https://4.bp.blogspot.com/_jH77WNrMVRA/SUBVIJTJdgI/AAAAAAAAFNo/cOn0lvCAhPg/s400/murphys_law_poster.jpg)  
"Como decía Murphy: Todo lo
que puede salir mal, saldrá mal"

Para poder hacer uso de los Unit Tests, necesitas instalar el framework apropiado dependiendo del lenguaje de programación que estes utilizando para desarrollar y a veces este ya viene integrado en los IDEs o entornos de desarrollo. Por ejemplo, Java hace uso de la librería [JUnit](https://www.junit.org/), la cual ya viene integrada en [Netbeans](https://www.netbeans.org/), Python hace uso de [PyUnit](https://pyunit.sourceforge.net/), Microsoft .Net utiliza [NUnit](https://www.nunit.org/index.php), el cual se puede integrar con el IDE [SharpDevelop](https://www.icsharpcode.net/OpenSource/SD/), entre otros.

![image](https://4.bp.blogspot.com/_jH77WNrMVRA/SUBb1ATdcxI/AAAAAAAAFNw/tnFVnw_VhAU/s400/NetbeansShot.png)  
"Captura de pantalla del IDE Netbeans, mostrando los resultados de
la ejecución de los Unit Tests en un proyecto de Java (Clic para agrandar)"