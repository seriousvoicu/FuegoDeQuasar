## Challenge de Mercadolibre (Fuego de Quasar)

#### Dependencias necesarias:


$ go get github.com/gorilla/mux 

$ go get github.com/mattn/go-sqlite3 (se requiere GCC https://jmeubank.github.io/tdm-gcc/articles/2020-03/9.2.0-release)



#### Proyecto publicado en:


https://topsecret-dot-challengeml-298601.rj.r.appspot.com


El siguiente JSON devuelve triangulacion y mensaje OK
```
POST /topsecret
{
   "satellites":[
      {
         "name":"kenobi",
         "distance":500.0,
         "message":[
         "",
         "",
            "este",
            "",
            "",
            "mensaje",
            ""
         ]
      },
      {
         "name":"skywalker",
         "distance":140,
         "message":[
            "",
            "es",
            "",
            "",
            "secreto"
         ]
      },
      {
         "name":"sato",
         "distance":582.06,
         "message":[
           "",
            "este",
            "",
            "un",
            "",
            ""
         ]
      }
   ]
}
```

La respuesta es

```
{
   "position":    {
      "x": 0,
      "y": -198
   },
   "message": "este es un mensaje secreto"
}
```

####Observaciones

El programa desde un principio fue planteado para aceptar N satelites aunque no fue probado

Las respuestas en error llegan como http 404 y muestran un mensaje indicando el error y el .go en el cual ocurrió (esto fue util al realizar los testings)