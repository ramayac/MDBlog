---
title: Encriptar y "Hashing".
date: 2007-09-14
author: Rodrigo Amaya
tags: debian, hashing, desencriptar, md5, descifrar, cifrado, gcipher, encriptar
post_id: blog-3515952828243908885.post-7428682289194336522
---

Cansado de que tanto los docentes, como los alumnos de mi universidad, usen casi CONSTANTEMENTE el termino de encriptar y hashing como si fueran lo mismo, me propongo (y espero) aclarar los términos (a groso modo), para el publico confundido en general :)

¿Qué es encriptar? En primer lugar la palabra "encriptar" ni siquiera debería de ser usada en español, el termino es una traducción literal del ingles, para este oscuro proceso siempre se había empleado el termino de CIFRADO. Pero bien, le seguiremos diciendo "Encriptado". En la ciencia de la criptografía, la encriptación es el proceso de TRANSFORMAR información para hacerla completamente ILEGIBLE para cualquier persona, excepto para aquel que tengan un conocimiento especial ... a este conocimiento especial le llamaremos LLAVE (de encriptación).

![image](https://bp2.blogger.com/_ayvorITawE4/Ruqm6Gy6MpI/AAAAAAAAAcc/tS_6OU2Zbbw/s320/mortal.png)    
"¿Se recuerdan de esto en
Mortal Kombat?"

El resultado de este proceso es información encriptada. En muuuuchos contextos, la palabra encriptar también se refiere al proceso inverso: desencriptar, para poder hacer que la información que antes era ILEGIBLE, vuelva a su condición original (gracias a la LLAVE).

En resumen, al encriptar se convierte o transforma la información, pero este proceso es reversible. Veamos un ejemplo:

> "Iholc gld gh od lqghshqghfld ohfwruhv
> Vdoydgruhñrv"
El texto anterior esta encriptado con el "[código Cesar](https://es.wikipedia.org/wiki/Cifrado_C%C3%A9sar)
". Este código cesar es la llave para desencriptar el texto: Se cuenta que Julio Cesar invento esta forma de encriptacion por si capturaban a un mensajero que lleva ordenes militares, los capturadores no pudieran leer el mensaje (y claro, solo los generales del Cesar tenian el conocimiento para hacerlo). ¿Y cual es este gran proceso para desencriptar el mensaje anterior? Sencillo, tomemos una letra: h, ahora nos vamos tres letras atras en el abecedario, eso nos da la letra e.

![image](https://bp3.blogger.com/_ayvorITawE4/RuqooWy6MrI/AAAAAAAAAcs/u6NNmtQDDF4/s320/320px-Caesar3.svg.png)    
"código de cifrado del
César"

El mismo proceso se repite para las letras mayusculas y minusculas.

Los programadores también podrían hacer algo así como (en Mono C#):

char LetraACesar(char letra) {

return System.Convert.ToChar(char.GetNumericValue(letra) + 3); }

Asi, la primera palabra seria:

> Iholc
> = Feliz
Y el resto del mensaje.... lo pueden descifrar con GCipher ;)
# apt-get install gcipher

¿Y que es Hashing? Hashing es la tecnica para convertir algun tipo de dato en un (relativo) numero pequeño, parece lo mismo que encriptar... pero la diferencia radica en que ese número, es la "huella digital" o "firma digital" de los datos. El algoritmo de hashing ([MD5](https://es.wikipedia.org/wiki/MD5) de un algoritmo de hashing) puede "cortar y mezclar" (substituye y superpone) los datos para crear la huella. Estas huellas del archivo pueden ser llamadas sumatorias hash, valores hash, codigos hash o simplemente hashes.

Por ejemplo, cuando bajas un DVD de Debian (o cualquier DVD de internet en general) y necesitas comprobar la integridad del mismo antes de quemarlo, simplemente usas MD5 para comprobar que los DVD que bajaste se encuentrar en perfecto estado, listos para ser "tostados". A continuacion, la firma o huella digital MD5 de cada DVD de Debian Testing (28 Mayo):

```
b2c4ac6cd5a7ef02f462851401ef91f7 debian-testing-i386-DVD-1.iso c60aa22abd48dc08a977179cc2bd430c debian-testing-i386-DVD-2.iso cb41104941fd2a1132a0b7c31e5f1c1c debian-testing-i386-DVD-3.iso 3a49cb2a95485a08b646f044b8417f92 debian-testing-i386-DVD-4.iso
```
Pero observen que, a partir del hash MD5 es CASI IMPOSIBLE generar un DVD completo de Debian Testing.

> ¡El proceso no es
> reversible!
A menos que .... Generemos un archivo de 4.3G, con datos completamente aleatorios y generemos su hash MD5 para ver si concuerda con el primer hash MD5 del primer DVD y repitamos este proceso las veces que sean necesarias (con brute force attack).... las probabilidades, son bastante remotas. Así que dejemos lo en CASI IMPOSIBLE... o podríamos terminar asi:

![image](https://bp3.blogger.com/_ayvorITawE4/Ruqm_Wy6MqI/AAAAAAAAAck/SjnoEtN5rME/s320/death.png)    
"..."

Esa aproximación jamas funcionaria para un archivo de 4.3 GB.... pero si para.... algo más