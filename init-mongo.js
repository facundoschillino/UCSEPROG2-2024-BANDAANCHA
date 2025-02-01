db = db.getSiblingDB("DBP");  //Cambiamos a la base de datos empleada para el programa, si no existe la crea.

//Creamos las colecciones
db.createCollection("tipo_producto");
db.createCollection("producto"); 
db.createCollection("camiones");


db.tipo_producto.insertMany([
    { nombre: 'golosinas', descripcion: '',fecha_creacion: new Date, fecha_actualizacion: new Date, estado: 'activo'},
    { nombre: 'bebidas', descripcion: '',fecha_creacion: new Date, fecha_actualizacion: new Date, estado: 'activo'},
    { nombre: 'cigarrillos', descripcion: '',fecha_creacion: new Date, fecha_actualizacion: new Date, estado: 'activo'},
    { nombre: 'comestibles', descripcion: '',fecha_creacion: new Date, fecha_actualizacion: new Date, estado: 'activo'},
    { nombre: 'higiene y salud', descripcion: '',fecha_creacion: new Date, fecha_actualizacion: new Date, estado: 'activo'},
]
);
